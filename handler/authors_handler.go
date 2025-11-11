package handler

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/juanplagos/bubble/model"
	"github.com/juanplagos/bubble/usecase"
)

type AuthorHandler struct {
	useCase usecase.AuthorUseCase
}

func NewAuthorHandler(useCase usecase.AuthorUseCase) *AuthorHandler {
	return &AuthorHandler{
		useCase: useCase,
	}
}

func (h *AuthorHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	authors, err := h.useCase.GetAllAuthors()
	if err != nil {
		WriteError(w, http.StatusInternalServerError, err, "não foi possível obter os autores")
		return
	}
	WriteSuccess(w, http.StatusOK, authors, "authors retrieved successfully")
}

func (h *AuthorHandler) GetByUsername(w http.ResponseWriter, r *http.Request) {
	username := strings.TrimPrefix(r.URL.Path, "/authors/")
	if username == "" {
		WriteError(w, http.StatusBadRequest, nil, "username is required")
		return
	}

	author, err := h.useCase.GetAuthorByUsername(username)
	if err != nil {
		WriteError(w, http.StatusNotFound, err, "author not found")
		return
	}
	WriteSuccess(w, http.StatusOK, author, "author retrieved successfully")
}

func (h *AuthorHandler) GetByEmail(w http.ResponseWriter, r *http.Request) {
	email := strings.TrimPrefix(r.URL.Path, "/authors/email/")
	if email == "" {
		WriteError(w, http.StatusBadRequest, nil, "email is required")
		return
	}

	author, err := h.useCase.GetAuthorByEmail(email)
	if err != nil {
		WriteError(w, http.StatusNotFound, err, "author not found")
		return
	}
	WriteSuccess(w, http.StatusOK, author, "author retrieved successfully")
}

func (h *AuthorHandler) Create(w http.ResponseWriter, r *http.Request) {
	var author model.Author
	if err := json.NewDecoder(r.Body).Decode(&author); err != nil {
		WriteError(w, http.StatusBadRequest, err, "invalid request body")
		return
	}

	if err := h.useCase.CreateAuthor(&author); err != nil {
		WriteError(w, http.StatusInternalServerError, err, "failed to create author")
		return
	}
	WriteSuccess(w, http.StatusCreated, author, "author created successfully")
}

func (h *AuthorHandler) Update(w http.ResponseWriter, r *http.Request) {
	username := strings.TrimPrefix(r.URL.Path, "/authors/")
	if username == "" {
		WriteError(w, http.StatusBadRequest, nil, "username is required")
		return
	}

	var author model.Author
	if err := json.NewDecoder(r.Body).Decode(&author); err != nil {
		WriteError(w, http.StatusBadRequest, err, "invalid request body")
		return
	}

	if err := h.useCase.UpdateAuthor(username, &author); err != nil {
		WriteError(w, http.StatusInternalServerError, err, "failed to update author")
		return
	}
	WriteSuccess(w, http.StatusOK, author, "author updated successfully")
}

func (h *AuthorHandler) Delete(w http.ResponseWriter, r *http.Request) {
	username := strings.TrimPrefix(r.URL.Path, "/authors/")
	if username == "" {
		WriteError(w, http.StatusBadRequest, nil, "username is required")
		return
	}

	if err := h.useCase.DeleteAuthor(username); err != nil {
		WriteError(w, http.StatusInternalServerError, err, "failed to delete author")
		return
	}
	WriteSuccess(w, http.StatusOK, nil, "author deleted successfully")
}
