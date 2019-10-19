package log

import (
	"bytes"
	"fmt"
	"html/template"
	"os"
	"time"

	"github.com/octogo/log/internal/gid"
	"github.com/octogo/log/pkg/color"
	"github.com/octogo/log/pkg/level"
)

// Entry is defined as
type Entry interface {
	Message() string
	Color() string
	BoldColor() string
	NoColor() string
	Date() string
	Time() string
	Milli() string
	Nano() string
	PID() string
	PPID() string
	GID() string
	LID() string
	Logger() string
	Level() string
	Func() string
	File() string
	Line() string
	Formatted(f string, disableColors bool) string
	LevelLevel() level.Level
}

type entryStruct struct {
	timestamp     time.Time
	level         level.Level
	message       string
	logger        string
	gid, lid      uint64
	caller        string
	file          string
	line          int
	disableColors bool
}

func newEntry(
	msg string,
	logger *Logger,
	lvl level.Level,
	caller, file string,
	line int,
) Entry {
	return &entryStruct{
		timestamp: time.Now(),
		level:     lvl,
		message:   msg,
		logger:    logger.Name,
		gid:       gid.Next(),
		lid:       logger.uid.Next(),
		caller:    caller,
		file:      file,
		line:      line,
	}
}

func (e entryStruct) Message() string {
	return e.message
}

func (e entryStruct) Color() string {
	if e.disableColors {
		return ""
	}
	return e.LevelLevel().Color().String()
}

func (e entryStruct) BoldColor() string {
	if e.disableColors {
		return ""
	}
	e.LevelLevel().Color().SetAttribute(color.Bold)
	return e.LevelLevel().Color().String()
}

func (e entryStruct) NoColor() string {
	if e.disableColors {
		return ""
	}
	return color.New(color.NormalDisplay).String()
}

func (e entryStruct) Date() string {
	return e.timestamp.Format("2006/01/02")
}

func (e entryStruct) Time() string {
	return e.timestamp.Format("15:04:05")
}

func (e entryStruct) Milli() string {
	return e.timestamp.Format(".000")
}

func (e entryStruct) Nano() string {
	return e.timestamp.Format(".000000")
}

func (e entryStruct) PID() string {
	return fmt.Sprintf("%d", os.Getpid())
}

func (e entryStruct) PPID() string {
	return fmt.Sprintf("%d", os.Getppid())
}

func (e entryStruct) GID() string {
	return fmt.Sprintf("%d", e.gid)
}

func (e entryStruct) LID() string {
	return fmt.Sprintf("%d", e.lid)
}

func (e entryStruct) Logger() string {
	return fmt.Sprintf("%s", e.logger)
}

func (e entryStruct) Level() string {
	return e.level.String()
}

func (e entryStruct) LevelLevel() level.Level {
	return e.level
}

func (e entryStruct) File() string {
	return e.file
}

func (e entryStruct) Func() string {
	return e.caller
}

func (e entryStruct) Line() string {
	return fmt.Sprintf("%d", e.line)
}

func (e entryStruct) Formatted(f string, disableColors bool) string {
	e.disableColors = disableColors
	tmpl, err := template.New("octolog/entry").Parse(f)
	if err != nil {
		panic(err)
	}
	buf := new(bytes.Buffer)
	err = tmpl.Execute(buf, e)
	if err != nil {
		panic(err)
	}
	return buf.String()
}
