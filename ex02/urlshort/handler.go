package urlshort

import (
	"net/http"

	"gopkg.in/yaml.v2"
)

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

func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	y := []struct {
		Path string `yaml:"path"`
		Url  string `yaml:"url"`
	}{}

	err := yaml.Unmarshal(yml, &y)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)

	for _, v := range y {
		m[v.Path] = v.Url
	}

	return MapHandler(m, fallback), nil
}
