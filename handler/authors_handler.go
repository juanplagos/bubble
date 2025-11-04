package handler

import (
	"encoding/json"
	"net/http"

	"github.com/juanplagos/bubble/mock"
)

func GetAuthors(w http.ResponseWriter, r *http.Request) {
	response := map[string]any{
		"message": "not far from fetched",
		"authors": mock.Author,
	}
	json.NewEncoder(w).Encode(response)
}