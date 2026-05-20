package engine

import (
	"sync"
	"testing"

	"github.com/mohamedlamineallal/MrLeanStorage/internal/config"
	"github.com/mohamedlamineallal/MrLeanStorage/internal/scanner"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

type MockScanner struct{}

func (m *MockScanner) Scan(target scanner.Target, ignorePatterns []string) (*scanner.Result, error) {
	return &scanner.Result{TargetName: target.Name, Files: []string{"/f1"}, FileSizes: []int64{10}}, nil
}

type MockCleaner struct{ dryRun bool }

func (m *MockCleaner) Clean(paths []string, hook func(path string, freed int64, err error)) (int, int64, error) {
	if hook != nil {
		hook(paths[0], 10, nil)
	}
	return 1, 10, nil
}
func (m *MockCleaner) DryRun() bool { return m.dryRun }

func TestEngine_HookSequence(t *testing.T) {
	logger := zap.NewNop()
	e := New(logger, &MockScanner{}, &MockCleaner{}, nil)

	var mu sync.Mutex
	var sequence []string

	targets := []config.TargetConfig{{Name: "T1", Path: "/p1"}}
	hooks := Hooks{
		OnTargetScanStart: func(name, path string) { mu.Lock(); sequence = append(sequence, "start"); mu.Unlock() },
		OnTargetScanEnd: func(name string, res *scanner.Result, err error) {
			mu.Lock()
			sequence = append(sequence, "end")
			mu.Unlock()
		},
		OnFileCleaned: func(path string, freed int64, err error) {
			mu.Lock()
			sequence = append(sequence, "cleaned")
			mu.Unlock()
		},
		OnTargetCleaned: func(name string) { mu.Lock(); sequence = append(sequence, "target_cleaned"); mu.Unlock() },
	}

	e.ScanAndClean(targets, hooks)

	expected := []string{"start", "end", "cleaned", "target_cleaned"}
	assert.Equal(t, expected, sequence)
}
