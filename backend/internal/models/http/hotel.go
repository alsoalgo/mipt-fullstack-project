package httpmodels

type Hotel struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	City        string `json:"city"`
	Description string `json:"description"`
	ImageURL    string `json:"imageUrl"`
}

type GetHotelRequest struct {
	ID int64 `json:"hotelId"`
}

func (r *GetHotelRequest) Valid() (bool, string) {
	if r.ID <= 0 {
		return false, "hotelId must be positive"
	}

	return true, ""
}

type GetHotelResponse struct {
	Hotel *Hotel `json:"hotel"`
}
