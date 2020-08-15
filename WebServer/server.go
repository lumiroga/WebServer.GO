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

//Listen is a function that watis for requests
func (serv *Server) Listen() error {
	http.Handle("/", serv.router)
	err := http.ListenAndServe(serv.port, nil)

	if err != nil {
		return err
	}

	return nil

}
