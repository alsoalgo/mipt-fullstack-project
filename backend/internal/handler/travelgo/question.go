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

func (h *TravelGoHandler) CreateQuestion(w http.ResponseWriter, r *http.Request) {
	ctx := context.WithValue(r.Context(), "opName", "CreateQuestion")

	data, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		w.Write(createResponse(Failed, err.Error(), nil))
		return
	}
	defer r.Body.Close()

	req := httpmodels.CreateQuestionRequest{}
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

	resp, err := h.service.CreateQuestion(ctx, &req)
	if err != nil {
		log.Println(err)
		w.Write(createResponse(Failed, err.Error(), nil))
		return
	}

	w.Write(createResponse(Success, "", resp))
}

func (h *TravelGoHandler) GetQuestions(w http.ResponseWriter, r *http.Request) {
	ctx := context.WithValue(r.Context(), "opName", "GetQuestions")

	req := httpmodels.GetQuestionsRequest{}
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

	resp, err := h.service.GetQuestions(ctx, &req)
	if err != nil {
		log.Println(err)
		w.Write(createResponse(Failed, err.Error(), nil))
		return
	}

	w.Write(createResponse(Success, "", resp))
}
