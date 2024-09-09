package dbmodels

import (
	"time"

	httpmodels "travelgo/internal/models/http"
)

type HotelOrder struct {
	ID        int64     `db:"id"`
	UserID    int64     `db:"user_id"`
	HotelID   int64     `db:"hotel_id"`
	DateFrom  time.Time `db:"date_from"`
	DateTo    time.Time `db:"date_to"`
	FirsName  string    `db:"first_name"`
	LastName  string    `db:"last_name"`
	SurName   string    `db:"sur_name"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func HotelOrderHTTP(hotel *Hotel, ho *HotelOrder) *httpmodels.HotelOrder {
	return &httpmodels.HotelOrder{
		Hotel: httpmodels.Hotel{
			Title:       hotel.Title,
			City:        hotel.City,
			Description: hotel.Description,
			ImageURL:    hotel.ImageURL,
		},
		Details: httpmodels.OrderDetails{
			DateFrom:  ho.DateFrom.Format(time.DateOnly),
			DateTo:    ho.DateTo.Format(time.DateOnly),
			FirstName: ho.FirsName,
			LastName:  ho.LastName,
			SurName:   ho.SurName,
		},
	}
}
