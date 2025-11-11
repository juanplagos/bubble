package usecase

import (
	"github.com/juanplagos/bubble/model"
	"github.com/juanplagos/bubble/repository"
)

type entryUseCase struct {
	repo repository.EntryRepo
}

func NewEntryUseCase(repo repository.EntryRepo) EntryUseCase {
	return &entryUseCase{
		repo: repo,
	}
}

func (eu *entryUseCase) GetAllEntries() ([]model.Entry, error) {
	return eu.repo.GetAllEntries()
}

func (eu *entryUseCase) GetEntryById(id int) (model.Entry, error) {
	return eu.repo.GetEntryById(id)
}

func (eu *entryUseCase) GetEntryBySlug(slug string) (model.Entry, error) {
	return eu.repo.GetEntryBySlug(slug)
}

func (eu *entryUseCase) CreateEntry(entry *model.Entry) error {
	return eu.repo.CreateEntry(entry)
}

func (eu *entryUseCase) UpdateEntry(id int, entry *model.Entry) error {
	return eu.repo.UpdateEntry(id, entry)
}

func (eu *entryUseCase) DeleteEntry(id int) error {
	return eu.repo.DeleteEntry(id)
}
