package engine

import (
	"testing"

	"github.com/mohamedlamineallal/MacosLeanStorage/internal/config"
	"github.com/mohamedlamineallal/MacosLeanStorage/internal/scanner"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestEngine_CleanParallel(t *testing.T) {
	logger := zap.NewNop()
	e := New(logger, nil, false)

	targets := []config.TargetConfig{
		{Name: "T1", Path: "/path/1"},
		{Name: "T2", Path: "/path/2"},
	}

	resMap := map[string]*scanner.Result{
		"T1": {TargetName: "T1", Files: []string{"/f1"}, FileSizes: []int64{100}},
		"T2": {TargetName: "T2", Files: []string{"/f2"}, FileSizes: []int64{200}},
	}

	hooks := Hooks{
		OnFileCleaned: func(path string, freed int64, err error) {},
	}

	count, totalSize, err := e.Clean(resMap, targets, hooks)

	assert.NoError(t, err)
	assert.Equal(t, 2, count)
	assert.Equal(t, int64(300), totalSize)
}
