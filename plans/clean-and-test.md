# Plan: Clean Process Concurrency & Engine Testing

## Objective
1. Refactor `Engine.Clean()` to process targets in parallel.
2. Develop a comprehensive test suite for `internal/engine/engine.go` and `internal/engine/command_handler.go`.

## Proposed Changes - Part 1: Parallel Cleanup
- Modify `Engine.Clean()` in `internal/engine/engine.go` to use a worker pool (similar to `Scan`).
- Ensure the `ResultAggregator` and `Hooks` remain thread-safe during parallel execution.
- If multiple targets share common cleanup tasks, keep the thread-safe aggregation.

## Proposed Changes - Part 2: Testing Strategy
- **Engine Tests:**
  - Mock the `Cleaner` and `Scanner` to isolate engine logic.
  - Test `ScanAndClean` flow, ensuring hooks are called in the correct order.
- **CommandHandler Tests:**
  - Introduce an interface for command execution to allow mocking `exec.Command`.
  - Test `Handle` logic, including `scheduler` interaction and hook callbacks.

## Verification
- Run tests with `-race` detection to ensure concurrency safety.
- Verify cleanup performance improvement on multi-target configurations.
