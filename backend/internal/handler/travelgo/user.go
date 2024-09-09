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

func (h *TravelGoHandler) EditProfile(w http.ResponseWriter, r *http.Request) {
	ctx := context.WithValue(r.Context(), "opName", "EditProfile")

	data, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(errors.New("readall: " + err.Error()))
		w.Write(createResponse(Failed, err.Error(), nil))
		return
	}
	defer r.Body.Close()

	req := httpmodels.EditProfileRequest{}

	err = jsoniter.Unmarshal(data, &req)
	if err != nil {
		log.Println(errors.New("unmarshal: " + err.Error()))
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

	resp, err := h.service.EditProfile(ctx, &req)
	if err != nil {
		log.Println(errors.New("login: " + err.Error()))
		w.Write(createResponse(Failed, err.Error(), nil))
		return
	}

	w.Write(createResponse(Success, "", resp))
}

func (h *TravelGoHandler) GetProfile(w http.ResponseWriter, r *http.Request) {
	ctx := context.WithValue(r.Context(), "opName", "EditProfile")

	req := httpmodels.GetProfileRequest{}

	userID, err := getUserIdFromRequest(r)
	if err != nil {
		log.Println(err)
		w.Write(createResponse(Failed, err.Error(), nil))
		return
	}
	req.UserID = userID

	ok, msg := req.Valid()
	if !ok {
		log.Println(errors.New("valid: " + msg))
		w.Write(createResponse(Failed, fmt.Sprintf("request invalid: %s", msg), nil))
		return
	}

	resp, err := h.service.GetProfile(ctx, &req)
	if err != nil {
		log.Println(errors.New("login: " + err.Error()))
		w.Write(createResponse(Failed, err.Error(), nil))
		return
	}

	w.Write(createResponse(Success, "", resp))
}
