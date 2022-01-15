package updater

import "content-updater/domain/model"

type ContentUpdater interface {
	UpdateEntryList(model.EntryList) error
}
