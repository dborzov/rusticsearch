package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/argusdusty/Ferret"
	_ "github.com/go-sql-driver/mysql"
)

const WORD_KEY = "name"

var entryMap map[string]interface{}
var jsonDict interface{}
var jsonString []byte

func Updater() {
	// Here be syncing DB with the index

}

func loadSearchItems() {
	db, err := sql.Open("mysql", *MySQLAddress)
	if err != nil {
		fmt.Println("Unable to connect to that DB address :(")
		fmt.Println("-----------------------------------")
		panic(err.Error())
	}
	defer db.Close()

	rows, err := db.Query(`SELECT 
								product.id, 
								product.name,
								product.category_id,
								category.name,
								vendor.name,
								vendor_inventory.regular_price
							FROM 
							    product,
							    category,
							    vendor_inventory, 
							    vendor
							WHERE
							    product.category_id=category.id 
							  AND 
							    product.content_status="published"
						      AND
							    vendor_inventory.product_id=product.id
						      AND
						        vendor_inventory.vendor_id=vendor.id
							    ;`)

	if err != nil {
		fmt.Println("Unable to connect to that DB address")
		fmt.Println("-----------------------------------")
		panic(err.Error())
	}

	// populating entries cycle
	Words := make([]string, 0)
	Values := make([]interface{}, 0)
	for rows.Next() {
		var id, name, category_id, category_name, vendor, price string
		if err := rows.Scan(&id, &name, &category_id, &category_name, &vendor, &price); err != nil {
			panic(err.Error())
		}

		entry := map[string]string{
			"id":          id,
			"name":        name,
			"category_id": category_id,
			"category":    category_name,
		}

		keyWord := vendor + name
		Words = append(Words, keyWord)
		Values = append(Values, 10) // to add some fancy prioritizing here

		jsonString, _ = json.Marshal(entry)
		ValueIds[keyWord] = entry
	}

	SearchEngine = ferret.New(Words, Words, Values, Converter)
}
