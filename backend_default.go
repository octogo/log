package octolog

// DefaultBackends returns the default backends.
func DefaultBackends() []Backend {
	format := DefaultLogFormat

	backendOut, err := NewStdoutBackend(
		format,
		LevelSlice{DEBUG, INFO, NOTICE},
	)
	if err != nil {
		panic(err)
	}

	backendErr, err := NewStderrBackend(
		format,
		LevelSlice{ERROR, WARNING, ALERT},
	)
	if err != nil {
		panic(err)
	}

	return []Backend{backendOut, backendErr}
}
