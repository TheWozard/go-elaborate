package steps

import (
	"context"
	"fmt"

	"gopkg.in/yaml.v3"
)

var (
	StepLookup = map[string]ContentProvider{
		"log": func() Content { return &LogStep{} },
	}
)

type Step struct {
	Name    string  `json:"name" yaml:"name"`
	Content Content `json:"content,omitempty" yaml:"content,omitempty"`
}

type Content interface {
	Run(context.Context) error
}

type ContentProvider func() Content

func (s *Step) UnmarshalYAML(node *yaml.Node) error {
	temp := struct {
		Name    string    `yaml:"name"`
		Content yaml.Node `yaml:"content"`
	}{}
	err := node.Decode(&temp)
	if err != nil {
		fmt.Println(err)
		return err
	}
	s.Name = temp.Name
	if provider, ok := StepLookup[temp.Name]; ok {
		s.Content = provider()
		return temp.Content.Decode(s.Content)
	} else {
		return fmt.Errorf("could not find step of name '%s'", temp.Name)
	}
}

func (s *Step) UnmarshalJSON(data []byte) error {
	// temp := struct {
	// 	Version string          `json:"version"`
	// 	Job     json.RawMessage `json:"job"`
	// }{}
	// err := json.Unmarshal(data, &temp)
	// if err != nil {
	// 	return err
	// }
	// m.Version = temp.Version
	// if provider, ok := JobLookup[temp.Version]; ok {
	// 	m.Job = provider()
	// } else {
	// 	return fmt.Errorf("could not find json job of version '%s'", m.Version)
	// }
	// if len(temp.Job) == 0 {
	// 	return fmt.Errorf("missing required json job content")
	// }
	// return json.Unmarshal(temp.Job, &m.Job)
	return nil
}
