package main

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"time"
)

type HatenaResponse struct {
	XMLName xml.Name `xml:"feed"`
	Text    string   `xml:",chardata"`
	Xmlns   string   `xml:"xmlns,attr"`
	App     string   `xml:"app,attr"`
	Link    []struct {
		Text string `xml:",chardata"`
		Rel  string `xml:"rel,attr"`
		Href string `xml:"href,attr"`
	} `xml:"link"`
	Title   string `xml:"title"`
	Updated string `xml:"updated"`
	Author  struct {
		Text string `xml:",chardata"`
		Name string `xml:"name"`
	} `xml:"author"`
	Generator struct {
		Text    string `xml:",chardata"`
		URI     string `xml:"uri,attr"`
		Version string `xml:"version,attr"`
	} `xml:"generator"`
	ID    string `xml:"id"`
	Entry []struct {
		Text string `xml:",chardata"`
		ID   string `xml:"id"`
		Link []struct {
			Text string `xml:",chardata"`
			Rel  string `xml:"rel,attr"`
			Href string `xml:"href,attr"`
			Type string `xml:"type,attr"`
		} `xml:"link"`
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
	} `xml:"entry"`
}

func GetHatenaResponse(ID string, blogID string, userName string, password string) ([]byte, error) {
	url := "https://blog.hatena.ne.jp/" + ID + "/" + blogID + "/atom/entry"

	client := &http.Client{Timeout: time.Duration(30) * time.Second}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(userName, password)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	re, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return re, nil
}
