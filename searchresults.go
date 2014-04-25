package main

type SearchItem struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type SearchPageItem struct {
	Id     string   `json:"id"`
	Name   string   `json:"name"`
	Price  string   `json:"price"`
	Images []string `json:"images"`
}

type SearchResult struct {
	Results []interface{} `json:"products"`
}

type SearchPageResult struct {
	Results []SearchPageItem `json:"products"`
}

type SearchEntries struct {
	Entries []interface{} `json:"products"`
}
