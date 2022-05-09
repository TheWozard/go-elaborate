package document

type Document interface {
	// Get returns t
	Get(path string) interface{}
	GetDefault(path string, def interface{}) interface{}
}

// Extractor is a generically callable implementation of pulling a particular variable from the data
type Extractor interface {
	Get(interface{}) interface{}
}

type Operation interface {
	Do() interface{}
}

// Schema is used to validate a certain interface matches expectations
type Schema interface {
	// Verify returns if the raw data matches the expectations of the schema
	Verify(raw interface{}) bool
}

type Elaborative interface {
	Describe()
}
