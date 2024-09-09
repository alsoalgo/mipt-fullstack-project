package travelgo

import (
	"log"

	jsoniter "github.com/json-iterator/go"
)

type Status string

const (
	Success Status = "success"
	Failed  Status = "failed"
)

type Response struct {
	Status  Status `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func createResponse(status Status, msg string, data any) []byte {
	resp, err := jsoniter.Marshal(
		&Response{
			Status:  status,
			Message: msg,
			Data:    data,
		},
	)
	if err != nil {
		log.Printf("marshal: %v", err)
		return []byte{}
	}
	return resp
}
