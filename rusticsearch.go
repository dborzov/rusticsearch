package main

import (
	"bytes"
	"fmt"
	"github.com/argusdusty/Ferret"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

var Correction = func(b []byte) [][]byte { return ferret.ErrorCorrect(b, ferret.LowercaseLetters) }
var LengthSorter = func(s string, v interface{}, l int, i int) float64 { return -float64(l + i) }
var FreqSorter = func(s string, v interface{}, l int, i int) float64 { return float64(v.(uint64)) }
var Converter = ferret.UnicodeToLowerASCII
var SearchEngine *ferret.InvertedSuffix

func main() {
	t := time.Now()
	Data, err := ioutil.ReadFile("search_index.csv")
	if err != nil {
		panic(err)
	}
	Words := make([]string, 0)
	Values := make([]interface{}, 0)
	for _, Vals := range bytes.Split(Data, []byte("\n")) {
		Vals = bytes.TrimSpace(Vals)
		WordFreq := bytes.Split(Vals, []byte("&&&"))
		if len(WordFreq) != 2 {
			continue
		}
		Freq, err := strconv.ParseUint(string(WordFreq[1]), 10, 64)
		if err != nil {
			continue
		}
		Words = append(Words, string(WordFreq[0]))
		Values = append(Values, Freq)
	}
	fmt.Println("Loaded dictionary in:", time.Now().Sub(t))
	t = time.Now()

	SearchEngine = ferret.New(Words, Words, Values, Converter)
	fmt.Println("Created index in:", time.Now().Sub(t))
	t = time.Now()
	fmt.Println(SearchEngine.Query("ar", 5))
	fmt.Println("Performed search in:", time.Now().Sub(t))
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ANOTHER REQUEST: /", r.URL.Path[1:])
	results, _ := SearchEngine.Query(r.URL.Path[1:], 5)
	fmt.Fprintf(w, "{[\n")
	for _, word := range results {
		fmt.Println("~~~~~~~~~~~~ ", word)
		fmt.Fprintf(w, "    \"%s\",\n", word)
	}
	fmt.Fprintf(w, "]}")
}
