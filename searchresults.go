package main

import (
	"bytes"
	"fmt"
)

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

func SearchQuery(query string, count int) []interface{} {
	// ideally, the result would exactly match
	results, _ := SearchEngine.Query(query, count)
	output := make([]interface{}, len(results))
	for i, item := range results {
		output[i] = ValueIds[string(item)]
	}
	if len(results) != 0 {
		return output
	}

	// exact match did not work, try word by word
	words := bytes.Split([]byte(query), []byte(" "))
	fmt.Printf("The words: ", words, " \n")
	for _, word := range words {
		results, _ := SearchEngine.Query(string(word), count)
		for _, item := range results {
			output = append(output, ValueIds[string(item)])
		}
	}

	return output
}
