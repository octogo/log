package log

import (
	"testing"
	"time"

	"github.com/octogo/log/pkg/level"
)

var (
	testLogger = NewLogger("TEST-LOGGER", nil)
	msg        = "TEST MESSAGE"
	testEntry  = newEntry(msg, testLogger, level.INFO, "caller", "file", 42)
)

func TestEntryDate(t *testing.T) {
	var (
		expectedDate  = time.Now().Format("2006/01/02")
		date          = testEntry.Date()
		formattedDate = testEntry.Formatted("{{.Date}}", true)
	)
	if date != expectedDate {
		t.Errorf("expected %v, got %v", expectedDate, date)
	}
	if formattedDate != expectedDate {
		t.Errorf("expected %v, got %v", expectedDate, formattedDate)
	}
}

func TestEntryTime(t *testing.T) {
	var (
		expectedTime  = time.Now().Format("15:04:05")
		_time         = testEntry.Time()
		formattedTime = testEntry.Formatted("{{.Time}}", true)
	)
	if _time != expectedTime {
		t.Errorf("expected %v, got %v", expectedTime, _time)
	}
	if formattedTime != expectedTime {
		t.Errorf("expected %v, got %v", expectedTime, formattedTime)
	}
}
func TestEntryMilli(t *testing.T) {
	var (
		expectedMilli  = 4
		milli          = len(testEntry.Milli())
		formattedMilli = len(testEntry.Formatted("{{.Milli}}", true))
	)
	if milli != expectedMilli {
		t.Errorf("expected %v, got %v", expectedMilli, milli)
	}
	if formattedMilli != expectedMilli {
		t.Errorf("expected %v, got %v", expectedMilli, formattedMilli)
	}
}
func TestEntryNano(t *testing.T) {
	var (
		expectedNano  = 7
		nano          = len(testEntry.Nano())
		formattedNano = len(testEntry.Formatted("{{.Nano}}", true))
	)
	if nano != expectedNano {
		t.Errorf("expected %v, got %v", expectedNano, nano)
	}
	if formattedNano != expectedNano {
		t.Errorf("expected %v, got %v", expectedNano, formattedNano)
	}
}
func TestEntryMessage(t *testing.T) {
	var (
		expectedMessage = msg
		message         = testEntry.Message()
		formattedMsg    = testEntry.Formatted("{{.Message}}", true)
	)
	if message != expectedMessage {
		t.Errorf("expected %v, got %v", expectedMessage, message)
	}
	if formattedMsg != expectedMessage {
		t.Errorf("expected %v, got %v", expectedMessage, formattedMsg)
	}
}
func TestEntryLevel(t *testing.T) {
	var (
		expectedLevel  = "INFO"
		lvl            = testEntry.Level()
		formattedLevel = testEntry.Formatted("{{.Level}}", true)
	)
	if lvl != expectedLevel {
		t.Errorf("expected %v, got %v", expectedLevel, lvl)
	}
	if formattedLevel != expectedLevel {
		t.Errorf("expected %v, got %v", expectedLevel, formattedLevel)
	}
}
func TestEntryLogger(t *testing.T) {
	var (
		expectedLogger  = "TEST-LOGGER"
		logger          = testEntry.Logger()
		formattedLogger = testEntry.Formatted("{{.Logger}}", true)
	)
	if logger != expectedLogger {
		t.Errorf("expected %v, got %v", expectedLogger, logger)
	}
	if formattedLogger != expectedLogger {
		t.Errorf("expected %v, got %v", expectedLogger, formattedLogger)
	}
}

func TestEntryFunc(t *testing.T) {
	var (
		expectedFunc  = "caller"
		Func          = testEntry.Func()
		formattedFunc = testEntry.Formatted("{{.Func}}", true)
	)
	if Func != expectedFunc {
		t.Errorf("expected %v, got %v", expectedFunc, Func)
	}
	if formattedFunc != expectedFunc {
		t.Errorf("expected %v, got %v", expectedFunc, formattedFunc)
	}
}

func TestEntryFile(t *testing.T) {
	var (
		expectedFile  = "file"
		file          = testEntry.File()
		formattedFile = testEntry.Formatted("{{.File}}", true)
	)
	if file != expectedFile {
		t.Errorf("expected %v, got %v", expectedFile, file)
	}
	if formattedFile != expectedFile {
		t.Errorf("expected %v, got %v", expectedFile, formattedFile)
	}
}

func TestEntryLine(t *testing.T) {
	var (
		expectedLine  = "42"
		line          = testEntry.Line()
		formattedLine = testEntry.Formatted("{{.Line}}", true)
	)
	if line != expectedLine {
		t.Errorf("expected %v, got %v", expectedLine, line)
	}
	if formattedLine != expectedLine {
		t.Errorf("expected %v, got %v", expectedLine, formattedLine)
	}
}
