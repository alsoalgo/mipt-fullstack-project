package dbmodels

import (
	"database/sql"
	"time"

	httpmodels "travelgo/internal/models/http"
)

type User struct {
	ID           int64          `db:"id"`
	Email        string         `db:"email"`
	PasswordHash string         `db:"password_hash"`
	Role         string         `db:"user_role"`
	FirstName    sql.NullString `db:"first_name"`
	LastName     sql.NullString `db:"last_name"`
	SurName      sql.NullString `db:"sur_name"`
	CreatedAt    time.Time      `db:"created_at"`
	UpdatedAt    time.Time      `db:"updated_at"`
}

func (u *User) ToHTTP() *httpmodels.User {
	firstName := ""
	if u.FirstName.Valid {
		firstName = u.FirstName.String
	}

	lastName := ""
	if u.LastName.Valid {
		lastName = u.LastName.String
	}

	surName := ""
	if u.SurName.Valid {
		surName = u.SurName.String
	}

	return &httpmodels.User{
		FirstName: firstName,
		LastName:  lastName,
		SurName:   surName,
	}
}
