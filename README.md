[![License](https://img.shields.io/badge/License-BSD%203--Clause-blue.svg)](https://opensource.org/licenses/BSD-3-Clause)
[![GoDoc](https://godoc.org/github.com/octogo/logrouter?status.svg)](https://godoc.org/github.com/octogo/log)
[![Build Status](https://travis-ci.org/octogo/logrouter.svg?branch=master)](https://travis-ci.org/octogo/log) 

# OctoLog

*Octolog* is a drop-in replacement for Go's built-in *log* package.

----

## Features

- logging in and filtering by log-levels
- colors *(that are automatically disabled when the output is not a
  terminal)*
- POSIX compliant routing of WARNINGS and ERRORS to STDERR by default
- customizable outputs & loggers
- can be used across multiple goroutines

----

## Quickstart

Using *OctoLog* in your code is easy:

```bash
go get -u github.com/octogo/log
```

```go
package main

import "github.com/octogo/log"

func main() {
  /// ...
}
```

Before using OctoLog it has to initialize itself. The `Init()` func
will ensure that the default outputs and standard logger are made
available and that the configuration is loaded correctly.

```go
log.Init()
```

Then you can simply use it like you would use the builtin `log`
package:

```go
log.Println("Hello world!")
log.Fatal("FATALITY")
```

Above code produces this output:

```text
2019/10/13 04:02:19 main Hello world!
2019/10/13 04:02:19 main FATALITY
exit status 1
```

If you want more granular control over the log-levels of your messages, simply
use the standard logger or initialize your own `log.Logger{}`.

```go
logger := log.New("myapp", nil)
logger.Debug("Debug message...")
logger.Info("Info...")
logger.Notice("Notification...")
logger.Warning("Warning...")
logger.Error("Error...")
```

```stdout
2019/10/13 04:09:33 myapp Info...
2019/10/13 04:09:33 myapp Notification...
2019/10/13 04:09:33 myapp Warning...
2019/10/13 04:09:33 myapp Error...
```

**Note:**
The log-entry with log-level DEBUG is not shown in the output. That's
because none of the default outputs is configured to log DEBUG level.

----

## Configuration

*Octolog* can easily be configured during run-time.

```go
import "github.com/octogo/log/pkg/config"
config := &config.Config{
  // all internal variables can be set here and then passed
  // to octolog during initialization phase.
}
log.InitWithConfig(config)
```

### Configuration via Simple Textfile

*Octolog* can be configured with a simple *YAML* file that is placed
in the current working directory. This repository includes a tool for
creating configuration file templates.

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
