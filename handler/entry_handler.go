package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/juanplagos/bubble/model"
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
		WriteError(w, http.StatusInternalServerError, err, "não foi possível obter os registros")
		return
	}
	WriteSuccess(w, http.StatusOK, entries, "entries retrieved successfully")
}

func (h *EntryHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/entries/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		WriteError(w, http.StatusBadRequest, err, "invalid entry ID")
		return
	}

	entry, err := h.useCase.GetEntryById(id)
	if err != nil {
		WriteError(w, http.StatusNotFound, err, "entry not found")
		return
	}
	WriteSuccess(w, http.StatusOK, entry, "entry retrieved successfully")
}

func (h *EntryHandler) GetBySlug(w http.ResponseWriter, r *http.Request) {
	slug := strings.TrimPrefix(r.URL.Path, "/entries/slug/")
	if slug == "" {
		WriteError(w, http.StatusBadRequest, nil, "slug is required")
		return
	}

	entry, err := h.useCase.GetEntryBySlug(slug)
	if err != nil {
		WriteError(w, http.StatusNotFound, err, "entry not found")
		return
	}
	WriteSuccess(w, http.StatusOK, entry, "entry retrieved successfully")
}

func (h *EntryHandler) Create(w http.ResponseWriter, r *http.Request) {
	var entry model.Entry
	if err := json.NewDecoder(r.Body).Decode(&entry); err != nil {
		WriteError(w, http.StatusBadRequest, err, "invalid request body")
		return
	}

	if err := h.useCase.CreateEntry(&entry); err != nil {
		WriteError(w, http.StatusInternalServerError, err, "failed to create entry")
		return
	}
	WriteSuccess(w, http.StatusCreated, entry, "entry created successfully")
}

func (h *EntryHandler) Update(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/entries/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		WriteError(w, http.StatusBadRequest, err, "invalid entry ID")
		return
	}

	var entry model.Entry
	if err := json.NewDecoder(r.Body).Decode(&entry); err != nil {
		WriteError(w, http.StatusBadRequest, err, "invalid request body")
		return
	}

	if err := h.useCase.UpdateEntry(id, &entry); err != nil {
		WriteError(w, http.StatusInternalServerError, err, "failed to update entry")
		return
	}
	WriteSuccess(w, http.StatusOK, entry, "entry updated successfully")
}

func (h *EntryHandler) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/entries/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		WriteError(w, http.StatusBadRequest, err, "invalid entry ID")
		return
	}

	if err := h.useCase.DeleteEntry(id); err != nil {
		WriteError(w, http.StatusInternalServerError, err, "failed to delete entry")
		return
	}
	WriteSuccess(w, http.StatusOK, nil, "entry deleted successfully")
}
