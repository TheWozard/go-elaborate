package resolver

// Memo wraps the resolver into a memo to resolve repeated calls.
func Memo(resolver Resolver) Resolver {
	return &memo{
		Typ:      resolverTypeMemo,
		Resolver: resolver,
	}
}

type memo struct {
	Typ      string   `json:"type" yaml:"type"`
	Resolver Resolver `json:"resolver" yaml:"resolver"`
	Requests int      `json:"requests" yaml:"requests"`
	data     interface{}
	err      error
}

func (m *memo) Get() (interface{}, error) {
	m.Requests += 1
	if m.data == nil && m.err == nil {
		m.data, m.err = m.Resolver.Get()
	}
	return m.data, m.err
}
