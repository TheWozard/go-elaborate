package step

import (
	"context"
	"encoding/json"
	"fmt"
	"go-elaborate/pkg/logging"
	"go-elaborate/pkg/util"
	"reflect"

	"go.uber.org/zap"
	"gopkg.in/yaml.v3"
)

const (
	LogStepName = "log"
)

var (
	// StepLookup provides bi-directional lookup of the relationship between implementations of Content and a string name.
	// To add new Content outside of this package call StepLookup.Add durring init.
	StepLookup = util.NewLookup(map[string]func() Content{
		LogStepName: func() Content { return &Log{} },
	})
)

// Content is the actual operations of a particular step
type Content interface {
	// Applies the operations of the step to a
	Run(context.Context) error
	// Converts a Content into a Step, wrapping it in standard per-step handling
	Step() Step
	// Used durring the decoding procees to determin if the struct should me marshaled
	IsEmpty() bool
}

// Step is a wrapping struct that handles all per-step operations
// TODO: Automatic tagging of step name as part of data lineage output.
// TODO: Wrap logger in with current step name.
type Step struct {
	Name    string  `json:"name" yaml:"name"`
	Content Content `json:"content,omitempty" yaml:"content,omitempty"`
}

func (s Step) Run(ctx context.Context) error {
	return s.Content.Run(logging.NewContext(ctx, zap.String("step", s.Name)))
}

func (s Step) MarshalYAML() (interface{}, error) {
	output := struct {
		Name    string  `yaml:"name"`
		Content Content `yaml:"content,omitempty"`
	}{
		Name: s.Name,
	}
	if s.Content != nil && !s.Content.IsEmpty() {
		output.Content = s.Content
	}
	return output, nil
}

func (s Step) MarshalJSON() ([]byte, error) {
	output := struct {
		Name    string  `json:"name"`
		Content Content `json:"content,omitempty"`
	}{
		Name: s.Name,
	}
	if s.Content != nil && !s.Content.IsEmpty() {
		output.Content = s.Content
	}
	return json.Marshal(output)
}

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
	if content, ok := StepLookup.LookupByName(temp.Name); ok {
		s.Content = content
		return temp.Content.Decode(s.Content)
	}
	return fmt.Errorf("could not find yaml step of name '%s'", temp.Name)
}

func (s *Step) UnmarshalJSON(data []byte) error {
	temp := struct {
		Name    string          `json:"name"`
		Content json.RawMessage `json:"content"`
	}{}
	err := json.Unmarshal(data, &temp)
	if err != nil {
		return err
	}
	if content, ok := StepLookup.LookupByName(temp.Name); ok {
		s.Content = content
		return json.Unmarshal(temp.Content, s.Content)
	}
	return fmt.Errorf("could not find json step of name '%s'", temp.Name)
}

func NewStep(c Content) (Step, error) {
	name, ok := StepLookup.LookupByType(c)
	if !ok {
		return Step{}, fmt.Errorf("could not find name for step %s", reflect.TypeOf(c))
	}
	return Step{
		Name:    name,
		Content: c,
	}, nil
}
