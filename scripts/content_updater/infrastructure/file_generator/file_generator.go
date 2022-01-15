package file_generator

import (
	"content-updater/config"
	"content-updater/domain/model"
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

func UpdateEntryList(el model.EntryList) error {
	el.SortEntries()

	var data []model.Entry
	for i := 0; i < config.Conf.NumOfEntry; i++ {
		data = append(data, el.Entries[i])
	}
	_data, err := json.Marshal(data)
	if err != nil {
		return err
	}

	filePath := config.Conf.EntryListFilePath + el.Source + config.Conf.EntryListFileFormat

	var _backupData []byte
	_backupData, _ = os.ReadFile(filePath)

	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(_data)
	if err != nil {
		if _, err = file.Write(_backupData); err != nil {
			errstr := fmt.Sprintf("can't write new and even original data to the file... %v", err)
			return errors.New(errstr)
		}
	}

	return nil
}
