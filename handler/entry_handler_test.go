package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/juanplagos/bubble/model"
)

type mockEntryUseCase struct {
	entries   []model.Entry
	entry     model.Entry
	err       error
	createErr error
	updateErr error
	deleteErr error
}

func (m *mockEntryUseCase) GetAllEntries() ([]model.Entry, error) {
	return m.entries, m.err
}

func (m *mockEntryUseCase) GetEntryById(id int) (model.Entry, error) {
	return m.entry, m.err
}

func (m *mockEntryUseCase) GetEntryBySlug(slug string) (model.Entry, error) {
	return m.entry, m.err
}

func (m *mockEntryUseCase) CreateEntry(entry *model.Entry) error {
	return m.createErr
}

func (m *mockEntryUseCase) UpdateEntry(id int, entry *model.Entry) error {
	return m.updateErr
}

func (m *mockEntryUseCase) DeleteEntry(id int) error {
	return m.deleteErr
}

func TestEntryHandler_GetAll(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		entries := []model.Entry{
			{ID: 1, Title: "Test", Slug: "test", Body: "Body", Author: "author", CreatedAt: time.Now()},
		}
		mockUC := &mockEntryUseCase{entries: entries}
		handler := NewEntryHandler(mockUC)

		req := httptest.NewRequest("GET", "/entries", nil)
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
		mockUC := &mockEntryUseCase{err: errors.New("database error")}
		handler := NewEntryHandler(mockUC)

		req := httptest.NewRequest("GET", "/entries", nil)
		w := httptest.NewRecorder()

		handler.GetAll(w, req)

		if w.Code != http.StatusInternalServerError {
			t.Errorf("Expected status %d, got %d", http.StatusInternalServerError, w.Code)
		}
	})
}

func TestEntryHandler_GetByID(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		entry := model.Entry{ID: 1, Title: "Test", Slug: "test", Body: "Body", Author: "author", CreatedAt: time.Now()}
		mockUC := &mockEntryUseCase{entry: entry}
		handler := NewEntryHandler(mockUC)

		req := httptest.NewRequest("GET", "/entries/1", nil)
		w := httptest.NewRecorder()

		handler.GetByID(w, req)

		if w.Code != http.StatusOK {
			t.Errorf("Expected status %d, got %d", http.StatusOK, w.Code)
		}
	})

	t.Run("invalid id", func(t *testing.T) {
		handler := NewEntryHandler(&mockEntryUseCase{})

		req := httptest.NewRequest("GET", "/entries/abc", nil)
		w := httptest.NewRecorder()

		handler.GetByID(w, req)

		if w.Code != http.StatusBadRequest {
			t.Errorf("Expected status %d, got %d", http.StatusBadRequest, w.Code)
		}
	})

	t.Run("not found", func(t *testing.T) {
		mockUC := &mockEntryUseCase{err: errors.New("not found")}
		handler := NewEntryHandler(mockUC)

		req := httptest.NewRequest("GET", "/entries/999", nil)
		w := httptest.NewRecorder()

		handler.GetByID(w, req)

		if w.Code != http.StatusNotFound {
			t.Errorf("Expected status %d, got %d", http.StatusNotFound, w.Code)
		}
	})
}

func TestEntryHandler_GetBySlug(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		entry := model.Entry{ID: 1, Title: "Test", Slug: "test", Body: "Body", Author: "author", CreatedAt: time.Now()}
		mockUC := &mockEntryUseCase{entry: entry}
		handler := NewEntryHandler(mockUC)

		req := httptest.NewRequest("GET", "/entries/slug/test", nil)
		w := httptest.NewRecorder()

		handler.GetBySlug(w, req)

		if w.Code != http.StatusOK {
			t.Errorf("Expected status %d, got %d", http.StatusOK, w.Code)
		}
	})

	t.Run("empty slug", func(t *testing.T) {
		handler := NewEntryHandler(&mockEntryUseCase{})

		req := httptest.NewRequest("GET", "/entries/slug/", nil)
		w := httptest.NewRecorder()

		handler.GetBySlug(w, req)

		if w.Code != http.StatusBadRequest {
			t.Errorf("Expected status %d, got %d", http.StatusBadRequest, w.Code)
		}
	})
}

func TestEntryHandler_Create(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockUC := &mockEntryUseCase{}
		handler := NewEntryHandler(mockUC)

		entry := model.Entry{Title: "Test", Slug: "test", Body: "Body", Author: "author"}
		body, _ := json.Marshal(entry)
		req := httptest.NewRequest("POST", "/entries", bytes.NewBuffer(body))
		w := httptest.NewRecorder()

		handler.Create(w, req)

		if w.Code != http.StatusCreated {
			t.Errorf("Expected status %d, got %d", http.StatusCreated, w.Code)
		}
	})

	t.Run("invalid body", func(t *testing.T) {
		handler := NewEntryHandler(&mockEntryUseCase{})

		req := httptest.NewRequest("POST", "/entries", bytes.NewBufferString("invalid json"))
		w := httptest.NewRecorder()

		handler.Create(w, req)

		if w.Code != http.StatusBadRequest {
			t.Errorf("Expected status %d, got %d", http.StatusBadRequest, w.Code)
		}
	})

	t.Run("create error", func(t *testing.T) {
		mockUC := &mockEntryUseCase{createErr: errors.New("database error")}
		handler := NewEntryHandler(mockUC)

		entry := model.Entry{Title: "Test", Slug: "test", Body: "Body", Author: "author"}
		body, _ := json.Marshal(entry)
		req := httptest.NewRequest("POST", "/entries", bytes.NewBuffer(body))
		w := httptest.NewRecorder()

		handler.Create(w, req)

		if w.Code != http.StatusInternalServerError {
			t.Errorf("Expected status %d, got %d", http.StatusInternalServerError, w.Code)
		}
	})
}

func TestEntryHandler_Update(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockUC := &mockEntryUseCase{}
		handler := NewEntryHandler(mockUC)

		entry := model.Entry{Title: "Updated", Slug: "updated", Body: "Body", Author: "author"}
		body, _ := json.Marshal(entry)
		req := httptest.NewRequest("PUT", "/entries/1", bytes.NewBuffer(body))
		w := httptest.NewRecorder()

		handler.Update(w, req)

		if w.Code != http.StatusOK {
			t.Errorf("Expected status %d, got %d", http.StatusOK, w.Code)
		}
	})

	t.Run("invalid id", func(t *testing.T) {
		handler := NewEntryHandler(&mockEntryUseCase{})

		req := httptest.NewRequest("PUT", "/entries/abc", nil)
		w := httptest.NewRecorder()

		handler.Update(w, req)

		if w.Code != http.StatusBadRequest {
			t.Errorf("Expected status %d, got %d", http.StatusBadRequest, w.Code)
		}
	})
}

func TestEntryHandler_Delete(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockUC := &mockEntryUseCase{}
		handler := NewEntryHandler(mockUC)

		req := httptest.NewRequest("DELETE", "/entries/1", nil)
		w := httptest.NewRecorder()

		handler.Delete(w, req)

		if w.Code != http.StatusOK {
			t.Errorf("Expected status %d, got %d", http.StatusOK, w.Code)
		}
	})

	t.Run("invalid id", func(t *testing.T) {
		handler := NewEntryHandler(&mockEntryUseCase{})

		req := httptest.NewRequest("DELETE", "/entries/abc", nil)
		w := httptest.NewRecorder()

		handler.Delete(w, req)

		if w.Code != http.StatusBadRequest {
			t.Errorf("Expected status %d, got %d", http.StatusBadRequest, w.Code)
		}
	})
}
