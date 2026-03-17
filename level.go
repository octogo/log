package log

import (
	"os"
	"sync"
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

var levelColors = map[Level]AnsiiColor{
	ERROR:   RED,
	WARNING: YELLOW,
	NOTICE:  GREEN,
	INFO:    WHITE,
	DEBUG:   CYAN,
}

func AddLevel(level string, color AnsiiColor) Level {
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

func ChangeColor(level Level, color AnsiiColor) {
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

func ColorOf(level Level) AnsiiColor {
	lvlMu.Lock()
	defer lvlMu.Unlock()

	color, ok := levelColors[level]
	if ok {
		return color
	}

	return BLACK
}
