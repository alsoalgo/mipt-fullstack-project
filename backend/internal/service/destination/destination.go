package destination

import (
	"context"

	dbmodels "travelgo/internal/models/db"
	destinationrepo "travelgo/internal/repository/destination"
)

type Service struct {
	repo *destinationrepo.Repository
}

func New(repo *destinationrepo.Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) CreateDestination(ctx context.Context, city string, cost int32, imageUrl string) (int64, error) {
	return s.repo.CreateDestination(ctx, city, cost, imageUrl)
}

func (s *Service) GetPopularDestinations(ctx context.Context) ([]*dbmodels.Destination, error) {
	return s.repo.GetPopularDestinations(ctx)
}
