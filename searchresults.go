package main

type SearchItem struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type SearchResult struct {
	Results []SearchItem `json:"products"`
}
