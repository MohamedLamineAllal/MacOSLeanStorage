# MacosLeanStorage (mls) Design Specifications

This document outlines the low-level design decisions for the `mls` tool, focusing on performance, efficiency, and safety.

## 1. Concurrent Scanner Architecture

### Worker Pool Model
To scan the filesystem without hitting OS file descriptor limits:
- **Task Producer**: Recursively traverses the root targets and pushes directories to a buffered channel.
- **Worker Pool**: A fixed number of workers (e.g., `runtime.NumCPU()`) pull directories from the channel, use `os.ReadDir`, calculate sizes, and check `atime`.
- **Aggregator**: Collects results and builds the cleanup manifest.

### Performance Optimizations
- **Fast Path Selection**: Use `Lstat` to quickly identify symlinks (skip) or sockets.
- **Batched Stats**: Use `os.ReadDir` instead of `filepath.Walk` to reduce system calls by reading multiple directory entries at once.
- **Early Exit**: If a directory's `mtime` hasn't changed since the last "Clean Scan," skip deep inspection (optional, requires state file).

## 2. Daemon & Scheduler Logic

### Systemd / Launchd Integration
Instead of a long-running Go process that idles:
- **Launchd**: On macOS, we will provide a `com.mls.daemon.plist` to trigger the tool at specific intervals (e.g., every 24 hours) or when the system is idle.
- **Go "Daemon" Mode**: A lightweight mode that runs as a persistent process if launchd is not preferred, using `robfig/cron`.

### "Check-in" Logic
- **State File**: Store the last scan timestamp in `~/Library/Application Support/mls/state.json`.
- **Notification**: Integrate with macOS `terminal-notifier` or native `NSUserNotification` to alert the user before/after large cleanups.

## 3. Safety Mechanisms

### Safe-Delete (Trash vs. Unlink)
- **Default**: Use `os.RemoveAll` for 100% safe cache directories.
- **Option**: Provide a `--trash` flag to move items to the macOS Trash (`~/.Trash`) instead of permanent deletion.

### Process Locking
- **FLock**: Ensure only one instance of the scanner is running to prevent race conditions on shared cache files.

## 4. Library Selections (Refined)

- **Scanner**: Standard Library (`os`, `path/filepath`, `sync`).
- **Daemon**: `robfig/cron/v3`.
- **Notifications**: `github.com/0xAX/notificator`.
- **Progress Bar**: `github.com/schollz/progressbar/v3` (for CLI feedback during scan).
- **Concurrency**: `golang.org/x/sync/errgroup` for easier error handling in worker pools.
