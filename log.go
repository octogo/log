package log

// Println is the drop-in replacement for the builtin's log.Println()
func Println(s ...any) {
	DefaultLogger.Println(s...)
}

// Printf is the drop-in replacement for the builtin's log.Printf()
func Printf(p string, s ...any) {
	DefaultLogger.Printf(p, s...)
}

// Fatal is the drop-in replacement for the builtin's log.Fatal()
func Fatal(s ...any) {
	DefaultLogger.Fatal(s...)
}

// Fatalf is the drop-in replacement for the builtin's log.Fatalf()
func Fatalf(p string, s ...any) {
	DefaultLogger.Fatalf(p, s...)
}

// Debug works like Println, but writes the message with log-level DEBUG.
func Debug(s ...any) {
	DefaultLogger.Debug(s...)
}

// Debugf works like Printf, but writes the message with log-level DEBUG.
func Debugf(p string, s ...any) {
	DefaultLogger.Debugf(p, s...)
}

// Notice works like Println, but writes the message with log-level NOTICE.
func Notice(s ...any) {
	DefaultLogger.Notice(s...)
}

// Noticef works like Printf, but writes the message with log-level NOTICE.
func Noticef(p string, s ...any) {
	DefaultLogger.Noticef(p, s...)
}

// Warning works like Println, but writes the message with log-level WARNING.
func Warning(s ...any) {
	DefaultLogger.Warning(s...)
}

// Warningf works like Printf, but writes the message with log-level WARNING.
func Warningf(p string, s ...any) {
	DefaultLogger.Warningf(p, s...)
}

// Error works like Println, but writes the message with log-level ERROR.
func Error(s ...any) {
	DefaultLogger.Error(s...)
}

// Errorf works like Printlf, but writes the message with log-level ERROR.
func Errorf(p string, s ...any) {
	DefaultLogger.Errorf(p, s...)
}
