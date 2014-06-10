package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/argusdusty/Ferret"
	_ "github.com/go-sql-driver/mysql"
	"strings"
	"time"
)

const WORD_KEY = "name"

var entryMap map[string]interface{}
var jsonDict interface{}
var jsonString []byte

func Updater() {
	// Here be periodic syncing with the index
	for {
		time.Sleep(time.Duration(*refresh_time) * time.Minute)
		fmt.Printf("TIME TO REFRESH THE SEARCH INDEX\n")
		loadSearchItems()
	}
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
								primary_category.name,
								product.subcategory_id,
								subcategory.name,
								vendor.name,
								vendor_inventory.regular_price,
								GROUP_CONCAT(image.url_src)
							FROM 
							    product,
							    products_to_images,
							    image,
							    category AS primary_category,
							    category AS subcategory,
							    vendor_inventory, 
							    vendor
							WHERE
							    products_to_images.product_id = product.id
							  AND
							    products_to_images.image_id = image.id
							  AND
							    product.category_id=primary_category.id 
							  AND
							    product.subcategory_id=subcategory.id
							  AND 
							    product.content_status="published"
						      AND
							    vendor_inventory.product_id=product.id
						      AND
						        vendor_inventory.vendor_id=vendor.id
						      AND
						        vendor_inventory.is_published=1
						    GROUP BY
						        product.name
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
