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
	"sync"

	"github.com/octogo/log/v2"
)

func main() {
package main

import (
	"fmt"
	"sync"

	"github.com/octogo/log/v2"
	"github.com/octogo/log/v2/ansii"
)

func main() {
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
	CUSTOM, err := log.AddLevel("CUSTOM", ansii.MAGENTA)
	if err != nil {
		log.Fatal(err)
	}

	l := log.DefaultLogger
	l.SetWants(log.WantsAllLevels())
	l.Sinks[0].SetWants(log.WantsAllLevels()) // sink[0] is STDOUT
	l.Log(CUSTOM, "This is a log-entry with a CUSTOM log-level.")

	// You can route logs to files and filter for log-levels, while doing so.
	l.AddSinks(
		log.WithSink("test.custom", log.WantsLevels(CUSTOM)),
		log.WithSink("test.log", log.WantsLevels(log.INFO, log.NOTICE)),
		log.WithSink("test.err", log.WantsLevels(log.WARNING, log.ERROR)),
	)
	l.Log(CUSTOM, "This is a log-entry with a CUSTOM log-level on STDOUT and in `test.custom`.")
	l.Log(log.INFO, "This is an INFO on STDOUT and in `test.log`.")
	l.Log(log.ERROR, "This is an ERROR on STDOUT and in `test.err`.")

	// Use Loggers to log concurrent code.
	l = log.New(
		"MyApp",
		log.WantsAllLevels(),
		log.DefaultFormat,
		log.DefaultSinks()...,
	)
	l.Println("Hello from MyApp!")

	wg := sync.WaitGroup{}

	wg.Go(func() {
		l := l.NewLogger("Child A")
		l.Println("Hello from A!")
	})

	wg.Go(func() {
		l := l.NewLogger("Child B")
		l.Println("Hello from B!")
	})

	wg.Wait()
}
```
