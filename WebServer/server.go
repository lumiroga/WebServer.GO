package main

import "net/http"

//Server struct
type Server struct {
	port   string
	router *Router
}

//NewServer creates a new Server instance
func NewServer(port string) *Server {
	return &Server{
		port:   port,
		router: NewRouter(),
	}
}

//Handle maps a path to correct handler
func (serv *Server) Handle(path string, handler http.HandlerFunc) {
	serv.router.rules[path] = handler
}

//AddMiddleware - adds middlewares in the path
func (serv *Server) AddMiddleware(f http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, m := range middlewares {
		f = m(f)
	}

	return f
}

//Listen is a function that watis for requests
func (serv *Server) Listen() error {
	http.Handle("/", serv.router)
	err := http.ListenAndServe(serv.port, nil)

	if err != nil {
		return err
	}

	return nil

}
