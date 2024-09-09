package dbmodels

import (
	"time"

	httpmodels "travelgo/internal/models/http"
)

type Hotel struct {
	ID          int64     `db:"id"`
	Title       string    `db:"title"`
	City        string    `db:"city"`
	Description string    `db:"description"`
	ImageURL    string    `db:"image_url"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}

func (h *Hotel) ToHTTP() *httpmodels.Hotel {
	return &httpmodels.Hotel{
		ID:          h.ID,
		Title:       h.Title,
		City:        h.City,
		Description: h.Description,
		ImageURL:    h.ImageURL,
	}
}
