# MacosLeanStorage (mls) - Project Instructions

## Tech Stack
- **Language:** Go 1.26+
- **CLI Framework:** Cobra
- **Configuration:** Viper
- **Logging:** Uber-zap
- **Testing:** Standard library `testing` with `testify` for assertions.

## Code Style
- Follow standard Go conventions (Uber Go Style Guide as a reference).
- Use `gofmt` for formatting.
- **Naming:** 
    - Use CamelCase for public symbols.
    - Keep names concise but descriptive.
- **Error Handling:**
    - Always handle errors.
    - Wrap errors with context where helpful using `fmt.Errorf("context: %w", err)`.
- **Concurrency:**
    - Use goroutines and channels for parallel scanning.
    - Ensure thread safety in the `Cleaner` and `Scanner`.
- **Logging:**
    - Use `zap.Logger` for structured logging.
    - Log levels: `Info` for general progress, `Debug` for detailed info, `Error` for failures.
- **Documentation:**
    - All public functions and types must have comments.
    - Maintain `ARCHITECTURE.md` for high-level design.

## Development Workflow
- **Research -> Strategy -> Execution** lifecycle.
- **Plan -> Act -> Validate** for each sub-task.
- Incremental commits: One feature/fix per commit.
- Run tests before committing.

## Linter Configuration
- Use `golangci-lint`.
- Configured in `.golangci.yml`.
