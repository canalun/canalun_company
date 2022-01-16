package repository

import "content-updater/domain/model"

type EntryRepository interface {
	GetLatestEntryList() (*model.EntryList, error)
}

type ContentUpdater interface {
	UpdateEntryList(model.EntryList) error
}
