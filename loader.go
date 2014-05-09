package main

import (
	"encoding/json"
	"fmt"
	"github.com/argusdusty/Ferret"
	"io/ioutil"
)

const WORD_KEY = "name"

var entryMap map[string]interface{}
var jsonDict interface{}
var jsonString []byte

func loadSearchItems() {
	Data, err := ioutil.ReadFile(*input_file)
	if err != nil {
		fmt.Println("search_index.json not found :(")
		panic(err)
	}

	json.Unmarshal(Data, &jsonDict)
	entries := jsonDict.([]interface{})

	// populating entries cycle
	Words := make([]string, 0)
	Values := make([]interface{}, 0)
	for _, entry := range entries {
		entryMap = entry.(map[string]interface{})
		keyWord := entryMap["vendor"].(string) + entryMap[WORD_KEY].(string)
		Words = append(Words, keyWord)
		Values = append(Values, 10)

		jsonString, _ = json.Marshal(entry)
		ValueIds[keyWord] = entry
	}

	SearchEngine = ferret.New(Words, Words, Values, Converter)
}
