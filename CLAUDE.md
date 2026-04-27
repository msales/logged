@~/CLAUDE.md

## Project Description

**Language:** Go 1.22  
**Module:** `github.com/msales/logged`  
**Purpose:** Low-level structured logging library providing log levels, formatters (JSON, logfmt), handlers (buffered stream, filtered, discard), and a buffer pool. Used as the engine behind `github.com/msales/log`.  
**Key dependencies:** `github.com/stretchr/testify` (testing).

## Directory Structure

```
(root)          Logger, Handler, Formatter, Buffer, and Pool implementations -- all in a single package
.github/        GitHub Actions CI workflow (test, staticcheck)
```

## Build, Test & Lint

```bash
# Build
go build ./...

# Run tests
go test ./...

# Run tests with verbose output
go test -v ./...

# Vet
go vet ./...

# Static analysis
staticcheck ./...
```

Makefile includes shared targets from `github.com/msales/make/golang`. CI runs on push via GitHub Actions (`test.yaml`): staticcheck (latest) on `ubuntu-latest`, race detection disabled. Docker build (`Dockerfile`) uses `msales/go-builder:1.22-base-1.0.1` for compilation verification.

## Code Style

### Imports

Standard library first, then third-party. Two groups separated by a blank line.

### Naming Conventions

- **Packages:** single package `logged` at root level
- **Interfaces:** `Logger`, `Handler`, `Formatter` -- noun-based, exported
- **Exported/unexported:** Public API includes types (`Logger`, `Handler`, `Formatter`, `Level`, `Format`), constructors (`New`, `BufferedStreamHandler`, `StreamHandler`), and filter functions. Internal types are unexported (`logger`, `bufStreamHandler`, `buffer`, `pool`, `closeHandler`).

### Error Handling

- `LevelFromString` and `FormatterFromString` return descriptive errors for invalid inputs via `fmt.Errorf`.
- `Close()` uses type-assertion checks to conditionally close handlers implementing `io.Closer`.
- Odd-length context slices are auto-normalized with an appended `LOGGED_ERROR` key.

### Testing

- **Framework:** standard `testing` package with `testify/assert`
- **Patterns:** table-driven tests in `format_test.go`, internal tests (`buffer_internal_test.go`), handler tests, logger tests, benchmark tests
- **Mocks:** hand-written test helpers (`CloseableHandler` in `mocks_test.go`)

## Branching & Git

- **Never** work directly on `main` or `master`. Always create a new branch off the latest `main`/`master`.
- Before creating a branch, **always pull** the latest changes from the remote (`git pull origin master`).
- If the user provides a task identifier in the format `TRK-XXXX`, use it as the branch name (e.g., `TRK-1234`). Otherwise, **ask the user** for a task ID or branch name before proceeding.
- **Never push** to `main` or `master` directly.
- Commit changes after each meaningful iteration. Use your judgment to decide when progress should be saved -- prefer smaller, atomic commits over large monolithic ones.
- Write clear, concise commit messages. Use conventional commit format when appropriate (e.g., `feat:`, `fix:`, `refactor:`, `test:`).
