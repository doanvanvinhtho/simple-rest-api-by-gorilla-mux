package main

import (
	"log"
	"net/http"

	"github.com/doanvanvinhtho/simple-rest-api-by-gorilla-mux/handle"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", handle.HomeLink)

	router.HandleFunc("/events/{id}", handle.GetOneEvent).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))
}
