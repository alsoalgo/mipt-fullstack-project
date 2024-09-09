package httpmodels

import "time"

type SearchRequest struct {
	City     string
	DateFrom time.Time
	DateTo   time.Time
}

func (r *SearchRequest) Valid() (bool, string) {
	if len(r.City) < 2 {
		return false, "city is short"
	}

	now := time.Now().UTC().Truncate(24 * time.Hour)
	if r.DateFrom.UTC().Before(now) {
		return false, "dateFrom can't be in past"
	}
	if r.DateTo.UTC().Before(now) {
		return false, "dateTo can't be in past"
	}
	return true, ""
}

type SearchResponse struct {
	Hotels []*Hotel `json:"hotels"`
}
