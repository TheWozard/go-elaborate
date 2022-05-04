package pipeline

import (
	"fmt"
	"go-elaborate/pkg/resolver"
)

type ExamplePipeline1 struct {
}

func Run(context map[string]interface{}, output map[string]interface{}) error {
	// Setup
	documentA, err := resolver.DataWithBackup(context["documentA"], resolver.Go(map[string]string{
		"fieldA": "a1",
	})).Get()
	if err != nil {
		return fmt.Errorf("failed to get documentA %v", err)
	}

	output.Set()
}
