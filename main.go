package main

import (
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gorilla/handlers"
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

	handler := handlers.LoggingHandler(log.Writer(), handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
	)(r))
	r.HandleFunc("/services", getServices).Methods("GET", "OPTIONS")
	r.HandleFunc("/service_detail/{id}", getServiceDetails).Methods("GET", "OPTIONS")
	r.HandleFunc("/service_create", createService).Methods("POST", "OPTIONS")
	// http.Handle("/", r)

	r.Use(mux.CORSMethodMiddleware(r))

	newServer := &http.Server{
		Handler: handler,
		Addr:    ":8080",
	}
	log.Println("Start Web Server...")
	/*
		http.ListenAndServe(":8080", r)
	*/
	log.Fatal(newServer.ListenAndServe())
}
