package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strconv"
)

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

// get random id between 1 and 1000000
// toy case
func getRandomID() string {
	id := rand.Intn(1000000)
	// convert to string
	return strconv.Itoa(id)
}
