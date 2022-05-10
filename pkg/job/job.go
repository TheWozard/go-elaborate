package job

import (
	"context"
	"encoding/json"
	"fmt"
	"go-elaborate/pkg/util"
	"reflect"

	"gopkg.in/yaml.v3"
)

const (
	MessageV1 = "1"
)

var (
	// JobLookup provides bi-directional lookup of the relationship between implementations of Job and a string name.
	// To add new Jobs outside of this package call JobLookup.Add durring init.
	JobLookup = util.NewLookup(map[string]func() Job{
		MessageV1: func() Job { return &V1{} },
	})
)

// Job high level interface to define a runable thing
// A job should be json or yaml marshalable to provide feedback about the execution of the job
type Job interface {
	// Run executes the pipeline steps
	Run(context.Context) error

	// TODO: to make it easier to programatically edit a Job. It should support some minimum form of modification
	// AddDestination()
	// SetResolver()

	// Converts a Job to a Message. This is required to write our a versioned message.
	Message() Message
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
	if job, ok := JobLookup.LookupByName(temp.Version); ok {
		m.Job = job
		return temp.Job.Decode(m.Job)
	}
	return fmt.Errorf("could not find yaml job of version '%s'", m.Version)
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
	if job, ok := JobLookup.LookupByName(temp.Version); ok {
		m.Job = job
		return json.Unmarshal(temp.Job, m.Job)
	}
	return fmt.Errorf("could not find json job of version '%s'", m.Version)

}

func NewMessage(j Job) (Message, error) {
	version, ok := JobLookup.LookupByType(j)
	if !ok {
		return Message{}, fmt.Errorf("could not find version for message %s", reflect.TypeOf(j))
	}
	return Message{
		Version: version,
		Job:     j,
	}, nil
}
