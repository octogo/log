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
	Name     string
	wants    []level.Level
	Outputs  []string
	uid      *uid.UID
	_outputs []Output
	mu       *sync.Mutex
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

	if l._outputs == nil {
		l._outputs = make([]Output, len(l.Outputs))
		for i := range l.Outputs {
			l._outputs[i] = loadOutput(l.Outputs[i])
		}
	}

	entry := newEntry(msg, l, lvl, caller, file, line)
	for i := range l._outputs {
		_, err := l._outputs[i].Log(entry)
		if err != nil {
			l.Outputs = append(l.Outputs[:i], l.Outputs[i+1:]...)
			l._outputs = append(l._outputs[:i], l._outputs[i+1:]...)
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

// Debug logs the given string with log-level DEBUG.
func (l *Logger) Debug(v interface{}) {
	l.log(fmt.Sprintf("%s", v), level.DEBUG)
}

// Info logs the given string with log-level INFO.
func (l *Logger) Info(v interface{}) {
	l.log(fmt.Sprintf("%s", v), level.INFO)
}

// Notice logs the given string with log-level NOTICE.
func (l *Logger) Notice(v interface{}) {
	l.log(fmt.Sprintf("%s", v), level.NOTICE)
}

// Warning logs the given string with log-level WARNING.
func (l *Logger) Warning(v interface{}) {
	l.log(fmt.Sprintf("%s", v), level.WARNING)
}

// Error logs the given string with log-level ERROR.
func (l *Logger) Error(v interface{}) {
	l.log(fmt.Sprintf("%s", v), level.ERROR)
}
