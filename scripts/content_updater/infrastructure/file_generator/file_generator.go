package file_generator

import (
	"content-updater/config"
	"content-updater/domain/model"
	"encoding/json"
	"os"
)

func UpdateEntryListFile(el model.EntryList) error {
	el.SortEntries()

	var data []model.Entry
	for i := 0; i < config.Conf.NumOfEntry; i++ {
		data = append(data, el.Entries[i])
	}
	_data, err := json.Marshal(data)
	if err != nil {
		return err
	}

	fileName := config.Conf.EntryListFilePath + el.Source + config.Conf.EntryListFileFormat
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(_data)
	if err != nil {
		//TODO: ここで前のファイルを復活させたい
		return err
	}

	return nil
}
