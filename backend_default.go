package log

// DefaultBackends returns the default backends.
func DefaultBackends() []Backend {
	format := DefaultLogFormat
	return []Backend{
		NewStdoutBackend(format, LevelSlice{DEBUG, INFO, NOTICE}),
		NewStderrBackend(format, LevelSlice{ALERT, WARNING, ERROR}),
	}
}
