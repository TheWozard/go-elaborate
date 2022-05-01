package resolver

import (
	"os"
)

// JSONFile creates a resolver for a JSON formated file
func JSONFile(path string) Resolver {
	return fileResolver{
		Typ:    resolverTypeFile,
		Path:   path,
		Format: jsonUnmarshal{},
	}
}

// YAMLFile creates a resolver for a YAML formated file
func YAMLFile(path string) Resolver {
	return fileResolver{
		Typ:    resolverTypeFile,
		Path:   path,
		Format: yamlUnmarshal{},
	}
}

type fileResolver struct {
	Typ    string      `json:"type" yaml:"type"`
	Path   string      `json:"path" yaml:"path"`
	Format unmarshaler `json:"format" yaml:"format"`
}

func (fr fileResolver) Get() (interface{}, error) {
	raw, err := os.ReadFile(fr.Path)
	if err != nil {
		return nil, err
	}
	var data interface{}
	err = fr.Format.Unmarshal(raw, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
