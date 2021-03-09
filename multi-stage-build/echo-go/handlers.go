package main

import (
	"fmt"
	"html"
	"net/http"
	"strings"
)

// Perform string operations based on the `op` parameter
func convertString(str string, op string) (string, error) {
	result := ""

	switch {
	case op == "None":
		result = str
	case op == "lower":
		result = strings.ToLower(str)
	case op == "upper":
		result = strings.ToUpper(str)
	case op == "title":
		result = strings.Title(strings.ToLower(str))
	default:
		return "", fmt.Errorf("Invalid operator: %s", op)
	}

	return result, nil
}

// Safe way to read a value in an array and return a default value if
// the array does not have enough elements
func readSplit(splits []string, index int, defString string) string {
	if len(splits) > index {
		return splits[index]
	}
	return defString
}

// StringEchoHandler will read the url path and echo a string back in response.
// It can perform some string transformations before returning the string back to you:
// - /<text>/upper => string to upper case
// - /<text>/lower => string to lower case
// - /<text>/title => title capitalisation for the given string
func StringEchoHandler(w http.ResponseWriter, r *http.Request) {
	splits := strings.Split(r.URL.Path, "/")

	str := readSplit(splits, 1, "")
	op := readSplit(splits, 2, "None")

	converted, err := convertString(str, op)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		fmt.Fprintf(w, "Echo: %q", html.EscapeString(converted))
	}
}

// HealthCheckHandler returns the service status.
func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Service is healthy :)")
}
