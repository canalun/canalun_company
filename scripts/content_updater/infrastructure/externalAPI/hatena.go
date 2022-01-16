package externalAPI

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"

	"content-updater/domain/model"
	"content-updater/infrastructure/env_setter"
)

const (
	relWithEntryLink = "alternate"
)

var hatenaEnv = env_setter.HatenaEnv{}

func InitHatenaEnv() {
	hatenaEnv = env_setter.GetHatenaEnvFromOSEnv()
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
		UserID:   hatenaEnv.User_name,
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
	Control struct {
		Text  string `xml:",chardata"`
		Draft string `xml:"draft"`
	} `xml:"control"`
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

func (a HatenaRepository) GetLatestEntryList() (*model.EntryList, error) {
	erd, err := a.getLatestEntryRelatedData()
	if err != nil {
		return nil, err
	}
	entryList := a.createEntryListFromEntryRelatedData(*erd)
	return &entryList, nil
}

func (a HatenaRepository) getLatestEntryRelatedData() (*hatenaEntryRelatedData, error) {
	url := "https://blog.hatena.ne.jp/" + a.ID + "/" + a.BlogID + "/atom/entry"
	fmt.Println(url)

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

	re, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(re))

	var hatenaEntryRelatedData hatenaEntryRelatedData
	xml.Unmarshal(re, &hatenaEntryRelatedData)
	return &hatenaEntryRelatedData, nil
}

func (a HatenaRepository) createEntryListFromEntryRelatedData(erd hatenaEntryRelatedData) model.EntryList {
	entryList := model.EntryList{
		Source: model.HatenaSource,
	}
	for _, entry := range erd.Entries {
		var linkToEntry string
		for _, link := range entry.Links {
			if link.Rel == relWithEntryLink {
				linkToEntry = link.Href
			}
		}
		if linkToEntry == "" {
			fmt.Printf("can't get link for `%#v`\n", entry.Title)
		}
		entryList.Entries = append(entryList.Entries, model.Entry{
			Title: entry.Title,
			URL:   linkToEntry,
			//TODO:  LastUpdatedAtをとるようにする
		})
	}
	return entryList
}
