package main

import (
	"database/sql"
	"flag"
	"fmt"
	"net/http"
	"os"

	_ "github.com/libsql/libsql-client-go/libsql"
	"github.com/pserghei/gophercises/ex02/urlshort"
	_ "modernc.org/sqlite"
)

var dbUrl = "file:urls.db"

func main() {
	db, err := sql.Open("libsql", dbUrl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to open db %s: %s", dbUrl, err)
		os.Exit(1)
	}

	f := flag.String("f", "", "The yaml file to be opened")
	flag.Parse()

	mux := defaultMux()

	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	yaml := make([]byte, 0)
	if *f != "" {
		var err error
		yaml, err = os.ReadFile(*f)
		if err != nil {
			panic(err)
		}
	} else {
		yaml = []byte(`
- path: /urlshort
  url: https://github.com/gophercises/urlshort
- path: /urlshort-final
  url: https://github.com/gophercises/urlshort/tree/solution
`)
	}

	yamlHandler, err := urlshort.YAMLHandler(yaml, mapHandler)
	if err != nil {
		panic(err)
	}

	json := []byte(`
    [
    {
        "path": "/example",
        "url": "https://example.com"
    },
    {
        "path": "/gophercises",
        "url": "https://gophercises.com"
    }
    ]
    `)

	jsonHandler, err := urlshort.JSONHandler(json, mapHandler)
	if err != nil {
		panic(err)
	}

	dbHandler, err := urlshort.DBHandler(db, mapHandler)
	if err != nil {
		panic(err)
	}

    fmt.Println("Starting servers on :8000, :8080 and :8888")
	go http.ListenAndServe(":8000", jsonHandler)
    go http.ListenAndServe(":8888", dbHandler)
	http.ListenAndServe(":8080", yamlHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
