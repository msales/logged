# Code Quality

## Test Coverage
- Good coverage across all major components: logger, handlers, formatters, and buffer.
- Internal tests (`buffer_internal_test.go`) verify unexported buffer pool behavior.
- `format_test.go` covers JSON and logfmt formatters with various value types.
- `handler_test.go` covers buffered stream handler, filter handler, level filter handler, discard handler, and stream handler.
- `logger_test.go` covers all log levels, context normalization, close behavior, and closeable handler delegation.
- `bench_test.go` includes benchmarks for JSON/logfmt formatting and buffered/unbuffered handler throughput.

## Code Patterns
- **Functional adapters:** `HandlerFunc` and `FormatterFunc` allow plain functions to satisfy interfaces.
- **sync.Pool-based buffer pooling:** Separate pools for JSON and logfmt formatters minimize allocation pressure.
- **Buffered async I/O:** `BufferedStreamHandler` uses a background goroutine with channel-based buffer swapping for high-throughput logging without blocking callers.
- **Context normalization:** Odd-length context slices are auto-corrected with a `LOGGED_ERROR` sentinel key.
- **unsafe.Pointer for zero-copy string conversion:** `buffer.String()` uses `unsafe.Pointer` to avoid allocation.

## Error Handling
- `LevelFromString` and `FormatterFromString` return descriptive errors for unrecognized inputs.
- `Close()` checks for `io.Closer` via type assertion before attempting cleanup.
- `closeHandler` wrapper ensures close propagation through handler chains.

## Documentation
- All exported types, interfaces, and functions have godoc comments.
- Constants (`LevelKey`, `MessageKey`, time format) are documented.
- `LICENCE` file present (MIT-style).
- `README.md` provides usage examples and API overview.
