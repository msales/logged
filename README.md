# logged

[![Go Report Card](https://goreportcard.com/badge/github.com/msales/logged)](https://goreportcard.com/report/github.com/msales/logged)
[![Build Status](https://travis-ci.org/msales/logged.svg?branch=master)](https://travis-ci.org/msales/logged)
[![Coverage Status](https://coveralls.io/repos/github/msales/logged/badge.svg?branch=master)](https://coveralls.io/github/msales/logged?branch=master)
[![GoDoc](https://godoc.org/github.com/msales/logged?status.svg)](https://godoc.org/github.com/msales/logged)
[![GitHub release](https://img.shields.io/github/release/msales/logged.svg)](https://github.com/msales/logged/releases)
[![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/msales/logged/master/LICENSE)

A fast logger for Go.

## Overview

Install with:

```shell
go get github.com/msales/logged
```

## Examples

```go
// Composable handlers
h := logged.LevelFilterHandler(
    logged.Info,
    logged.StreamHandler(os.Stdout, logged.LogfmtFormat()),
)

// The logger can have an initial context
l := logged.New(h, "env", "prod")

// All messages can have a context
l.Warn("connection error", "redis", conn.Name(), "timeout", conn.Timeout())
```

Will log the message

```
lvl=warn msg="connection error" redis=dsn_1 timeout=0.500
```

## License

MIT-License. As is. No warranties whatsoever. Mileage may vary. Batteries not included.