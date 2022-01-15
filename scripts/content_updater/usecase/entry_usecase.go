package usecase

import (
	"content-updater/domain/repository"
	"content-updater/infrastructure/file_generator"
)

type EntryUsecase struct {
	EntryRepositories []repository.EntryRepository
}

func NewEntryUsecase(
	ers ...repository.EntryRepository,
) EntryUsecase {
	return EntryUsecase{
		EntryRepositories: ers,
	}
}

func (a EntryUsecase) UpdateList() error {
	for _, entryRepository := range a.EntryRepositories {
		entryList, err := entryRepository.GetEntryList()
		if err != nil {
			return err
		}
		file_generator.UpdateEntryList(*entryList)
	}
	return nil
}
