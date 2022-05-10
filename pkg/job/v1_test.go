package job_test

import (
	"encoding/json"
	"go-elaborate/pkg/job"
	"go-elaborate/pkg/job/step"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

func TestMarshalV1(t *testing.T) {
	testCases := []struct {
		desc string
		job  job.V1
		yaml string
		json string
	}{
		{
			desc: "empty job",
			job:  job.V1{},
			yaml: "steps: []\n",
			json: "{\"steps\":null}",
		},
		{
			desc: "initalized steps",
			job:  job.V1{Steps: []step.Step{}},
			yaml: "steps: []\n",
			json: "{\"steps\":[]}",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			y, err := yaml.Marshal(tC.job)
			assert.Nil(t, err)
			assert.Equal(t, tC.yaml, string(y))
			j, err := json.Marshal(tC.job)
			assert.Nil(t, err)
			assert.Equal(t, tC.json, string(j))
		})
	}
}
