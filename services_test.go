package main

import (
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
	contains := `{"id":"1","name":"a_Service","description":"a blah, blah, blah","versionCount":1,"url":"https://example.com/a_Service"}`
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
	const contains = `{"id":"2","name":"b_Service","description":"b blah, blah, blah","versionCount":2,"url":"https://example.com/b_Service"}`
	if !strings.Contains(recorder.Body.String(), contains) {
		t.Errorf("Response does not contains expected content: %v is missing %v",
			recorder.Body.String(), contains)
	}
}
