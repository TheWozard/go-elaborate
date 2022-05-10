package logging

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestLoggerUpdatesOnContext(t *testing.T) {
	var ctx context.Context
	assert.Equal(t, logger, From(ctx))
	assert.Same(t, logger, From(ctx))
	assert.Equal(t, logger, From(context.Background()))
	assert.Same(t, logger, From(ctx))
	ctx = NewContext(context.Background())
	assert.Equal(t, logger, From(ctx))
	assert.Same(t, logger, From(ctx))
	ctx = NewContext(context.Background(), zap.String("some", "addition"))
	assert.NotEqual(t, logger, From(ctx))
	assert.NotSame(t, logger, From(ctx))
}
