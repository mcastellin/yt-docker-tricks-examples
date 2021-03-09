package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthCheck(t *testing.T) {
	request, _ := http.NewRequest("GET", "/healthz", nil)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HealthCheckHandler)

	handler.ServeHTTP(rr, request)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestUpperCaseString(t *testing.T) {
	request, _ := http.NewRequest("GET", "/should-be-upper/upper", nil)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(StringEchoHandler)

	handler.ServeHTTP(rr, request)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("StringEchoHandler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := "Echo: \"SHOULD-BE-UPPER\""
	if body := rr.Body.String(); body != expected {
		t.Errorf("StringEchoHandler returned wrong body: got %v want %v",
			body, expected)
	}
}

func TestDefaultStringEcho(t *testing.T) {
	request, _ := http.NewRequest("GET", "/shouldJustBeEcho", nil)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(StringEchoHandler)

	handler.ServeHTTP(rr, request)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("StringEchoHandler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := "Echo: \"shouldJustBeEcho\""
	if body := rr.Body.String(); body != expected {
		t.Errorf("StringEchoHandler returned wrong body: got %v want %v",
			body, expected)
	}
}

func TestInvalidStringEchoCommand(t *testing.T) {
	request, _ := http.NewRequest("GET", "/should-be-upper/invalid", nil)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(StringEchoHandler)

	handler.ServeHTTP(rr, request)

	if status := rr.Code; status != http.StatusInternalServerError {
		t.Errorf("StringEchoHandler returned wrong status code: got %v want %v",
			status, http.StatusInternalServerError)
	}
}
