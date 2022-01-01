package main

import (
	"fmt"
	"log"

	"github.com/kelseyhightower/envconfig"
)

func main() {
	var goenv Env
	envconfig.Process("HATENA", &goenv)

	hatenaErd, err := GetEntryRelatedDataFromHatena(goenv.Id, goenv.Blog_id, goenv.User_name, goenv.Password)
	if err != nil {
		log.Fatal(err)
	}
	hatenaEntries := EnumTitleAndLinkOfHatenaEntries(*hatenaErd)

	for _, item := range hatenaEntries {
		fmt.Printf("%#v, %#v\n", item.Title, item.Links[0].Href)
	}
}
