package usecase

import (
	"github.com/juanplagos/bubble/repository"
	"github.com/juanplagos/bubble/model"
)

type EntryUseCase struct {
	repo repository.PostgresEntryRepo
}

func NewEntryUseCase(repo repository.PostgresEntryRepo) EntryUseCase {
	return EntryUseCase{
		repo: repo,
	}
}

func (eu *EntryUseCase) GetAllEntries() ([]model.Entry, error) {
	return eu.repo.GetAllEntries()
}