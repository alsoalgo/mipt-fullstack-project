package userrepo

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

func (r *Repository) CreateUser(
	ctx context.Context,
	email, passwordHash string,
) (int64, error) {
	ctx = context.WithValue(ctx, "opName", "CreateUser")

	q := `INSERT INTO users (email, password_hash) VALUES ($1, $2) RETURNING id`

	id := int64(0)
	err := r.db.GetContext(ctx, &id, q, email, passwordHash)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *Repository) GetUserByEmail(ctx context.Context, email string) (*dbmodels.User, error) {
	ctx = context.WithValue(ctx, "opName", "GetUserByEmail")

	q := `SELECT ` + getColumns() + ` FROM users WHERE email = $1`

	user := dbmodels.User{}
	err := r.db.GetContext(ctx, &user, q, email)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *Repository) GetUserByID(ctx context.Context, id int64) (*dbmodels.User, error) {
	ctx = context.WithValue(ctx, "opName", "GetUserByID")

	q := `SELECT ` + getColumns() + ` FROM users WHERE id = $1`

	user := dbmodels.User{}
	err := r.db.GetContext(ctx, &user, q, id)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *Repository) UpdateUser(ctx context.Context, userId int64, firstName, lastName, surName string) (*dbmodels.User, error) {
	ctx = context.WithValue(ctx, "opName", "UpdateUser")

	q := `
		UPDATE users 
		SET first_name = $1, 
			last_name = $2, 
			sur_name = $3
		WHERE id = $4 
		RETURNING ` + getColumns()

	nUser := dbmodels.User{}
	err := r.db.GetContext(ctx, &nUser, q, firstName, lastName, surName, userId)
	if err != nil {
		return nil, err
	}

	return &nUser, nil
}

func (r *Repository) ExistsUser(ctx context.Context, email string) (bool, error) {
	ctx = context.WithValue(ctx, "opName", "ExistsUser")

	q := `SELECT EXISTS(SELECT ` + getColumns() + ` FROM users WHERE email = $1)`

	exists := bool(false)

	err := r.db.GetContext(ctx, &exists, q, email)
	if err != nil {
		return false, err
	}

	return exists, nil
}

func (r *Repository) GetUserByToken(ctx context.Context, token string) (*dbmodels.User, error) {
	ctx = context.WithValue(ctx, "opName", "GetUserByToken")

	q := `
		SELECT ` + getColumns() + ` 
		FROM users 
		WHERE id = (
			SELECT user_id 
			FROM tokens 
			WHERE token = $1
		)`

	user := dbmodels.User{}

	err := r.db.GetContext(ctx, &user, q, token)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *Repository) DeleteUser(ctx context.Context, id int64) error {
	ctx = context.WithValue(ctx, "opName", "DeleteUser")

	q := `DELETE FROM users WHERE id = $1`

	_, err := r.db.ExecContext(ctx, q, id)
	if err != nil {
		return err
	}

	return nil
}

func getColumns() string {
	return strings.Join(columns, ", ")
}

var columns = []string{
	"id",
	"email",
	"password_hash",
	"user_role",
	"first_name",
	"last_name",
	"sur_name",
	"created_at",
	"updated_at",
}
