package externalAPI

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"content-updater/domain/model"
)

const (
	relWithEntryLink = "alternate"
)

type HatenaRepository struct {
	ID       string
	BlogID   string
	UserName string
	Password string
}

func NewHatenaRepository(ID string, blogID string, userName string, password string) HatenaRepository {
	return HatenaRepository{
		ID:       ID,
		BlogID:   blogID,
		UserName: userName,
		Password: password,
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
	Title     string `xml:"title"`
	Updated   string `xml:"updated"`
	Published string `xml:"published"`
	Edited    string `xml:"edited"`
	Summary   struct {
		Text string `xml:",chardata"`
		Type string `xml:"type,attr"`
	} `xml:"summary"`
	Content struct {
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

func (a HatenaRepository) GetArticles() ([]model.Article, error) {
	erd, err := a.getEntryRelatedData()
	if err != nil {
		return nil, err
	}
	var articles []model.Article
	for _, entry := range erd.Entries {
		var linkToEntry string
		for _, link := range entry.Links {
			if link.Rel == relWithEntryLink {
				linkToEntry = link.Href
			}
		}
		url, err := url.Parse(linkToEntry)
		if err != nil {
			fmt.Printf("%#v; %#v\n", entry.Title, err)
		}
		articles = append(articles, model.Article{
			Title: entry.Title,
			URL:   url,
		})
	}
	return articles, nil
}

func (a HatenaRepository) getEntryRelatedData() (*hatenaEntryRelatedData, error) {
	url := "https://blog.hatena.ne.jp/" + a.ID + "/" + a.BlogID + "/atom/entry"

	client := &http.Client{
		Timeout: 30000000000,
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(a.UserName, a.Password)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	re, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var hatenaEntryRelatedData hatenaEntryRelatedData
	xml.Unmarshal(re, &hatenaEntryRelatedData)
	return &hatenaEntryRelatedData, nil
}
