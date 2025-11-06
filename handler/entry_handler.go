package handler

import (
	"encoding/json"
	"net/http"

	"github.com/juanplagos/bubble/usecase"
)

type EntryHandler struct {
	useCase usecase.EntryUseCase
}	

func NewEntryHandler(useCase usecase.EntryUseCase) *EntryHandler {
	return &EntryHandler{
		useCase: useCase,
	}
}

func (h *EntryHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	entries, err := h.useCase.GetAllEntries()

	if err != nil {
		http.Error(w, "não foi possível obter os registros", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(entries)
}