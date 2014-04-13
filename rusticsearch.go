package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/argusdusty/Ferret"
	"io/ioutil"
	"net/http"
	"strconv"
)

type SearchItem struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func (i SearchItem) String() string {
	return "id: " + string(i.Id) + ", name:" + string(i.Name)
}

type SearchResult struct {
	Results []SearchItem `json:"products"`
}

var Correction = func(b []byte) [][]byte { return ferret.ErrorCorrect(b, ferret.LowercaseLetters) }
var LengthSorter = func(s string, v interface{}, l int, i int) float64 { return -float64(l + i) }
var FreqSorter = func(s string, v interface{}, l int, i int) float64 { return float64(v.(uint64)) }
var Converter = ferret.UnicodeToLowerASCII
var ValueIds = map[string]SearchItem{}
var SearchEngine *ferret.InvertedSuffix

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
	for i, Vals := range bytes.Split(Data, []byte("\n")) {
		WordFreq := bytes.Split(Vals, []byte("----------> "))
		if len(WordFreq) != 2 {
			fmt.Printf("Bollocks! search_index.csv line: %v breaks everything: \n \"%v\" \n I quit! \n", i, string(Vals))
			panic(Vals)
		}

		Words = append(Words, string(WordFreq[0]))
		// to add some priority mechanism in here in the future
		Values = append(Values, 10)
		ValueIds[string(WordFreq[0])] = SearchItem{string(WordFreq[0]), string(WordFreq[1])}
	}

	fmt.Println("Created index...")
	SearchEngine = ferret.New(Words, Words, Values, Converter)

	http.HandleFunc("/", handler)
	port := flag.Int("port", 8080, "a serving TCP port")
	flag.Parse()
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~")
	fmt.Printf("   Starting server at port %v... \n", ":"+strconv.Itoa(*port))
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~")
	err = http.ListenAndServe(":"+strconv.Itoa(*port), nil)
	if err != nil {
		fmt.Printf("Dang it! Error at ListenAndServe: %v \n", err)
		panic(err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	h := w.Header()
	h.Set("Access-Control-Allow-Origin", "*")
	h.Set("Access-Control-Allow-Methods", "POST, GET, PUT, PATCH, DELETE, OPTIONS")
	h.Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, API-Date, Host, Authorization, Key, User-Token")
	h.Set("Access-Control-Max-Age", "1728000")

	fmt.Println("ANOTHER REQUEST: %s", r.URL.Path)
	results, _ := SearchEngine.Query(r.URL.Path[1:], 5)
	output := make([]SearchItem, 0)
	for _, word := range results {
		fmt.Println("~~~~~~~~~~~~: %v ", string(word))
		output = append(output, ValueIds[string(word)])
		fmt.Println("Here is the value added: ", ValueIds[string(word)].String())
	}
	searchResults, err := json.Marshal(SearchResult{output})
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w, string(searchResults))
}
