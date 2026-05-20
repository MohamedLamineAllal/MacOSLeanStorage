package cmd

import (
	"sync"
)

// ResultAggregator provides a thread-safe way to track unique files and aggregate scan stats.
// It uses a map to deduplicate paths, preventing memory-intensive path duplication.
type ResultAggregator struct {
	mu          sync.RWMutex
	uniquePaths map[string]int64
	totalSize   int64
}

// NewResultAggregator creates a new ResultAggregator.
func NewResultAggregator() *ResultAggregator {
	return &ResultAggregator{
		uniquePaths: make(map[string]int64),
	}
}

// Add appends results from a scanner, deduplicating paths and adding size.
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

// GetStats returns the unique file count and aggregated size.
func (ra *ResultAggregator) GetStats() (int, int64) {
	ra.mu.RLock()
	defer ra.mu.RUnlock()
	return len(ra.uniquePaths), ra.totalSize
}

// GetUniquePaths returns the list of all unique file paths.
func (ra *ResultAggregator) GetUniquePaths() []string {
	ra.mu.RLock()
	defer ra.mu.RUnlock()
	paths := make([]string, 0, len(ra.uniquePaths))
	for path := range ra.uniquePaths {
		paths = append(paths, path)
	}
	return paths
}
