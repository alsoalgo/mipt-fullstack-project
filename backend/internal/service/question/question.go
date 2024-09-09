package question

import (
	"context"

	dbmodels "travelgo/internal/models/db"
	questionrepo "travelgo/internal/repository/question"
)

type Service struct {
	repo *questionrepo.Repository
}

func New(repo *questionrepo.Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) CreateQuestion(ctx context.Context, userId int64, topic, question string) error {
	return s.repo.CreateQuestion(ctx, userId, topic, question)
}

func (s *Service) GetQuestionsByStatus(ctx context.Context, status string) ([]*dbmodels.UserQuestion, error) {
	return s.repo.GetQuestionsByStatus(ctx, status)
}

func (s *Service) GetQuestionsByUserID(ctx context.Context, userId int64) ([]*dbmodels.UserQuestion, error) {
	return s.repo.GetQuestionsByUserID(ctx, userId)
}
