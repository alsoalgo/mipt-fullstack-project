package hotelrepo

import (
	"context"
	"strings"
	"time"

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

func (r *Repository) CreateHotel(ctx context.Context, title, description, imageUrl string) (int64, error) {
	ctx = context.WithValue(ctx, "opName", "CreateHotel")

	q := `INSERT INTO hotel (title, description, image_url) VALUES ($1, $2, $3) RETURNING id`

	id := int64(0)
	err := r.db.GetContext(ctx, &id, q, title, description, imageUrl)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *Repository) GetHotelByID(ctx context.Context, id int64) (*dbmodels.Hotel, error) {
	ctx = context.WithValue(ctx, "opName", "GetHotelByID")

	q := `SELECT ` + getHotelColumns() + ` FROM hotel WHERE id = $1`

	hotel := dbmodels.Hotel{}
	err := r.db.GetContext(ctx, &hotel, q, id)
	if err != nil {
		return nil, err
	}

	return &hotel, nil
}

func (r *Repository) UpdateHotel(ctx context.Context, id int64, hotel *dbmodels.Hotel) (*dbmodels.Hotel, error) {
	ctx = context.WithValue(ctx, "opName", "UpdateHotel")

	q := `UPDATE hotel SET title = $1, description = $2 WHERE id = $3 RETURNING ` + getHotelColumns()

	nHotel := dbmodels.Hotel{}
	err := r.db.GetContext(ctx, &nHotel, q, hotel.Title, hotel.Description, id)
	if err != nil {
		return nil, err
	}

	return &nHotel, nil
}

func (r *Repository) DeleteHotel(ctx context.Context, id int64) error {
	ctx = context.WithValue(ctx, "opName", "DeleteHotel")

	q := `DELETE FROM hotel WHERE id = $1`

	_, err := r.db.ExecContext(ctx, q, id)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) GetHotels(ctx context.Context) ([]*dbmodels.Hotel, error) {
	ctx = context.WithValue(ctx, "opName", "GetHotels")

	q := `SELECT ` + getHotelColumns() + ` FROM hotel`

	hotels := make([]*dbmodels.Hotel, 0)
	err := r.db.SelectContext(ctx, &hotels, q)
	if err != nil {
		return nil, err
	}

	return hotels, nil
}

func (r *Repository) GetAvailableHotels(
	ctx context.Context,
	city string,
	dateFrom time.Time,
	dateTo time.Time,
) ([]*dbmodels.Hotel, error) {
	ctx = context.WithValue(ctx, "opName", "GetAvailableHotels")

	q := `
	WITH available_hotels as (
			SELECT a.hotel_id, ($3::date - $2::date) + 1 AS date_diff
			FROM available_room a
			LEFT JOIN hotel h ON a.hotel_id = h.id
			WHERE h.city = $1::text
			AND a.available_date >= $2::date
			AND a.available_date <= $3::date
			AND a.room_count > 0
	), available_days AS (
		SELECT DISTINCT hotel_id, count(hotel_id) AS count FROM available_hotels
		GROUP BY hotel_id
	)
	SELECT ahi.id, h.city, h.title, h.description, h.image_url, h.created_at, h.updated_at FROM (
		SELECT DISTINCT b.hotel_id AS id, d.count AS count
		FROM available_hotels b LEFT JOIN available_days d ON b.hotel_id = d.hotel_id
		WHERE b.date_diff = d.count
	) AS ahi
	LEFT JOIN hotel h ON ahi.id = h.id
	`
	// $1 - city, $2 - from, $3 - to

	hotels := make([]*dbmodels.Hotel, 0)
	err := r.db.SelectContext(ctx, &hotels, q, city, dateFrom, dateTo)
	if err != nil {
		return nil, err
	}

	return hotels, nil
}

func (r *Repository) Book(
	ctx context.Context,
	hotelId int64,
	dateFrom time.Time,
	dateTo time.Time,
) error {
	ctx = context.WithValue(ctx, "opName", "Book")

	q := `
	UPDATE available_room
	SET
		room_count = room_count - 1
	WHERE
		available_date >= $1 AND available_date <= $2
		AND hotel_id = $3
	`

	_, err := r.db.ExecContext(ctx, q, dateFrom, dateTo, hotelId)
	if err != nil {
		return err
	}

	return nil
}

func getHotelColumns() string {
	return strings.Join(hotelColumns, ", ")
}

var hotelColumns = []string{
	"id",
	"title",
	"city",
	"description",
	"image_url",
	"created_at",
	"updated_at",
}

/*
func getRoomColumns() string {
	return strings.Join(roomColumns, ",")
}

var roomColumns = []string{
	"id",
	"hotel_id",
	"room_count",
	"available_date",
	"created_at",
	"updated_at",
}
*/
