package document

// Registrar registers a global document name to a resolver and a schema
// the registered document immediately returns a Reference that can be used.
type Registrar interface {
	// Load begins loading the document immediately.
	// This will not block the current thread but start loading the document in the background.
	// Used when you know the document will always be needed.
	Load(name string, resolver Resolver, schema Schema) Reference
	// Lazy waits until the document is used to do a first time load.
	// Used when the document may not be needed.
	Lazy(name string, resolver Resolver, schema Schema) Reference
}

// Reference represents both a Document and a new Registrar that can register sub documents
// Any sub document will only begin loading after the parent document has been resolved
type Reference interface {
	Document
	Registrar
}

type Resolver interface {
	Get() interface{}
}

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
