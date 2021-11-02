package main

import (
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func Test_services(t *testing.T) {
	req, err := http.NewRequest("GET", "/services", nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(getServices)
	handler.ServeHTTP(recorder, req)
	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Confirm the response contains an expected substring?
	contains := `{"id":"1","name":"a_Service","description":"a blah, blah, blah","versionCount":1,"url":"https://example.com/a_Service","versions":[{"id":"0","name":"version_0"}]}`
	//contains := `{"id":"1","name":"a_Service","description":"a blah, blah, blah","versionCount":1,"url":"https://example.com/a_Service"}`
	if !strings.Contains(recorder.Body.String(), contains) {
		t.Errorf("Response does not contains expected content: %v is missing %v",
			recorder.Body.String(), contains)
	}
}

// Acceptance test for /services/{id}
// Note its checking the service has a URL making navigation to the service possible as requested by Product Owner
func Test_services_id(t *testing.T) {
	req, err := http.NewRequest("GET", "/service_detail/2", nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(getServiceDetails)
	handler.ServeHTTP(recorder, req)
	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Confirm the response has the expected data?
	//const contains = `{ "id": "ccb5fa2c-d6da-4d68-a901-3d81682e3a2c", "name": "o_Service", "description": "o blah, blah, blah",`
	const contains = `{"id":"2","name":"b_Service","description":"b blah, blah, blah","versionCount":2,"url":"https://example.com/b_Service","versions":[{"id":"0","name":"version_0"},{"id":"1","name":"version_1"}]}`
	if !strings.Contains(recorder.Body.String(), contains) {
		t.Errorf("Response does not contains expected content: %v is missing %v",
			recorder.Body.String(), contains)
	}
}

func Test_service_create(t *testing.T) {

	body := strings.NewReader(`{"name":"Test Create","description":"create","id":"234","versionCount":33,"url":"https://example.com/created_service"}`)
	req, err := http.NewRequest("POST", "/services", body)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(createService)
	handler.ServeHTTP(recorder, req)
	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

// Unit tests
func Test_id_creation(t *testing.T) {
	id := getRandomID()
	if id == "" {
		t.Errorf("ID is empty")
	}
	log.Printf("New ID: `%s` is the new ID.\n", id)
}
