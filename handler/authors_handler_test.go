package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/juanplagos/bubble/model"
)

type mockAuthorUseCase struct {
	authors   []model.Author
	author    model.Author
	err       error
	createErr error
	updateErr error
	deleteErr error
}

func (m *mockAuthorUseCase) GetAllAuthors() ([]model.Author, error) {
	return m.authors, m.err
}

func (m *mockAuthorUseCase) GetAuthorByUsername(username string) (model.Author, error) {
	return m.author, m.err
}

func (m *mockAuthorUseCase) GetAuthorByEmail(email string) (model.Author, error) {
	return m.author, m.err
}

func (m *mockAuthorUseCase) CreateAuthor(author *model.Author) error {
	return m.createErr
}

func (m *mockAuthorUseCase) UpdateAuthor(username string, author *model.Author) error {
	return m.updateErr
}

func (m *mockAuthorUseCase) DeleteAuthor(username string) error {
	return m.deleteErr
}

func TestAuthorHandler_GetAll(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		authors := []model.Author{
			{Username: "user1", Email: "user1@test.com", Password: "pass1"},
		}
		mockUC := &mockAuthorUseCase{authors: authors}
		handler := NewAuthorHandler(mockUC)

		req := httptest.NewRequest("GET", "/authors", nil)
		w := httptest.NewRecorder()

		handler.GetAll(w, req)

		if w.Code != http.StatusOK {
			t.Errorf("Expected status %d, got %d", http.StatusOK, w.Code)
		}

		var response Response
		if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
			t.Fatalf("Failed to unmarshal response: %v", err)
		}

		if !response.Success {
			t.Error("Expected success to be true")
		}
	})

	t.Run("error", func(t *testing.T) {
		mockUC := &mockAuthorUseCase{err: errors.New("database error")}
		handler := NewAuthorHandler(mockUC)

		req := httptest.NewRequest("GET", "/authors", nil)
		w := httptest.NewRecorder()

		handler.GetAll(w, req)

		if w.Code != http.StatusInternalServerError {
			t.Errorf("Expected status %d, got %d", http.StatusInternalServerError, w.Code)
		}
	})
}

func TestAuthorHandler_GetByUsername(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		author := model.Author{Username: "user1", Email: "user1@test.com", Password: "pass1"}
		mockUC := &mockAuthorUseCase{author: author}
		handler := NewAuthorHandler(mockUC)

		req := httptest.NewRequest("GET", "/authors/user1", nil)
		w := httptest.NewRecorder()

		handler.GetByUsername(w, req)

		if w.Code != http.StatusOK {
			t.Errorf("Expected status %d, got %d", http.StatusOK, w.Code)
		}
	})

	t.Run("empty username", func(t *testing.T) {
		handler := NewAuthorHandler(&mockAuthorUseCase{})

		req := httptest.NewRequest("GET", "/authors/", nil)
		w := httptest.NewRecorder()

		handler.GetByUsername(w, req)

		if w.Code != http.StatusBadRequest {
			t.Errorf("Expected status %d, got %d", http.StatusBadRequest, w.Code)
		}
	})

	t.Run("not found", func(t *testing.T) {
		mockUC := &mockAuthorUseCase{err: errors.New("not found")}
		handler := NewAuthorHandler(mockUC)

		req := httptest.NewRequest("GET", "/authors/nonexistent", nil)
		w := httptest.NewRecorder()

		handler.GetByUsername(w, req)

		if w.Code != http.StatusNotFound {
			t.Errorf("Expected status %d, got %d", http.StatusNotFound, w.Code)
		}
	})
}

