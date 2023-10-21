package filesystem

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testDecodeResult struct {
	Value bool `yaml:"yamlValue" json:"jsonValue"`
}

func TestContent_JSONDecode(t *testing.T) {
	t.Parallel()

	// GIVEN
	testYaml := Content(`{"jsonValue": true}`)

	// WHEN
	var dto testDecodeResult
	err := testYaml.JSONDecode(&dto)

	// THEN
	assert.NoError(t, err)
	assert.True(t, dto.Value)
}

func TestContent_YAMLDecode(t *testing.T) {
	t.Parallel()

	// GIVEN
	testYaml := Content("yamlValue: true")

	// WHEN
	var dto testDecodeResult
	err := testYaml.YAMLDecode(&dto)

	// THEN
	assert.NoError(t, err)
	assert.True(t, dto.Value)
}
