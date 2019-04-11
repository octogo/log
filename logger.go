package log

import (
	"fmt"
	"os"
	"runtime"
	"sync/atomic"
	"time"
)

// Logger defines the interface of a octolog logger.
type Logger struct {
	name    string
	nextLID uint64
}

// NewLogger returns a *Logger with the given name.
// Logger names can increase the readability of log-traces, when configured to
// be included in the backend's output format.
//
// Note that this implementation doesn't check for duplicate names, as names
// have no other functionality or logic behind them other than being displayed.
//
// A child Logger can be obtained by calling the NewLogger method on any
// already existing Logger. Child loggers have the name of their parent
// prepended to their own names.
func NewLogger(name string) *Logger {
	return &Logger{name: name}
}

func (logger *Logger) log(level Level, format string, args ...interface{}) {
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
		gid:       NextGID(),
		lid:       atomic.AddUint64(&logger.nextLID, 1),
		logger:    *logger,
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

	Log(entry)
}

// Name returns the name of this logger.
func (logger *Logger) Name() string {
	return logger.name
}

// Debug logs a debug message.
func (logger *Logger) Debug(args ...interface{}) {
	logger.log(DEBUG, "", args...)
}

// Debugf logs a formatted debug message.
func (logger *Logger) Debugf(format string, args ...interface{}) {}

// Info logs an info message.
func (logger *Logger) Info(args ...interface{}) {
	logger.log(INFO, "", args...)
}

// Infof logs a formatted info message.
func (logger *Logger) Infof(format string, args ...interface{}) {
	logger.log(INFO, format, args...)
}

// Notice logs a notice.
func (logger *Logger) Notice(args ...interface{}) {
	logger.log(NOTICE, "", args...)
}

// Noticef logs a formatted notice.
func (logger *Logger) Noticef(format string, args ...interface{}) {
	logger.log(NOTICE, format, args...)
}

// Alert logs an alert.
func (logger *Logger) Alert(args ...interface{}) {
	logger.log(ALERT, "", args...)
}

// Alertf logs a formatted alert.
func (logger *Logger) Alertf(format string, args ...interface{}) {
	logger.log(ALERT, format, args...)
}

// Warning logs a warning message.
func (logger *Logger) Warning(args ...interface{}) {
	logger.log(WARNING, "", args...)
}

// Warningf logs a formatted warning message.
func (logger *Logger) Warningf(format string, args ...interface{}) {
	logger.log(WARNING, format, args...)
}

// Error logs an error message.
func (logger *Logger) Error(args ...interface{}) {
	logger.log(ERROR, "", args...)
}

// Errorf logs a formatted error message.
func (logger *Logger) Errorf(format string, args ...interface{}) {
	logger.log(ERROR, format, args...)
}

// Panic logs a critical message and then panics.
func (logger *Logger) Panic(args ...interface{}) {
	logger.log(ERROR, "", args...)
	panic(fmt.Sprint(args...))
}

// Panicf logs a formatted critical message and then panics.
func (logger *Logger) Panicf(format string, args ...interface{}) {
	logger.log(ERROR, format, args...)
	panic(fmt.Sprintf(format, args...))
}

// Fatal logs a critical log message and then exists with return-code 1.
func (logger *Logger) Fatal(args ...interface{}) {
	logger.log(ERROR, "", args...)
	os.Exit(1)
}

// Fatalf logs a formatted critical message and then exists with return-code 1.
func (logger *Logger) Fatalf(format string, args ...interface{}) {
	logger.log(ERROR, format, args...)
	os.Exit(1)
}

// NewLogger returns a new logger with the given name and this Logger's name
// as a name prefix (i.e. thisLoggersName + "." + name).
//
// Note that this implementation doesn't check for duplicate names, as names
// have no other functionality or logic behind them other than being displayed.
//
// Child loggers automatically have the name of their parent prepended to their
// own name.
func (logger *Logger) NewLogger(name string) *Logger {
	return &Logger{name: logger.name + "." + name}
}
