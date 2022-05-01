package resolver

import (
	"encoding/json"

	"gopkg.in/yaml.v2"
)

// unmarshaler provides a self labeling unmarshaler
type unmarshaler interface {
	json.Marshaler
	yaml.Marshaler
	Unmarshal([]byte, interface{}) error
}

type jsonUnmarshal struct {
}

func (ju jsonUnmarshal) Unmarshal(raw []byte, target interface{}) error {
	return json.Unmarshal(raw, target)
}

func (ju jsonUnmarshal) MarshalJSON() ([]byte, error) {
	return []byte(`"JSON"`), nil
}

func (ju jsonUnmarshal) MarshalYAML() (interface{}, error) {
	return "JSON", nil
}

type yamlUnmarshal struct {
}

func (yu yamlUnmarshal) Unmarshal(raw []byte, target interface{}) error {
	return yaml.Unmarshal(raw, target)
}

func (yu yamlUnmarshal) MarshalJSON() ([]byte, error) {
	return []byte(`"YAML"`), nil
}

func (yu yamlUnmarshal) MarshalYAML() (interface{}, error) {
	return "YAML", nil
}
