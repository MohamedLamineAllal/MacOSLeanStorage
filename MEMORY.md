# MEMORY.md

## Current Phase
- **Phase 1: Project Initialization & Deep Research**

## Summary of Completed Tasks
- [x] Researched macOS storage locations for Arc, VSCode, Chrome, Discord, Cursor, Antigravity, OpenAI Atlas.
- [x] Defined safety thresholds based on `atime` (Access Time) and application type.
- [x] Planned Go project architecture with a focus on concurrency (Worker Pools) and safety (FLock, Dry-run).
- [x] Initialized Go module and project directory structure.
- [x] Configured `.vscode` with recommended settings and installed Go tools.
- [x] Transitioned to `AGENTS.md` workflow and configured `.gemini/settings.json`.
- [x] Organized and expanded documentation under `docs/`.
- [x] Initialized `Prompts.log` and logged all major directives.

## Pending Tasks for Current Phase
- [ ] Implement concurrent directory scanner in `internal/scanner` (Worker Pool + Channel).
- [ ] Implement multi-profile discovery logic for Chromium-based browsers.
- [ ] Implement cleanup logic in `internal/cleaner` with `os.RemoveAll` and optional `--trash` support.
- [ ] Implement basic CLI commands (`scan`, `clean`) using Cobra.
- [ ] Set up `robfig/cron` in `internal/scheduler` for daemon mode.

## Brainstormed Items & Decisions
- **Decision**: Use `atime` (Access Time) as the primary metric for staleness.
- **Decision**: Multi-profile applications will be handled by auto-discovering directories in `User Data`.
- **Decision**: The tool will prefer `os.ReadDir` for better performance on large directory trees.
- **Decision**: Background tasks will be manageable via `launchd` plist on macOS.
