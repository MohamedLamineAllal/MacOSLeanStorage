package engine

import (
	"testing"

	"github.com/mohamedlamineallal/MacosLeanStorage/internal/config"
	"github.com/mohamedlamineallal/MacosLeanStorage/internal/scheduler"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

// MockScheduler implements the scheduler interface for testing
type MockScheduler struct{}

func (m *MockScheduler) ShouldRunCommand(name string, intervalDays int) bool {
	return true
}
func (m *MockScheduler) UpdateCommandRunTime(name string) {}

func TestCommandHandler_Handle(t *testing.T) {
	logger := zap.NewNop()
	e := New(logger, nil, false)
	s := &scheduler.Scheduler{} // Note: In a real scenario, mock this properly if possible
	ch := NewCommandHandler(e, s, logger)

	target := config.TargetConfig{
		Name:         "TestTarget",
		Command:      "echo hello",
		IntervalDays: 1,
	}

	hooks := CommandHooks{
		BeforeExecutingCommand: func(name, command string) {
			assert.Equal(t, "TestTarget", name)
		},
	}

	ch.Handle(target, hooks)
}
