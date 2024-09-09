package travelgo

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"

	httpmodels "travelgo/internal/models/http"
)

func (h *TravelGoHandler) GetPopularDestinations(w http.ResponseWriter, r *http.Request) {
	ctx := context.WithValue(r.Context(), "opName", "GetPopularDestinations")

	req := httpmodels.GetPopularDestinationsRequest{}
	ok, msg := req.Valid()
	if !ok {
		log.Println(errors.New("valid: " + msg))
		w.Write(createResponse(Failed, fmt.Sprintf("request invalid: %s", msg), nil))
		return
	}

	resp, err := h.service.GetPopularDestinations(ctx, &req)
	if err != nil {
		log.Println(err)
		w.Write(createResponse(Failed, err.Error(), nil))
		return
	}

	w.Write(createResponse(Success, "", resp))
}
