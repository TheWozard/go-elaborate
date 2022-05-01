package resolver_test

import (
	"encoding/json"
	"fmt"
	"go-elaborate/pkg/resolver"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"
)

func TestFileResolvers(t *testing.T) {

	tests := []struct {
		name     string
		path     string
		resolver func(path string) resolver.Resolver
		output   interface{}
		err      error
	}{
		{
			name:     "JSON_File",
			path:     "./testdata/data.json",
			resolver: resolver.JSONFile,
			output:   map[string]interface{}{"id": float64(1), "data": "basic json file"},
		},
		{
			name:     "YAML_File",
			path:     "./testdata/data.yaml",
			resolver: resolver.YAMLFile,
			output:   map[interface{}]interface{}{"id": 1, "data": "basic yaml file"},
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("%d_%s", i, test.name), func(t *testing.T) {
			res := test.resolver(test.path)
			result, err := res.Get()
			if test.err == nil {
				require.Nil(t, err)
			} else {
				require.EqualError(t, err, test.err.Error())
			}
			assert.Equal(t, test.output, result)
		})
	}
}

func TestFileResolversMarshal(t *testing.T) {

	tests := []struct {
		name     string
		path     string
		resolver func(path string) resolver.Resolver
		json     string
		yaml     string
	}{
		{
			name:     "JSON_File",
			path:     "./testdata/data.json",
			resolver: resolver.JSONFile,
			json:     `{"type":"file","path":"./testdata/data.json","format":"JSON"}`,
			yaml:     "type: file\npath: ./testdata/data.json\nformat: JSON\n",
		},
		{
			name:     "YAML_File",
			path:     "./testdata/data.yaml",
			resolver: resolver.YAMLFile,
			json:     `{"type":"file","path":"./testdata/data.yaml","format":"YAML"}`,
			yaml:     "type: file\npath: ./testdata/data.yaml\nformat: YAML\n",
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("%d_%s", i, test.name), func(t *testing.T) {
			res := test.resolver(test.path)
			j, err := json.Marshal(res)
			require.Nil(t, err)
			assert.Equal(t, test.json, string(j))
			y, err := yaml.Marshal(res)
			require.Nil(t, err)
			assert.Equal(t, test.yaml, string(y))
		})
	}
}
