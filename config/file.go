package config

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"

	"github.com/ghodss/yaml"
)

// ReadJSONFromFile read a JSON format file or a YAML format
// file then convert to JSON
func ReadJSONFromFile(fp string) (b []byte, err error) {
	if b, err = ioutil.ReadFile(fp); err != nil {
		return
	}
	switch ext := filepath.Ext(fp); ext {
	case ".yml", ".yaml":
		return yaml.YAMLToJSON(b)
	default:
		return
	}
}

// Wrap ReadJSONFromFile, unmarhsal JSON to object automatically
func FromFile(fp string, cfg interface{}) (err error) {
	b, err := ReadJSONFromFile(fp)
	if err != nil {
		return
	}
	if err = json.Unmarshal(b, cfg); err != nil {
		return
	}
	return
}
