package travelgo

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"

	jsoniter "github.com/json-iterator/go"

	httpmodels "travelgo/internal/models/http"
)

func (h *TravelGoHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	ctx := context.WithValue(r.Context(), "opName", "CreateOrder")

	data, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		w.Write(createResponse(Failed, err.Error(), nil))
		return
	}
	defer r.Body.Close()

	req := httpmodels.CreateOrderRequest{}
	err = jsoniter.Unmarshal(data, &req)
	if err != nil {
		log.Println(err)
		w.Write(createResponse(Failed, err.Error(), nil))
		return
	}

	req.UserID, err = getUserIdFromRequest(r)
	if err != nil {
		log.Println(err)
		w.Write(createResponse(Failed, err.Error(), nil))
		return
	}

	ok, msg := req.Valid()
	if !ok {
		log.Println(errors.New("valid: " + msg))
		w.Write(createResponse(Failed, fmt.Sprintf("request invalid: %s", msg), nil))
		return
	}

	resp, err := h.service.CreateOrder(ctx, &req)
	if err != nil {
		log.Println(err)
		w.Write(createResponse(Failed, err.Error(), nil))
		return
	}

	w.Write(createResponse(Success, "", resp))
}

func (h *TravelGoHandler) GetOrders(w http.ResponseWriter, r *http.Request) {
	ctx := context.WithValue(r.Context(), "opName", "GetOrders")

	req := httpmodels.GetOrdersRequest{}
	userId, err := getUserIdFromRequest(r)
	if err != nil {
		log.Println(err)
		w.Write(createResponse(Failed, err.Error(), nil))
		return
	}
	req.UserID = userId

	ok, msg := req.Valid()
	if !ok {
		log.Println(errors.New("valid: " + msg))
		w.Write(createResponse(Failed, fmt.Sprintf("request invalid: %s", msg), nil))
		return
	}

	resp, err := h.service.GerOrders(ctx, &req)
	if err != nil {
		log.Println(err)
		w.Write(createResponse(Failed, err.Error(), nil))
		return
	}

	w.Write(createResponse(Success, "", resp))
}
