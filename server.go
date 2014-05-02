package main

import (
	"net/http"
)

type Server struct {
	headers   map[string]string
	endpoints []Endpoint
}

var Header = http.Header{
	"Access-Control-Allow-Origin":  []string{"*"},
	"Access-Control-Allow-Methods": []string{"POST, GET, PUT, PATCH, DELETE, OPTIONS"},
	"Access-Control-Allow-Headers": []string{"Origin, X-Requested-With, Content-Type, Accept, API-Date, Host, Authorization, Key, User-Token"},
	"Access-Control-Max-Age":       []string{"1728000"},
}

type Filler func(r *http.Request) []byte

type Endpoint struct {
	path    string
	reactor Filler
}

func (this Endpoint) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h := w.Header()
}

func SetItUp() {
	// searchpage := Endpoint{
	// 	path:    "/searchpage/",
	// 	reactor: handler_searchpage,
	// }

	// autocomplete := Endpoint{
	// 	path:    "/autocomplete/",
	// 	reactor: handler_autocomplete,
	// }

	//    return Server{
	//        headers
	//    }
}
