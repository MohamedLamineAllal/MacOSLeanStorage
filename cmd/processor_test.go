package cmd

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestProcessor(t *testing.T) {
	logger := zap.NewNop()
	tp := NewTargetProcessor(logger, nil, true)
	assert.NotNil(t, tp)
}
