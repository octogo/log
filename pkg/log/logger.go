package log

import (
	"fmt"
	"runtime"
	"strings"
	"sync"

	"github.com/octogo/log/pkg/level"
	"github.com/octogo/log/pkg/uid"
)

var defaultLogger *Logger

// Logger is the primary interface for using octolog in other packages.
type Logger struct {
	Name    string
	wants   []level.Level
	Outputs []string
	uid     *uid.UID
	outputs []Output
	mu      *sync.Mutex
}

// NewLogger returns an initialized Logger.
func NewLogger(name string, wants []level.Level, Outputs ...string) *Logger {
	if name == "" {
		name = LoggerName
	}
	if Outputs == nil || len(Outputs) == 0 {
		Outputs = DefaultOutputs
	}
	l := &Logger{
		Name:    name,
		wants:   wants,
		Outputs: Outputs,
		uid:     &uid.UID{},
		mu:      &sync.Mutex{},
	}
	return RegisterLogger(name, l)
}

// NewLogger returns a new child logger.
// ChildLoger will have the name of its parent prepended to its own name.
func (l *Logger) NewLogger(name string) *Logger {
	if name == "" {
		return l
	}
	name = strings.Join([]string{l.Name, name}, ".")
	newLogger := NewLogger(name, l.wants)
	newLogger.Outputs = l.Outputs
	return newLogger
}

func (l *Logger) log(msg string, lvl level.Level) {
	l.mu.Lock()
	defer l.mu.Unlock()
	var (
		caller string
		file   string
		line   int
		fpcs   = make([]uintptr, 1)
		n      = runtime.Callers(3, fpcs)
	)
	if n != 0 {
		f := runtime.FuncForPC(fpcs[0] - 1)
		if f != nil {
			file, line = f.FileLine(fpcs[0] - 1)
			caller = f.Name()
		}
	}

	if l.outputs == nil {
		l.outputs = make([]Output, len(l.Outputs))
		for i := range l.Outputs {
			l.outputs[i] = loadOutput(l.Outputs[i])
		}
	}

	entry := newEntry(msg, l, lvl, caller, file, line)
	for i := range l.outputs {
		_, err := l.outputs[i].Log(entry)
		if err != nil {
			l.Outputs = append(l.Outputs[:i], l.Outputs[i+1:]...)
			l.outputs = append(l.outputs[:i], l.outputs[i+1:]...)
		}
	}
}

// SetWants configures this logger to only accept entries of the given log-level.
func (l *Logger) SetWants(wants []level.Level) {
	l.wants = wants
}

// Wants returns true if this logger is configured to log the given log-level.
func (l *Logger) Wants(lvl level.Level) bool {
	if l.wants == nil {
		return true
	}
	for i := range l.wants {
		if l.wants[i] == lvl {
			return true
		}
	}
	return false
}

// Log logs the given value with the given log-level.
func (l *Logger) Log(lvl level.Level, v interface{}) {
	if redacted, ok := v.(Redacted); ok {
		v = redacted.Redacted()
	}
	l.log(fmt.Sprintf("%s", v), lvl)
}

// Logf logs the given values under the given log-level after formatting them.
func (l *Logger) Logf(lvl level.Level, format string, args ...interface{}) {
	l.log(fmt.Sprintf(l.formatArgs(format, args...)), lvl)
}

// Debug logs the given string with log-level DEBUG.
func (l *Logger) Debug(v interface{}) {
	if redacted, ok := v.(Redacted); ok {
		v = redacted.Redacted()
	}
	l.log(fmt.Sprintf("%s", v), level.DEBUG)
}

// Debugf logs the given values with log-level DEBUG after formatting them.
func (l *Logger) Debugf(format string, args ...interface{}) {
	l.log(fmt.Sprintf(l.formatArgs(format, args...)), level.DEBUG)
}

// Info logs the given string with log-level INFO.
func (l *Logger) Info(v interface{}) {
	l.log(fmt.Sprintf("%s", v), level.INFO)
}

// Infof logs the given values with log-level INFO after formatting them.
func (l *Logger) Infof(format string, args ...interface{}) {
	l.log(fmt.Sprintf(l.formatArgs(format, args...)), level.INFO)
}

// Notice logs the given string with log-level NOTICE.
func (l *Logger) Notice(v interface{}) {
	if redacted, ok := v.(Redacted); ok {
		v = redacted.Redacted()
	}
	l.log(fmt.Sprintf("%s", v), level.NOTICE)
}

// Noticef logs the given values with log-level NOTICE after formatting them.
func (l *Logger) Noticef(format string, args ...interface{}) {
	l.log(fmt.Sprintf(l.formatArgs(format, args...)), level.NOTICE)
}

// Warning logs the given string with log-level WARNING.
func (l *Logger) Warning(v interface{}) {
	if redacted, ok := v.(Redacted); ok {
		v = redacted.Redacted()
	}
	l.log(fmt.Sprintf("%s", v), level.WARNING)
}

// Warningf logs the given values with log-level WARNING after formatting them.
func (l *Logger) Warningf(format string, args ...interface{}) {
	l.log(fmt.Sprintf(l.formatArgs(format, args...)), level.WARNING)
}

// Error logs the given string with log-level ERROR.
func (l *Logger) Error(v interface{}) {
	if redacted, ok := v.(Redacted); ok {
		v = redacted.Redacted()
	}
	l.log(fmt.Sprintf("%s", v), level.ERROR)
}

// Errorf logs the given values with log-level ERROR after formatting them.
func (l *Logger) Errorf(format string, args ...interface{}) {
	l.log(fmt.Sprintf(l.formatArgs(format, args...)), level.ERROR)
}

func (l Logger) formatArgs(format string, args ...interface{}) string {
	for i := range args {
		if redacted, ok := args[i].(Redacted); ok {
			args[i] = redacted.Redacted()
		}
	}
	return fmt.Sprintf(format, args...)
}
