package filesystem

import (
	"encoding/json"

	"gopkg.in/yaml.v3"
)

const (
	ContentOperationAppend ContentOperation = iota
	ContentOperationOverwrite
)

var validContentOperations = map[ContentOperation]bool{
	ContentOperationAppend:    true,
	ContentOperationOverwrite: true,
}

type (
	ContentOperation uint8

	// Content represents []byte output from file with additional helper functions.
	Content []byte
)

func (o ContentOperation) Is(val ContentOperation) bool {
	return o == val
}

func (o ContentOperation) assetValid() error {
	if validContentOperations[o] {
		return nil
	}
	return ErrUnsupportedContentOperation
}

// JSONDecode is an equivalent of using `json.Unmarshal` on []byte containing JSON.
func (c Content) JSONDecode(dto any) error {
	return json.Unmarshal(c, dto)
}

// YAMLDecode is an equivalent of using `yaml.Unmarshal` on []byte containing YAML.
func (c Content) YAMLDecode(dto any) error {
	return yaml.Unmarshal(c, dto)
}

func (c Content) Length() int {
	return len(c)
}

func (c Content) String() string {
	return string(c)
}
