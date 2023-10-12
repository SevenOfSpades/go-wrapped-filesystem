package filesystem

import (
	"encoding/json"

	"gopkg.in/yaml.v3"
)

// Content represents []byte output from file with additional helper functions.
type Content []byte

// JSONDecode is an equivalent of using `json.Unmarshal` on []byte containing JSON.
func (c Content) JSONDecode(dto any) error {
	return json.Unmarshal(c, &dto)
}

// YAMLDecode is an equivalent of using `yaml.Unmarshal` on []byte containing YAML.
func (c Content) YAMLDecode(dto any) error {
	return yaml.Unmarshal(c, &dto)
}

func (c Content) Length() int {
	return len(c)
}

func (c Content) String() string {
	return string(c)
}
