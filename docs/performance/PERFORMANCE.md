# Performance Architecture

## Overview
To improve the performance of the scanning operation, especially when dealing with large numbers of targets or deep directory trees, I have implemented a worker pool pattern for parallel scanning.

## Design Choices
- **Worker Pool Pattern**: Instead of scanning targets sequentially, I have introduced a dispatcher/worker pattern using Go routines and channels.
- **Why Worker Pool**:
    - **Resource Management**: Limits the number of concurrent I/O operations, preventing the tool from overwhelming the OS or hitting file descriptor limits.
    - **Concurrency**: Leverages Go's lightweight routines for efficient parallel processing of multiple target paths.
    - **Backpressure**: Prevents memory spikes by controlling the ingestion of jobs.

## Implementation Details
1. **Job Queue**: A channel `jobs := make(chan scanner.Target, len(targets))` handles target distribution.
2. **Result Aggregation**: A buffered channel `results := make(chan scanner.Result, len(targets))` aggregates results from workers.
3. **Workers**: A pool of workers (configurable, currently set to a reasonable number based on CPU cores) consumes the jobs and performs the scanning.
4. **Synchronization**: `sync.WaitGroup` is used to ensure all workers finish before closing result channels and summarizing findings.

## Directory Traversal & Globbing Strategy

### How We Walk the Tree
The `mls` tool uses a combination of pattern-based globbing and recursive depth-first traversal to identify stale files.

1.  **Globbing Phase**:
    - We use `github.com/bmatcuk/doublestar/v4` to parse target paths.
    - This allows for powerful patterns like `**/Cache` or `.../Service Worker/**`.
    - Globbing handles the high-level path expansion. Once the glob matches the initial directory structure, the scanner takes over.

2.  **Recursive Traversal (`walkFiles`)**:
    - For each matched path, `mls` recursively crawls the subdirectories.
    - **Optimization - Fast Exit**: During the scan, if a folder is marked as "stale" (based on its mtime or its contents' staleness), `mls` treats it as a single deleteable unit. This prevents the scanner from wasting I/O by walking thousands of files inside a directory that will be deleted anyway.
    - **Ignore Pattern Enforcement**: At every step of the walk, the scanner checks the entry name against `ignorePatterns` (e.g., `.DS_Store`, `.git`, `.fseventsd`). If an entry is ignored, the scanner skips it and any of its children entirely.

## Parallel Deletion Implementation

To match the high-performance scanning, the `clean` operation now also utilizes a parallel worker pool.

### How Parallel Deletion Works
1. **Worker Distribution**: When `Clean(paths)` is called, the tool calculates the available CPU cores using `runtime.NumCPU()`.
2. **Path Channel**: A buffered channel distributes the list of files and directories to be deleted to a pool of worker goroutines.
3. **Parallel I/O**: Each worker concurrently executes `os.RemoveAll(path)` and calculates file sizes, effectively parallelizing the I/O-bound cleanup process.
4. **Result Aggregation**: A result channel safely collects the number of deleted items and the space freed from all workers, which the main thread then aggregates into the final cleanup summary.

### Performance Benefits
- **I/O Saturating**: By deleting files in parallel, the tool can better saturate the filesystem's I/O bandwidth, which is particularly beneficial when deleting thousands of small cache files across different directories.
- **Latency Reduction**: Total deletion time is reduced from a sequential crawl to near-parallel hardware limits.
