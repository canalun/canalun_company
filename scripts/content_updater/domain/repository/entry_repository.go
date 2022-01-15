package repository

import "content-updater/domain/model"

type EntryRepository interface {
	GetEntryList() (*model.EntryList, error)
}

type ContentUpdater interface {
	UpdateEntryList(model.EntryList) error
}
