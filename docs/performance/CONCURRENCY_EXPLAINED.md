# Concurrency, Parallelism, and Goroutines in Go

This document explains how the `mls` tool leverages Go's concurrency model to achieve high-performance filesystem scanning.

## 1. The Core Concept: Goroutines vs. OS Threads

### Traditional Multi-threading (OS Threads)
In languages like C++ or Java, "threading" typically refers to **OS Threads**. These are managed directly by the Operating System.
- **Heavyweight**: Each thread requires significant memory (stack size, often 1MB+) and context switching between threads is expensive (CPU registers, state saving/restoring).
- **Limited**: Creating thousands of OS threads will quickly exhaust system resources.

### Go Goroutines
Goroutines are **User-space Threads** managed by the Go runtime, not the OS.
- **Lightweight**: They start with a tiny stack (a few KB) that grows and shrinks as needed.
- **Fast Switching**: Context switching a goroutine is much faster than an OS thread because it involves saving/restoring fewer registers in user-space.

## 2. Are Goroutines "Real Parallelism"?

**Yes.** 

Go uses an **M:N Scheduler** (often called the "Go runtime scheduler"). 
- It maps `M` goroutines onto `N` OS threads.
- If you have 8 CPU cores, the Go runtime will generally attempt to keep 8 OS threads active (the `GOMAXPROCS` setting, which defaults to the number of CPU cores).
- **The "Magic"**: When a goroutine performs a blocking operation (like waiting for I/O from the disk), the Go scheduler moves other ready goroutines to a different, non-blocked OS thread, keeping your CPU cores busy.

So, while goroutines are designed for concurrency (managing many tasks at once, like async I/O), they **achieve true parallelism** by distributing those tasks across all available CPU cores when those tasks are computationally intensive or when they can run in parallel.

## 3. How the Worker Pool Works in `mls`

In the `TargetProcessor`, we use this model to parallelize the scanning of different cleanup targets:

```go
numWorkers := runtime.NumCPU() // Get the number of logical cores
jobs := make(chan config.TargetConfig, len(targets))
// ...
for i := 0; i < numWorkers; i++ {
    go func() { /* Worker goroutine logic */ }()
}
```

1. **Alignment with Cores**: By setting `numWorkers := runtime.NumCPU()`, we ensure that we don't create an arbitrary amount of CPU contention. We match the number of workers to the number of physical/logical cores.
2. **Channel-based Distribution**: We feed `targets` into a `jobs` channel. Multiple goroutines (workers) wait on this channel.
3. **True Parallelism**: Because these targets involve intensive file system I/O (which is often buffered and handled by the kernel) and CPU-bound staleness calculations (mtime checks, dir size recursion), having one worker per core ensures that all cores are working on different branches of the filesystem simultaneously.

## Parallelism Safety and Correctness

One might wonder: *How can we scan multiple filesystems simultaneously without causing data corruption or race conditions?*

The `mls` worker pool achieves thread-safe parallelism through three key architectural decisions:

### 1. Target Independence (Data Immutability)
The primary reason this parallelization works is that each cleanup target is **independent**. 
- A worker picks up a `config.TargetConfig` job from the channel.
- It scans the filesystem starting from that specific root.
- It does not modify global state or shared variables.
- Because no two workers are ever operating on the same root directory simultaneously, there is no risk of data contention or conflicting state updates within the `Scanner`.

### 2. Synchronization with `sync.WaitGroup`
To ensure the program doesn't terminate before scanning is finished:
- We use a `sync.WaitGroup` to track active workers.
- The main goroutine calls `wg.Wait()` to block until all workers have signaled `wg.Done()`. 
- This guarantees that the system only attempts to process or clean results once the entire parallel scan phase has reached a consistent "finalized" state.

### 3. Channel-based Aggregation (Communicating by Sharing)
We follow the Go mantra: *"Do not communicate by sharing memory; instead, share memory by communicating."*
- Workers do not write to a shared result slice (which would require a `sync.Mutex` and kill performance).
- Instead, each worker sends its own `scanner.Result` through a **buffered results channel**.
- The main goroutine collects these results from the channel once the workers are finished. This is inherently thread-safe because only the main goroutine reads from the results channel, while only the workers write to it.

### Why this approach is robust:
- **No Race Conditions**: Because we have no shared mutable state across workers, the code is effectively "race-free" by design.
- **Deterministic Summaries**: Because the order of processing doesn't impact the calculation of `TotalSize` or the list of `Files`, we achieve perfectly reproducible output regardless of which worker picks up which job.

