package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
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
	//http.HandleFunc("/service_cre")
}

///////////////////////////////////////////////////////////////////////////////
// main, entry point
///////////////////////////////////////////////////////////////////////////////

func main() {
	services = loadServices() // loadServices() handles error checking

	// start the server
	handleRequests()
	log.Println("Start Web Server...")
	http.ListenAndServe(":8080", nil)
}

///////////////////////////////////////////////////////////////////////////////
// utility functions
///////////////////////////////////////////////////////////////////////////////

// Load services from a file
func loadServices() Services {
	if services.Services != nil {
		// services = loadServices()
		return services
	}
	// read the file and unmarshal the json into the services struct
	var services Services
	content, err := readDBFile("./data/sample.json")
	if err != nil {
		log.Fatal("Error reading file: ", err)
		return services
	}
	err = json.Unmarshal([]byte(content), &services)
	if err != nil {
		log.Fatal("Error unmarshalling services JSON: ", err)
		return services
	}
	log.Println("Read in services, count: ", len(services.Services))
	return services
}

// read a file and return the content
func readDBFile(fileName string) (string, error) {
	// read file  fd
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		return "", err
	}
	return string(content), nil
}
