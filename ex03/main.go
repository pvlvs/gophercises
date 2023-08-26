package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	cmdLine := flag.Bool("c", false, "Start the adventure in the terminal instead")
	flag.Parse()

	if *cmdLine {
		startTerminal()
	} else {
		mux := routes()

		log.Println("Starting server on :8080")
		http.ListenAndServe(":8080", mux)
	}
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

