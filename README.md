[![License](https://img.shields.io/badge/License-BSD%203--Clause-blue.svg)](https://opensource.org/licenses/BSD-3-Clause)
[![GoDoc](https://godoc.org/github.com/octogo/logrouter?status.svg)](https://godoc.org/github.com/octogo/log)
[![Build Status](https://travis-ci.org/octogo/logrouter.svg?branch=master)](https://travis-ci.org/octogo/log) 

# OctoLog

*Octolog* is a Go package providing an interface for
application-level logging.

## Features

- works out-of-the-box
- POSIX compliant routing of warnings and errors to STDERR
- logging in and filtering by log-levels
- straight-forward configuration, optionally via YAML file
- colors *(that are automatically disabled when the output is not a
  terminal)*
- customizable outputs & loggers

----

## Quickstart

Use the included CLI tool to create a Go source file:

```bash
go install github.com/octogo/log/cmd/octolog

octolog gensrc -h
```

## Installation

```bash
go get github.com/octogo/log
```

## Usage

```go
package main

import "github.com/octogo/log"

func main() {
  /// ...
}
```

Before using OctoLog it has to initialize itself. The `Init()` func
will ensure that the default outputs and standard logger are made
available and that the configuration is loaded correctly, if present.

```go
log.Init()
```

Then you can simply use it like you would use the builtin `log`
package:

```go
log.Log("Hello world!")
log.Fatal("FATALITY")
```

Above code produces this output:

```text
2019/10/13 04:02:19 main INFO Hello world!
2019/10/13 04:02:19 main ERROR FATALITY
exit status 1
```

If you want more granular control over the log-levels of your messages, simply
use the standard logger or initialize your own `log.Logger{}`.

```go
logger := log.New(
  "myapp",  // unique name of the logger
  nil,      // []level.Level of log-levels to whitelist (nil implies *all*)
  // if no Outputs are specified, the logger will be initialized with the
  // DefaultOutputs set as its outputs. But you could easily configure the
  // Logger to log into a custom log file by specifying 'file://my.log' here.
  )
logger.Debug("Debug message...")
logger.Info("Info...")
logger.Notice("Notification...")
logger.Warning("Warning...")
logger.Error("Error...")
```

```stdout
2019/10/13 04:09:33 myapp INFO Info...
2019/10/13 04:09:33 myapp NOTICE Notification...
2019/10/13 04:09:33 myapp WARNING Warning...
2019/10/13 04:09:33 myapp ERROR Error...
```

**Note:**
The log-entry with log-level DEBUG is not shown in the output. That's
because none of the default outputs is configured to log DEBUG level.
See *Configuration section* below for more details.

### Gotta log them ALL

The function signatures do not force you to log strings:

```go
// signature of log.Log
func Log(interface{}) {}
// signature of log.Logf
func Logf(string, ...interface{}) {}
// signature of log.Logger.Error
func (Logger) Error(interface{}) {}
// signature of log.Logger.Errorf
func (Logger) Errorf(string ...interface{}) {}
// and so on...
```

Of course, you can log simple strings too, if you like.

### Redaction

Sometimes it is desireable to redact certain parts of the log-messages,
because they would perhaps disclose secrets.
If a logged value satisfies the `log.Redactor` interface, the
return-value of its `Redacted()` function will be logged instead of its native
string representation.

```go
// Redactor is defined as something providing a Redacted() function.
type Redactor interface {
  Redacted() string
}
```

See `examples/redacted/main.go` for more information.

----

## Configuration

*OctoLog* can easily be configured during run-time.

There is a special initialization phase during start-up that takes care of
loading a possibly existing configuration file, but amost everything can
easiliy be configured during run-time, even after initialization phase.

```go
import "github.com/octogo/log/pkg/config"

log.InitWithConfig(&config.Config{
  // all internal variables can be set here and then passed
  // to octolog during initialization phase.
})
```

See `pkg/config/config.go` for more information.

----

### Configuration via Simple Textfile

*OctoLog* can be configured with a simple *YAML* file that is placed
in the current working directory. This enables users of prebuilt binaries
to configure logging without having to rebuild the Go source.

This repository includes a tool for creating configuration file templates with
lots of comments and examples that work out-of-the-box.

Simply run:

```text
go install github.com/octogo/log/cmd/octolog
```

You can then easily create a configuration file in your current
working directory by calling:

```
octolog genconf
```

*See `octolog -h` for usage details.*
