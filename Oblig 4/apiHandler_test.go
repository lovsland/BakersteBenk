package main

import (
	"testing"
	"net/http/httptest"
	"net/http"
)

func TestApi1Handler(t *testing.T) {
	req, err := http.NewRequest("GET", "/fotball", nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(api1handler)

	handler.ServeHTTP(recorder, req)

	status := recorder.Code
	if  status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestApi2Handler(t *testing.T) {
	req, err := http.NewRequest("GET", "/ManUnited", nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(api2handler)

	handler.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	//if recorder.Body.String() != expected {
	//	t.Errorf("handler returned unexpected body: got %v want %v",
	//		recorder.Body.String(), expected)
	//}
}

func TestApi3Handler(t *testing.T) {
	req, err := http.NewRequest("GET", "/ManUnitedSquad", nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(api3handler)

	handler.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}