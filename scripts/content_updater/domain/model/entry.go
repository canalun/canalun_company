package model

import (
	"net/url"
	"sort"
	"time"
)

const (
	HatenaSource = "Hatena"
	ZennSource   = "Zenn"
)

type EntryList struct {
	Source  string
	Entries []Entry
}

type Entry struct {
	Title       string
	URL         *url.URL
	PublishedAt time.Time
}

func (a EntryList) SortEntries() {
	sort.Slice(a.Entries, func(i, j int) bool { return a.Entries[i].PublishedAt.After(a.Entries[j].PublishedAt) })
}
