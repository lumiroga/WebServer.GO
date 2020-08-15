package main

import (
	"fmt"
	"net/http"
)

//Router - struct that contains http routes
type Router struct {
	rules map[string]http.HandlerFunc
}

//NewRouter - ceates a new Router instance
func NewRouter() *Router {
	return &Router{
		rules: make(map[string]http.HandlerFunc),
	}
}

func (r *Router) ServeHTTP(writter http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writter, "Hola mundo!")
}
