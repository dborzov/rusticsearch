package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

type Handler func(w http.ResponseWriter, r *http.Request)
type Filler func(r *http.Request) []byte

func papaHandler(a Filler) Handler {
	rh := func(w http.ResponseWriter, r *http.Request) {
		//setting the desired header
		h := w.Header()
		h.Set("Access-Control-Allow-Origin", "*")
		h.Set("Access-Control-Allow-Methods", "POST, GET, PUT, PATCH, DELETE, OPTIONS")
		h.Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, API-Date, Host, Authorization, Key, User-Token")
		h.Set("Access-Control-Max-Age", "1728000")
		response := string(a(r))
		fmt.Fprintf(w, response)
		// now log the fucker
		file, err := os.OpenFile("history.log", os.O_RDWR|os.O_APPEND, 0660)
		defer file.Close()
		if err != nil {
			fmt.Printf("Having trouble with Logger file here \n")
			return
		}
		file.WriteString("\n\n---------------------------\n")
		file.WriteString(fmt.Sprintf("Time is %s:\n", time.Now()))
		file.WriteString("Yet another request:\n\n")

		jr, _ := json.MarshalIndent(*r, "    ", "   ")
		file.Write(jr)
		file.WriteString("\n\nOur response shall be:\n\n")
		file.WriteString(response)
	}
	return rh
}

func handler_autocomplete(r *http.Request) []byte {
	search_query = r.URL.Path[14:]
	fmt.Printf("AUTOCOMPLETE REQUEST: %s \n", search_query)
	results, _ := SearchEngine.Query(search_query, 5)
	output := make([]interface{}, 0)
	for _, word := range results {
		fmt.Printf("~~~~~~~~~~~~: %v \n", string(word))
		output = append(output, ValueIds[string(word)])
		fmt.Println("Here is the value added: ", ValueIds[string(word)])
	}
	searchResults, err := json.Marshal(SearchResult{output})
	if err != nil {
		fmt.Printf("Error with jsonifying %s \n ", err)
	}
	return searchResults
}

func handler_welcome(r *http.Request) []byte {
	fmt.Printf("HELLO REQUEST:\n")
	return []byte("Welcome to Rustic Search!")
}

func handler_searchpage(r *http.Request) []byte {
	search_query = r.URL.Path[12:]
	fmt.Printf("SEARCHPAGE REQUEST: %s \n", search_query)
	results, _ := SearchEngine.Query(search_query, 100)
	output := make([]interface{}, 0)
	for _, word := range results {
		fmt.Printf("~~~~~~~~~~~~: %v \n", string(word))
		output = append(output, ValueIds[string(word)])
		fmt.Println("Here is the value added: ", ValueIds[string(word)])
	}
	searchResults, err := json.Marshal(SearchResult{output})
	if err != nil {
		panic(err)
	}
	return searchResults
}
