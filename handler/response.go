package handler

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
	Message string      `json:"message,omitempty"`
}

func WriteJSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

func WriteSuccess(w http.ResponseWriter, statusCode int, data interface{}, message string) {
	response := Response{
		Success: true,
		Data:    data,
		Message: message,
	}
	WriteJSON(w, statusCode, response)
}

func WriteError(w http.ResponseWriter, statusCode int, err error, message string) {
	response := Response{
		Success: false,
		Message: message,
	}
	if err != nil {
		response.Error = err.Error()
	}
	WriteJSON(w, statusCode, response)
}
