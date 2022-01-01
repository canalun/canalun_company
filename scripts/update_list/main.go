package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/kelseyhightower/envconfig"
)

type Env struct {
	Id        string
	Blog_id   string
	User_name string
	Password  string
}

type HatenaResponse struct {
	Feed struct {
		Link []struct {
			Href string `xml:"href,attr"`
		} `xml:"link"`
		Title    string    `xml:"title"`
		Subtitle string    `xml:"subtitle"`
		Updated  time.Time `xml:"updated"`
		Author   struct {
			Name string `xml:"name"`
		} `xml:"author"`
		Generator string `xml:"generator"`
		ID        string `xml:"id"`
		Entry     []struct {
			ID     string   `xml:"id"`
			Link   []string `xml:"link"`
			Author struct {
				Name string `xml:"name"`
			} `xml:"author"`
			Title            string    `xml:"title"`
			Updated          time.Time `xml:"updated"`
			Published        time.Time `xml:"published"`
			Edited           time.Time `xml:"edited"`
			Summary          string    `xml:"summary"`
			Content          string    `xml:"content"`
			FormattedContent string    `xml:"formatted-content,omitempty"`
			Category         []string  `xml:"category,omitempty"`
			Control          struct {
				Draft string `xml:"draft"`
			} `xml:"control"`
		} `xml:"entry"`
	} `xml:"feed"`
}

func main() {
	var goenv Env
	envconfig.Process("HATENA", &goenv)

	url := "https://blog.hatena.ne.jp/" + goenv.Id + "/" + goenv.Blog_id + "/atom/entry"

	client := &http.Client{Timeout: time.Duration(30) * time.Second}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.SetBasicAuth(goenv.User_name, goenv.Password)

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var hatenaResponse HatenaResponse
	xml.Unmarshal(b, &hatenaResponse)
	fmt.Printf("%#v\n", hatenaResponse)
}
