package model

import (
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
	Title         string     `json:"title"`
	URL           string     `json:"url"`
	LastUpdatedAt *time.Time `json:"lastUpdatedAt"`
}

func (a EntryList) SortEntries() {
	for _, item := range a.Entries {
		if item.LastUpdatedAt == nil {
			return
		}
	}
	sort.Slice(a.Entries, func(i, j int) bool { return a.Entries[i].LastUpdatedAt.After(*a.Entries[j].LastUpdatedAt) })
}
