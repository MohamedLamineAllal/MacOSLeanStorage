package engine

import (
	"testing"
	"time"

	"github.com/mohamedlamineallal/MrLeanStorage/internal/config"
	"github.com/mohamedlamineallal/MrLeanStorage/internal/scanner"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

type SlowCleaner struct{}
func (m *SlowCleaner) Clean(paths []string, hook func(path string, freed int64, err error)) (int, int64, error) {
	time.Sleep(100 * time.Millisecond)
	return 1, 10, nil
}
func (m *SlowCleaner) DryRun() bool { return false }

func TestEngine_CleanParallelStress(t *testing.T) {
	logger := zap.NewNop()
	e := New(logger, &MockScanner{}, &SlowCleaner{}, nil)

	targets := []config.TargetConfig{
		{Name: "T1", Path: "/p1"},
		{Name: "T2", Path: "/p2"},
		{Name: "T3", Path: "/p3"},
	}

	resMap := map[string]*scanner.Result{
		"T1": {TargetName: "T1", Files: []string{"/f1"}, FileSizes: []int64{10}},
		"T2": {TargetName: "T2", Files: []string{"/f2"}, FileSizes: []int64{10}},
		"T3": {TargetName: "T3", Files: []string{"/f3"}, FileSizes: []int64{10}},
	}

	start := time.Now()
	_, _, err := e.Clean(resMap, targets, Hooks{})
	elapsed := time.Since(start)

	assert.NoError(t, err)
	// If it was sequential, it would take ~300ms. If parallel, it should be significantly less.
	assert.Less(t, elapsed, 250*time.Millisecond)
}
