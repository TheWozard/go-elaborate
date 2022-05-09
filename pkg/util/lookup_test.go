package util_test

import (
	"go-elaborate/pkg/util"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLookupByName(t *testing.T) {
	type CustomStruct struct{}
	type CustomWithPointer struct{}
	lookup := util.NewLookup(map[string]func() interface{}{
		"one":            func() interface{} { return 1 },
		"name":           func() interface{} { return "John Doe" },
		"colors":         func() interface{} { return []string{"red", "green", "blue"} },
		"custom":         func() interface{} { return CustomStruct{} },
		"custom-pointer": func() interface{} { return &CustomWithPointer{} },
	})
	testCases := []struct {
		name   string
		ok     bool
		result interface{}
	}{
		{
			name: "one", ok: true, result: 1,
		},
		{
			name: "name", ok: true, result: "John Doe",
		},
		{
			name: "colors", ok: true, result: []string{"red", "green", "blue"},
		},
		{
			name: "custom", ok: true, result: CustomStruct{},
		},
		{
			name: "custom-pointer", ok: true, result: &CustomWithPointer{},
		},
		{
			name: "missing", ok: false, result: nil,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			result, ok := lookup.LookupByName(tC.name)
			assert.Equal(t, tC.ok, ok)
			assert.Equal(t, tC.result, result)
		})
	}
}

func TestLookupByType(t *testing.T) {
	type CustomStruct struct {
		data string
	}
	type CustomWithPointer struct {
		other string
	}
	lookup := util.NewLookup(map[string]func() interface{}{
		"one":            func() interface{} { return 1 },
		"name":           func() interface{} { return "John Doe" },
		"colors":         func() interface{} { return []string{"red", "green", "blue"} },
		"custom":         func() interface{} { return CustomStruct{} },
		"custom-pointer": func() interface{} { return &CustomWithPointer{} },
		"map":            func() interface{} { return map[string]int{} },
	})
	testCases := []struct {
		desc  string
		input interface{}
		ok    bool
		name  string
	}{
		{
			desc: "int", input: 1, ok: true, name: "one",
		},
		{
			desc: "diffrent int", input: 2, ok: true, name: "one",
		},
		{
			desc: "string", input: "John Doe", ok: true, name: "name",
		},
		{
			desc: "diffrent string", input: "Just some text", ok: true, name: "name",
		},
		{
			desc: "string slice", input: []string{"red", "green", "blue"}, ok: true, name: "colors",
		},
		{
			desc: "empty string slice", input: []string{}, ok: true, name: "colors",
		},
		{
			desc: "custom struct", input: CustomStruct{}, ok: true, name: "custom",
		},
		{
			desc: "modified custom struct", input: CustomStruct{data: "this one has non-zero data"}, ok: true, name: "custom",
		},
		{
			desc: "custom pointer generator struct", input: CustomWithPointer{}, ok: true, name: "custom-pointer",
		},
		{
			desc: "empty int slice", input: []int{}, ok: false, name: "",
		},
		{
			desc: "generic struct", input: struct{}{}, ok: false, name: "",
		},
		{
			desc: "generic map", input: map[string]string{}, ok: false, name: "",
		},
		{
			desc: "generic map", input: map[string]int{}, ok: true, name: "map",
		},
		{
			desc: "generic function", input: func() {}, ok: false, name: "",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			result, ok := lookup.LookupByType(tC.input)
			assert.Equal(t, tC.ok, ok)
			assert.Equal(t, tC.name, result)
		})
	}
}
