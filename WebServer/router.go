package main

import (
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
	handler, exists := r.FindHandler(request.URL.Path)

	if !exists {
		writter.WriteHeader(http.StatusNotFound)
		return
	}

	handler(writter, request)
}
