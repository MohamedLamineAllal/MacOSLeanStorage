package engine

import (
	"os"
	"runtime"
	"sync"
	"time"

	"github.com/mohamedlamineallal/MacosLeanStorage/internal/cleaner"
	"github.com/mohamedlamineallal/MacosLeanStorage/internal/config"
	"github.com/mohamedlamineallal/MacosLeanStorage/internal/scanner"
	"go.uber.org/zap"
)

// LogEvent represents a structured event for logging callbacks.
type LogEvent struct {
	Type    string
	Message string
	Path    string
	Size    int64
}

// Engine orchestrates scanning and cleaning targets.
type Engine struct {
	scanner *scanner.Scanner
	cleaner *cleaner.Cleaner
	logger  *zap.Logger
}

// NewEngine creates a new Engine.
func NewEngine(logger *zap.Logger, ignorePatterns []string, dryRun bool) *Engine {
	return &Engine{
		scanner: scanner.New(logger, ignorePatterns),
		cleaner: cleaner.New(logger, dryRun, ignorePatterns),
		logger:  logger,
	}
}

// RunOptions configures the execution of the Engine.
type RunOptions struct {
	IsClean  bool
	DryRun   bool
	LogFile  *os.File
	OnEvent  func(event LogEvent) // Callback for UI/CLI feedback
}

// ResultAggregator tracks unique file stats.
type ResultAggregator struct {
	mu           sync.RWMutex
	uniquePaths  map[string]int64
	totalSize    int64
}

func (ra *ResultAggregator) Add(files []string, sizes []int64) {
	ra.mu.Lock()
	defer ra.mu.Unlock()
	for i, file := range files {
		if _, exists := ra.uniquePaths[file]; !exists {
			ra.uniquePaths[file] = sizes[i]
			ra.totalSize += sizes[i]
		}
	}
}

// Run executes the full scan/clean pipeline.
func (e *Engine) Run(targets []config.TargetConfig, opts RunOptions) (int, int64, error) {
	e.cleaner.SetDryRun(opts.DryRun)
	if opts.LogFile != nil {
		e.cleaner.SetLogFile(opts.LogFile)
	}

	resultMap := e.ScanTargets(targets)
	aggregator := &ResultAggregator{uniquePaths: make(map[string]int64)}

	for _, t := range targets {
		if t.Command != "" {
			continue // Command execution handled at processor level
		}

		res, ok := resultMap[t.Name]
		if !ok {
			continue
		}

		if opts.OnEvent != nil {
			opts.OnEvent(LogEvent{Type: "target_start", Message: t.Name, Path: t.Path})
		}

		aggregator.Add(res.Files, res.FileSizes)

		if opts.IsClean && len(res.Files) > 0 {
			_, _, err := e.cleaner.Clean(res.Files)
			if err != nil {
				e.logger.Error("Clean failed", zap.String("target", t.Name), zap.Error(err))
			}
		}
	}

	uniqueCount := len(aggregator.uniquePaths)
	return uniqueCount, aggregator.totalSize, nil
}

// ScanTargets processes multiple targets in parallel and returns scan results.
func (e *Engine) ScanTargets(targets []config.TargetConfig) map[string]*scanner.Result {
	numWorkers := runtime.NumCPU()
	jobs := make(chan config.TargetConfig, len(targets))
	results := make(chan struct {
		Name string
		Res  *scanner.Result
		Err  error
	}, len(targets))

	var wg sync.WaitGroup
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for t := range jobs {
				target := scanner.Target{
					Name:        t.Name,
					Path:        t.Path,
					Threshold:   time.Duration(t.Threshold) * 24 * time.Hour,
					SafetyLevel: t.SafetyLevel,
					Type:        t.Type,
				}
				res, err := e.scanner.Scan(target, t.IgnorePatterns)
				results <- struct {
					Name string
					Res  *scanner.Result
					Err  error
				}{t.Name, res, err}
			}
		}()
	}

	for _, t := range targets {
		if t.Command == "" {
			jobs <- t
		}
	}
	close(jobs)
	wg.Wait()
	close(results)

	resultMap := make(map[string]*scanner.Result)
	for res := range results {
		if res.Err == nil {
			resultMap[res.Name] = res.Res
		}
	}
	return resultMap
}

// Cleaner returns the underlying Cleaner instance.
func (e *Engine) Cleaner() *cleaner.Cleaner {
	return e.cleaner
}
