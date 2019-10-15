# OctoLog

*Octolog* is a drop-in replacement for Go's built-in *log* package.

----

## Features

- logging in and filtering by log-levels
- colors *(that are automatically disabled when the output is not a
  terminal)*
- out-of-the-box routing of errors and warnings to STDERR
- customizable outputs & loggers
- out-of-the-box support for concurrent use

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
logger := log.New(
    "myapp", // unique name of logger
    nil,     // nil implies all
)
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
