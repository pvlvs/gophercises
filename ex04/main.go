package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/pserghei/gophercises/ex04/parser"
)

func main() {
	f, err := os.ReadFile("./html/2.html")
	if err != nil {
		log.Fatal(err)
	}

	r := strings.NewReader(string(f))

    links, err := parser.Parse(r)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(links)
}
