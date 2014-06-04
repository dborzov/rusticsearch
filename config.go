package main

import (
	"flag"
)

var port = flag.Int("port", 8080, "a serving TCP port")
var input_file = flag.String("datafile", "search_index.json", "input data filepath")
var MySQLAddress = flag.String("DBaddr", "root@tcp(localhost:3306)/peter", "Database connection address")

func config() {
	flag.Parse()
}
