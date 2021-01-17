package log

import (
	"os"
	"sync"

	"github.com/octogo/log/internal/lib"
	"github.com/octogo/log/pkg/level"
	"github.com/octogo/log/pkg/log/terminal"
)

// FileOutput implements an output that writes the logs to a file.
type FileOutput struct {
	File   *os.File
	wants  []level.Level
	format string
	mu     *sync.Mutex
}

// NewFileOutput returns an initialized FileOutput.
func NewFileOutput(file *os.File, wants []level.Level, format string) Output {
	if format == "" {
		format = DefaultLogFormat
	}
	output := &FileOutput{
		File:   file,
		wants:  wants,
		format: format,
		mu:     &sync.Mutex{},
	}
	return RegisterOutput(output.URL(), output)
}

// Type returns the type of this output (i.e. file).
func (fOut FileOutput) Type() string {
	return "file"
}

// URI returns the name of the underlying file.
func (fOut FileOutput) URI() string {
	return fOut.File.Name()
}

// URL returns the URL of this output.
func (fOut FileOutput) URL() string {
	return lib.URL(fOut.Type(), fOut.URI())
}

// Log writes the given Entry to the underlying file.
func (fOut FileOutput) Log(e Entry) (n int, err error) {
	fOut.mu.Lock()
	defer fOut.mu.Unlock()
	if fOut.Wants(e.LevelLevel()) {
		return fOut.File.WriteString(e.Formatted(fOut.format, !terminal.IsTerminal(int(fOut.File.Fd()))) + "\n")
	}
	return 0, nil
}

// SetFormat sets the format of this backend to the given string.
func (fOut *FileOutput) SetFormat(f string) {
	fOut.mu.Lock()
	defer fOut.mu.Unlock()
	fOut.format = f
}

// SetWants configures this backend to only log entries of the given levels.
func (fOut *FileOutput) SetWants(wants []level.Level) {
	fOut.mu.Lock()
	defer fOut.mu.Unlock()
	fOut.wants = wants
}

// Wants returns true if this backend is configured to log the given level.
func (fOut FileOutput) Wants(lvl level.Level) bool {
	if fOut.wants == nil {
		return true
	}
	for i := range fOut.wants {
		if fOut.wants[i] == lvl {
			return true
		}
	}
	return false
}
