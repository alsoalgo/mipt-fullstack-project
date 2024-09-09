package destinationrepo

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

func (r *Repository) CreateDestination(ctx context.Context, city string, cost int32, imageUrl string) (int64, error) {
	ctx = context.WithValue(ctx, "opName", "CreateDestination")

	q := `INSERT INTO popular_destination (city, cost, image_url) VALUES ($1, $2, $3) RETURNING id`

	id := int64(0)
	err := r.db.GetContext(ctx, &id, q, city, cost, imageUrl)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *Repository) GetPopularDestinations(ctx context.Context) ([]*dbmodels.Destination, error) {
	ctx = context.WithValue(ctx, "opName", "GetPopularDestinations")

	q := `SELECT ` + getColumns() + ` FROM popular_destination`

	destinations := make([]*dbmodels.Destination, 0)
	err := r.db.SelectContext(ctx, &destinations, q)
	if err != nil {
		return nil, err
	}

	return destinations, nil
}

func getColumns() string {
	return strings.Join(columns, ", ")
}

var columns = []string{
	"id",
	"city",
	"cost",
	"image_url",
	"created_at",
	"updated_at",
}
