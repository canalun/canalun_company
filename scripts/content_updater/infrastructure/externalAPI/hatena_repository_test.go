package externalAPI

import (
	"content-updater/domain/model"
	"fmt"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/jarcoal/httpmock"
)

func TestMain(m *testing.M) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(
		"GET",
		"https://blog.hatena.ne.jp/test/test/atom/entry",
		httpmock.NewStringResponder(200, httpmock.File("./testdata/entry_related_data_test.xml").String()),
	)
	runTest := m.Run()
	os.Exit(runTest)
}

func TestHatenaRepository_getLatestEntryRelatedData(t *testing.T) {
	type fields struct {
		ID       string
		BlogID   string
		UserID   string
		Password string
	}
	tests := []struct {
		name    string
		fields  fields
		want    *hatenaEntryRelatedData
		wantErr bool
	}{
		{
			name: "データを正しくパースできる",
			fields: fields{
				ID:       "test",
				BlogID:   "test",
				UserID:   "test",
				Password: "test",
			},
			want: &hatenaEntryRelatedData{
				Entries: []hatenaEntry{
					{
						Title: "title-1",
						Links: []link{
							{
								Rel:  "edit",
								Href: "edit-link-1",
							},
							{
								Rel:  "alternate",
								Type: "text/html",
								Href: "alternate-link-1",
							},
						},
						Summary: summary{
							Type: "text",
							Text: "summary-1",
						},
						Control: control{
							Draft: "no",
						},
					},
					{
						Title: "title-2",
						Links: []link{
							{
								Rel:  "edit",
								Href: "edit-link-2",
							},
							{
								Rel:  "alternate",
								Type: "text/html",
								Href: "alternate-link-2",
							},
						},
						Summary: summary{
							Type: "text",
							Text: "summary-2",
						},
						Control: control{
							Draft: "no",
						},
					},
					{
						Title: "title-3",
						Links: []link{
							{
								Rel:  "edit",
								Href: "edit-link-3",
							},
							{
								Rel:  "alternate",
								Type: "text/html",
								Href: "alternate-link-3",
							},
						},
						Summary: summary{
							Type: "text",
							Text: "summary-3",
						},
						Control: control{
							Draft: "yes",
						},
					},
					{
						Title: "title-4",
						Links: []link{
							{
								Rel:  "edit",
								Href: "edit-link-4",
							},
							{
								Rel:  "alternate",
								Type: "text/html",
								Href: "alternate-link-4",
							},
						},
						Summary: summary{
							Type: "text",
							Text: "summary-4",
						},
						Control: control{
							Draft: "no",
						},
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := HatenaRepository{
				ID:       tt.fields.ID,
				BlogID:   tt.fields.BlogID,
				UserID:   tt.fields.UserID,
				Password: tt.fields.Password,
			}
			got, err := a.getLatestEntryRelatedData()
			if (err != nil) != tt.wantErr {
				t.Errorf("HatenaRepository.getLatestEntryRelatedData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			opt := cmp.Options{
				cmpopts.IgnoreFields(hatenaEntryRelatedData{}, "XMLName", "Text", "Xmlns", "App", "Links", "Title", "Generator", "ID"),
				cmpopts.IgnoreFields(hatenaEntry{}, "Text", "ID", "Author"),
				cmpopts.IgnoreFields(control{}, "Text"),
			}
			diff := cmp.Diff(got, tt.want, opt)
			fmt.Println(got)
			if diff != "" {
				t.Errorf("HatenaRepository.getLatestEntryRelatedData()  %v", diff)
			}
		})
	}
}

func TestHatenaRepository_createEntryListFromEntryRelatedData(t *testing.T) {
	type args struct {
		erd hatenaEntryRelatedData
	}
	tests := []struct {
		name string
		args args
		want model.EntryList
	}{
		{
			name: "正常系: 下書き中はスキップ",
			args: args{
				erd: hatenaEntryRelatedData{
					Entries: []hatenaEntry{
						{
							Title: "title-1",
							Links: []link{
								{
									Rel:  "edit",
									Href: "edit-link-1",
								},
								{
									Rel:  "alternate",
									Type: "text/html",
									Href: "alternate-link-1",
								},
							},
							Summary: summary{
								Type: "text",
								Text: "summary-1",
							},
							Control: control{
								Draft: "no",
							},
						},
						{
							Title: "title-2",
							Links: []link{
								{
									Rel:  "edit",
									Href: "edit-link-2",
								},
								{
									Rel:  "alternate",
									Type: "text/html",
									Href: "alternate-link-2",
								},
							},
							Summary: summary{
								Type: "text",
								Text: "summary-2",
							},
							Control: control{
								Draft: "yes",
							},
						},
						{
							Title: "title-3",
							Links: []link{
								{
									Rel:  "edit",
									Href: "edit-link-3",
								},
								{
									Rel:  "alternate",
									Type: "text/html",
									Href: "alternate-link-3",
								},
							},
							Summary: summary{
								Type: "text",
								Text: "summary-3",
							},
							Control: control{
								Draft: "no",
							},
						},
					},
				},
			},
			want: model.EntryList{
				Source: "Hatena",
				Entries: []model.Entry{
					{
						Title: "title-1",
						URL:   "alternate-link-1",
					},
					{
						Title: "title-3",
						URL:   "alternate-link-3",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := HatenaRepository{}
			got := a.createEntryListFromEntryRelatedData(tt.args.erd)
			diff := cmp.Diff(got, tt.want)
			if diff != "" {
				t.Errorf("HatenaRepository.createEntryListFromEntryRelatedData(); diff = %v", diff)
			}
		})
	}
}
