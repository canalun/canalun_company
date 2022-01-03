package repository

import "content-updater/domain/model"

type EntryRepository interface {
	GetEntryList() (*model.EntryList, error)
}

type ContentUpdater interface {
	UpdateEntryListFile(model.EntryList) error
}
