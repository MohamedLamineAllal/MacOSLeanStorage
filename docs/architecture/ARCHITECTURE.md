# Architecture: MacosLeanStorage

## 1. Project Goal
Build a high-performance, safe, and efficient storage cleanup tool for macOS, focused on developer and browser data.

## 2. Go Project Structure
```text
/cmd
  /mls          - Main CLI entry point
/internal
  /scanner      - Logic for traversing directories and analyzing file metadata
  /cleaner      - Safe deletion logic with dry-run support
  /config       - Configuration management (YAML)
  /scheduler    - Background execution and cron-like logic
  /models       - Shared data structures (Rule, Target, etc.)
/pkg
  /utils        - Shared utilities (size formatting, path expansion)
```

## 3. Design Choices

### Performance & Efficiency
*   **Concurrency**: Use Go routines to scan multiple directories in parallel. A worker pool will handle large directory trees to avoid OS file descriptor limits.
*   **Minimal Footprint**: Avoid loading entire file lists into memory. Use streaming or batching if necessary.
*   **I/O Optimization**: Extensively uses `os.ReadDir` (Go 1.16+) recursively for maximum performance and lower memory overhead compared to `filepath.Walk`.

### Safety
*   **Dry Run**: Every operation defaults to a dry run. The user must explicitly pass a flag to delete.
*   **Immutable Rules**: Rules for "100% safe" directories are hardcoded or verified by a strict schema.
*   **Exclusion List**: Built-in protection for critical system folders (`/System`, `/Library/CoreServices`).

### Optimization
*   **Caching**: Cache the results of a scan if running frequently, though for storage cleanup, fresh data is usually better.
*   **Incremental Scanning**: Use `atime` to quickly skip files that were recently accessed.

## 4. Library Choices
*   **CLI**: `github.com/spf13/cobra` - Industry standard for Go CLIs.
*   **Config**: `github.com/spf13/viper` - Excellent for YAML/JSON and env vars.
*   **Scheduling**: `github.com/robfig/cron/v3` - For background daemon tasks.
*   **Formatting**: `github.com/dustin/go-humanize` - For readable bytes (e.g., "1.2 GB").
*   **Logging**: `go.uber.org/zap` - For high-performance structured logging.

## 5. File/Directory Monitoring
While `fsnotify` is great for real-time, storage cleanup is better suited for **Periodic Polling** (Cron) because:
1.  Cache files are created/deleted constantly; real-time monitoring would be too noisy.
2.  We care about "staleness" over time, not instantaneous change.
3.  Polling every 24 hours is much lighter on CPU/Battery than keeping a file watcher active on `~/Library`.
