package main

import (
	"fmt"
	"github.com/argusdusty/Ferret"
	"net/http"
	"strconv"
)

var Correction = func(b []byte) [][]byte { return ferret.ErrorCorrect(b, ferret.LowercaseLetters) }
var LengthSorter = func(s string, v interface{}, l int, i int) float64 { return -float64(l + i) }
var FreqSorter = func(s string, v interface{}, l int, i int) float64 { return float64(v.(uint64)) }
var Converter = ferret.UnicodeToLowerASCII

var ValueIds = make(map[string]interface{})
var SearchEngine *ferret.InvertedSuffix
var search_query string

func main() {
	fmt.Println("Hi, I am Rustic Search Server!")
	config()
	loadSearchItems()

	mux := http.NewServeMux()
	mux.HandleFunc("/searchpage/", papaHandler(handler_searchpage))
	mux.HandleFunc("/autocomplete/", papaHandler(handler_autocomplete))
	mux.HandleFunc("/", papaHandler(handler_welcome))
	mux.HandleFunc("/reset", papaHandler(handler_reset))

	fmt.Println("Created index...")
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~")
	fmt.Printf("   Starting server at port %v... \n", ":"+strconv.Itoa(*port))
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~")
	err := http.ListenAndServe(":"+strconv.Itoa(*port), mux)
	if err != nil {
		fmt.Printf("Dang it! Error at ListenAndServe: %v \n", err)
		panic(err)
	}
}
