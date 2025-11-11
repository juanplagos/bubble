package usecase

import (
	"errors"
	"testing"
	"time"

	"github.com/juanplagos/bubble/model"
)

type mockEntryRepo struct {
	entries   []model.Entry
	entry     model.Entry
	err       error
	createErr error
	updateErr error
	deleteErr error
}

func (m *mockEntryRepo) GetAllEntries() ([]model.Entry, error) {
	return m.entries, m.err
}

func (m *mockEntryRepo) GetEntryById(id int) (model.Entry, error) {
	return m.entry, m.err
}

func (m *mockEntryRepo) GetEntryBySlug(slug string) (model.Entry, error) {
	return m.entry, m.err
}

func (m *mockEntryRepo) CreateEntry(entry *model.Entry) error {
	return m.createErr
}

func (m *mockEntryRepo) UpdateEntry(id int, entry *model.Entry) error {
	return m.updateErr
}

func (m *mockEntryRepo) DeleteEntry(id int) error {
	return m.deleteErr
}

func TestEntryUseCase_GetAllEntries(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		entries := []model.Entry{
			{ID: 1, Title: "Test", Slug: "test", Body: "Body", Author: "author", CreatedAt: time.Now()},
		}
		repo := &mockEntryRepo{entries: entries}
		uc := NewEntryUseCase(repo)

		result, err := uc.GetAllEntries()

		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if len(result) != 1 {
			t.Errorf("Expected 1 entry, got %d", len(result))
		}
	})

	t.Run("error", func(t *testing.T) {
		repo := &mockEntryRepo{err: errors.New("database error")}
		uc := NewEntryUseCase(repo)

		_, err := uc.GetAllEntries()

		if err == nil {
			t.Error("Expected error, got nil")
		}
	})
}

func TestEntryUseCase_GetEntryById(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		entry := model.Entry{ID: 1, Title: "Test", Slug: "test", Body: "Body", Author: "author", CreatedAt: time.Now()}
		repo := &mockEntryRepo{entry: entry}
		uc := NewEntryUseCase(repo)

		result, err := uc.GetEntryById(1)

		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if result.ID != 1 {
			t.Errorf("Expected ID 1, got %d", result.ID)
		}
	})

	t.Run("error", func(t *testing.T) {
		repo := &mockEntryRepo{err: errors.New("not found")}
		uc := NewEntryUseCase(repo)

		_, err := uc.GetEntryById(999)

		if err == nil {
			t.Error("Expected error, got nil")
		}
	})
}

func TestEntryUseCase_GetEntryBySlug(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		entry := model.Entry{ID: 1, Title: "Test", Slug: "test", Body: "Body", Author: "author", CreatedAt: time.Now()}
		repo := &mockEntryRepo{entry: entry}
		uc := NewEntryUseCase(repo)

		result, err := uc.GetEntryBySlug("test")

		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if result.Slug != "test" {
			t.Errorf("Expected slug 'test', got %s", result.Slug)
		}
	})
}

func TestEntryUseCase_CreateEntry(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		repo := &mockEntryRepo{}
		uc := NewEntryUseCase(repo)

		entry := &model.Entry{Title: "Test", Slug: "test", Body: "Body", Author: "author"}

		err := uc.CreateEntry(entry)

		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
	})

	t.Run("error", func(t *testing.T) {
		repo := &mockEntryRepo{createErr: errors.New("database error")}
		uc := NewEntryUseCase(repo)

		entry := &model.Entry{Title: "Test", Slug: "test", Body: "Body", Author: "author"}

		err := uc.CreateEntry(entry)

		if err == nil {
			t.Error("Expected error, got nil")
		}
	})
}

func TestEntryUseCase_UpdateEntry(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		repo := &mockEntryRepo{}
		uc := NewEntryUseCase(repo)

		entry := &model.Entry{Title: "Updated", Slug: "updated", Body: "Body", Author: "author"}

		err := uc.UpdateEntry(1, entry)

		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
	})

	t.Run("error", func(t *testing.T) {
		repo := &mockEntryRepo{updateErr: errors.New("database error")}
		uc := NewEntryUseCase(repo)

		entry := &model.Entry{Title: "Updated", Slug: "updated", Body: "Body", Author: "author"}

		err := uc.UpdateEntry(1, entry)

		if err == nil {
			t.Error("Expected error, got nil")
		}
	})
}

func TestEntryUseCase_DeleteEntry(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		repo := &mockEntryRepo{}
		uc := NewEntryUseCase(repo)

		err := uc.DeleteEntry(1)

		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
	})

	t.Run("error", func(t *testing.T) {
		repo := &mockEntryRepo{deleteErr: errors.New("database error")}
		uc := NewEntryUseCase(repo)

		err := uc.DeleteEntry(1)

		if err == nil {
			t.Error("Expected error, got nil")
		}
	})
}

