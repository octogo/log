package log

import (
	"fmt"
	"log"
	"runtime"
	"strings"
	"sync"
	"time"
)

func init() {
	// Always initialize a DefaultLogger logging to STDOUT/STDERR.
	DefaultLogger = New(
		"main",
		WantsAllLevels(),
		DefaultFormat,
		DefaultSinks()...,
	)
}

// DefaultLogger holds a reference to the DefaultLogger.
var DefaultLogger *Logger

type Logger struct {
	// Name contains the name of this Logger.
	Name string

	// Wants contains the list of log-levels this logger is interested in.
	Wants []Level

	// Sinks contains the list of sinks to forward all log-messages to having
	// any of the wanted log-levels.
	Sinks []*Sink

	// format is raw template-string for formatting log-messages.
	format string

	// fmt is the initialized template.
	fmt Formatter

	// mu is used as synchronization semaphore.
	mu *sync.Mutex
}

// New returns an initialized *Logger.
func New(name string, wants []Level, format string, sinks ...*Sink) *Logger {
	if wants == nil {
		wants = []Level{ERROR, WARNING, NOTICE, INFO}
	}

	if sinks == nil {
		sinks = []*Sink{}
	}

	return &Logger{
		Name:   name,
		Wants:  wants,
		Sinks:  sinks,
		format: format,
		fmt:    NewFormatter(format),
		mu:     &sync.Mutex{},
	}
}

// Println logs a line with log-level INFO.
func (l *Logger) Println(s ...any) {
	l.Log(INFO, s...)
}

// Printf logs a line with log-level INFO.
func (l *Logger) Printf(p string, s ...any) {
	l.Log(INFO, fmt.Sprintf(p, s...))
}

// Debug logs a line with log-level DEBUG.
func (l *Logger) Debug(s ...any) {
	l.Log(DEBUG, s...)
}

// Debugf logs a line with log-level DEBUG.
func (l *Logger) Debugf(p string, s ...any) {
	l.Log(DEBUG, fmt.Sprintf(p, s...))
}

// Notice logs a line with log-level NOTICE.
func (l *Logger) Notice(s ...any) {
	l.Log(NOTICE, s...)
}

// Noticef logs a line with log-level NOTICE.
func (l *Logger) Noticef(p string, s ...any) {
	l.Log(NOTICE, fmt.Sprintf(p, s...))
}

// Warning logs a line with log-level WARNING.
func (l *Logger) Warning(s ...any) {
	l.Log(WARNING, s...)
}

// Warningf logs a line with log-level WARNING.
func (l *Logger) Warningf(p string, s ...any) {
	l.Log(WARNING, fmt.Sprintf(p, s...))
}

// Error logs a line with log-level ERROR.
func (l *Logger) Error(s ...any) {
	l.Log(ERROR, s...)
}

// Errorf logs a line with log-level ERROR.
func (l *Logger) Errorf(p string, s ...any) {
	l.Log(ERROR, fmt.Sprintf(p, s...))
}

// Fatal logs a line with log-level ERROR and then exists the program.
func (l *Logger) Fatal(s ...any) {
	l.Log(ERROR, s...)
	exit(1)
}

// Fatalf logs a line with log-level ERROR and then exists the program.
func (l *Logger) Fatalf(p string, s ...any) {
	l.Log(ERROR, fmt.Sprintf(p, s...))
	exit(1)
}

// Log sends any message with the given log-level to all sinks of this logger.
func (logger *Logger) Log(level Level, s ...any) {
	segments := make([]string, len(s))
	for i, p := range s {
		switch t := p.(type) {
		case Redacter:
			segments[i] = t.Redacted()
		default:
			segments[i] = fmt.Sprintf("%s", p)
		}
	}

	msg := Message{
		timestamp: time.Now(),
		logger:    logger,
		level:     level,
		msg:       strings.Join(segments, " "),
	}

	for _, sink := range logger.Sinks {
		if sink.Wants(level) {
			logger.log(sink, msg)
		}
	}
}

// log is a helper for Log().
func (logger *Logger) log(sink *Sink, msg Message) {
	logger.mu.Lock()
	defer logger.mu.Unlock()

	var (
		stackDepth = 4
		fpcs       = make([]uintptr, 1)
		n          = runtime.Callers(stackDepth, fpcs)
	)
	if n != 0 {
		f := runtime.FuncForPC(fpcs[0] - 1)
		if f != nil {
			msg.file, msg.line = f.FileLine(fpcs[0] - 1)
			msg.caller = f.Name()
		}
	}

	if _, err := sink.Log(msg, logger.fmt); err != nil {
		log.Fatal(err)
	}
}

// SetName sets this Logger's name.
func (logger *Logger) SetName(name string) {
	logger.mu.Lock()
	defer logger.mu.Unlock()
	logger.Name = name
}

// SetWants sets the []Level this Logger is interested in.
func (logger *Logger) SetWants(wants []Level) {
	logger.mu.Lock()
	defer logger.mu.Unlock()
	logger.Wants = wants
}

// SetFormat sets the template-string of this Logger for formatting
// log-messages.
func (logger *Logger) SetFormat(format string) {
	logger.format = format
	logger.fmt = NewFormatter(format)
}

// AddSinks adds the given sinks to this Logger.
func (logger *Logger) AddSinks(sinks ...*Sink) {
	for _, s := range sinks {
		for _, _s := range logger.Sinks {
			if _s.f == s.f {
				return
			}
		}

		logger.Sinks = append(logger.Sinks, s)
	}
}

// SetSinks sets this Logger's sinks.
func (logger *Logger) SetSinks(sinks ...*Sink) {
	logger.Sinks = sinks
}

// NewLogger returns a new *Logger
// The name of the new logger is automatically be prefixed with
// fmt.Sprintf("%s.%s", logger.Name, name).
func (logger *Logger) NewLogger(name string) *Logger {
	return New(
		strings.Join([]string{logger.Name, name}, "."),
		logger.Wants,
		logger.format,
		logger.Sinks...,
	)
}
