package resolver_test

import (
	"go-elaborate/pkg/resolver"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMemoizerResolver(t *testing.T) {
	testData := "testdata"

	counterA := &testCountingResolver{resolver: resolver.Go(testData)}
	counterB := &testCountingResolver{resolver: resolver.Go(testData)}
	res := resolver.Memo(counterA)

	for i := 0; i < 10; i++ {
		result, err := res.Get()
		assert.Nil(t, err)
		assert.Equal(t, testData, result)
		assert.Equal(t, counterA.count, 1)

		result2, err := counterB.Get()
		assert.Nil(t, err)
		assert.Equal(t, testData, result2)
		assert.Equal(t, counterB.count, i+1)
	}
}

type testCountingResolver struct {
	count    int
	resolver resolver.Resolver
}

func (tcr *testCountingResolver) Get() (interface{}, error) {
	tcr.count++
	return tcr.resolver.Get()
}
