package octolog

import "os"

// NewStdoutBackend returns a new *FileBackend logging to stdout.
func NewStdoutBackend(format string, levels LevelSlice) (*FileBackend, error) {
	return NewFileBackend(format, levels, os.Stdout, true)
}

// NewStderrBackend returns a new *FileBackend logging to stderr.
func NewStderrBackend(format string, levels LevelSlice) (*FileBackend, error) {
	return NewFileBackend(format, levels, os.Stderr, true)
}
