package log

import (
	"os"
	"sync"

	"github.com/octogo/log/v2/ansii"
)

type Level string

var lvlMu = &sync.Mutex{}

const (
	ERROR   Level = "ERROR"
	WARNING Level = "WARNING"
	NOTICE  Level = "NOTICE"
	INFO    Level = "INFO"
	DEBUG   Level = "DEBUG"
	INVALID Level = ""
)

var levelColors = map[Level]ansii.Color{
	ERROR:   ansii.RED,
	WARNING: ansii.YELLOW,
	NOTICE:  ansii.GREEN,
	INFO:    ansii.WHITE,
	DEBUG:   ansii.CYAN,
}

func AddLevel(level string, color ansii.Color) Level {
	if level == string(INVALID) {
		return Level(level)
	}

	lvlMu.Lock()
	defer lvlMu.Unlock()

	for known := range levelColors {
		if string(known) == level {
			return Level(level)
		}
	}

	// always add level to STDOUT sink
	FindSink(os.Stdout).AddWants(Level(level))

	// always add to DefaultLogger
	DefaultLogger.Wants = append(DefaultLogger.Wants, Level(level))

	levelColors[Level(level)] = color
	return Level(level)
}

func ChangeColor(level Level, color ansii.Color) {
	lvlMu.Lock()
	defer lvlMu.Unlock()
	levelColors[level] = color
}

func AllLevels() []Level {
	lvlMu.Lock()
	defer lvlMu.Unlock()

	levels := []Level{}

	for lvl := range levelColors {
		levels = append(levels, lvl)
	}

	return levels
}

func ColorOf(level Level) ansii.Color {
	lvlMu.Lock()
	defer lvlMu.Unlock()

	color, ok := levelColors[level]
	if ok {
		return color
	}

	return ansii.BLACK
}
