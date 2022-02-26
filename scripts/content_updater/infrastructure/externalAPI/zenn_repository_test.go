package externalAPI

import (
	"content-updater/domain/model"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestZennRepository_getLatestEntryRelatedData(t *testing.T) {
	type fields struct {
		UserID string
	}
	tests := []struct {
		name    string
		fields  fields
		want    *zennEntryRelatedData
		wantErr bool
	}{
		{
			name: "データを正しくパースできる",
			fields: fields{
				UserID: "test",
			},
			want: &zennEntryRelatedData{
				Channel: channel{
					Entries: []zennEntry{
						{
							Title:   "title-1",
							PubDate: "Thu, 17 Feb 2022 16:34:52 GMT",
							Link:    "link-1",
						},
						{
							Title:   "title-2",
							PubDate: "Thu, 17 Feb 2022 16:34:52 GMT",
							Link:    "link-2",
						},
						{
							Title:   "title-3",
							PubDate: "Thu, 17 Feb 2022 16:34:52 GMT",
							Link:    "link-3",
						},
						{
							Title:   "title-4",
							PubDate: "Thu, 17 Feb 2022 16:34:52 GMT",
							Link:    "link-4",
						},
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := ZennRepository{
				UserID: tt.fields.UserID,
			}
			got, err := a.getLatestEntryRelatedData()
			if (err != nil) != tt.wantErr {
				t.Errorf("ZennRepository.getLatestEntryRelatedData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			opt := cmp.Options{
				cmpopts.IgnoreFields(zennEntryRelatedData{}, "XMLName", "Text"),
				cmpopts.IgnoreFields(channel{}, "Text"),
				cmpopts.IgnoreFields(zennEntry{}, "Text"),
			}
			diff := cmp.Diff(got, tt.want, opt)
			if diff != "" {
				t.Errorf("ZennRepository.getLatestEntryRelatedData()  %v", diff)
			}
		})
	}
}

func TestZennRepository_createEntryListFromEntryRelatedData(t *testing.T) {
	time, err := time.Parse(time.RFC1123, "Thu, 17 Feb 2022 16:34:52 GMT")
	if err != nil {
		t.Errorf("cannot parse time")
	}
	type fields struct {
		UserID string
	}
	type args struct {
		erd zennEntryRelatedData
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   model.EntryList
	}{
		{
			name: "正常系: 下書き中はスキップ",
			args: args{
				erd: zennEntryRelatedData{
					Channel: channel{
						Entries: []zennEntry{
							{
								Title:   "title-1",
								PubDate: "Thu, 17 Feb 2022 16:34:52 GMT",
								Link:    "link-1",
							},
							{
								Title:   "title-2",
								PubDate: "Thu, 17 Feb 2022 16:34:52 GMT",
								Link:    "link-2",
							},
						},
					},
				},
			},
			want: model.EntryList{
				Source: "Zenn",
				Entries: []model.Entry{
					{
						Title:         "title-1",
						URL:           "link-1",
						LastUpdatedAt: &time,
					},
					{
						Title:         "title-2",
						URL:           "link-2",
						LastUpdatedAt: &time,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := ZennRepository{
				UserID: tt.fields.UserID,
			}
			got := a.createEntryListFromEntryRelatedData(tt.args.erd)
			diff := cmp.Diff(got, tt.want)
			if diff != "" {
				t.Errorf("ZennRepository.createEntryListFromEntryRelatedData() = %v, want %v", got, tt.want)
			}
		})
	}
}
