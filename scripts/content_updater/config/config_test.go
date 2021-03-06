package config

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestInitConfig(t *testing.T) {
	tests := []struct {
		name           string
		configFilePath string
		want           Config
	}{
		{
			name:           "正しくconfigよみこむ",
			configFilePath: "../config.yml",
			want: Config{
				NumOfEntry:          5,
				EntryListFilePath:   "../../company_home/materials/entry_list/",
				EntryListFileFormat: ".json",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := InitConfig(tt.configFilePath); err != nil {
				t.Error(err)
			}
		})
		if diff := cmp.Diff(Conf, tt.want); diff != "" {
			t.Errorf(diff)
		}
	}
}
