package tokenrepo

import (
	"context"
	"strings"

	"github.com/jmoiron/sqlx"
)

type Repository struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) AddToken(ctx context.Context, userId int64, token string) (int64, error) {
	ctx = context.WithValue(ctx, "opName", "AddToken")

	q := `INSERT INTO tokens(user_id, token) VALUES ($1, $2) RETURNING id`

	id := int64(0)
	err := r.db.GetContext(ctx, &id, q, userId, token)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *Repository) IsTokenExpired(ctx context.Context, token string) (bool, error) {
	ctx = context.WithValue(ctx, "opName", "IsTokenExpired")

	q := `SELECT COUNT(id) as count FROM tokens WHERE token = $1`

	rs := 0
	err := r.db.GetContext(ctx, &rs, q, token)
	if err != nil {
		return false, err
	}

	return rs > 0, nil
}

func (r *Repository) CleanExpiredTokens(ctx context.Context) error {
	ctx = context.WithValue(ctx, "opName", "CleanExpiredTokens")

	q := `DELETE FROM tokens WHERE expires_at < now()`

	_, err := r.db.ExecContext(ctx, q)
	if err != nil {
		return err
	}

	return nil
}

func getColumns() string {
	return strings.Join(columns, ",")
}

var columns = []string{
	"id",
	"token",
	"user_id",
	"created_at",
	"updated_at",
	"expires_at",
}
