# Rigorous Test Plan for Engine and CommandHandler

## 1. Engine Testing Coverage
- **New()**: Verify initialization of `scanner` and `cleaner`.
- **Scan()**:
    - [ ] Sequential vs Parallel target scans (validate hook timing).
    - [ ] Hook order: `OnTargetScanStart` -> `OnTargetScanEnd` for each target.
    - [ ] Error handling: What happens if `scanner.Scan` fails for one target? (Should it continue or abort?)
    - [ ] Empty target list handling.
    - [ ] Partial/Full results verification.
- **Clean()**:
    - [ ] Verify concurrency safety (run with `-race`).
    - [ ] Hook order: `OnFileCleaned` and `OnTargetCleaned` callback sequence.
    - [ ] Handle skipped targets (when `len(res.Files) == 0`).
    - [ ] Verify `ResultAggregator` aggregates unique paths correctly across parallel threads.
- **ProcessCommands()**:
    - [ ] Verify command dispatch order.
    - [ ] Hook order: `BeforeHandleCommand`, `BeforeExecutingCommand`, `AfterExecutingCommand`, `AfterHandleCommand`.
    - [ ] Verify dry-run logic (ensure commands are NOT executed if `DryRun == true`).
- **ScanAndClean()**:
    - [ ] End-to-end flow integration test.

## 2. CommandHandler Testing Coverage
- **Handle()**:
    - [ ] Verify `scheduler` interaction (`ShouldRunCommand` check).
    - [ ] Verify `UpdateCommandRunTime` behavior.
    - [ ] Verify hook sequence and error propagation.
- **ExecuteCommand()**:
    - [ ] Dry-run mode verification (returns `nil` error).
    - [ ] Real execution simulation (mocking `exec.Command` if possible or using `/bin/true`/`/bin/false`).
    - [ ] Error handling for failed command execution.

## 3. Implementation Strategy
- Step 1: Create a mock interface for `Cleaner` and `Scanner` to isolate engine logic.
- Step 2: Implement Test Cases in phases (Phase A: Hook sequence, Phase B: Concurrency/Race tests, Phase C: Error/Edge cases).
- Step 3: Run all with `go test -race`.
