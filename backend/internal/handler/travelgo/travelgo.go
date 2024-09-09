package travelgo

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	jsoniter "github.com/json-iterator/go"

	httpmodels "travelgo/internal/models/http"
	"travelgo/internal/service/auth"
	"travelgo/internal/service/travelgo"
)

type TravelGoHandler struct {
	service *travelgo.TravelGoService
}

func NewTravelGoHandler(srv *travelgo.TravelGoService) *TravelGoHandler {
	return &TravelGoHandler{
		service: srv,
	}
}

func (h *TravelGoHandler) Login(w http.ResponseWriter, r *http.Request) {
	ctx := context.WithValue(r.Context(), "opName", "Login")

	data, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(errors.New("readall: " + err.Error()))
		w.Write(createResponse(Failed, err.Error(), nil))
		return
	}
	defer r.Body.Close()

	req := httpmodels.LoginRequest{}

	err = jsoniter.Unmarshal(data, &req)
	if err != nil {
		log.Println(errors.New("unmarshal: " + err.Error()))
		w.Write(createResponse(Failed, err.Error(), nil))
		return
	}

	ok, msg := req.Valid()
	if !ok {
		log.Println(errors.New("valid: " + msg))
		w.Write(createResponse(Failed, fmt.Sprintf("request invalid: %s", msg), nil))
		return
	}

	resp, err := h.service.Login(ctx, &req)
	if err != nil {
		log.Println(errors.New("login: " + err.Error()))
		w.Write(createResponse(Failed, err.Error(), nil))
		return
	}

	http.SetCookie(w, &http.Cookie{Name: "token", Value: resp.Token})
	w.Write(createResponse(Success, "", resp))
}

func (h *TravelGoHandler) Register(w http.ResponseWriter, r *http.Request) {
	ctx := context.WithValue(r.Context(), "opName", "Register")

	data, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		w.Write(createResponse(Failed, err.Error(), nil))
		return
	}
	defer r.Body.Close()

	req := httpmodels.RegisterRequest{}

	err = jsoniter.Unmarshal(data, &req)
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

	resp, err := h.service.Register(ctx, &req)
	if err != nil {
		log.Println(err)
		w.Write(createResponse(Failed, err.Error(), nil))
		return
	}

	w.Write(createResponse(Success, "", resp))
}

func (h *TravelGoHandler) Check(w http.ResponseWriter, r *http.Request) {
	ctx := context.WithValue(r.Context(), "opName", "Search")

	data, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		w.Write(createResponse(Failed, err.Error(), nil))
		return
	}
	defer r.Body.Close()

	req := httpmodels.CheckTokenRequest{}
	err = jsoniter.Unmarshal(data, &req)
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

	resp, err := h.service.CheckToken(ctx, &req)
	if err != nil {
		log.Println(err)
		w.Write(createResponse(Failed, err.Error(), nil))
		return
	}

	w.Write(createResponse(Success, "", resp))
}

func (h *TravelGoHandler) Search(w http.ResponseWriter, r *http.Request) {
	ctx := context.WithValue(r.Context(), "opName", "Search")

	req := httpmodels.SearchRequest{}
	err := parseSearchRequest(r, &req)
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

	resp, err := h.service.Search(ctx, &req)
	if err != nil {
		log.Println(err)
		w.Write(createResponse(Failed, err.Error(), nil))
		return
	}

	w.Write(createResponse(Success, "", resp))
}

func (h *TravelGoHandler) GetHotel(w http.ResponseWriter, r *http.Request) {
	ctx := context.WithValue(r.Context(), "opName", "GetHotel")

	data, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		w.Write(createResponse(Failed, err.Error(), nil))
		return
	}
	defer r.Body.Close()

	req := httpmodels.GetHotelRequest{}
	err = jsoniter.Unmarshal(data, &req)
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

	resp, err := h.service.GetHotel(ctx, &req)
	if err != nil {
		log.Println(err)
		w.Write(createResponse(Failed, err.Error(), nil))
		return
	}

	w.Write(createResponse(Success, "", resp))
}

func parseSearchRequest(r *http.Request, req *httpmodels.SearchRequest) error {
	city := r.URL.Query().Get("city")

	dateFrom, err := time.Parse(time.DateOnly, r.URL.Query().Get("from"))
	if err != nil {
		return errors.New("date from time parse: " + err.Error())
	}

	dateTo, err := time.Parse(time.DateOnly, r.URL.Query().Get("to"))
	if err != nil {
		return errors.New("date to time parse: " + err.Error())
	}

	req.City = city
	req.DateFrom = dateFrom
	req.DateTo = dateTo

	return nil
}

func getUserIdFromRequest(r *http.Request) (int64, error) {
	userId, err := auth.ParseToken(r)
	if err != nil {
		return 0, errors.New("ParseToken: " + err.Error())
	}
	return userId, nil
}
