package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHealthCheck(t *testing.T) {
	request, _ := http.NewRequest("GET", "/healthz", nil)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HealthCheckHandler)

	handler.ServeHTTP(rr, request)

	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestUpperCaseString(t *testing.T) {
	request, _ := http.NewRequest("GET", "/should-be-upper/upper", nil)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(StringEchoHandler)

	handler.ServeHTTP(rr, request)

	assert.Equal(t, rr.Code, http.StatusOK)

	expected := "Echo: \"SHOULD-BE-UPPER\""
	assert.Equal(t, expected, rr.Body.String())
}

func TestDefaultStringEcho(t *testing.T) {
	request, _ := http.NewRequest("GET", "/shouldJustBeEcho", nil)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(StringEchoHandler)

	handler.ServeHTTP(rr, request)

	assert.Equal(t, rr.Code, http.StatusOK)

	expected := "Echo: \"shouldJustBeEcho\""
	assert.Equal(t, expected, rr.Body.String())
}

func TestInvalidStringEchoCommand(t *testing.T) {
	request, _ := http.NewRequest("GET", "/should-be-upper/invalid", nil)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(StringEchoHandler)

	handler.ServeHTTP(rr, request)

	assert.Equal(t, rr.Code, http.StatusInternalServerError)
}
