package usecase

import (
	"content-updater/domain/repository"
	"content-updater/infrastructure/file_generator"

	"github.com/pkg/errors"
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
		entryList, err := entryRepository.GetLatestEntryList()
		if err != nil {
			return errors.WithStack(err)
		}
		if err := file_generator.UpdateEntryList(*entryList); err != nil {
			return errors.WithStack(err)
		}
	}
	return nil
}
