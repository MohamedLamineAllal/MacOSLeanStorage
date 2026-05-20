package cmd

import (
	"testing"

	"github.com/mohamedlamineallal/MrLeanStorage/internal/config"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestProcessor_Run(t *testing.T) {
	logger := zap.NewNop()
	tp := NewTargetProcessor(logger, nil, true)

	targets := []config.TargetConfig{
		{
			Name: "Test Target",
			Path: "/tmp",
			Threshold: 30,
			Type: "file",
		},
	}

	err := tp.Run(targets, false, false)
	assert.NoError(t, err)
}
