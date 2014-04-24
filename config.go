package main

import (
	"flag"
	"fmt"
)

var port = flag.Int("port", 8080, "a serving TCP port")
var input_file = flag.String("datafile", "search_index.csv", "input data filepath")

func config() {
	fmt.Println("Call rusticsearch -h for help with parameters")
	flag.Parse()
}
