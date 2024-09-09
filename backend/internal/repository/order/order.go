package orderrepo

import (
	"context"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"

	dbmodels "travelgo/internal/models/db"
)

type Repository struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) CreateOrder(
	ctx context.Context,
	userId, hotelId int64, dateFrom, dateTo time.Time,
	firtsName, lastName, surName string,
) (int64, error) {
	ctx = context.WithValue(ctx, "opName", "CreateOrder")

	q := `INSERT INTO hotel_order (
			user_id, hotel_id, date_from, date_to, 
			first_name, last_name, sur_name
		) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`

	id := int64(0)
	err := r.db.GetContext(ctx, &id, q, userId, hotelId, dateFrom, dateTo, firtsName, lastName, surName)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *Repository) GetOrderByID(ctx context.Context, id int64) (*dbmodels.HotelOrder, error) {
	ctx = context.WithValue(ctx, "opName", "GetOrderByID")

	q := `SELECT ` + getColumns() + ` FROM hotel_order WHERE id = $1`

	order := &dbmodels.HotelOrder{}
	err := r.db.GetContext(ctx, &order, q, id)
	if err != nil {
		return nil, err
	}

	return order, nil
}

func (r *Repository) GetOrdersByUserID(ctx context.Context, id int64) ([]*dbmodels.HotelOrder, error) {
	ctx = context.WithValue(ctx, "opName", "GetOrdersByUserID")

	q := `SELECT ` + getColumns() + ` FROM hotel_order WHERE user_id = $1`

	orders := make([]*dbmodels.HotelOrder, 0)
	err := r.db.SelectContext(ctx, &orders, q, id)
	if err != nil {
		return nil, err
	}

	return orders, nil
}

func getColumns() string {
	return strings.Join(columns, ",")
}

var columns = []string{
	"id",
	"user_id",
	"hotel_id",
	"date_from",
	"date_to",
	"first_name",
	"last_name",
	"sur_name",
	"created_at",
	"updated_at",
}
