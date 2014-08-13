package main

import (
	"flag"
)

var port = flag.Int("port", 8080, "a serving TCP port")
var refreshTime = flag.Int("refresh_time", 10, "time period (in minutes) when the search index is refreshed")
var mySQLAddress = flag.String("DBaddr", "root@tcp(localhost:3306)/peter", "Database connection address")

func config() {
	flag.Parse()
}
