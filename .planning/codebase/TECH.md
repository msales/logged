# Technology Stack

## Language & Runtime
Go 1.22

## Dependencies
| Dependency | Purpose |
|---|---|
| `github.com/stretchr/testify` v1.9.0 | Testing assertions |

This library has minimal dependencies by design -- it is a foundational logging package.

## Build System
- **Makefile:** Includes shared targets from `github.com/msales/make/golang`.
- **Dockerfile:** `msales/go-builder:1.22-base-1.0.1` -- compiles to verify the library builds. Uses `GITHUB_TOKEN` for private module access.
- No binary output -- this is a library.

## Testing
- Standard `testing` package with `testify/assert`.
- Internal tests (`buffer_internal_test.go`) for unexported types.
- Table-driven tests for formatter output (`format_test.go`).
- Handler tests (`handler_test.go`) for buffered stream, filter, and level filter handlers.
- Logger tests (`logger_test.go`) for all log levels and context normalization.
- Benchmark tests (`bench_test.go`) for formatter and handler performance.

## CI/CD
GitHub Actions workflow (`.github/workflows/test.yaml`):
- Triggers on every push.
- Runs on `ubuntu-latest`.
- Steps: checkout, external `msales/github-actions` go-test action with staticcheck (latest). Race detection disabled.
