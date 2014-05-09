package main

import (
	"bytes"
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

	mentions := make(map[string]int)
	inner_bound := make([]string, 0)
	// outer_bound := ma([]string, 0)
	for _, word := range words {
		results, _ := SearchEngine.Query(string(word), 100)
		for _, item := range results {
			mentions[string(item)] = mentions[string(item)] + 1
			inner_bound = append(inner_bound, string(item))
		}
	}

	// first comes inner sum
	for key, value := range mentions {
		if len(output) < count && value > 1 {
			output = append(output, ValueIds[key])
		}
	}

	for _, term := range inner_bound {
		if len(output) < count {
			output = append(output, ValueIds[term])
		}
	}

	return output
}
