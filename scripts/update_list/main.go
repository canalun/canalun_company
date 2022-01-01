package main

import (
	"encoding/xml"
	"fmt"
	"log"

	"github.com/kelseyhightower/envconfig"
)

type Env struct {
	Id        string
	Blog_id   string
	User_name string
	Password  string
}

const (
	EnumNumOfArticle = 5
)

func main() {
	var goenv Env
	envconfig.Process("HATENA", &goenv)

	var hatenaResponse HatenaResponse
	hr, err := GetHatenaResponse(goenv.Id, goenv.Blog_id, goenv.User_name, goenv.Password)
	if err != nil {
		log.Fatal(err)
	}

	xml.Unmarshal(hr, &hatenaResponse)

	for idx := range hr {
		fmt.Printf("%#v, %#v\n", idx, hr[idx])
	}
}
