package job_test

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"

	"go-elaborate/pkg/job"
	"go-elaborate/pkg/job/step"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

func TestParse(t *testing.T) {
	tests := []struct {
		name        string
		data        string
		unmarshaler func([]byte, interface{}) error
		message     job.Message
		err         error
	}{
		{
			// This is an interesting case. Technically not an error when there is nothing
			name:        "empty yaml",
			data:        ``,
			unmarshaler: yaml.Unmarshal,
		},
		{
			name:        "empty json",
			data:        ``,
			unmarshaler: json.Unmarshal,
			err:         fmt.Errorf("unexpected end of JSON input"),
		},
		{
			name:        "bad yaml version",
			data:        `version: bad`,
			unmarshaler: yaml.Unmarshal,
			err:         fmt.Errorf("could not find yaml job of version 'bad'"),
			message:     job.Message{Version: "bad"},
		},
		{
			name:        "bad json version",
			data:        `{"version":"bad"}`,
			unmarshaler: json.Unmarshal,
			err:         fmt.Errorf("could not find json job of version 'bad'"),
			message:     job.Message{Version: "bad"},
		},
		{
			name:        "yaml missing job",
			data:        `version: 1`,
			unmarshaler: yaml.Unmarshal,
			message:     job.Message{Version: "1", Job: &job.V1{}},
		},
		{
			name:        "json missing job",
			data:        `{"version": "1"}`,
			unmarshaler: json.Unmarshal,
			err:         fmt.Errorf("unexpected end of JSON input"),
			message:     job.Message{Version: "1", Job: &job.V1{}},
		},
		{
			name:        "yaml load v1",
			data:        "version: 1\njob: {}\n",
			unmarshaler: yaml.Unmarshal,
			message:     job.Message{Version: "1", Job: &job.V1{}},
		},
		{
			name:        "yaml loads job with data v1",
			data:        "version: 1\njob:\n  steps:\n    - name: log\n",
			unmarshaler: yaml.Unmarshal,
			message: job.Message{Version: "1", Job: &job.V1{
				Steps: []step.Step{
					{Name: "log", Content: &step.Log{}},
				},
			}},
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("%d/%s", i, test.name), func(t *testing.T) {
			message := job.Message{}
			err := test.unmarshaler([]byte(test.data), &message)
			if test.err != nil {
				assert.EqualError(t, err, test.err.Error())
			} else {
				assert.Nil(t, err)
			}
			assert.Equal(t, test.message, message)
		})
	}
}

// BadJob is a struct for a job not found in the standard lookup
type BadJob struct{}

func (bj BadJob) Run(ctx context.Context) error {
	return nil
}

func (bj BadJob) Message() job.Message {
	message, _ := job.NewMessage(bj)
	return message
}

func TestNewMessage(t *testing.T) {

	tests := []struct {
		name string
		job  job.Job
		err  error
		yaml string
		json string
	}{
		{
			name: "invalid Job",
			job:  BadJob{},
			err:  fmt.Errorf("could not find version for message job_test.BadJob"),
			yaml: "version: \"\"\njob: null\n",
			json: `{"version":"","job":null}`,
		},
		{
			name: "empty Job",
			job:  job.V1{},
			yaml: "version: \"1\"\njob:\n    steps: []\n",
			json: `{"version":"1","job":{"steps":null}}`,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("%d/%s", i, test.name), func(t *testing.T) {
			message, err := job.NewMessage(test.job)
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
