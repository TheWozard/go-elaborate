package resolver

const (
	resolverTypeSchema  = "schema"
	resolverTypeMemo    = "memo"
	resolverTypeTracker = "tracker"
	resolverTypeFile    = "file"
	resolverTypeGo      = "go"
)

// Standard interface for geting some generic data chunk to be transformed
type Resolver interface {
	Get() (interface{}, error)
}
