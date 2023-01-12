package utils

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func SendErrorResponse(w http.ResponseWriter, status int, message string) {
	response := &Response{
		Status:  status,
		Message: message,
		Data:    nil,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(response)
}

func SendSuccessResponse(w http.ResponseWriter, status int, data interface{}) {
	response := &Response{
		Status:  status,
		Message: "Success",
		Data:    data,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(response)
}
