package urlShort

import (
	"net/http"
        "fmt"
        "gopkg.in/yaml.v2"
)

func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        if val, ok := pathsToUrls[r.URL.Path]; ok {
            http.Redirect(w, r, val, 301)
        }
        fallback.ServeHTTP(w, r)
    }
}

type yamlStruct struct {
        PATH string `yaml:"path"`
        URL  string `yaml:"url"`
}

func bytesToString(data []byte) string {
        return string(data[:])
}


// Shortest way to unmarshalling multiple yaml lines
func parseYAML(yml []byte) ([]yamlStruct, error) {
    var yamlArray []yamlStruct
    err := yaml.Unmarshal(yml, &yamlArray)
    if err != nil {
        return nil, err
    }
    return yamlArray, nil
}


func buildMap(yamlArray []yamlStruct) map[string]string {
        result := make(map[string]string)
        for _, yamlStruct := range yamlArray {
                result[yamlStruct.PATH] = yamlStruct.URL
        }
        return result
}

func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
  parsedYaml, err := parseYAML(yml)
  if err != nil {
    return nil, err
  }
  pathMap := buildMap(parsedYaml)
  fmt.Println(pathMap)
  return MapHandler(pathMap, fallback), nil
}
