package usecase

import (
	"github.com/juanplagos/bubble/model"
	"github.com/juanplagos/bubble/repository"
)

type authorUseCase struct {
	repo repository.AuthorRepo
}

func NewAuthorUseCase(repo repository.AuthorRepo) AuthorUseCase {
	return &authorUseCase{
		repo: repo,
	}
}

func (au *authorUseCase) GetAllAuthors() ([]model.Author, error) {
	return au.repo.GetAllAuthors()
}

func (au *authorUseCase) GetAuthorByUsername(username string) (model.Author, error) {
	return au.repo.GetAuthorByUsername(username)
}

func (au *authorUseCase) GetAuthorByEmail(email string) (model.Author, error) {
	return au.repo.GetAuthorByEmail(email)
}

func (au *authorUseCase) CreateAuthor(author *model.Author) error {
	return au.repo.CreateAuthor(author)
}

func (au *authorUseCase) UpdateAuthor(username string, author *model.Author) error {
	return au.repo.UpdateAuthor(username, author)
}

func (au *authorUseCase) DeleteAuthor(username string) error {
	return au.repo.DeleteAuthor(username)
}
