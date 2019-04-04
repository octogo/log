[![License](https://img.shields.io/badge/License-BSD%203--Clause-blue.svg)](https://opensource.org/licenses/BSD-3-Clause)
[![GoDoc](https://godoc.org/github.com/octogo/logrouter?status.svg)](https://godoc.org/github.com/octogo/log)

# OctoLog

OctoLog is a very thin and flexible asynchronous logging library for Golang.

----

**Features:**

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

import "github.com/octogo/octolog"

func main() {
    router := octolog.New() // Initialize new router
    defer router.Drain()    // Drain all queued logs before terminating

    logger := router.NewLogger("example") // Initialize new logger
    logger.Info("Hello world!")           // Log "Hello World!"
}
```

*For more examples see the
[examples directory](https://github.com/octogo/octolog/blob/master/examples).*

----

## A Hitchhikers Guide to Logging with OctoLog

Using *octolog* in your Go project is easy.
To get started, simply import it.

```go
import "github.com/octogo/octolog"
```

Now you can control the default configuration values.

Remove all metadata fields from log-lines like this:

```go
func main() {
    octolog.DefaultLogFormat = "{{.Body}}"
}
```

### The Router

Next, you can obtain a new *Router*.
The router is the core unit or engine of *octolog*.
It essentiall spawns a dedicated go-routine for logging and provides
methods for interfacing with it, such as:

- managing backends
- spawning new loggers
- routing entries from the loggers to the backends

```go
func main() {
    router := octolog.New()
    defer router.Drain()
}
```

Use the `New()` function to obtain a newly initialized router.
Always remember to defer the `Drain()` method to ensure that all
entries are logged before the main go-routine exists.

### The Backends

By default the router will spawn with two FileBackends:

- `os.Stdout` - logs levels `DEBUG`, `INFO` and `NOTICE`
- `os.Stderr` - logs levels `ALERT`, `WARNING` and `ERROR`

Use the router's `SetBackends()` and `AddBackends()` methods to
reconfigure the backends to your needs.

Adding a `FileBackend` that will additionally log all entries to a
dedicated log-file is easy:

```go
func main() {
    fileBackend, err := octolog.NewFileBackend(
        format: octolog.DefaultLogFormat,
        levels: octolog.AllLevelSlice(),
        file: *os.File, // Replace with an actual *os.File value
        colorOutput: false,
    )
    if err != nil {
        panic(err)
    }

    router.AddBackends(fileBackend)
}
```

Note the use of the `AddBackends()` method to `add` the backend and
not overwrite the default backends like `SetBackends()` would do.

### The Loggers

Once all backends have been set up, logs can be written in a variety
of different levels using a `Logger`.

```go
func main() {
    logger := router.NewLogger(name: "tutorial")

    logger.Debug("debug")
    logger.Info("info")
    logger.Notice("notice")
    logger.Alert("alert")
    logger.Warning("warning")
    logger.Error("error")
}
```

There is also an equivalent formatter method for each of the above:

```go
func main() {
    logger.Debugf("Debug: %s", "debug")
    logger.Infof("Info: %s", "info")
    logger.Noticef("Notice: %s", "notice")
    logger.Alertf("Alert: %s", "alert")
    logger.Warningf("Warning: %s", "warning")
    logger.Errorf("Error: %s", "error")
}
```

----

### Levels

The supported log-levels and their intended uses are:

- **ERROR** - fatal condition
- **WARNING** - fatal condition approaching
- **ALERT** - heads-up alert
- **NOTICE** - heads-up
- **INFO** - plain output
- **DEBUG** - debug information

Additional log-levels may be introduced

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
- **{{.Body}}** - the actual log message

----
