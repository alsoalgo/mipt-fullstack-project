package httpmodels

import (
	"fmt"
	"time"
)

type CreateOrderRequest struct {
	UserID    int64  `json:"userId,omitempty"`
	HotelID   int64  `json:"hotelId"`
	DateFrom  string `json:"dateFrom"`
	DateTo    string `json:"dateTo"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	SurName   string `json:"surName"`
}

func (r *CreateOrderRequest) Valid() (bool, string) {
	if r.UserID <= 0 {
		return false, "userId must be positive"
	}

	if r.HotelID <= 0 {
		return false, "hotelId must be positive"
	}

	_, err := time.Parse(time.DateOnly, r.DateFrom)
	if err != nil {
		return false, fmt.Sprintf("Can't parse dateFrom: %v", err.Error())
	}

	_, err = time.Parse(time.DateOnly, r.DateTo)
	if err != nil {
		return false, fmt.Sprintf("Can't parse dateTo: %v", err.Error())
	}

	if len(r.FirstName) == 0 {
		return false, "firstName is empty"
	}

	if len(r.LastName) == 0 {
		return false, "lastName is empty"
	}

	if len(r.SurName) == 0 {
		return false, "surName is empty"
	}

	return true, ""
}

type CreateOrderResponse struct {
	Created bool `json:"created"`
}

type GetOrdersRequest struct {
	UserID int64
}

func (r *GetOrdersRequest) Valid() (bool, string) {
	if r.UserID <= 0 {
		return false, "userId must be positive"
	}

	return true, ""
}

type GetOrdersResponse struct {
	Orders []*HotelOrder `json:"orders"`
}

type OrderDetails struct {
	DateFrom string `json:"from"`
	DateTo   string `json:"to"`

	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	SurName   string `json:"surName"`
}
