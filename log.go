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
	exit(1)
}

// Fatalf is the drop-in replacement for the builtin's log.Fatalf()
func Fatalf(p string, s ...any) {
	DefaultLogger.Fatalf(p, s...)
	exit(1)
}

// PrintDebug works like Println, but writes the message with log-level DEBUG.
func PrintDebug(s ...any) {
	DefaultLogger.PrintDebug(s...)
}

// PrintDebugf works like Printf, but writes the message with log-level DEBUG.
func PrintDebugf(p string, s ...any) {
	DefaultLogger.PrintDebugf(p, s...)
}

// PrintNotice works like Println, but writes the message with log-level NOTICE.
func PrintNotice(s ...any) {
	DefaultLogger.PrintNotice(s...)
}

// PrintNoticef works like Printf, but writes the message with log-level NOTICE.
func PrintNoticef(p string, s ...any) {
	DefaultLogger.PrintNoticef(p, s...)
}

// PrintWarning works like Println, but writes the message with log-level WARNING.
func PrintWarning(s ...any) {
	DefaultLogger.PrintWarning(s...)
}

// PrintWarningf works like Printf, but writes the message with log-level WARNING.
func PrintWarningf(p string, s ...any) {
	DefaultLogger.PrintWarningf(p, s...)
}

// PrintErr works like Println, but writes the message with log-level ERROR.
func PrintErr(s ...any) {
	DefaultLogger.PrintErr(s...)
}

// PrintErrf works like Printlf, but writes the message with log-level ERROR.
func PrintErrf(p string, s ...any) {
	DefaultLogger.PrintErrf(p, s...)
}
