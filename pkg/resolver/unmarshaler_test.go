package resolver

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"
)

func TestUnmarshalers(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		unmarshaler unmarshaler
		output      interface{}
		err         error
	}{
		{
			name:        "JSON_Unmarshal_Success",
			input:       `{"data":"some basic data"}`,
			unmarshaler: jsonUnmarshal{},
			output:      map[string]interface{}{"data": "some basic data"},
		},
		{
			name:        "JSON_Unmarshal_Failure",
			input:       `}{`,
			unmarshaler: jsonUnmarshal{},
			err:         fmt.Errorf("invalid character '}' looking for beginning of value"),
		},
		{
			name:        "YAML_Unmarshal_Success",
			input:       `data: some basic data`,
			unmarshaler: yamlUnmarshal{},
			output:      map[interface{}]interface{}{"data": "some basic data"},
		},
		{
			name:        "YAML_Unmarshal_Success",
			input:       "---\ndata: some basic data",
			unmarshaler: yamlUnmarshal{},
			output:      map[interface{}]interface{}{"data": "some basic data"},
		},
		{
			name:        "YAML_Unmarshal_Failure",
			input:       `}{`,
			unmarshaler: yamlUnmarshal{},
			err:         fmt.Errorf("yaml: did not find expected node content"),
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("%d_%s", i, test.name), func(t *testing.T) {
			var data interface{}
			err := test.unmarshaler.Unmarshal([]byte(test.input), &data)
			if test.err == nil {
				require.Nil(t, err)
			} else {
				require.EqualError(t, err, test.err.Error())
			}
			require.Equal(t, test.output, data)
		})
	}
}

func TestUnmarshalersMarshal(t *testing.T) {
	tests := []struct {
		name        string
		unmarshaler unmarshaler
		json        string
		yaml        string
	}{
		{
			name:        "JSON",
			unmarshaler: jsonUnmarshal{},
			json:        `"JSON"`,
			yaml:        "JSON\n",
		},
		{
			name:        "YAML",
			unmarshaler: yamlUnmarshal{},
			json:        `"YAML"`,
			yaml:        "YAML\n",
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("%d_%s", i, test.name), func(t *testing.T) {
			j, err := json.Marshal(test.unmarshaler)
			require.Nil(t, err)
			assert.Equal(t, test.json, string(j))
			y, err := yaml.Marshal(test.unmarshaler)
			require.Nil(t, err)
			assert.Equal(t, test.yaml, string(y))
		})
	}
}
