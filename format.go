package log

import (
	"bytes"
	"fmt"
	"html/template"
	"os"
	"time"
)

var (
	// DefaultFormat mimics the format of the builtin log package.
	DefaultFormat = "{{.Message}}"

	// LogFormat templates messages for log-files, including a timestamp, the logger and log-level.
	LogFormat = "{{.Date}} {{.Time}} [{{.BoldColor}}{{.Level}}{{.NoColor}}] {{.Color}}{{.Logger}}{{.NoColor}} {{.Message}}"

	// DebugFormat is an even more verbose format, than the DefaultFormat.
	DebugFormat = "{{.Date}} {{.Time}}{{.Nano}} {{.BoldColor}}{{.Logger}}{{.NoColor}} {{.Color}}{{.Message}}{{.NoColor}} [{{.File}}:{{.Line}} ({{.Caller}})]"

	// predefinedNow is a helper for testing with predefined time.Now()
	predefinedNow *time.Time
)

// Formatter is a helper for formatting log messages.
type Formatter struct {
	tmpl     *template.Template
	colorize bool
	msg      Message
}

// NewFormatter returns an initialized Formatter.
func NewFormatter(format string) Formatter {
	tmpl, err := template.New("octolog.message").Parse(format)
	if err != nil {
		panic(err)
	}

	return Formatter{
		tmpl: tmpl,
	}
}

// Format takes a *Message and returns it as formatted string.
func (f Formatter) Format(msg Message, colorize bool) string {
	f.msg = msg
	f.colorize = colorize
	buf := new(bytes.Buffer)
	err := f.tmpl.Execute(buf, f)
	if err != nil {
		panic(err)
	}
	return buf.String() + "\n"
}

// Date returns the string-formatted date of the log-message's timestamp.
func (f Formatter) Date() string {
	var date time.Time
	if predefinedNow != nil {
		date = *predefinedNow
	} else {
		date = f.msg.Timestamp
	}
	return date.Format("2006/01/02")
}

// Time returns the string-formatted time of the log-message's timestamp.
func (f Formatter) Time() string {
	var t time.Time
	if predefinedNow != nil {
		t = *predefinedNow
	} else {
		t = f.msg.Timestamp
	}
	return t.Format("15:04:05")
}

// Milli returns the string-formatted millisecond of the log-message's timestamp.
func (f Formatter) Milli() string {
	var t time.Time
	if predefinedNow != nil {
		t = *predefinedNow
	} else {
		t = f.msg.Timestamp
	}
	return t.Format(".000")
}

// Nano returns the string.formatted nanosecond of the log-message's timestamp.
func (f Formatter) Nano() string {
	var t time.Time
	if predefinedNow != nil {
		t = *predefinedNow
	} else {
		t = f.msg.Timestamp
	}
	return t.Format(".000000")
}

// PID returns the string-formatted PID of this process.
func (f Formatter) PID() string {
	return fmt.Sprintf("%d", os.Getpid())
}

// Logger returns the name of the log-message's Logger.
func (f Formatter) Logger() string {
	return fmt.Sprintf("%s", f.msg.Logger.Name)
}

// Level returns the string-formatted log-level of the log-message.
func (f Formatter) Level() string {
	return string(f.msg.Level)
}

// Message returns the actual text to be logged.
func (f Formatter) Message() string {
	return f.msg.Msg
}

// File returns the filename of the calling function.
func (f Formatter) File() string {
	return f.msg.File
}

// Caller returns the name of the caller function.
func (f Formatter) Caller() string {
	return f.msg.Caller
}

// Line returns the line in the file of the caller function.
func (f Formatter) Line() string {
	return fmt.Sprintf("%d", f.msg.Line)
}

// NoColor returns the ANSII escape sequence for resetting all state.
func (f Formatter) NoColor() string {
	if !f.colorize {
		return ""
	}
	return ansiiReset()
}

// Color returns the ANSII escape sequence for the color defined by the
// log-level.
func (f Formatter) Color() string {
	if !f.colorize {
		return ""
	}
	return AnsiiSequence(Normal, ColorOf(f.msg.Level))
}

// BoldColor returns the ANSII escape-sequence for the color defined by the
// log-level, in bold text.
func (f Formatter) BoldColor() string {
	if !f.colorize {
		return ""
	}
	return AnsiiSequence(Bold, ColorOf(f.msg.Level))
}
