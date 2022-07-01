package file_generator

import (
	"content-updater/config"
	"content-updater/domain/model"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

func TestUpdateEntryList(t *testing.T) {
	type args struct {
		el model.EntryList
	}
	days := []time.Time{
		time.Date(2022, 1, 3, 0, 0, 0, 0, time.UTC),
		time.Date(2022, 1, 2, 0, 0, 0, 0, time.UTC),
		time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
	}
	tests := []struct {
		name       string
		initConfFn func() config.Config
		args       args
		assertFn   func() string
		wantErr    bool
	}{
		{
			name: "create valid output file",
			initConfFn: func() config.Config {
				return config.Config{
					NumOfEntry:          2,
					EntryListFilePath:   "./",
					EntryListFileFormat: ".json",
				}
			},
			args: args{
				model.EntryList{
					Entries: []model.Entry{
						{
							Title:         "title-1",
							URL:           "http://test.com/1",
							LastUpdatedAt: &days[0],
						},
						{
							Title:         "title-2",
							URL:           "http://test.com/2",
							LastUpdatedAt: &days[1],
						},
						{
							Title:         "title-3",
							URL:           "http://test.com/3",
							LastUpdatedAt: &days[2],
						},
					},
					Source: "test",
				},
			},
			assertFn: func() string {
				assertData := []byte(`[{"title":"title-1","url":"http://test.com/1","lastUpdatedAt":"2022-01-03T00:00:00Z"},{"title":"title-2","url":"http://test.com/2","lastUpdatedAt":"2022-01-02T00:00:00Z"}]`)
				data, _ := os.ReadFile(config.Conf.EntryListFilePath + "test" + config.Conf.EntryListFileFormat)
				return cmp.Diff(data, assertData)
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		config.Conf = tt.initConfFn()
		t.Run(tt.name, func(t *testing.T) {
			if err := UpdateEntryList(tt.args.el); (err != nil) != tt.wantErr {
				t.Errorf("UpdateEntryList() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
		if re := tt.assertFn(); re != "" {
			t.Errorf("UpdateEntryList() fail: %v", re)
		}
		if err := deleteContentFile(); err != nil {
			fmt.Println("could not delete the generated file. please delete manually.")
		}
	}
}

func deleteContentFile() error {
	if err := os.Remove(config.Conf.EntryListFilePath + "test" + config.Conf.EntryListFileFormat); err != nil {
		return err
	}
	return nil
}
