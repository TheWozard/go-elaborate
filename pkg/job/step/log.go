package step

import "context"

type Log struct {
}

func (ls Log) Run(ctx context.Context) error {
	return nil
}

func (ls Log) Step() Step {
	step, _ := NewStep(ls)
	return step
}

// No data is stored in Log so it is always empty
func (ls Log) IsEmpty() bool {
	return true
}
