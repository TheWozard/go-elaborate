package steps

import "context"

type LogStep struct {
}

func (ls LogStep) Run(ctx context.Context) error {
	return nil
}

func (ls LogStep) Step() Step {
	return Step{
		Name:    "log",
		Content: ls,
	}
}
