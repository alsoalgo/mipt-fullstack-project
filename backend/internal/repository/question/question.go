package question

import (
	"context"
	"strings"

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

func (r *Repository) CreateQuestion(ctx context.Context, userId int64, title string, question string) error {
	ctx = context.WithValue(ctx, "opName", "CreateQuestion")

	q := `INSERT INTO user_question (user_id, title, question) VALUES ($1, $2, $3) RETURNING id`

	id := int64(0)
	err := r.db.GetContext(ctx, &id, q, userId, title, question)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) GetQuestionsByStatus(ctx context.Context, status string) ([]*dbmodels.UserQuestion, error) {
	ctx = context.WithValue(ctx, "opName", "GetQuestionsByStatus")

	q := `SELECT ` + getColumns() + ` FROM user_question WHERE status = $1`

	questions := make([]*dbmodels.UserQuestion, 0)
	err := r.db.SelectContext(ctx, &questions, q, status)
	if err != nil {
		return nil, err
	}

	return questions, nil
}

func (r *Repository) GetQuestionsByUserID(ctx context.Context, userId int64) ([]*dbmodels.UserQuestion, error) {
	ctx = context.WithValue(ctx, "opName", "GetQuestionsByUserID")

	q := `SELECT ` + getColumns() + ` FROM user_question WHERE user_id = $1`

	questions := make([]*dbmodels.UserQuestion, 0)
	err := r.db.SelectContext(ctx, &questions, q, userId)
	if err != nil {
		return nil, err
	}

	return questions, nil
}

func getColumns() string {
	return strings.Join(columns, ", ")
}

var columns = []string{
	"id",
	"user_id",
	"title",
	"question",
	"status",
	"created_at",
	"updated_at",
}
