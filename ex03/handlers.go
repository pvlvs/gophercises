package main

import (
	"fmt"
	"net/http"
	"strings"
	"text/template"
)

func routeHandler(w http.ResponseWriter, r *http.Request) {
	sm := storyMap()
	path := strings.TrimPrefix(r.URL.Path, "/")

	if path == "" {
		http.Redirect(w, r, "/intro", http.StatusPermanentRedirect)
		return
	}

	if s, ok := sm[path]; ok {
		t, err := template.ParseFiles("story.tmpl.html")
        check(err)

		t.Execute(w, s)
		return
	}

	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "Not a valid arc")
}
