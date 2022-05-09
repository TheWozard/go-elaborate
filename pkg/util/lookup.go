package util

import "reflect"

// NewReflectLookup creates a new Lookup where the names are automatically generated using the reflected type name.
func NewReflectLookup[T any](items ...T) *Lookup[T] {
	lookup := &Lookup[T]{
		nameLookup:    map[string]T{},
		genericLookup: map[string]string{},
	}
	for _, item := range items {
		name := reflect.TypeOf(item).Name()
		lookup.nameLookup[name] = item
		lookup.genericLookup[name] = name
	}
	return lookup
}

// NewNamedLookup creates a new Lookup where the names of the generics are provided as keys in the map.
func NewNamedLookup[T any](items map[string]T) *Lookup[T] {
	lookup := &Lookup[T]{
		nameLookup:    map[string]T{},
		genericLookup: map[string]string{},
	}
	for name, item := range items {
		typ := reflect.TypeOf(item).Name()
		lookup.nameLookup[name] = item
		lookup.genericLookup[typ] = name
	}
	return lookup
}

// Lookup generic bi-directional lookup for names and their associated generic.
// This should not be manually instantiated, and instead use NewReflectLookup or NewNamedLookup.
type Lookup[T any] struct {
	nameLookup    map[string]T
	genericLookup map[string]string
}

// LookupByName uses a string name to lookup the relevant generic and return a new instance of it.
// In the event of a pointer, the pointer will still point to the original interfaces.
func (l Lookup[T]) LookupByName(name string) (T, bool) {
	if item, ok := l.nameLookup[name]; ok {
		return item, true
	}
	var zero T
	return zero, false
}

// LookupByStruct uses the reflected type of the passed item to lookup the name this generic
// has been given. This may just be the reflected name of the type or a unique string, depending
// on how the lookup was created.
func (l Lookup[T]) LookupByStruct(item T) (string, bool) {
	name, ok := l.genericLookup[reflect.TypeOf(item).Name()]
	return name, ok
}
