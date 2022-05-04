package resolver

const (
	resolverTypeSchema    = "schema"
	resolverTypeMemo      = "memo"
	resolverTypeTracker   = "tracker"
	resolverTypeFile      = "file"
	resolverTypeGo        = "go"
	resolverTypeIntercept = "intercept"
)

// Standard interface for geting some generic data chunk to be transformed
type Resolver interface {
	Get() (interface{}, error)
}

func (resolvers []Resolver) Get() (interface{}, error)
