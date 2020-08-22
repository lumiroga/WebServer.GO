package main

import (
	"net/http"
)

//Router - struct that contains http routes
type Router struct {
	rules map[string]map[string]http.HandlerFunc
}

//NewRouter - ceates a new Router instance
func NewRouter() *Router {
	return &Router{
		rules: make(map[string]map[string]http.HandlerFunc),
	}
}

//FindHandler finds a path a returns the handler function
func (r *Router) FindHandler(path string, method string) (http.HandlerFunc, bool, bool) {
	_, exist := r.rules[path]

	handler, methodExist := r.rules[path][method]

	return handler, exist, methodExist
}

func (r *Router) ServeHTTP(writter http.ResponseWriter, request *http.Request) {
	handler, exists, methodExist := r.FindHandler(request.URL.Path, request.Method)

	if !exists {
		writter.WriteHeader(http.StatusNotFound)
		return
	}
	if !methodExist {
		writter.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	handler(writter, request)
}
