package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type Services struct {
	Services []Service `json:"services"`
}
type Service struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	VersionCount int    `json:"versionCount"`
	URL          string `json:"url"`
}

// not used yet
type Meta struct {
	Page     int
	PageSize int
	//total int
}

var services Services

func commonAPIWrapper(call func(http.ResponseWriter, *http.Request)) {
	//return func(w http.ResponseWriter, r *http.Request) {
	// TODO:
	// check if the services are loaded
	// do autherization and authentication
	services = loadServices()
	//call(w, r)
	//}
}

// Handle routing of requests
func handleRequests() {
	//http.HandleFunc("/services", commonAPIWrapper(getServices))
	http.HandleFunc("/services", getServices)
	http.HandleFunc("/service_detail/", getServiceDetails)
	http.HandleFunc("/service_create/", createService)
}

///////////////////////////////////////////////////////////////////////////////
// main, entry point
///////////////////////////////////////////////////////////////////////////////

func mainX() {
	services = loadServices() // loadServices() handles error checking

	// start the server
	handleRequests()
	log.Println("Start Web Server...")
	http.ListenAndServe(":8080", nil)
}

func main() {
	services = loadServices() // loadServices() handles error checking

	r := mux.NewRouter()
	r.HandleFunc("/services", getServices).Methods("GET")
	r.HandleFunc("/service_detail/{id}", getServiceDetails).Methods("GET")
	r.HandleFunc("/service_create", createService).Methods("POST")
	http.Handle("/", r)
	log.Println("Start Web Server...")
	http.ListenAndServe(":8080", r)
}

///////////////////////////////////////////////////////////////////////////////
// utility functions
///////////////////////////////////////////////////////////////////////////////

// Load services from a file
func loadServices() Services {
	if services.Services != nil {
		// already loaded no need to do it again
		return services
	}
	// read the file and unmarshal the json into the services struct
	var services Services
	content := readDBFile()
	err := json.Unmarshal([]byte(content), &services)
	if err != nil {
		log.Fatal("Error unmarshalling services JSON: ", err)
	}
	log.Println("Read in services, count: ", len(services.Services))
	return services
}

// read a file and return the content
func readDBFile() string {
	// prefer to read in services.json but if not found use the mock data
	sampleData := "./data/sample.json"
	servicesData := "./data/services.json"

	var content []byte
	// Check if file exists
	if _, err := os.Stat(servicesData); err == nil {
		log.Println("Using services data")
		content, err = ioutil.ReadFile(servicesData)
		if err != nil {
			log.Fatal("Error reading file: ", err)
		}
	} else {
		// Fallback to sample data
		log.Println("Using sample data")
		content, err = ioutil.ReadFile(sampleData)
		if err != nil {
			log.Fatal("Error reading file: ", err)
		}
	}
	return string(content)
}

func writeDBFile() {
	// write the content to the file
	jsonString, _ := json.Marshal(services)
	err := ioutil.WriteFile("./data/services.json", jsonString, 0644)
	if err != nil {
		log.Fatal("Error writing file: ", err)
	}
}
