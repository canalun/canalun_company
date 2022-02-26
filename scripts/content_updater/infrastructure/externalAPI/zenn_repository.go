package externalAPI

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"time"

	"content-updater/domain/model"
	"content-updater/infrastructure/env_setter"
)

const (
	zennBaseUrlPrefix = "https://zenn.dev/"
	zennBaseUrlSuffix = "/feed"
)

var zennEnv = env_setter.ZennEnv{}

func InitZennEnv() {
	zennEnv = env_setter.GetZennEnvFromOSEnv()
}

type ZennRepository struct {
	UserID string
}

func NewZennRepository() ZennRepository {
	return ZennRepository{
		UserID: zennEnv.User_id,
	}
}

type zennEntryRelatedData struct {
	XMLName xml.Name `xml:"rss"`
	Text    string   `xml:",chardata"`
	Dc      string   `xml:"dc,attr"`
	Content string   `xml:"content,attr"`
	Atom    string   `xml:"atom,attr"`
	Version string   `xml:"version,attr"`
	Channel channel  `xml:"channel"`
}

type channel struct {
	Text        string `xml:",chardata"`
	Title       string `xml:"title"`
	Description string `xml:"description"`
	Link        struct {
		Text string `xml:",chardata"`
		Href string `xml:"href,attr"`
		Rel  string `xml:"rel,attr"`
		Type string `xml:"type,attr"`
	} `xml:"link"`
	Image struct {
		Text  string `xml:",chardata"`
		URL   string `xml:"url"`
		Title string `xml:"title"`
		Link  string `xml:"link"`
	} `xml:"image"`
	Generator     string      `xml:"generator"`
	LastBuildDate string      `xml:"lastBuildDate"`
	Language      string      `xml:"language"`
	Entries       []zennEntry `xml:"item"`
}

type zennEntry struct {
	Text        string `xml:",chardata"`
	Title       string `xml:"title"`
	Description string `xml:"description"`
	Link        string `xml:"link"`
	Guid        struct {
		Text        string `xml:",chardata"`
		IsPermaLink string `xml:"isPermaLink,attr"`
	} `xml:"guid"`
	PubDate   string `xml:"pubDate"`
	Enclosure struct {
		Text   string `xml:",chardata"`
		URL    string `xml:"url,attr"`
		Length string `xml:"length,attr"`
		Type   string `xml:"type,attr"`
	} `xml:"enclosure"`
	Creator string `xml:"creator"`
}

func (a ZennRepository) GetLatestEntryList() (*model.EntryList, error) {
	erd, err := a.getLatestEntryRelatedData()
	if err != nil {
		return nil, err
	}
	entryList := a.createEntryListFromEntryRelatedData(*erd)
	return &entryList, nil
}

func (a ZennRepository) getLatestEntryRelatedData() (*zennEntryRelatedData, error) {
	url := zennBaseUrlPrefix + a.UserID + zennBaseUrlSuffix

	client := &http.Client{
		Timeout: 30000000000, //nano sec
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	re, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var zennEntryRelatedData zennEntryRelatedData
	xml.Unmarshal(re, &zennEntryRelatedData)
	return &zennEntryRelatedData, nil
}

func (a ZennRepository) createEntryListFromEntryRelatedData(erd zennEntryRelatedData) model.EntryList {
	entryList := model.EntryList{
		Source: model.ZennSource,
	}
	for _, entry := range erd.Channel.Entries {
		t, _ := time.Parse(time.RFC1123, entry.PubDate)
		entryList.Entries = append(entryList.Entries, model.Entry{
			Title:         entry.Title,
			URL:           entry.Link,
			LastUpdatedAt: &t,
		})
	}
	return entryList
}
