package main

import (
	"net/http"
)

type Server struct {
	headers   map[string]string
	endpoints []Endpoint
}

type Filler func(r *http.Request) []byte

type Endpoint struct {
	path    string
	reactor Filler
}

func (this Endpoint) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}

func SetItUp() {
	searchpage := Endpoint{
		path:    "/searchpage/",
		reactor: handler_searchpage,
	}

	autocomplete := Endpoint{
		path:    "/autocomplete/",
		reactor: handler_autocomplete,
	}
}
