package main

import (
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// GLOBAL
var services Services

///////////////////////////////////////////////////////////////////////////////
// main, entry point
///////////////////////////////////////////////////////////////////////////////

func main() {
	loadServices()                   // loadServices() handles error checking
	rand.Seed(time.Now().UnixNano()) // used to generate random ids for new services

	r := mux.NewRouter()
	r.HandleFunc("/services", getServices).Methods("GET")
	r.HandleFunc("/service_detail/{id}", getServiceDetails).Methods("GET")
	r.HandleFunc("/service_create", createService).Methods("POST")
	http.Handle("/", r)
	log.Println("Start Web Server...")
	http.ListenAndServe(":8080", nil)
}
