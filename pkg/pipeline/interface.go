package pipeline

type Pipeline = func(context map[string]interface{}) error
