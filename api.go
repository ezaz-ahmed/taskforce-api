package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type APIServer struct {
	addr  string
	store Store
}

func newAPIServer(addr string, store Store) *APIServer {
	return &APIServer{addr, store}
}

func (s *APIServer) Serve() {
	router := mux.NewRouter()
	subRouter := router.PathPrefix("/api/v1").SubRouter()

	log.Println("Starting the Api server at", s.addr)

	log.Fatal(http.ListenAndServe(s.addr, subRouter))
}
