package main

import (
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
	fmt.Println(string(b))
}
