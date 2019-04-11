package log

import (
	"os"

	"github.com/pkg/errors"
	"golang.org/x/sys/unix"
)

// FileBackend is a backend logging to a file.
type FileBackend struct {
	BaseBackend
	file        *os.File
	colorOutput bool
}

// NewFileBackend returns a newly initialized FileBackend
func NewFileBackend(format string, levels LevelSlice, file *os.File, colorOutput bool) *FileBackend {
	if file == nil {
		panic(errors.New("file can not be nil"))
	}

	return &FileBackend{
		BaseBackend: BaseBackend{
			format: format,
			levels: levelSliceToSet(levels),
		},
		file:        file,
		colorOutput: colorOutput,
	}
}

// isTerminal returns a bool indicating if the log-file is a TTY.
func (fileBackend FileBackend) isTerminal() bool {
	_, err := unix.IoctlGetTermios(int(fileBackend.file.Fd()), unix.TCGETS)
	return err == nil
}

// Log takes a Record and logs it by appending it to the file.
func (fileBackend FileBackend) Log(entry Entry) {
	line := fileBackend.FormattedLogLine(entry)

	if fileBackend.colorOutput && fileBackend.isTerminal() {
		line = entry.Color() + line + entry.ColorReset()
	}

	fileBackend.file.WriteString(line + "\n")
}
