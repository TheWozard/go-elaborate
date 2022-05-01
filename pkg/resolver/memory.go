package resolver

// Go creates a resolver for in memory data
func Go(data interface{}) Resolver {
	return goResolver{
		Typ:  resolverTypeGo,
		data: data,
	}
}

type goResolver struct {
	Typ  string `json:"type" yaml:"type"`
	data interface{}
}

func (gr goResolver) Get() (interface{}, error) {
	return gr.data, nil
}
