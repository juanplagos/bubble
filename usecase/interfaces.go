package usecase

import "github.com/juanplagos/bubble/model"

type EntryUseCase interface {
	GetAllEntries() ([]model.Entry, error)
	GetEntryById(id int) (model.Entry, error)
	GetEntryBySlug(slug string) (model.Entry, error)
	CreateEntry(entry *model.Entry) error
	UpdateEntry(id int, entry *model.Entry) error
	DeleteEntry(id int) error
}

type AuthorUseCase interface {
	GetAllAuthors() ([]model.Author, error)
	GetAuthorByUsername(username string) (model.Author, error)
	GetAuthorByEmail(email string) (model.Author, error)
	CreateAuthor(author *model.Author) error
	UpdateAuthor(username string, author *model.Author) error
	DeleteAuthor(username string) error
}
