package job

import (
	"context"

	"go-elaborate/pkg/job/steps"
)

type JobV1 struct {
	// Resolvers
	Steps []steps.Step `json:"steps" yaml:"steps"`
	// Outputs
}

func (j JobV1) Run(ctx context.Context) error {
	// Create output object
	// Iterate over steps
	// Validate?
	// Output
	// Profit
	return nil
}
