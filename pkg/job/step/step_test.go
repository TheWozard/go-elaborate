package step_test

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"

	"go-elaborate/pkg/job/step"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

func TestUnmarshalStep(t *testing.T) {
	tests := []struct {
		name        string
		data        string
		unmarshaler func([]byte, interface{}) error
		step        step.Step
		err         error
	}{
		{
			// This is an interesting case. Technically not an error when there is nothing
			name:        "empty yaml",
			data:        ``,
			unmarshaler: yaml.Unmarshal,
			step:        step.Step{},
		},
		{
			name:        "empty content",
			data:        `name: log`,
			unmarshaler: yaml.Unmarshal,
			step:        step.Step{Name: "log", Content: &step.Log{}},
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("%d/%s", i, test.name), func(t *testing.T) {
			step := step.Step{}
			err := test.unmarshaler([]byte(test.data), &step)
			if test.err != nil {
				assert.EqualError(t, err, test.err.Error())
			} else {
				assert.Nil(t, err)
			}
			assert.Equal(t, test.step, step)
		})
	}
}

// BadContent is a struct for a Content not found in the standard lookup
type BadContent struct{}

func (bc BadContent) Run(ctx context.Context) error {
	return nil
}

func (bc BadContent) Step() step.Step {
	message, _ := step.NewStep(bc)
	return message
}

func (bc BadContent) IsEmpty() bool {
	return true
}

func TestNewStep(t *testing.T) {

	tests := []struct {
		name    string
		content step.Content
		err     error
		yaml    string
		json    string
	}{
		{
			name:    "invalid content",
			content: BadContent{},
			err:     fmt.Errorf("could not find name for step step_test.BadContent"),
			yaml:    "name: \"\"\n",
			json:    `{"name":""}`,
		},
		{
			name:    "log",
			content: step.Log{},
			yaml:    "name: log\n",
			json:    `{"name":"log"}`,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("%d/%s", i, test.name), func(t *testing.T) {
			message, err := step.NewStep(test.content)
			if test.err != nil {
				assert.EqualError(t, err, test.err.Error())
			} else {
				assert.Nil(t, err)
			}
			yaml, err := yaml.Marshal(message)
			assert.Nil(t, err)
			assert.Equal(t, test.yaml, string(yaml))
			json, err := json.Marshal(message)
			assert.Nil(t, err)
			assert.Equal(t, test.json, string(json))
		})
	}
}
