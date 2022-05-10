package job

import (
	"context"
	"go-elaborate/pkg/job/step"
)

type V1 struct {
	// Resolvers
	Steps []step.Step `json:"steps" yaml:"steps"`
	// Outputs
}

func (v V1) Run(ctx context.Context) error {
	// Create output object
	// Iterate over steps
	// Validate?
	// Output
	// Profit
	return nil
}

func (v V1) Message() Message {
	message, _ := NewMessage(v)
	return message
}
