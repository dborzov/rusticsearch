package main

import (
	"bytes"
	"fmt"
	"github.com/argusdusty/Ferret"
	"io/ioutil"
	"net/http"
)

type SearchItem struct {
	id   []byte
	name []byte
}

var Correction = func(b []byte) [][]byte { return ferret.ErrorCorrect(b, ferret.LowercaseLetters) }
var LengthSorter = func(s string, v interface{}, l int, i int) float64 { return -float64(l + i) }
var FreqSorter = func(s string, v interface{}, l int, i int) float64 { return float64(v.(uint64)) }
var Converter = ferret.UnicodeToLowerASCII
var SearchEngine *ferret.InvertedSuffix
var ValueIds map[string]SearchItem

func main() {
	fmt.Println("Hi, I am Rustic Search Server!")
	Data, err := ioutil.ReadFile("search_index.csv")
	if err != nil {
		fmt.Println("search_index.csv not found :(")
		panic(err)
	}

	fmt.Println("Parsing search_index.csv...")
	Words := make([]string, 0)
	Values := make([]interface{}, 0)
	ValueIds = make(map[string]SearchItem)
	for i, Vals := range bytes.Split(Data, []byte("\n")) {
		WordFreq := bytes.Split(Vals, []byte("----------> "))
		if len(WordFreq) != 2 {
			fmt.Printf("Bollocks! search_index.csv line: %v breaks everything: \n \"%v\" \n I quit! \n", i, string(Vals))
			panic(Vals)
		}

		Words = append(Words, string(WordFreq[0]))
		// to add some priority mechanism in here in the future
		Values = append(Values, 10)
		ValueIds[string(WordFreq[0])] = SearchItem{WordFreq[0], WordFreq[1]}
	}

	fmt.Println("Created index...")
	SearchEngine = ferret.New(Words, Words, Values, Converter)

	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~")
	fmt.Println("   Starting server...")
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	h := w.Header()
	h.Set("Access-Control-Allow-Origin", "*")
	h.Set("Access-Control-Allow-Methods", "POST, GET, PUT, PATCH, DELETE, OPTIONS")
	h.Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, API-Date, Host, Authorization, Key, User-Token")
	h.Set("Access-Control-Max-Age", "1728000")

	fmt.Println("ANOTHER REQUEST: /", r.URL.Path[1:])
	results, _ := SearchEngine.Query(r.URL.Path[1:], 5)
	fmt.Fprintf(w, "{\"products\":[")
	for i, word := range results {
		fmt.Println("~~~~~~~~~~~~ ", word)
		fmt.Fprintf(w, "  {\"id\": \"%d\", \"name\": \"%s\"}", i, word)
		if i != len(results)-1 {
			fmt.Fprintf(w, ",")
		}
	}
	fmt.Fprintf(w, "]}")
}
