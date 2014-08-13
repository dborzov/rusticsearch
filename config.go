package main

import (
	"flag"
	"fmt"
	"os"
)

var port = flag.Int("port", 8080, "a serving TCP port")
var refreshTime = flag.Int("refreshTime", 10, "time period (in minutes) when the search index is refreshed")
var mySQLAddress = flag.String("DBaddr", "root@tcp(localhost:3306)/peter", "Database connection address")
var configFilePath = flag.String("configFile", defaultConfigFilePath, "configuration json filepath")

const defaultConfigFilePath = "rusticsearch.config"
const configFileTemplate = `{
			\"port\":8080,
			\"refreshTime\":10,
}`


func config() {
	flag.Parse()
	if _, err := os.Stat(*configFilePath); os.IsNotExist(err) {
		if *configFilePath == defaultConfigFilePath {
			fmt.Printf("No config file found, creating a blank one: %s \n", defaultConfigFilePath)
		} else {
			fmt.Printf("File %s does not seem to exist \n", *configFilePath)
		}
	}
}
