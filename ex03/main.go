package main

import (
	"log"
	"net/http"
)

func main() {
	mux := routes()

	log.Println("Starting server on :8080")
	http.ListenAndServe(":8080", mux)
}
