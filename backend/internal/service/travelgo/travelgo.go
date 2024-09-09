package travelgo

import (
	"context"
	"errors"
	"math/rand"
	"time"

	"github.com/jmoiron/sqlx"

	dbmodels "travelgo/internal/models/db"
	httpmodels "travelgo/internal/models/http"
	"travelgo/internal/service/auth"
	"travelgo/internal/service/destination"
	"travelgo/internal/service/hotel"
	"travelgo/internal/service/order"
	question "travelgo/internal/service/question"
	"travelgo/internal/service/user"
)

type TravelGoService struct {
	db *sqlx.DB

	question    *question.Service
	user        *user.Service
	hotel       *hotel.Service
	order       *order.Service
	auth        *auth.Service
	destination *destination.Service
}

func NewTravelGoService(
	db *sqlx.DB,
	question *question.Service,
	user *user.Service,
	hotel *hotel.Service,
	order *order.Service,
	auth *auth.Service,
	destination *destination.Service,
) *TravelGoService {
	return &TravelGoService{
		db:          db,
		question:    question,
		user:        user,
		hotel:       hotel,
		order:       order,
		auth:        auth,
		destination: destination,
	}
}

func (s *TravelGoService) Login(ctx context.Context, req *httpmodels.LoginRequest) (*httpmodels.LoginResponse, error) {
	ctx = context.WithValue(ctx, "opName", "Login")

	token, err := s.auth.Login(ctx, req)
	if err != nil {
		return nil, errors.New("Login: " + err.Error())
	}

	return &httpmodels.LoginResponse{Token: token}, nil
}

func (s *TravelGoService) Register(ctx context.Context, req *httpmodels.RegisterRequest) (*httpmodels.RegisterResponse, error) {
	ctx = context.WithValue(ctx, "opName", "Register")

	id, err := s.auth.Register(ctx, req)
	if err != nil {
		return nil, errors.New("Register: " + err.Error())
	}

	return &httpmodels.RegisterResponse{Registered: id != int64(0)}, nil
}

func (s *TravelGoService) CheckToken(ctx context.Context, req *httpmodels.CheckTokenRequest) (*httpmodels.CheckTokenResponse, error) {
	ctx = context.WithValue(ctx, "opName", "Register")

	exists, err := s.auth.CheckToken(ctx, req)
	if err != nil {
		return nil, errors.New("CheckToken: " + err.Error())
	}

	return &httpmodels.CheckTokenResponse{Exists: exists}, nil
}

func (s *TravelGoService) CreateQuestion(ctx context.Context, req *httpmodels.CreateQuestionRequest) (*httpmodels.CreateQuestionResponse, error) {
	ctx = context.WithValue(ctx, "opName", "CreateQuestion")

	err := s.question.CreateQuestion(ctx, req.UserID, req.Title, req.Question)
	if err != nil {
		return nil, errors.New("CreateQuestion: " + err.Error())
	}

	return &httpmodels.CreateQuestionResponse{Created: true}, nil
}

func (s *TravelGoService) GetQuestions(ctx context.Context, req *httpmodels.GetQuestionsRequest) (*httpmodels.GetQuestionsResponse, error) {
	ctx = context.WithValue(ctx, "opName", "GetQuestions")

	dbQuestions, err := s.question.GetQuestionsByUserID(ctx, req.UserID)
	if err != nil {
		return nil, errors.New("GetOrdersByUserID: " + err.Error())
	}

	questions := make([]*httpmodels.UserQuestion, 0)
	for _, q := range dbQuestions {
		questions = append(questions, q.ToHTTP())
	}

	return &httpmodels.GetQuestionsResponse{
		Questions: questions,
	}, nil
}

func (s *TravelGoService) Search(ctx context.Context, req *httpmodels.SearchRequest) (*httpmodels.SearchResponse, error) {
	ctx = context.WithValue(ctx, "opName", "Search")

	dbHotels, err := s.hotel.GetAvailableHotels(ctx, req.City, req.DateFrom, req.DateTo)
	if err != nil {
		return nil, errors.New("GetAvailableHotels: " + err.Error())
	}

	hotels := make([]*httpmodels.Hotel, 0)
	for _, h := range dbHotels {
		hotels = append(hotels, h.ToHTTP())
	}

	return &httpmodels.SearchResponse{
		Hotels: hotels,
	}, nil
}

