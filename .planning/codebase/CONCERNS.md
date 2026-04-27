# Concerns & Risks

## Technical Debt
- **unsafe.Pointer usage:** `buffer.String()` uses `unsafe.Pointer` for zero-copy string conversion. While performant, this relies on Go runtime internals and could break with future Go versions.
- **No context.Context support:** The logging API does not accept `context.Context`, which limits integration with tracing and cancellation patterns.
- **Hardcoded time format:** ISO8601 format is hardcoded in formatters with no option to customize.

## Missing Coverage
- No tests for concurrent access to `BufferedStreamHandler` (race detection is disabled in CI).
- No tests for the `Close()` behavior of `bufStreamHandler` under concurrent writes.
- No tests for edge cases in `escapeString` (e.g., multi-byte UTF-8 characters -- the function casts runes to `byte` which may truncate non-ASCII characters).

## Security
- **GITHUB_TOKEN in Dockerfile:** Standard private module pattern but token could leak into Docker layer cache.
- No secrets or credentials in source code.
- Log output should not contain PII -- this is the caller's responsibility, not enforced by the library.

## Performance
- Buffer pooling (`sync.Pool`) minimizes GC pressure for high-throughput logging.
- `BufferedStreamHandler` amortizes I/O with configurable flush size and interval.
- The `escapeString` function processes one rune at a time which may be slower than bulk operations for large strings.
- JSON and logfmt formatters each maintain their own buffer pool (512-byte initial size).
