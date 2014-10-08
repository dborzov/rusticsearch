package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	ferret "github.com/argusdusty/Ferret"
	_ "github.com/go-sql-driver/mysql"
)

var entryMap map[string]interface{}
var jsonDict interface{}
var jsonString []byte

// Updater requeries and rebuilds search index periodically (every *refreshTime)
//thus syncing all the changes
func Updater() {
	// Here be periodic syncing with the index
	for {
		time.Sleep(time.Duration(*refreshTime) * time.Minute)
		fmt.Printf("TIME TO REFRESH THE SEARCH INDEX\n")
		loadSearchItems()
	}
}

func loadSearchItems() {
	db, err := sql.Open("mysql", *mySQLAddress)
	if err != nil {
		fmt.Println("Unable to connect to that DB address :(")
		fmt.Println("-----------------------------------")
		panic(err.Error())
	}
	defer db.Close()

	rows, err := db.Query(config.SQLQuery)

	if err != nil {
		fmt.Println("Unable to connect to that DB address")
		fmt.Println("-----------------------------------")
		panic(err.Error())
	}

	// populating entries cycle
	Words := make([]string, 0)
	Values := make([]interface{}, 0)
	for rows.Next() {
		var id, name, category_id, category_name, subcategory_id, subcategory_name, vendor, price, images string
		if err := rows.Scan(&id, &name, &category_id, &category_name, &subcategory_id, &subcategory_name, &vendor, &price, &images); err != nil {
			panic(err.Error())
		}

		list_images := strings.Split(images, ",")

		entry := map[string]interface{}{
			"id":             id,
			"name":           name,
			"category_id":    category_id,
			"category":       category_name,
			"subcategory_id": subcategory_id,
			"subcategory":    subcategory_name,
			"price":          price,
			"vendor":         vendor,
			"images":         list_images,
		}

		keyWord := vendor + name
		Words = append(Words, keyWord)
		Values = append(Values, 10) // to add some fancy prioritizing here

		jsonString, _ = json.Marshal(entry)
		ValueIds[keyWord] = entry
	}

	SearchEngine = ferret.New(Words, Words, Values, Converter)
}
