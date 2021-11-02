package main

import (
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

var services Services

// Did not end up using this but towards authN and authZ
// NOT USED!!
func commonAPIWrapper(call func(http.ResponseWriter, *http.Request)) {
	//return func(w http.ResponseWriter, r *http.Request) {
	// TODO:
	// check if the services are loaded
	// do autherization and authentication
	services = loadServices()
	//call(w, r)
	//}
}

///////////////////////////////////////////////////////////////////////////////
// main, entry point
///////////////////////////////////////////////////////////////////////////////

func main() {
	services = loadServices() // loadServices() handles error checking
	rand.Seed(time.Now().UnixNano())

	r := mux.NewRouter()
	r.HandleFunc("/services", getServices).Methods("GET")
	r.HandleFunc("/service_detail/{id}", getServiceDetails).Methods("GET")
	r.HandleFunc("/service_create", createService).Methods("POST")
	http.Handle("/", r)
	log.Println("Start Web Server...")
	http.ListenAndServe(":8080", nil)
}
