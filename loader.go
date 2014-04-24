package main

import (
	"bytes"
	"fmt"
	"github.com/argusdusty/Ferret"
	"io/ioutil"
)

func loadSearchItems() {
	Data, err := ioutil.ReadFile(*input_file)
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
		ValueIds[string(WordFreq[0])] = SearchItem{string(WordFreq[1]), string(WordFreq[0])}
	}

	SearchEngine = ferret.New(Words, Words, Values, Converter)

}
