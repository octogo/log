package octolog

import (
	"fmt"
	"os"
	"runtime"
	"strings"
	"sync/atomic"
	"time"
)

// Logger defines the interface of a octolog logger.
type Logger interface {
	Router() *Router
	Name() string
	NewLogger(name string) Logger
	Debug(...interface{})
	Debugf(string, ...interface{})
	Info(...interface{})
	Infof(string, ...interface{})
	Notice(...interface{})
	Noticef(string, ...interface{})
	Alert(...interface{})
	Alertf(string, ...interface{})
	Warning(...interface{})
	Warningf(string, ...interface{})
	Error(...interface{})
	Errorf(string, ...interface{})
	Panic(...interface{})
	Panicf(string, ...interface{})
	Fatal(...interface{})
	Fatalf(string, ...interface{})
}

// l implements a Logger.
type l struct {
	router  *Router
	name    string
	nextLID uint64
}

func (logger *l) log(level Level, format string, args ...interface{}) {
	var (
		file   string
		line   int
		caller string
	)

	fpcs := make([]uintptr, 1)
	n := runtime.Callers(3, fpcs)
	if n != 0 {
		f := runtime.FuncForPC(fpcs[0] - 1)
		if f != nil {
			file, line = f.FileLine(fpcs[0] - 1)
			caller = f.Name()
		}
	}

	entry := &E{
		gid:       logger.router.NextGID(),
		lid:       atomic.AddUint64(&logger.nextLID, 1),
		logger:    logger,
		timestamp: time.Now(),
		level:     level,
		runtime: struct {
			File string
			Line int
			Func string
		}{
			File: file,
			Line: line,
			Func: caller,
		},
		args:       args,
		formatTmpl: format,
	}

	logger.router.Log(entry)
}

// Router returns the parent router of this logger.
func (logger *l) Router() *Router {
	return logger.router
}

// Name returns the name of this logger.
func (logger *l) Name() string {
	return logger.name
}

// Debug logs a debug message.
func (logger *l) Debug(args ...interface{}) {
	logger.log(DEBUG, "", args...)
}

// Debugf logs a formatted debug message.
func (logger *l) Debugf(format string, args ...interface{}) {}

// Info logs an info message.
func (logger *l) Info(args ...interface{}) {
	logger.log(INFO, "", args...)
}

// Infof logs a formatted info message.
func (logger *l) Infof(format string, args ...interface{}) {
	logger.log(INFO, format, args...)
}

// Notice logs a notice.
func (logger *l) Notice(args ...interface{}) {
	logger.log(NOTICE, "", args...)
}

// Noticef logs a formatted notice.
func (logger *l) Noticef(format string, args ...interface{}) {
	logger.log(NOTICE, format, args...)
}

// Alert logs an alert.
func (logger *l) Alert(args ...interface{}) {
	logger.log(ALERT, "", args...)
}

// Alertf logs a formatted alert.
func (logger *l) Alertf(format string, args ...interface{}) {
	logger.log(ALERT, format, args...)
}

// Warning logs a warning message.
func (logger *l) Warning(args ...interface{}) {
	logger.log(WARNING, "", args...)
}

// Warningf logs a formatted warning message.
func (logger *l) Warningf(format string, args ...interface{}) {
	logger.log(WARNING, format, args...)
}

// Error logs an error message.
func (logger *l) Error(args ...interface{}) {
	logger.log(ERROR, "", args...)
}

// Errorf logs a formatted error message.
func (logger *l) Errorf(format string, args ...interface{}) {
	logger.log(ERROR, format, args...)
}

// Panic logs a critical message and then panics.
func (logger *l) Panic(args ...interface{}) {
	logger.log(ERROR, "", args...)
	panic(fmt.Sprint(args...))
}

// Panicf logs a formatted critical message and then panics.
func (logger *l) Panicf(format string, args ...interface{}) {
	logger.log(ERROR, format, args...)
	panic(fmt.Sprintf(format, args...))
}

// Fatal logs a critical log message and then exists with return-code 1.
func (logger *l) Fatal(args ...interface{}) {
	logger.log(ERROR, "", args...)
	logger.router.Drain()
	os.Exit(1)
}

// Fatalf logs a formatted critical message and then exists with return-code 1.
func (logger *l) Fatalf(format string, args ...interface{}) {
	logger.log(ERROR, format, args...)
	logger.router.Drain()
	os.Exit(1)
}

// NewLogger returns a new child logger.
// The new logger will have the name of this logger automatically prepended and
// all of its configuration copied over to the new logger.
func (logger *l) NewLogger(name string) Logger {
	return &l{
		router: logger.router,
		name:   strings.Join([]string{logger.name, name}, "."),
	}
}
