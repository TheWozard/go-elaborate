package util_test

import (
	"go-elaborate/pkg/util"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// There is unusual behavior for Lookup as it is using reflected type.
// Things that might seem diffrent return the same values when looked up.

func TestLookupPrimitivesName(t *testing.T) {
	lookup := util.NewNamedLookup(map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	})
	testCases := []struct {
		desc   string
		name   string
		ok     bool
		result int
	}{
		{
			desc:   "one",
			name:   "one",
			ok:     true,
			result: 1,
		},
		{
			desc:   "two",
			name:   "two",
			ok:     true,
			result: 2,
		},
		{
			desc:   "three",
			name:   "three",
			ok:     true,
			result: 3,
		},
		{
			desc:   "four",
			name:   "four",
			ok:     false,
			result: 0,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			result, ok := lookup.LookupByName(tC.name)
			assert.Equal(t, tC.ok, ok)
			assert.Equal(t, tC.result, result)
		})
	}
}

func TestLookupPrimitivesType(t *testing.T) {
	lookup := util.NewNamedLookup(map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	})
	testCases := []struct {
		desc   string
		input  int
		ok     bool
		result string
	}{
		{
			desc:   "one",
			input:  1,
			ok:     true,
			result: "three",
		},
		{
			desc:   "two",
			input:  5,
			ok:     true,
			result: "three",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			result, ok := lookup.LookupByStruct(tC.input)
			assert.Equal(t, tC.ok, ok)
			assert.Equal(t, tC.result, result)
		})
	}
}

func TestLookupStructName(t *testing.T) {
	type StructA struct{}
	type StructB struct {
		data string
	}
	lookup := util.NewReflectLookup[interface{}](StructA{}, StructB{}, StructB{data: "test"}, &StructB{data: "other"})
	testCases := []struct {
		desc   string
		name   string
		ok     bool
		result interface{}
	}{
		{
			desc:   "A",
			name:   "StructA",
			ok:     true,
			result: StructA{},
		},
		{
			desc: "B",
			name: "StructB",
			ok:   true,
			result: StructB{
				data: "test",
			},
		},
		{
			desc: "Indirection",
			name: "",
			ok:   true,
			result: &StructB{
				data: "other",
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			result, ok := lookup.LookupByName(tC.name)
			assert.Equal(t, tC.ok, ok)
			assert.Equal(t, tC.result, result)
		})
	}
}

func TestLookupStructType(t *testing.T) {
	type StructA struct{}
	type StructB struct {
		data string
	}
	lookup := util.NewNamedLookup(map[string]interface{}{
		"a": StructA{},
		"b": StructB{},
		"c": StructB{data: "test"},
	})
	testCases := []struct {
		desc   string
		input  interface{}
		ok     bool
		result string
	}{
		{
			desc:   "A",
			input:  StructA{},
			ok:     true,
			result: "a",
		},
		{
			desc:   "B",
			input:  StructB{},
			ok:     true,
			result: "c",
		},
		{
			desc:   "A",
			input:  struct{}{},
			ok:     false,
			result: "",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			result, ok := lookup.LookupByStruct(tC.input)
			assert.Equal(t, tC.ok, ok)
			assert.Equal(t, tC.result, result)
		})
	}
}

func TestMutationDoNotPersist(t *testing.T) {
	type StructB struct {
		data string
	}
	lookup := util.NewNamedLookup(map[string]StructB{
		"a": {},
	})
	a, _ := lookup.LookupByName("a")
	aprime, _ := lookup.LookupByName("a")
	require.Equal(t, aprime, a)
	a.data = "modified"
	require.NotEqual(t, aprime, a)
}
