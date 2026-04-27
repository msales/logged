# Architecture

## Overview
`logged` is a structured logging engine providing the core building blocks for the `msales` logging stack. It defines log levels, message formatters, output handlers, and a high-performance buffer pool. The higher-level `github.com/msales/log` package wraps this library with a global logger API.

## Package Structure
Single package `logged` at the root. All types live in one package:

| File | Responsibility |
|---|---|
| `logger.go` | `Logger` interface, `Level`/`Format` types, `New()` constructor, context normalization/merging |
| `handler.go` | `Handler` interface, `BufferedStreamHandler`, `StreamHandler`, `FilterHandler`, `LevelFilterHandler`, `DiscardHandler` |
| `format.go` | `Formatter` interface, `JSONFormat()`, `LogfmtFormat()`, string escaping utilities |
| `buffer.go` | `buffer` type (byte slice wrapper with append helpers), `pool` (sync.Pool-based buffer pooling) |

## Key Interfaces
- **`Logger`** -- `Debug`, `Info`, `Warn`, `Error`, `Crit` methods + `Close()`. Created via `New(handler, ctx...)`.
- **`Handler`** -- `Log(msg, lvl, ctx)`. Implementations: `bufStreamHandler` (buffered async writer), `HandlerFunc` (adapter), `closeHandler` (wraps handler + closer).
- **`Formatter`** -- `Format(msg, lvl, ctx) []byte`. Implementations: `JSONFormat()`, `LogfmtFormat()`.

## Data Flow
1. Caller invokes `logger.Info("msg", "key", "val")`.
2. Logger normalizes context (ensures even number of key-value pairs), merges with logger-level context.
3. Logger delegates to its `Handler.Log(msg, lvl, ctx)`.
4. If using `BufferedStreamHandler`: message is formatted via the `Formatter`, written to a pooled buffer, and when the buffer exceeds `flushBytes` or the `flushInterval` ticks, the buffer is swapped and flushed to the underlying `io.Writer` via a background goroutine.
5. `LevelFilterHandler` wraps a handler and drops messages above the configured level.

## External Dependencies
No external services. This is a pure library with zero runtime dependencies beyond the Go standard library.
