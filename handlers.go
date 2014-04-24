package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func handler_autocomplete(w http.ResponseWriter, r *http.Request) {
	h := w.Header()
	h.Set("Access-Control-Allow-Origin", "*")
	h.Set("Access-Control-Allow-Methods", "POST, GET, PUT, PATCH, DELETE, OPTIONS")
	h.Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, API-Date, Host, Authorization, Key, User-Token")
	h.Set("Access-Control-Max-Age", "1728000")

	search_query = r.URL.Path[14:]
	fmt.Printf("AUTOCOMPLETE REQUEST: %s \n", search_query)
	results, _ := SearchEngine.Query(search_query, 5)
	output := make([]SearchItem, 0)
	for _, word := range results {
		fmt.Printf("~~~~~~~~~~~~: %v \n", string(word))
		output = append(output, ValueIds[string(word)])
		fmt.Println("Here is the value added: ", ValueIds[string(word)])
	}
	searchResults, err := json.Marshal(SearchResult{output})
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(w, string(searchResults))
}

func handler_welcome(w http.ResponseWriter, r *http.Request) {
	h := w.Header()
	h.Set("Access-Control-Allow-Origin", "*")
	h.Set("Access-Control-Allow-Methods", "POST, GET, PUT, PATCH, DELETE, OPTIONS")
	h.Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, API-Date, Host, Authorization, Key, User-Token")
	h.Set("Access-Control-Max-Age", "1728000")

	fmt.Printf("HELLO REQUEST:\n")
	fmt.Fprintf(w, "Welcome to SowingoSearchServer!")
}

func handler_searchpage(w http.ResponseWriter, r *http.Request) {
	h := w.Header()
	h.Set("Access-Control-Allow-Origin", "*")
	h.Set("Access-Control-Allow-Methods", "POST, GET, PUT, PATCH, DELETE, OPTIONS")
	h.Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, API-Date, Host, Authorization, Key, User-Token")
	h.Set("Access-Control-Max-Age", "1728000")

	search_query = r.URL.Path[14:]
	fmt.Printf("SEARCHPAGE REQUEST: %s \n", search_query)
	results, _ := SearchEngine.Query(search_query, 100)
	output := make([]SearchPageItem, 0)
	for _, word := range results {
		fmt.Printf("~~~~~~~~~~~~: %v \n", string(word))
		output = append(output, ValueIds[string(word)].Convert2SearchPageItem())
		fmt.Println("Here is the value added: ", ValueIds[string(word)])
	}
	searchResults, err := json.Marshal(SearchPageResult{output})
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w, string(searchResults))
}
