package model

import "net/url"

type Article struct {
	Title string
	URL   *url.URL
}
