package httpmodels

type Destination struct {
	City     string `json:"city"`
	Cost     int32  `json:"cost"`
	ImageURL string `json:"imageUrl"`
}

type GetPopularDestinationsRequest struct{}

func (r *GetPopularDestinationsRequest) Valid() (bool, string) {
	return true, ""
}

type GetPopularDestinationsResponse struct {
	Destinations []*Destination `json:"destinations"`
}
