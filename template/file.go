package template

import (
	"encoding/json"
	"path/filepath"

	"github.com/flosch/pongo2"
	"github.com/ghodss/yaml"
)

// ReadJSONFromFile read a JSON format file or a YAML format
// file then convert to JSON
func ReadJSONFromFile(fp string, vals M) (b []byte, err error) {
	tpl, err := pongo2.FromFile(fp)
	if err != nil {
		return
	}
	if b, err = tpl.ExecuteBytes(pongo2.Context(vals)); err != nil {
		return
	}
	switch ext := filepath.Ext(fp); ext {
	case ".yml", ".yaml":
		return yaml.YAMLToJSON(b)
	default:
		return
	}
}

func ReadYAMLFromFile(fp string, vals M) (b []byte, err error) {
	tpl, err := pongo2.FromFile(fp)
	if err != nil {
		return
	}
	if b, err = tpl.ExecuteBytes(pongo2.Context(vals)); err != nil {
		return
	}
	switch ext := filepath.Ext(fp); ext {
	case ".json":
		return yaml.JSONToYAML(b)
	default:
		return
	}
}

// Wrap ReadJSONFromFile, unmarhsal JSON to object automatically
func FromFile(fp string, vals M) (_ M, err error) {
	b, err := ReadJSONFromFile(fp, vals)
	if err != nil {
		return
	}
	doc := M{}
	if err = json.Unmarshal(b, &doc); err != nil {
		return
	}
	return doc, nil
}

func UnmarshalFile(fp string, vals M, cfg interface{}) (err error) {
	b, err := ReadJSONFromFile(fp, vals)
	if err != nil {
		return
	}
	if err = json.Unmarshal(b, cfg); err != nil {
		return
	}
	return
}
