package main

import (
	"flag"
)

var port = flag.Int("port", 8080, "a serving TCP port")
var input_file = flag.String("datafile", "search_index.json", "input data filepath")

func config() {
	flag.Parse()
}
