package httpmodels

type HotelOrder struct {
	Hotel   Hotel        `json:"hotel"`
	Details OrderDetails `json:"orderDetails"`
}
