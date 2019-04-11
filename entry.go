package log

import (
	"bytes"
	"fmt"
	"go/build"
	"html/template"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

// Entry defines the interface of a log entry.
type Entry interface {
	GID() string                // Global entry ID
	LID() string                // Logger entry ID
	PID() string                // Process ID
	PPID() string               // Parent's PID
	Logger() string             // Name of the logger
	Level() Level               // Log-level
	LevelNum() string           // Log-level as digit
	LevelLetters(n uint) string // First n letters of log-level name
	File() string               // Relative path to the source file of the caller
	FileLong() string           // Absolute path to the source file of the caller
	Line() string               // Line number in source file of caller
	Pkg() string                // Package name of the caller
	PkgLong() string            // Full package name of the caller
	Func() string               // The calling function
	Date() string               // Date in the format of YYYY-MM-DD
	Time() string               // Time in the format of HH:mm:ss
	TimeExact() string          // Time in the format of HH:mm:ss.S
	Msg() string                // Actual message to log
	Color() string              // ANSII color escape sequence
	ColorReset() string         // ANSII color reset escape sequence
}

// E implements an Entry.
type E struct {
	gid, lid  uint64
	logger    Logger
	timestamp time.Time
	level     Level
	runtime   struct {
		File string
		Line int
		Func string
	}
	args       []interface{}
	formatTmpl string
}

// SetGID sets the global entry ID of this entry.
func (entry *E) SetGID(gid uint64) {
	entry.gid = gid
}

// GID returns the global entry ID as a string.
func (entry E) GID() string {
	return strconv.Itoa(int(entry.gid))
}

// LID returns the local entry ID as a string.
func (entry E) LID() string {
	return strconv.Itoa(int(entry.lid))
}

// PID returns the process' ID as a string.
func (entry E) PID() string {
	return fmt.Sprintf("%d", os.Getpid())
}

// PPID returns the parent's PID.
func (entry E) PPID() string {
	return fmt.Sprintf("%d", os.Getppid())
}

// Logger returns the name of the corresponding logger.
func (entry E) Logger() string {
	return entry.logger.Name()
}

// Level returns the string representation of the log-level.
func (entry E) Level() Level {
	return entry.level
}

// LevelNum returns the log-level as digit.
func (entry E) LevelNum() string {
	return strconv.Itoa(int(entry.level))
}

// LevelLetters returns the first n letters of the string-representation of the log-level.
func (entry E) LevelLetters(n uint) string {
	if n == 0 {
		n = 3
	}

	levelName := levelNames[entry.level]
	if len(levelName) > int(n) {
		return levelName
	}
	return levelName[0:n]
}

// FileLong returns the absolute path to the source file of the caller.
func (entry E) FileLong() string {
	return entry.runtime.File

}

// File returns only the relative portion of E.FileLong()
func (entry E) File() string {
	path := entry.FileLong()
	gopath := build.Default.GOPATH + string(filepath.Separator)
	path = strings.Replace(path, gopath, "", 1)
	exploded := strings.Split(path, string(filepath.Separator))
	return filepath.Join(exploded[3:]...)
}

// Line returns the line number in the source file of the caller.
func (entry E) Line() string {
	return string(entry.runtime.Line)
}

// PkgLong returns the absolute path of the caller's package.
func (entry E) PkgLong() string {
	path := entry.FileLong()
	gopath := build.Default.GOPATH + string(filepath.Separator)
	path = strings.Replace(path, gopath, "", 1)
	exploded := strings.Split(path, string(filepath.Separator))
	return filepath.Join(exploded[1 : len(exploded)-1]...)
}

// Pkg returns the package path of the caller.
func (entry E) Pkg() string {
	pkg := entry.PkgLong()
	exploded := strings.Split(pkg, string(filepath.Separator))
	return filepath.Join(exploded[2:]...)
}

// Func returns the caller's function's name.
func (entry E) Func() string {
	return entry.runtime.Func
}

// Date returns the current date in the format of YYYY-MM-DD.
func (entry E) Date() string {
	return entry.timestamp.Format("2006-01-02")
}

// Time returns the current time in the format of HH:mm:ss.
func (entry E) Time() string {
	return entry.timestamp.Format("15:04:05")
}

// TimeExact returns the current time in the format of HH:mm:ss.S
func (entry E) TimeExact() string {
	return entry.timestamp.Format("15:04:05.000")
}

// Msg returns the actual message to log.
func (entry E) Msg() string {
	args := entry.args
	for i, arg := range args {
		if redactor, ok := arg.(Redactor); ok == true {
			args[i] = redactor.Redacted()
		}
	}

	var buf bytes.Buffer
	if entry.formatTmpl != "" {
		fmt.Fprintf(&buf, entry.formatTmpl, args...)
	} else {
		fmt.Fprintln(&buf, args...)
		buf.Truncate(buf.Len() - 1)
	}
	return buf.String()
}

// Color returns the ANSII color escape sequence.
func (entry E) Color() string {
	return Colors[entry.level]
}

// ColorReset returns the ANSII color reset escape sequence.
func (entry E) ColorReset() string {
	return colorSeqReset()
}

// FormattedEntry returns a string containing the entry templated to the format
// defined by the given format string.
func FormattedEntry(e Entry, format string) string {
	tmpl, err := template.New("octolog/entry").Parse(format)
	if err != nil {
		internalError(err)
		return ""
	}

	buf := new(bytes.Buffer)
	err = tmpl.Execute(buf, e)
	if err != nil {
		internalError(err)
		return ""
	}

	return buf.String()
}
