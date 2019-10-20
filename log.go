package log

import (
	"github.com/octogo/log/pkg/level"
	"github.com/octogo/log/pkg/log"
)

// New returns an initialized Logger with the given name.
// If a Logger with the given name has already been registered, then that
// Logger will be returned instead of initializing dupicate Loggers with the
// same name. This also ensures that the LID of a logger will always increase.
func New(name string, wants []level.Level, outputs ...string) *log.Logger {
	return log.NewLogger(name, wants, outputs...)
}

// Println logs the given value with log-level INFO.
func Println(v interface{}) {
	log.Println(v)
}

// Printf wraps Println and supports string formatting.
func Printf(f string, args ...interface{}) {
	log.Printf(f, args...)
}

// Log is an alias for Println.
func Log(v interface{}) {
	Println(v)
}

// Logf is an alias for Printf.
func Logf(f string, args ...interface{}) {
	Printf(f, args...)
}

// Fatal logs the given value with log-level ERROR and exits with RC-1.
func Fatal(v interface{}) {
	log.Fatal(v)
}

// Fatalf wraps Fatal() and supports string formatting.
func Fatalf(f string, args ...interface{}) {
	log.Fatalf(f, args...)
}
