# Memory Index

## Current Phase: Implementation & Refinement
The project is currently in the **Development & Feature Refinement** phase. We have successfully scaffolded the core components, implemented the scanner, cleaner, and scheduler, and recently integrated the full cleanup inventory.

## Completed Tasks
- **System Architecture**: Established Go project structure (`cmd`, `internal/scanner`, `internal/cleaner`, `internal/config`, `internal/scheduler`).
- **Configuration Management**: Implemented YAML-based config with `Viper`. Fixed config creation bug and added support for granular item types (`file`, `folder`, `both`) per cleanup target.
- **Scanner**: Implemented robust, concurrent directory traversal using `os.ReadDir` for maximum efficiency (compliant with `AGENTS.md`). Supports both individual file pruning and entire directory removal based on configuration.
- **Cleaner**: Implemented safe file deletion with dry-run support by default. Handles both file and recursive directory removal.
- **Logging**: Refined CLI output with concise summaries and `--verbose` support for detailed match listing.
- **Config CLI**: Enhanced config management with `mls config open` (default app) and `mls config reveal` (Finder).
- **Scheduler**: Implemented a periodic task scheduler with persistent last-run state tracking.
- **Cleanup Inventory**: Populated configuration with the full list of safe cleanup targets.
- **Ignore Patterns**: Implemented a performant glob-based ignore system (global and target-specific) to exclude macOS metadata. Optimized directory walking to avoid unnecessary errors and refactored logging.
- **Bug Fix**: Resolved size discrepancy between `scan` and `clean` commands. Prevented double-counting in `Scanner` for stale folders and implemented accurate recursive directory size calculation in `Cleaner`. Harmonized ignore patterns across both components for exact total size matching. Verified with `cmd/size_consistency_test.go`.
- **UI/UX Enhancement**: Integrated `fatih/color` for vibrant output. Implemented an "Ultra-Compact" CLI mode: individual matches are suppressed in the console, providing only target-level summaries. For Dry Runs, individual deletions are only listed if the total count is <= 20; otherwise, the user is redirected to the full audit log at `/tmp/mls-last-run.log`. This ensures maximum CLI performance while maintaining a high-fidelity audit trail.
- **Compliance**: Restored architectural compliance to `AGENTS.md` mandates, including `Prompts.log` and `ACTIONS.log` enforcement.
- **Documentation Pass**: Completed a comprehensive GoDoc documentation pass for core components:
    - `cmd/processor.go`, `cmd/colors.go`
    - `internal/scanner/scanner.go`
    - `internal/cleaner/cleaner.go`
    - `internal/config/config.go`
    - `internal/scheduler/scheduler.go`
    All public and private functions, types, and significant variables now have clear, concise GoDoc-compliant comments.

## Pending Tasks
- Add more robust error handling and telemetry (Zap).
- Perform stress testing on directory traversal (large caches).
- Future: Add automated release/build process (Cobra/Go).

## Decisions & Notes
- **Glob Support**: Scanner now uses `filepath.Glob` to handle paths like `.../User Data/*/Cache`.
- **Catch-up Logic**: Scheduler tracks last successful run via `~/.MacosLeanStorage.lastrun` and checks for execution on startup if the last run was > 23 hours ago.
- **Safety First**: Dry-run mode remains the default for all cleanup operations.
- **Recursive Globbing**: Enabled via `doublestar` library, allowing `**` patterns for deep directory recursion.

