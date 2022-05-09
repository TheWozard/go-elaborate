package util

import (
	"fmt"
	"reflect"
)

// NewLookup creates a new Lookup where the names of the generics are provided as keys in the map.
// In order to generate all the reverse lookups, all generators are called at least once to get the type of the
// generated object.
func NewLookup[T any](items map[string]func() T) *Lookup[T] {
	lookup := &Lookup[T]{
		nameLookup:    map[string]func() T{},
		genericLookup: map[string]string{},
	}
	for name, generator := range items {
		lookup.Add(name, generator)
	}
	return lookup
}

// Lookup generic bi-directional lookup for names and their associated generic.
// This should not be manually instantiated, and instead use NewLookup.
type Lookup[T any] struct {
	nameLookup    map[string]func() T
	genericLookup map[string]string
}

// LookupByName uses a string name to lookup the relevant generator and return a new instance of it.
func (l Lookup[T]) LookupByName(name string) (T, bool) {
	if generator, ok := l.nameLookup[name]; ok {
		return generator(), true
	}
	var zero T
	return zero, false
}

// LookupByType uses the reflected type of the passed item to lookup the name of its generator.
// In the event two generators generate the same type of object the last one added will be returned
func (l Lookup[T]) LookupByType(item T) (string, bool) {
	id := l.genericToId(item)
	name, ok := l.genericLookup[id]
	return name, ok
}

// HasName returns if the lookup has a generator for the passed name.
func (l Lookup[T]) HasName(name string) bool {
	_, ok := l.LookupByName(name)
	return ok
}

// HasType returns if the lookup has a name for the passed type.
func (l Lookup[T]) HasType(item T) bool {
	_, ok := l.LookupByType(item)
	return ok
}

// Add adds the generator and name to the lookup.
func (l Lookup[T]) Add(name string, generator func() T) {
	// TODO: currently this will accept multiple of the same typ and name
	// This will overwrite the previous holders of these values.
	// It might be worth erroring or panicing in the event of these.
	id := l.genericToId(generator())
	l.nameLookup[name] = generator
	l.genericLookup[id] = name
}

func (l Lookup[T]) genericToId(item T) string {
	typ := reflect.TypeOf(item)
	kind := typ.Kind()
	if kind == reflect.Pointer {
		return typ.Elem().Name()
	} else if kind == reflect.Array || kind == reflect.Slice {
		return fmt.Sprintf("[]%s", typ.Elem().Name())
	} else if kind == reflect.Map {
		return fmt.Sprintf("map[%s]%s", typ.Key().Name(), typ.Elem().Name())
	}
	return typ.Name()
}
