package resolver

// Go creates a resolver for in memory data
func Go(data interface{}) Resolver {
	return goResolver{
		Typ:  resolverTypeGo,
		Data: data,
	}
}

type goResolver struct {
	Typ  string      `json:"type" yaml:"type"`
	Data interface{} `json:"data,omitempty" yaml:"data"`
}

func (gr goResolver) Get() (interface{}, error) {
	return gr.Data, nil
}

// DataWithBackup will return a resolver to the data if it exists, otherwise it will return the backup Resolver
func DataWithBackup(data interface{}, backup Resolver) Resolver {
	if data != nil {
		return Go(data)
	}
	return backup
}
