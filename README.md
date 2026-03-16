[![License](https://img.shields.io/badge/License-BSD%203--Clause-blue.svg)](https://opensource.org/licenses/BSD-3-Clause)
[![GoDoc](https://godoc.org/github.com/octogo/logrouter?status.svg)](https://godoc.org/github.com/octogo/log)
[![Build Status](https://travis-ci.org/octogo/logrouter.svg?branch=master)](https://travis-ci.org/octogo/log)

# OctoLog

Go logging the way I like it.

## Features

- drop-in replacement of the builtin "log" package
- customizable and colorful log-levels
- POSIX compliant routing of warnings and errors to STDERR

----

## Installation

```bash
go get github.com/octogo/log/v2
```

## Quickstart

```go
package main

import "github.com/octogo/log/v2

// monkey-patch os.Exit to avoid pre-matutre end of demo code.
os.Exit = func(code int) {
  fmt.Println("EXIT-CODE:", code)
}

// Drop-in replacement for builtin "log" package.
log.Println("This is a normal log message.")
log.Printf("This is another normal log message via %s.", "Printf")
log.Fatal("This is a fatal error written to STDERR.")
log.Fatalf("This is another fatal error written to STDERR via %s.", "Fatalf")

// But there are several additional log-levels.
log.Debug("This is some debug output.")
log.Notice("This is a notification in GREEN.")
log.Warning("This is a warning in YELLOW on STDERR.")
log.Error("This is an error in RED on STDERR.")

// A log-level is simply a string. Add arbitrary custom log-levels.
log.Log(log.Level("CUSTOM"), "This is a log-entry with a CUSTOM log-level.")

// Use Loggers to log concurrent code.
logger := log.New(
  "My App",
  WantsAllLevels(),
  DefaultFormat,
  DefaultSinks(),
)

goroutineA := logger.NewLogger("")
```
