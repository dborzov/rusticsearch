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

func SearchQuery(query string, count int) []interface{} {
	// ideally, the result would exactly match
	results, _ := SearchEngine.Query(query, count)
	output := make([]interface{}, len(results))
	for i, word := range results {
		output[i] = ValueIds[string(word)]
	}
	if len(results) != 0 {
		return output
	}

	return output
}
