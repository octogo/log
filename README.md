[![License](https://img.shields.io/badge/License-BSD%203--Clause-blue.svg)](https://opensource.org/licenses/BSD-3-Clause)
[![GoDoc](https://godoc.org/github.com/octogo/logrouter?status.svg)](https://godoc.org/github.com/octogo/log)
[![Build Status](https://travis-ci.org/octogo/logrouter.svg?branch=master)](https://travis-ci.org/octogo/log) 

# OctoLog

Package `log` is a very thin and flexible asynchronous logging library for Golang.
It is thread-safe and can be used from many different goroutines simultaneously.

----

**Features:**

- [x] asychronous logging from multiple goroutines
- [x] logging to arbitrary files, to syslog, to custom backends
- [x] custom log formats
- [x] support for backend reconfiguration during run-time
- [x] ANSI terminal colors *(automaticly disabled if file is not a TTY)*

## Usage

### Installation

```bash
go get -u github.com/octogo/log
```

### Quickstart

```go
package main

import "github.com/octogo/log"

func main() {
    defer log.Close()   // Drain all queued logs before terminating

    log.Println("Hello world!")  // goes to stdout
    log.Fatal("bye")             // goes to stderr
}
```

*For more examples see the
[examples directory](https://github.com/octogo/log/blob/master/examples).*

----

## A Hitchhikers Guide to Logging with OctoLog

Using *octolog* in your Go project is easy.
To get started, simply import it.

```go
import "github.com/octogo/log"
```

## Logging

Simple logging is accomplished by using the `Println()`, `Printf()`, `Fatal()` and
`Fatalf()` functions.

```go
func main() {
    defer log.Close()

    log.Printf("Hello %s!", "world")
    log.Fatalf("kthx%s", "bye")
}
```

### The Backends

By default the library will spawn with two default backends:

- `os.Stdout` - logs levels `DEBUG`, `INFO` and `NOTICE`
- `os.Stderr` - logs levels `ALERT`, `WARNING` and `ERROR`

Use the `SetBackends()` and `AddBackends()` methods to reconfigure the backends
to your needs.

Adding a `FileBackend` that will additionally log all entries to a
dedicated logfile is easy:

```go
func main() {
    defer log.Close()

    logFile, err := os.OpenFile("out.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
    if err != nil {
        log.Fatal(err)
    }

    fileBackend, err := octolog.NewFileBackend(
        format: octolog.DefaultLogFormat,
        levels: octolog.AllLevelSlice(),
        file: logFile,
        colorOutput: false,
    )
    if err != nil {
        log.Fatal(err)
    }

    log.AddBackends(fileBackend)
}
```

Note the use of the `AddBackends()` method to `add` the backend and
not overwrite the default backends like `SetBackends()` would do.

### The Loggers

For more fine-grained logging and increasing the overall readability of
log-traces, it is possible to instantiate an arbitrary number of loggers.
A logger is an interface of which many can exist in separate goroutines at
the same time.

```go
func main() {
    logger := log.NewLogger(name: "tutorial")
    logger.Debug("debug")
    logger.Info("info")
    logger.Notice("notice")
    logger.Alert("alert")
    logger.Warning("warning")
    logger.Error("error")

    childLogger := logger.NewLogger("child")
    childLogger.Debug("debug")
    childLogger.Info("info")
    childLogger.Notice("notice")
    childLogger.Alert("alert")
    childLogger.Warning("warning")
    childLogger.Error("error")
}
```

There is also an equivalent formatted method for each of the above:

```go
    logger.Debugf("Debug: %s", "debug")
    logger.Infof("Info: %s", "info")
    logger.Noticef("Notice: %s", "notice")
    logger.Alertf("Alert: %s", "alert")
    logger.Warningf("Warning: %s", "warning")
    logger.Errorf("Error: %s", "error")
```

----

### Levels

The supported log-levels and their intended uses are:

- **ERROR** - fatal condition
- **WARNING** - fatal condition closing in
- **ALERT** - heads-up alert
- **NOTICE** - heads-up notice
- **INFO** - informational output
- **DEBUG** - debug information

### Formatting

The `LogFormat` is a Go template string with support for the
following fields:

- **{{.GID}}** - global router ID
- **{{.LID}}** - local logger ID
- **{{.PID}}** - Process ID
- **{{.PPID}}** - parent's PID
- **{{.Logger}}** - name of the logger the log came through
- **{{.Level}}** - the log level as upper-case string
- **{{.LevelNum}}** - the log level as logrouter.Level
- **{{.LevelLetters}}** - n letters of the log level
- **{{.File}}** - the shortened path of the file that logged
- **{{.LongFile}}** - the absolute path of the file that logged
- **{{.Line}}** - the number of the line in the file that logged
- **{{.Pkg}}** - the shortened package name
- **{{.LongPkg}}** - the long package name
- **{{.Func}}** - the name of the function that logged
- **{{.Date}}** - the date of the log (2006-01-02)
- **{{.Time}}** - the timestamp of the log (15:05:06)
- **{{.TimeExact}}** - the timestamp of the log (15:05:06.000)
- **{{.Msg}}** - the actual log message

----