func TestAuthorHandler_GetByEmail(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		author := model.Author{Username: "user1", Email: "user1@test.com", Password: "pass1"}
		mockUC := &mockAuthorUseCase{author: author}
		handler := NewAuthorHandler(mockUC)

		req := httptest.NewRequest("GET", "/authors/email/user1@test.com", nil)
		w := httptest.NewRecorder()

		handler.GetByEmail(w, req)

		if w.Code != http.StatusOK {
			t.Errorf("Expected status %d, got %d", http.StatusOK, w.Code)
		}
	})

	t.Run("empty email", func(t *testing.T) {
		handler := NewAuthorHandler(&mockAuthorUseCase{})

		req := httptest.NewRequest("GET", "/authors/email/", nil)
		w := httptest.NewRecorder()

		handler.GetByEmail(w, req)

		if w.Code != http.StatusBadRequest {
			t.Errorf("Expected status %d, got %d", http.StatusBadRequest, w.Code)
		}
	})
}

func TestAuthorHandler_Create(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockUC := &mockAuthorUseCase{}
		handler := NewAuthorHandler(mockUC)

		author := model.Author{Username: "user1", Email: "user1@test.com", Password: "pass1"}
		body, _ := json.Marshal(author)
		req := httptest.NewRequest("POST", "/authors", bytes.NewBuffer(body))
		w := httptest.NewRecorder()

		handler.Create(w, req)

		if w.Code != http.StatusCreated {
			t.Errorf("Expected status %d, got %d", http.StatusCreated, w.Code)
		}
	})

	t.Run("invalid body", func(t *testing.T) {
		handler := NewAuthorHandler(&mockAuthorUseCase{})

		req := httptest.NewRequest("POST", "/authors", bytes.NewBufferString("invalid json"))
		w := httptest.NewRecorder()

		handler.Create(w, req)

		if w.Code != http.StatusBadRequest {
			t.Errorf("Expected status %d, got %d", http.StatusBadRequest, w.Code)
		}
	})

	t.Run("create error", func(t *testing.T) {
		mockUC := &mockAuthorUseCase{createErr: errors.New("database error")}
		handler := NewAuthorHandler(mockUC)

		author := model.Author{Username: "user1", Email: "user1@test.com", Password: "pass1"}
		body, _ := json.Marshal(author)
		req := httptest.NewRequest("POST", "/authors", bytes.NewBuffer(body))
		w := httptest.NewRecorder()

		handler.Create(w, req)

		if w.Code != http.StatusInternalServerError {
			t.Errorf("Expected status %d, got %d", http.StatusInternalServerError, w.Code)
		}
	})
}

func TestAuthorHandler_Update(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockUC := &mockAuthorUseCase{}
		handler := NewAuthorHandler(mockUC)

		author := model.Author{Username: "user1", Email: "updated@test.com", Password: "newpass"}
		body, _ := json.Marshal(author)
		req := httptest.NewRequest("PUT", "/authors/user1", bytes.NewBuffer(body))
		w := httptest.NewRecorder()

		handler.Update(w, req)

		if w.Code != http.StatusOK {
			t.Errorf("Expected status %d, got %d", http.StatusOK, w.Code)
		}
	})

	t.Run("empty username", func(t *testing.T) {
		handler := NewAuthorHandler(&mockAuthorUseCase{})

		req := httptest.NewRequest("PUT", "/authors/", nil)
		w := httptest.NewRecorder()

		handler.Update(w, req)

		if w.Code != http.StatusBadRequest {
			t.Errorf("Expected status %d, got %d", http.StatusBadRequest, w.Code)
		}
	})
}

func TestAuthorHandler_Delete(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockUC := &mockAuthorUseCase{}
		handler := NewAuthorHandler(mockUC)

		req := httptest.NewRequest("DELETE", "/authors/user1", nil)
		w := httptest.NewRecorder()

		handler.Delete(w, req)

		if w.Code != http.StatusOK {
			t.Errorf("Expected status %d, got %d", http.StatusOK, w.Code)
		}
	})

	t.Run("empty username", func(t *testing.T) {
		handler := NewAuthorHandler(&mockAuthorUseCase{})

		req := httptest.NewRequest("DELETE", "/authors/", nil)
		w := httptest.NewRecorder()

		handler.Delete(w, req)

		if w.Code != http.StatusBadRequest {
			t.Errorf("Expected status %d, got %d", http.StatusBadRequest, w.Code)
		}
	})
}
