package hotel

import (
	"context"
	"strings"
	"time"

	dbmodels "travelgo/internal/models/db"
	hotelrepo "travelgo/internal/repository/hotel"
)

type Service struct {
	repo *hotelrepo.Repository
}

func New(repo *hotelrepo.Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) CreateHotel(ctx context.Context, name, description, imageUrl string) (int64, error) {
	return s.repo.CreateHotel(ctx, name, description, imageUrl)
}

func (s *Service) GetHotelByID(ctx context.Context, id int64) (*dbmodels.Hotel, error) {
	return s.repo.GetHotelByID(ctx, id)
}

func (s *Service) UpdateHotel(ctx context.Context, id int64, hotel *dbmodels.Hotel) (*dbmodels.Hotel, error) {
	return s.repo.UpdateHotel(ctx, id, hotel)
}

func (s *Service) DeleteHotel(ctx context.Context, id int64) error {
	return s.repo.DeleteHotel(ctx, id)
}

func (s *Service) GetAvailableHotels(ctx context.Context, city string, dateFrom, dateTo time.Time) ([]*dbmodels.Hotel, error) {
	city = strings.ToLower(city)
	return s.repo.GetAvailableHotels(ctx, city, dateFrom, dateTo)
}

func (s *Service) Book(
	ctx context.Context,
	hotelId int64, dateFrom, dateTo time.Time,
) error {
	return s.repo.Book(ctx, hotelId, dateFrom, dateTo)
}
