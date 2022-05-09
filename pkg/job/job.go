package job

import (
	"context"
	"encoding/json"
	"fmt"

	"gopkg.in/yaml.v3"
)

const (
	MessageV1 = "1"
)

var (
	lookup = map[string]func() Job{
		"1": func() Job { return &JobV1{} },
	}
)

// Job high level interface to define a runable thing
// A job should be json or yaml marshalable to provide feedback about the execution of the job
type Job interface {
	// Run executes the pipeline
	Run(context.Context) error

	// TODO: to make it easier to programatically edit a Job. It should support some minimum form of modification
	// AddDestination()
	// SetResolver()
}

// Message is a wrapping struct for versioning job content
type Message struct {
	Version string `json:"version" yaml:"version"`
	Job     Job    `json:"job" yaml:"job"`
}

func (m *Message) UnmarshalYAML(node *yaml.Node) error {
	temp := struct {
		Version string    `yaml:"version"`
		Job     yaml.Node `yaml:"job"`
	}{}
	err := node.Decode(&temp)
	if err != nil {
		return err
	}
	m.Version = temp.Version
	if provider, ok := lookup[temp.Version]; ok {
		m.Job = provider()
		return temp.Job.Decode(m.Job)
	} else {
		return fmt.Errorf("could not find yaml job of version '%s'", m.Version)
	}
}

func (m *Message) UnmarshalJSON(data []byte) error {
	temp := struct {
		Version string          `json:"version"`
		Job     json.RawMessage `json:"job"`
	}{}
	err := json.Unmarshal(data, &temp)
	if err != nil {
		return err
	}
	m.Version = temp.Version
	if provider, ok := lookup[temp.Version]; ok {
		m.Job = provider()
		return json.Unmarshal(temp.Job, m.Job)
	} else {
		return fmt.Errorf("could not find json job of version '%s'", m.Version)
	}
}
