package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestWriteJSON(t *testing.T) {
	w := httptest.NewRecorder()
	data := map[string]string{"test": "value"}

	WriteJSON(w, http.StatusOK, data)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}

	if w.Header().Get("Content-Type") != "application/json" {
		t.Errorf("Expected Content-Type application/json, got %s", w.Header().Get("Content-Type"))
	}

	var result map[string]string
	if err := json.Unmarshal(w.Body.Bytes(), &result); err != nil {
		t.Errorf("Failed to unmarshal JSON: %v", err)
	}

	if result["test"] != "value" {
		t.Errorf("Expected test=value, got %s", result["test"])
	}
}

func TestWriteSuccess(t *testing.T) {
	w := httptest.NewRecorder()
	data := map[string]string{"key": "value"}

	WriteSuccess(w, http.StatusCreated, data, "success message")

	if w.Code != http.StatusCreated {
		t.Errorf("Expected status code %d, got %d", http.StatusCreated, w.Code)
	}

	var response Response
	if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
		t.Errorf("Failed to unmarshal JSON: %v", err)
	}

	if !response.Success {
		t.Error("Expected Success to be true")
	}

	if response.Message != "success message" {
		t.Errorf("Expected message 'success message', got %s", response.Message)
	}

	if response.Data == nil {
		t.Error("Expected Data to be set")
	}
}

func TestWriteError(t *testing.T) {
	t.Run("with error", func(t *testing.T) {
		w := httptest.NewRecorder()
		err := &testError{msg: "test error"}

		WriteError(w, http.StatusBadRequest, err, "error message")

		if w.Code != http.StatusBadRequest {
			t.Errorf("Expected status code %d, got %d", http.StatusBadRequest, w.Code)
		}

		var response Response
		if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
			t.Errorf("Failed to unmarshal JSON: %v", err)
		}

		if response.Success {
			t.Error("Expected Success to be false")
		}

		if response.Error != "test error" {
			t.Errorf("Expected error 'test error', got %s", response.Error)
		}

		if response.Message != "error message" {
			t.Errorf("Expected message 'error message', got %s", response.Message)
		}
	})

	t.Run("with nil error", func(t *testing.T) {
		w := httptest.NewRecorder()

		WriteError(w, http.StatusBadRequest, nil, "error message")

		if w.Code != http.StatusBadRequest {
			t.Errorf("Expected status code %d, got %d", http.StatusBadRequest, w.Code)
		}

		var response Response
		if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
			t.Errorf("Failed to unmarshal JSON: %v", err)
		}

		if response.Success {
			t.Error("Expected Success to be false")
		}

		if response.Error != "" {
			t.Errorf("Expected empty error, got %s", response.Error)
		}

		if response.Message != "error message" {
			t.Errorf("Expected message 'error message', got %s", response.Message)
		}
	})
}

type testError struct {
	msg string
}

func (e *testError) Error() string {
	return e.msg
}
