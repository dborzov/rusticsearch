package main

import (
	"net/http"
)

type Server struct {
	headers   map[string]string
	endpoints []Endpoint
}

type Endpoint struct {
	path string
}

func (this Endpoint) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}
