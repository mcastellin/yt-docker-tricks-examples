package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func listen() {
	host := os.Getenv("ECHO_HOST")
	port := os.Getenv("ECHO_PORT")

	if len(port) == 0 {
		port = "8081"
	}

	log.Printf("Starting server at %s:%s\n", host, port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", host, port), nil))
}

func main() {

	http.HandleFunc("/", StringEchoHandler)

	http.HandleFunc("/healthz", HealthCheckHandler)

	listen()
}
