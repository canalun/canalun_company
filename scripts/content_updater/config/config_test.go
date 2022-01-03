package config

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestInitConfig(t *testing.T) {
	tests := []struct {
		name string
		want Config
	}{
		{
			name: "正しくconfigよみこむ",
			want: Config{
				NumOfEntry:          5,
				EntryListFilePath:   "../../company_home/entry_list/",
				EntryListFileFormat: ".json",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			InitConfig()
		})
		if diff := cmp.Diff(Conf, tt.want); diff != "" {
			t.Errorf(diff)
		}
	}
}
