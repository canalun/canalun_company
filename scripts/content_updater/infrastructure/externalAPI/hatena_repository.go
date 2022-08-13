package externalAPI

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"time"

	"content-updater/domain/model"
	"content-updater/infrastructure/env_setter"
)

const (
	relWithEntryLink    = "alternate"
	hatenaBaseUrlPrefix = "https://blog.hatena.ne.jp/"
	hatenaBaseUrlSuffix = "/atom/entry"
)

var hatenaEnv = env_setter.HatenaEnv{}

func InitHatenaEnv() error {
	_hatenaEnv, err := env_setter.GetHatenaEnvFromOSEnv()
	if err != nil {
		return err
	}
	hatenaEnv = _hatenaEnv
	return nil
}

type HatenaRepository struct {
	ID       string
	BlogID   string
	UserID   string
	Password string
}

func NewHatenaRepository() HatenaRepository {
	return HatenaRepository{
		ID:       hatenaEnv.Id,
		BlogID:   hatenaEnv.Blog_id,
		UserID:   hatenaEnv.User_id,
		Password: hatenaEnv.Password,
	}
}

type hatenaEntryRelatedData struct {
	XMLName xml.Name `xml:"feed"`
	Text    string   `xml:",chardata"`
	Xmlns   string   `xml:"xmlns,attr"`
	App     string   `xml:"app,attr"`
	Links   []link   `xml:"link"`
	Title   string   `xml:"title"`
	Updated string   `xml:"updated"`
	Author  struct {
		Text string `xml:",chardata"`
		Name string `xml:"name"`
	} `xml:"author"`
	Generator struct {
		Text    string `xml:",chardata"`
		URI     string `xml:"uri,attr"`
		Version string `xml:"version,attr"`
	} `xml:"generator"`
	ID      string        `xml:"id"`
	Entries []hatenaEntry `xml:"entry"`
}

type hatenaEntry struct {
	Text   string `xml:",chardata"`
	ID     string `xml:"id"`
	Links  []link `xml:"link"`
	Author struct {
		Text string `xml:",chardata"`
		Name string `xml:"name"`
	} `xml:"author"`
	Title     string  `xml:"title"`
	Updated   string  `xml:"updated"`
	Published string  `xml:"published"`
	Edited    string  `xml:"edited"`
	Summary   summary `xml:"summary"`
	Content   struct {
		Text string `xml:",chardata"`
		Type string `xml:"type,attr"`
	} `xml:"content"`
	FormattedContent struct {
		Text   string `xml:",chardata"`
		Type   string `xml:"type,attr"`
		Hatena string `xml:"hatena,attr"`
	} `xml:"formatted-content"`
	Control control `xml:"control"`
}

type link struct {
	Text string `xml:",chardata"`
	Rel  string `xml:"rel,attr"`
	Href string `xml:"href,attr"`
	Type string `xml:"type,attr"`
}
type summary struct {
	Text string `xml:",chardata"`
	Type string `xml:"type,attr"`
}

type control struct {
	Text  string `xml:",chardata"`
	Draft string `xml:"draft"`
}

func (a HatenaRepository) GetLatestEntryList() (*model.EntryList, error) {
	erd, err := a.getLatestEntryRelatedData()
	if err != nil {
		return nil, err
	}
	entryList := a.createEntryListFromEntryRelatedData(*erd)
	return &entryList, nil
}

func (a HatenaRepository) getLatestEntryRelatedData() (*hatenaEntryRelatedData, error) {
	url := hatenaBaseUrlPrefix + a.ID + "/" + a.BlogID + hatenaBaseUrlSuffix

	client := &http.Client{
		Timeout: 30000000000, //nano sec
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(a.UserID, a.Password)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	re, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var hatenaEntryRelatedData hatenaEntryRelatedData
	if err := xml.Unmarshal(re, &hatenaEntryRelatedData); err != nil {
		return nil, err
	}
	return &hatenaEntryRelatedData, nil
}

func (a HatenaRepository) createEntryListFromEntryRelatedData(erd hatenaEntryRelatedData) model.EntryList {
	entryList := model.EntryList{
		Source: model.HatenaSource,
	}
	for _, entry := range erd.Entries {
		if entry.Control.Draft != "no" {
			continue
		}
		var linkToEntry string
		for _, link := range entry.Links {
			if link.Rel == relWithEntryLink {
				linkToEntry = link.Href
			}
		}
		if linkToEntry == "" {
			fmt.Printf("can't get link for `%#v`\n", entry.Title)
		}
		t, _ := time.Parse(time.RFC3339, entry.Published)
		entryList.Entries = append(entryList.Entries, model.Entry{
			Title:         entry.Title,
			URL:           linkToEntry,
			LastUpdatedAt: &t,
		})
	}
	return entryList
}
