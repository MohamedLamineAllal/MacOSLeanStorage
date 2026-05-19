package cmd

import (
	"fmt"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResultAggregator_Add(t *testing.T) {
	agg := NewResultAggregator()

	// Add unique paths
	agg.Add([]string{"/a", "/b"}, []int64{50, 50})
	agg.Add([]string{"/c"}, []int64{50})

	count, size := agg.GetStats()
	assert.Equal(t, 3, count)
	assert.Equal(t, int64(150), size)

	// Add overlapping paths
	agg.Add([]string{"/a", "/d"}, []int64{50, 100})
	count, size = agg.GetStats()
	assert.Equal(t, 4, count) // /a is already in
	assert.Equal(t, int64(250), size)
}

func TestResultAggregator_Concurrency(t *testing.T) {
	agg := NewResultAggregator()
	var wg sync.WaitGroup

	// Simulate concurrent additions from multiple workers
	numWorkers := 10
	filesPerWorker := 100
	wg.Add(numWorkers)

	for i := 0; i < numWorkers; i++ {
		go func(workerID int) {
			defer wg.Done()
			paths := make([]string, filesPerWorker)
			sizes := make([]int64, filesPerWorker)
			for j := 0; j < filesPerWorker; j++ {
				paths[j] = fmt.Sprintf("/file-%d-%d", workerID, j)
				sizes[j] = 10
			}
			agg.Add(paths, sizes)
		}(i)
	}

	wg.Wait()

	count, size := agg.GetStats()
	assert.Equal(t, numWorkers*filesPerWorker, count)
	assert.Equal(t, int64(numWorkers*filesPerWorker*10), size)
}
