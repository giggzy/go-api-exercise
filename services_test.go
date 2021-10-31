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
	handler := http.HandlerFunc(returnServices)
	handler.ServeHTTP(recorder, req)
	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Confirm the response has the expected data?
	contains := `{ "id": "ccb5fa2c-d6da-4d68-a901-3d81682e3a2c", "name": "o_Service", "description": "o blah, blah, blah",`
	if !strings.Contains(recorder.Body.String(), contains) {
		t.Errorf("Response does not contains expected content: %v is missing %v",
			recorder.Body.String(), contains)
	}
}

// Acceptance test for /services/{id}
func Test_services_id(t *testing.T) {
	req, err := http.NewRequest("GET", "/services/ccb5fa2c-d6da-4d68-a901-3d81682e3a2c", nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(returnServiceDetails)
	handler.ServeHTTP(recorder, req)
	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Confirm the response has the expected data?
	const contains = `{ "id": "ccb5fa2c-d6da-4d68-a901-3d81682e3a2c", "name": "o_Service", "description": "o blah, blah, blah",`
	if !strings.Contains(recorder.Body.String(), contains) {
		t.Errorf("Response does not contains expected content: %v is missing %v",
			recorder.Body.String(), contains)
	}
}
