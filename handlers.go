package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func getServices(w http.ResponseWriter, r *http.Request) {
	log.Println("getServives requested")
	w.Header().Set("Content-Type", "application/json")

	// When running as a test services are not yet initialized
	services = loadServices()

	// Optional query string parameters
	search := r.URL.Query().Get("search")
	page := r.URL.Query().Get("page")
	pageSize := r.URL.Query().Get("pageSize")

	var foundServices []Service
	if search != "" {
		log.Println("getServives request for search: ", search)
		for _, service := range services.Services {
			if strings.Contains(service.Name, search) {
				foundServices = append(foundServices, service)
			}
		}
	} else {
		log.Println("getServives request for all services")
		foundServices = services.Services
	}

	// pagination
	// setting some defaults if the query string parameters are not set
	// assuming better approach to handle default values, TODO: learn more about this
	if page == "" {
		page = "0"
	}
	if pageSize == "" {
		pageSize = "12" // the design mock up has 12 items per page, 4 * 3
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
	if start >= len(foundServices) || start < 0 {
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
func getServiceDetails(w http.ResponseWriter, r *http.Request) {
	log.Println("getServiceDetails requested")
	w.Header().Set("Content-Type", "application/json")

	// When running as a test services are not yet initialized
	services = loadServices()

	id := strings.TrimPrefix(r.URL.Path, "/service_detail/")
	log.Println("getServiceDetails request for id2: ", id)

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