func (s *TravelGoService) CreateOrder(ctx context.Context, req *httpmodels.CreateOrderRequest) (*httpmodels.CreateOrderResponse, error) {
	ctx = context.WithValue(ctx, "opName", "CreateOrder")

	dateFrom, err := time.Parse(time.DateOnly, req.DateFrom)
	if err != nil {
		return nil, errors.New("Parse: " + err.Error())
	}

	dateTo, err := time.Parse(time.DateOnly, req.DateTo)
	if err != nil {
		return nil, errors.New("Parse: " + err.Error())
	}

	err = s.hotel.Book(ctx, req.HotelID, dateFrom, dateTo)
	if err != nil {
		return nil, errors.New("Book: " + err.Error())
	}

	_, err = s.order.CreateOrder(ctx, req.UserID, req.HotelID, dateFrom, dateTo, req.FirstName, req.LastName, req.SurName)
	if err != nil {
		return nil, errors.New("CreateOrder: " + err.Error())
	}

	return &httpmodels.CreateOrderResponse{
		Created: true,
	}, nil
}

func (s *TravelGoService) GerOrders(ctx context.Context, req *httpmodels.GetOrdersRequest) (*httpmodels.GetOrdersResponse, error) {
	ctx = context.WithValue(ctx, "opName", "GerOrders")

	dbOrders, err := s.order.GetOrdersByUserID(ctx, req.UserID)
	if err != nil {
		return nil, errors.New("GetOrdersByUserID: " + err.Error())
	}

	hotels := make(map[int64]*dbmodels.Hotel, 0)
	for _, o := range dbOrders {
		_, ok := hotels[o.HotelID]
		if ok {
			continue
		}

		hotel, err := s.hotel.GetHotelByID(ctx, o.HotelID)
		if err != nil {
			return nil, errors.New("GetHotelByID: " + err.Error())
		}

		hotels[o.HotelID] = hotel
	}

	orders := make([]*httpmodels.HotelOrder, 0)
	for _, o := range dbOrders {
		orders = append(orders, dbmodels.HotelOrderHTTP(hotels[o.HotelID], o))
	}

	return &httpmodels.GetOrdersResponse{
		Orders: orders,
	}, nil
}

func (s *TravelGoService) EditProfile(ctx context.Context, req *httpmodels.EditProfileRequest) (*httpmodels.EditProfileResponse, error) {
	ctx = context.WithValue(ctx, "opName", "EditProfile")

	edited, err := s.user.EditProfile(ctx, req.UserID, req.FirstName, req.LastName, req.SurName)
	if err != nil {
		return nil, err
	}

	return &httpmodels.EditProfileResponse{Edited: edited}, nil
}

func (s *TravelGoService) GetProfile(ctx context.Context, req *httpmodels.GetProfileRequest) (*httpmodels.GetProfileResponse, error) {
	ctx = context.WithValue(ctx, "opName", "EditProfile")

	user, err := s.user.GetUserByID(ctx, req.UserID)
	if err != nil {
		return nil, errors.New("GetUserByID: " + err.Error())
	}

	return &httpmodels.GetProfileResponse{
		Info: user,
	}, nil
}

func (s *TravelGoService) GetPopularDestinations(ctx context.Context, req *httpmodels.GetPopularDestinationsRequest) (*httpmodels.GetPopularDestinationsResponse, error) {
	ctx = context.WithValue(ctx, "opName", "GetPopularDestinations")

	dbDestinations, err := s.destination.GetPopularDestinations(ctx)
	if err != nil {
		return nil, errors.New("GetPopularDestinations: " + err.Error())
	}

	destinations := make([]*httpmodels.Destination, 0)
	for _, d := range dbDestinations {
		destinations = append(destinations, d.ToHTTP())
	}

	for i := range destinations {
		j := rand.Intn(i + 1)
		destinations[i], destinations[j] = destinations[j], destinations[i]
	}

	return &httpmodels.GetPopularDestinationsResponse{
		Destinations: destinations,
	}, nil
}

func (s *TravelGoService) GetHotel(ctx context.Context, req *httpmodels.GetHotelRequest) (*httpmodels.GetHotelResponse, error) {
	ctx = context.WithValue(ctx, "opName", "GetHotel")

	dbHotel, err := s.hotel.GetHotelByID(ctx, req.ID)
	if err != nil {
		return nil, errors.New("GetHotelByID: " + err.Error())
	}

	return &httpmodels.GetHotelResponse{
		Hotel: dbHotel.ToHTTP(),
	}, nil
}
