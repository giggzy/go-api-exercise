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

func returnServices(w http.ResponseWriter, r *http.Request) {
	log.Println("returnServices requested")
	w.Header().Set("Content-Type", "application/json")

	// Optional query string parameters
	search := r.URL.Query().Get("search")
	page := r.URL.Query().Get("page")
	pageSize := r.URL.Query().Get("pageSize")

	var foundServices []Service
	if search != "" {
		log.Println("returnServices request for search: ", search)
		for _, service := range services.Services {
			if strings.Contains(service.Name, search) {
				foundServices = append(foundServices, service)
			}
		}
	} else {
		log.Println("returnServices request for all services")
		foundServices = services.Services
	}

	// pagination
	// setting some defaults if the query string parameters are not set
	// assuming better approach to handle default values, TODO: learn more about this
	if page == "" {
		page = "0"
	}
	if pageSize == "" {
		pageSize = "10"
	}

	// convert the page and pageSize to ints
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		log.Println("Error converting page to int: ", err)
	}
	pageSizeInt, err := strconv.Atoi(pageSize)
	if err != nil {
		log.Println("Error converting pageSize to int: ", err)
	}

	// get the index of the first and last element to return
	start := pageInt * pageSizeInt
	end := start + pageSizeInt
	if end > len(foundServices) {
		end = len(foundServices)
	}

	// Need some bounds checking here
	if start >= len(foundServices) {
		// invalid page
		log.Println("Invalid page requested: ", page)
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "Invalid page requested: ` + page + `"}`))
		return
	}

	foundServices = foundServices[start:end]
	// Should I add page and pageSize to the response?
	//meta := `json: {"page": ` + page + `, "pageSize": ` + pageSize + `}`
	//meta := Meta{Page: pageInt, PageSize: pageSizeInt}
	//ID           string `json:"id"`

	// Intention here was to return some meta data about the response
	// e.g. total number of services, page number, page size, etc.
	//foundServices = append(foundServices, meta)

	// write the response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(foundServices)

}

// This uses the http.Request object to get the URL path
// expects /service_detail/{id} as the path
func returnServiceDetails(w http.ResponseWriter, r *http.Request) {
	log.Println("returnServiceDetails requested")
	w.Header().Set("Content-Type", "application/json")

	id := strings.TrimPrefix(r.URL.Path, "/service_detail/")
	log.Println("returnServiceDetails request for id2: ", id)

	var foundService Service
	for _, service := range services.Services {
		if service.ID == id {
			foundService = service
		}
	}
	if foundService != (Service{}) {
		json.NewEncoder(w).Encode(foundService)
		w.WriteHeader(http.StatusOK)
	} else {
		// invalid id
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "not found"}`))
	}
}

/*

// Handler that expects id to be a query parameter
// e.g. /service_detail?id=1

func returnServiceDetails(w http.ResponseWriter, r *http.Request) {
	log.Println("returnServiceDetails requested")
	w.Header().Set("Content-Type", "application/json")

	id := r.URL.Query().Get("id")
	log.Println("returnServiceDetails request for id: ", id)

	var foundService Service
	for _, service := range services.Services {
		if service.ID == id {
			foundService = service
		}
	}
	if foundService != (Service{}) {
		json.NewEncoder(w).Encode(foundService)
		w.WriteHeader(http.StatusOK)
	} else {
		// invalid id
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "not found"}`))
	}
}
*/

func handleRequests() {
	http.HandleFunc("/services", returnServices)
	http.HandleFunc("/service_detail/", returnServiceDetails)
}

///////////////////////////////////////////////////////////////////////////////
// main, entry point
///////////////////////////////////////////////////////////////////////////////

func main() {
	// read the file and unmarshal the json into the services struct
	content, err := readDBFile("./data/sample.json")
	if err != nil {
		log.Println(err)
	}
	//var payload Services
	err = json.Unmarshal([]byte(content), &services)
	if err != nil {
		log.Fatal("Error unmarshalling services JSON: ", err)
	}
	log.Println("Read in services: ", len(services.Services))

	// start the server
	handleRequests()
	log.Println("Start Web Server...")
	http.ListenAndServe(":8080", nil)
}

///////////////////////////////////////////////////////////////////////////////
// utility functions
///////////////////////////////////////////////////////////////////////////////

// read a file and return the content
func readDBFile(fileName string) (string, error) {
	// read file  fd
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		return "", err
	}
	return string(content), nil
}
