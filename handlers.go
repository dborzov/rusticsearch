package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"
)

type Handler func(w http.ResponseWriter, r *http.Request)

func papaHandler(a Filler) Handler {
	rh := func(w http.ResponseWriter, r *http.Request) {
		//setting the desired header
		h := w.Header()
		h.Set("Access-Control-Allow-Origin", "*")
		h.Set("Access-Control-Allow-Methods", "POST, GET, PUT, PATCH, DELETE, OPTIONS")
		h.Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, API-Date, Host, Authorization, Key, User-Token")
		h.Set("Access-Control-Max-Age", "1728000")
		preresponse, err := a(r)
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, err.Error())
			return
		}
		response := string(preresponse)
		fmt.Fprintf(w, response)

		// now logging the request
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

func handler_welcome(r *http.Request) ([]byte, error) {
	fmt.Printf("HELLO REQUEST!!\n")
	return []byte("Welcome to Rustic Search!"), nil
}

func handler_searchpage(r *http.Request) ([]byte, error) {
	search_query = r.URL.Path[12:]
	fmt.Printf("SEARCHPAGE REQUEST: %s \n", search_query)
	results, _ := SearchEngine.Query(search_query, 100)
	output := make([]interface{}, 0)
	for _, word := range results {
		output = append(output, ValueIds[string(word)])
	}
	searchResults, err := json.Marshal(SearchResult{output})
	if err != nil {
		return []byte{}, errors.New("json.Marshaling the response failed :-(")
	}
	return searchResults, nil
}
