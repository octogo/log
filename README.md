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

```go
import "github.com/octogo/log"

func main() {
    log.Init()
    log.Println("Hello world!")
    log.Fatal("FATALITY")
}
```

```stdout
2019/10/13 04:02:19 main Hello world!
2019/10/13 04:02:19 main FATALITY
exit status 1
```

Using a *Logger* it is possible to log in various levels.

```go
func main() {
    log.Init()
    logger := log.New(
        "myapp", // unique name of logger
        nil,        // nil implies all
    )
    logger.Debug("Debug message...")
    logger.Info("Info...")
    logger.Notice("Notification...")
    logger.Warning("Warning...")
    logger.Error("Error...")
}
```

```stdout
2019/10/13 04:09:33 myapp Debug message...
2019/10/13 04:09:33 myapp Info...
2019/10/13 04:09:33 myapp Notification...
2019/10/13 04:09:33 myapp Warning...
2019/10/13 04:09:33 myapp Error...
```

----

## Configuration

*Octolog* can be configured with a simple *YAML* file that is placed
in the current working directory.

**Example:** `logging.yml`

```yaml
# rootlogger defines the name of the standard logger
# default: main
rootlogger: 'main'

# default format defines the default log-format for all
# outputs that habe no expicit log-format in the
# configuration.
#
# supported formatting labels:
# {{.Date}}   - the date formatted as 2006/01/02
# {{.Time}}   - the time formatted as 15:04:05
# {{.Milli}}  - the time formatted as .000
# {{.Nano}}   - the time formatted as .000000
# {{.PID}}    - the process' ID
# {{.PPID}}   - the partent-process' ID
# {{.GID}}    - the global message ID
# {{.LID}}    - the logger message ID
# {{.Logger}} - the name of the logger
# {{.Level}}  - the log-level of the entry
# {{.Func}}   - the name of the calling function
# {{.File}}   - the source file of the calling function
# {{.Line}}   - the line in the above source file
#
# default: '{{.Date}} {{.Time}} {{.Level}} {{.Message}}'
defaultformat: '{{.Date}} {{.Time}} {{.Level}} {{.Message}}'

# defaultoutputs defines the default outputs for every logger
# that has no explicit outputs configured.
# default: [ 'file:///dev/stdout', 'file:///dev/stderr' ]
defaultoutputs:
  - 'file:///dev/stdout'
  - 'file:///dev/stderr'

# outputs defines the outputs that should automatically be
# initialize upon startup.
#
# an output is defined as:
#   {
#     url:      # the URL, i.e. file://octo.log
#     wants:    # the list of log-levels to log
#               # providing no log-levels implies `all`
#     format:   # log-format to use, if not the global default
#   }
outputs:
  # log INFO and NOTICE to STDOUT
  - url: 'file:///dev/stdout'
    wants: [ INFO, NOTICE ]
  # log WARNING and ERROR to STDERR
  - url: 'file:///dev/stderr'
    wants: [ WARNING, ERROR ]

# loggers defines the loggers that should automatically be
# initialized upon startup.
#
# a logger is defined as:
#   {
#     name:     # a unique name
#     wants:    # list of log-levels to log
#               # providing no log-levels implies `all`
#     outputs:  # list of output URLs to communicate with
#               # providing no URLs implies `defaultoutputs`
#   }
loggers:
  - name: main
```
