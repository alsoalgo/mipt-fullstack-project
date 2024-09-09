package user

import (
	"context"
	"errors"

	dbmodels "travelgo/internal/models/db"
	httpmodels "travelgo/internal/models/http"
	userrepo "travelgo/internal/repository/user"
)

type Service struct {
	userRepo *userrepo.Repository
}

func New(userRepo *userrepo.Repository) *Service {
	return &Service{
		userRepo: userRepo,
	}
}

func (s *Service) CreateUser(ctx context.Context, email, passwordHash string) (int64, error) {
	return s.userRepo.CreateUser(ctx, email, passwordHash)
}

func (s *Service) GetUserByEmail(ctx context.Context, email string) (*dbmodels.User, error) {
	return s.userRepo.GetUserByEmail(ctx, email)
}

func (s *Service) GetUserByID(ctx context.Context, id int64) (*httpmodels.User, error) {
	ctx = context.WithValue(ctx, "opName", "EditProfile")

	dbUser, err := s.userRepo.GetUserByID(ctx, id)
	if err != nil {
		return nil, errors.New("GetUserByID: " + err.Error())
	}

	return dbUser.ToHTTP(), nil
}

func (s *Service) EditProfile(ctx context.Context, id int64, firstName, lastName, surName string) (bool, error) {
	ctx = context.WithValue(ctx, "opName", "EditProfile")

	_, err := s.userRepo.UpdateUser(ctx, id, firstName, lastName, surName)
	if err != nil {
		return false, errors.New("UpdateUser: " + err.Error())
	}

	return true, nil
}
