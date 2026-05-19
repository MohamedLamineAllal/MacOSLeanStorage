package engine

import (
	"sync"
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

func TestCommandHandler_Handle_DryRun(t *testing.T) {
	logger := zap.NewNop()
	// MockCleaner with DryRun=true
	cleaner := &MockCleaner{dryRun: true}
	e := New(logger, &MockScanner{}, cleaner, nil)
	s := &scheduler.Scheduler{} 
	ch := NewCommandHandler(e, s, logger)

	target := config.TargetConfig{
		Name:    "DryRunTarget",
		Command: "false", // This would fail if executed
	}

	// Should not panic or error
	ch.Handle(target, CommandHooks{})
}

func TestCommandHandler_HookSequence(t *testing.T) {
	logger := zap.NewNop()
	e := New(logger, &MockScanner{}, &MockCleaner{dryRun: false}, nil)
	s := &scheduler.Scheduler{}
	ch := NewCommandHandler(e, s, logger)

	var sequence []string
	var mu sync.Mutex
	
	target := config.TargetConfig{
		Name:    "HookTarget",
		Command: "echo hello",
	}
	
	hooks := CommandHooks{
		BeforeHandleCommand:    func(name, cmd string, should bool) { mu.Lock(); sequence = append(sequence, "before_handle"); mu.Unlock() },
		BeforeExecutingCommand: func(name, cmd string) { mu.Lock(); sequence = append(sequence, "before_exec"); mu.Unlock() },
		AfterExecutingCommand:  func(name, cmd string, err error) { mu.Lock(); sequence = append(sequence, "after_exec"); mu.Unlock() },
		AfterHandleCommand:     func(name, cmd string, err error) { mu.Lock(); sequence = append(sequence, "after_handle"); mu.Unlock() },
	}

	ch.Handle(target, hooks)

	expected := []string{"before_handle", "before_exec", "after_exec", "after_handle"}
	assert.Equal(t, expected, sequence)
}
