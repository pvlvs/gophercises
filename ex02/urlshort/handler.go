package urlshort

import (
	"database/sql"
	json "encoding/json"
	"net/http"

	"gopkg.in/yaml.v2"
)

type urlS struct {
	Path string `yaml:"path" json:"path"`
	Url  string `yaml:"url" json:"url"`
}

func makeMap(s []urlS) map[string]string {
	m := make(map[string]string)
	for _, v := range s {
		m[v.Path] = v.Url
	}

	return m
}

func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := pathsToUrls[r.URL.Path]

		if path != "" {
			http.Redirect(w, r, path, http.StatusSeeOther)
		} else {
			fallback.ServeHTTP(w, r)
		}
	}
}

func makeYAML(yml []byte) []urlS {
	us := []urlS{}
	err := yaml.Unmarshal(yml, &us)
	if err != nil {
		panic(err)
	}

	return us
}

func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	us := makeYAML(yml)

	return MapHandler(makeMap(us), fallback), nil
}

func makeJSON(j []byte) []urlS {
	us := []urlS{}
	err := json.Unmarshal(j, &us)
	if err != nil {
		panic(err)
	}

	return us
}

func JSONHandler(j []byte, fallback http.Handler) (http.HandlerFunc, error) {
	us := makeJSON(j)

	return MapHandler(makeMap(us), fallback), nil
}

func readRows(rows *sql.Rows) []urlS {
	us := []urlS{}

	for rows.Next() {
		url := urlS{}

		err := rows.Scan(&url.Path, &url.Url)
		if err != nil {
			panic(err)
		}

		us = append(us, url)
	}

	return us
}

func DBHandler(db *sql.DB, fallback http.Handler) (http.HandlerFunc, error) {
	stmt := "select * from urls"

	rows, err := db.Query(stmt)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	us := readRows(rows)

	return MapHandler(makeMap(us), fallback), nil
}
