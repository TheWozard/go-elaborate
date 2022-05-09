package job_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"go-elaborate/pkg/job"
	"go-elaborate/pkg/job/steps"

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
			message:     job.Message{Version: "1", Job: &job.JobV1{}},
		},
		{
			name:        "json missing job",
			data:        `{"version": "1"}`,
			unmarshaler: json.Unmarshal,
			err:         fmt.Errorf("unexpected end of JSON input"),
			message:     job.Message{Version: "1", Job: &job.JobV1{}},
		},
		{
			name:        "yaml load v1",
			data:        "version: 1\njob: {}\n",
			unmarshaler: yaml.Unmarshal,
			message:     job.Message{Version: "1", Job: &job.JobV1{}},
		},
		{
			name:        "yaml loads job with data v1",
			data:        "version: 1\njob:\n  steps:\n    - name: log\n",
			unmarshaler: yaml.Unmarshal,
			message: job.Message{Version: "1", Job: &job.JobV1{
				Steps: []steps.Step{
					{Name: "log", Content: &steps.LogStep{}},
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
