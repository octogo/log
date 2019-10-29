package log

import (
	"fmt"
	"os"

	"github.com/octogo/log/pkg/level"
)

// Println logs the given values with log-level INFO.
func Println(v interface{}) {
	defaultLogger.log(fmt.Sprintf("%s", v), level.INFO)
}

// Printf formats and logs the given values with log-level INFO.
func Printf(f string, args ...interface{}) {
	defaultLogger.log(fmt.Sprintf(f, args...), level.INFO)
}

// Fatal logs the given value with log-level ERROR and exits with RC-1.
func Fatal(v interface{}) {
	defaultLogger.log(fmt.Sprintf("%s", v), level.ERROR)
	os.Exit(1)
}

// Fatalf formats and logs the given values with log-level ERROR and exits with
// RC-1.
func Fatalf(f string, args ...interface{}) {
	defaultLogger.log(fmt.Sprintf(f, args...), level.ERROR)
	os.Exit(1)
}
