package resolver_test

import (
	"fmt"
	"go-elaborate/pkg/resolver"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGoResolver(t *testing.T) {
	tests := []struct {
		name string
		data interface{}
		err  error
	}{
		{
			name: "string",
			data: "test data",
		},
		{
			name: "map",
			data: map[string]interface{}{
				"example": "data",
			},
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("%d_%s", i, test.name), func(t *testing.T) {
			res := resolver.Go(test.data)
			result, err := res.Get()
			if test.err == nil {
				require.Nil(t, err)
			} else {
				require.EqualError(t, err, test.err.Error())
			}
			assert.Equal(t, test.data, result)
		})
	}
}
