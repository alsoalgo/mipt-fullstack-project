package dbmodels

import (
	"time"

	httpmodels "travelgo/internal/models/http"
)

type Destination struct {
	ID        int64     `db:"id"`
	City      string    `db:"city"`
	Cost      int32     `db:"cost"`
	ImageURL  string    `db:"image_url"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func (d *Destination) ToHTTP() *httpmodels.Destination {
	return &httpmodels.Destination{
		City:     d.City,
		Cost:     d.Cost,
		ImageURL: d.ImageURL,
	}
}
