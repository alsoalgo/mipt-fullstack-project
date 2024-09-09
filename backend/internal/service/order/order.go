package order

import (
	"context"
	"time"

	dbmodels "travelgo/internal/models/db"
	orderrepo "travelgo/internal/repository/order"
)

type Service struct {
	repo *orderrepo.Repository
}

func New(repo *orderrepo.Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) CreateOrder(
	ctx context.Context, userId, hotelId int64, dateFrom, dateTo time.Time,
	firstName, lastName, surName string) (int64, error) {
	return s.repo.CreateOrder(ctx, userId, hotelId, dateFrom, dateTo, firstName, lastName, surName)
}

func (s *Service) GetOrdersByUserID(ctx context.Context, userId int64) ([]*dbmodels.HotelOrder, error) {
	return s.repo.GetOrdersByUserID(ctx, userId)
}
