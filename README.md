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

import (
	"fmt"

	"github.com/octogo/log/v2"
	"github.com/octogo/log/v2/ansii"
)

func main() {
	// set a custom ExitHandler to avoid pre-mature end of demo code.
	log.ExitHandler(func(code int) {
		fmt.Println("EXIT-CODE:", code)
	})

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

	// Log-level are simply strings. Add arbitrary custom log-levels and they
	// will automatically be picked up by the DefaultLogger and the STDOUT sink.
	CUSTOM := log.AddLevel("CUSTOM", ansii.MAGENTA)

	// Logging custom log-levels can only be done through a Logger.
	// The simplest case is to use the DefaultLogger, which automatically
	// regards all custom log-levels.
	logger := log.DefaultLogger
	logger.Log(CUSTOM, "This is a log-entry with a CUSTOM log-level.")

	// When creating custom Loggers, make sure to register your custom
	// log-levels, otherwise they will be disregarded.
	logger = log.New(
		"MyCustomLogger",
		log.WantsLevels(CUSTOM),
		log.DefaultSinks()...,
	)
	logger.Log(CUSTOM, "Hello from MyCustomLogger!")

	// Create arbitrary custom Loggers as you please...
	logger = log.New(
		"MyApp",
		 log.WantsAllLevels(),  // includes custom log-levels
	)
	logger.Println("Hello from MyApp!")

	// ... add child Loggers for safely logging concurrent code in a structured
	// hierarchical tree.
	childA := logger.NewLogger("A")
	childB := logger.NewLogger("B")
	childA.Println("Hello from child A!")
	childB.Println("Hello from child B!")

	// Children inherit Wants() and Sinks() from their parent, at creation.
	// Therefore it makes sense to configure the parent before spawning child
	// Loggers.
	logger = log.New(
		"MyApp",
		log.WantsAllLevels(),
		append(
			log.DefaultSinks(),
			log.WithSink("myapp.log", log.WantsLevels(log.INFO, log.NOTICE)),
			log.WithSink("myapp.err", log.WantsLevels(log.WARNING, log.ERROR)),
			log.WithSink("myapp.debug", log.WantsLevels(log.DEBUG)),
			log.WithSink("myapp.custom", log.WantsLevels(CUSTOM)),
		)...,
	)
	logger.Log(log.DEBUG, "Hello on STDOUT and in myapp.debug")
	logger.Log(log.INFO, "Hello on STDOUT and in myapp.log")
	logger.Log(log.NOTICE, "Hello on STDOUT and in myapp.log")
	logger.Log(log.WARNING, "Hello on STDERR and in myapp.err")
	logger.Log(log.ERROR, "Hello on STDERR and in myapp.err")
	logger.Log(CUSTOM, "Hello on STDOUT and in myapp.custom")
}

```

Outputs the following logs:

```
1970-01-01 00:00 main [INFO] This is a normal log message.
1970-01-01 00:00 main [INFO] This is another normal log message via Printf.
1970-01-01 00:00 main [ERROR] This is a fatal error written to STDERR.
EXIT-CODE: 1
1970-01-01 00:00 main [ERROR] This is another fatal error written to STDERR via Fatalf.
EXIT-CODE: 1
1970-01-01 00:00 main [DEBUG] This is some debug output.
1970-01-01 00:00 main [NOTICE] This is a notification in GREEN.
1970-01-01 00:00 main [WARNING] This is a warning in YELLOW on STDERR.
1970-01-01 00:00 main [ERROR] This is an error in RED on STDERR.
1970-01-01 00:00 main [CUSTOM] This is a log-entry with a CUSTOM log-level.
1970-01-01 00:00 MyCustomLogger [CUSTOM] Hello from MyCustomLogger!
1970-01-01 00:00 MyApp [INFO] Hello from MyApp!
1970-01-01 00:00 MyApp.A [INFO] Hello from child A!
1970-01-01 00:00 MyApp.B [INFO] Hello from child B!
1970-01-01 00:00 MyApp [DEBUG] Hello on STDOUT and in myapp.debug
1970-01-01 00:00 MyApp [INFO] Hello on STDOUT and in myapp.log
1970-01-01 00:00 MyApp [NOTICE] Hello on STDOUT and in myapp.log
1970-01-01 00:00 MyApp [WARNING] Hello on STDERR and in myapp.err
1970-01-01 00:00 MyApp [ERROR] Hello on STDERR and in myapp.err
1970-01-01 00:00 MyApp [CUSTOM] Hello on STDOUT and in myapp.custom
```

The following files are created:

```
File: myapp.custom
1970-01-01 00:00 MyApp [CUSTOM] Hello on STDOUT and in myapp.custom

File: myapp.debug
1970-01-01 00:00 MyApp [DEBUG] Hello on STDOUT and in myapp.debug

File: myapp.err
1970-01-01 00:00 MyApp [WARNING] Hello on STDERR and in myapp.err
1970-01-01 00:00 MyApp [ERROR] Hello on STDERR and in myapp.err

File: myapp.log
1970-01-01 00:00 MyApp [INFO] Hello on STDOUT and in myapp.log
1970-01-01 00:00 MyApp [NOTICE] Hello on STDOUT and in myapp.log
```