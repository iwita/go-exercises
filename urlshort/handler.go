package urlshort

import (
	"net/http"

	"github.com/ghodss/yaml"
)

func MapHandler(pathToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if dest, ok := pathToUrls[path]; ok {
			http.Redirect(w, r, dest, http.StatusFound)
			return
		}
		fallback.ServeHTTP(w, r)
	}
}

func yamlHandler(yamlBytes []byte, fallback http.Handler) (http.HandlerFunc, error) {

	// 1. Parse the yaml
	pathUrls, err := parseYAML(yamlBytes)
	if err != nil {
		return nil, err
	}

	// 2. Convert YAML array into map
	pathToUrls := buildMap(pathUrls)

	// 3. Return a mapHander
	return MapHandler(pathToUrls, fallback), nil
}

func parseYAML(data []byte) ([]pathUrl, error) {
	var pathUrls []pathUrl
	err := yaml.Unmarshal(data, &pathUrls)
	if err != nil {
		return nil, err
	}
	return pathUrls, nil
}
func buildMap(pathUrls []pathUrl) map[string]string {
	pathToUrls := make(map[string]string)
	for _, pu := range pathUrls {
		pathToUrls[pu.Path] = pu.URL
	}
	return pathToUrls
}

type pathUrl struct {
	Path string `yaml:"path,omitempty"`
	URL  string `yaml:"url,omitempty"`
}
