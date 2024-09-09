package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	chi "github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	config "travelgo/configs"
	handlers "travelgo/internal/handler/travelgo"
	destinationrepo "travelgo/internal/repository/destination"
	hotelrepo "travelgo/internal/repository/hotel"
	orderrepo "travelgo/internal/repository/order"
	questionrepo "travelgo/internal/repository/question"
	tokenrepo "travelgo/internal/repository/token"
	userrepo "travelgo/internal/repository/user"
	"travelgo/internal/service/auth"
	"travelgo/internal/service/destination"
	"travelgo/internal/service/hotel"
	"travelgo/internal/service/order"
	"travelgo/internal/service/question"
	"travelgo/internal/service/travelgo"
	"travelgo/internal/service/user"
)

func main() {
	_ = context.Background()

	env, ok := os.LookupEnv("ENV")
	mustBeOk("env not found", ok)

	cfg, err := config.Parse(env)
	mustInit("config", err)

	db := connectDB(cfg)
	defer db.Close()

	hrepo := hotelrepo.New(db)
	hotel := hotel.New(hrepo)

	orepo := orderrepo.New(db)
	order := order.New(orepo)

	urepo := userrepo.New(db)
	user := user.New(urepo)

	qrepo := questionrepo.New(db)
	question := question.New(qrepo)

	trepo := tokenrepo.New(db)
	auth := auth.New(urepo, trepo, "secret")

	/*
		cleaner := cleaner.New(trepo)
		go cleaner.Start(ctx, 1)
		defer cleaner.Stop()
	*/

	drepo := destinationrepo.New(db)
	destination := destination.New(drepo)

	travelGoService := travelgo.NewTravelGoService(
		db,
		question,
		user,
		hotel,
		order,
		auth,
		destination,
	)

	handlers := handlers.NewTravelGoHandler(travelGoService)

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))
	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	r.Route("/api/v1", func(r chi.Router) {
		r.Use(middleware.AllowContentType("application/json"))

		r.Post("/login", handlers.Login)
		r.Post("/register", handlers.Register)
		r.Post("/check", handlers.Check)

		r.Group(func(r chi.Router) {
			r.Use(auth.IsAuthorized())

			r.Post("/question", handlers.CreateQuestion)
			r.Post("/search", handlers.Search)
			r.Post("/order", handlers.CreateOrder)
			r.Post("/hotel", handlers.GetHotel)
			r.Post("/profile/edit", handlers.EditProfile)
		})

		r.Group(func(r chi.Router) {
			r.Use(auth.IsAuthorized())

			r.Get("/profile", handlers.GetProfile)
			r.Get("/orders", handlers.GetOrders)
			r.Get("/questions", handlers.GetQuestions)
			r.Get("/destinations/popular", handlers.GetPopularDestinations)
		})
	})

	log.Println("Server started on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func connectDB(cfg *config.Config) *sqlx.DB {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.DB.Postgres.Host,
		cfg.DB.Postgres.Port,
		cfg.DB.Postgres.User,
		cfg.DB.Postgres.Password,
		cfg.DB.Postgres.DBName,
	)

	db, err := sql.Open("postgres", psqlconn)
	mustInit("db", err)

	dbx := sqlx.NewDb(db, "postgres")
	mustInit("db ping", dbx.Ping())

	return dbx
}

func mustBeOk(msg string, ok bool) {
	if !ok {
		panic(msg)
	}
}

func mustInit(msg string, err error) {
	if err != nil {
		panic(msg + err.Error())
	}
}
