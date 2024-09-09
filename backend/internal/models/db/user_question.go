package dbmodels

import (
	"time"

	httpmodels "travelgo/internal/models/http"
)

type UserQuestion struct {
	ID        int64     `db:"id"`
	UserID    string    `db:"user_id"`
	Title     string    `db:"title"`
	Question  string    `db:"question"`
	Status    string    `db:"status"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func (uq *UserQuestion) ToHTTP() *httpmodels.UserQuestion {
	return &httpmodels.UserQuestion{
		Title:    uq.Title,
		Question: uq.Question,
	}
}
