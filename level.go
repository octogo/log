package log

import (
	"sync"

	"github.com/octogo/log/v2/ansii"
)

type Level string

var mu = &sync.Mutex{}

const (
	ERROR   Level = "ERROR"
	WARNING Level = "WARNING"
	NOTICE  Level = "NOTICE"
	INFO    Level = "INFO"
	DEBUG   Level = "DEBUG"
)

var levelColors = map[Level]ansii.Color{
	ERROR:   ansii.RED,
	WARNING: ansii.YELLOW,
	NOTICE:  ansii.GREEN,
	INFO:    ansii.WHITE,
	DEBUG:   ansii.CYAN,
}

func AddLevel(identifier string, color ansii.Color) {
	mu.Lock()
	defer mu.Unlock()
	levelColors[Level(identifier)] = color
}

func ChangeColor(level Level, color ansii.Color) {
	mu.Lock()
	defer mu.Unlock()
	levelColors[level] = color
}

func AllLevels() []Level {
	mu.Lock()
	defer mu.Unlock()

	levels := []Level{}

	for lvl := range levelColors {
		levels = append(levels, lvl)
	}

	return levels
}

func ColorOf(level Level) ansii.Color {
	mu.Lock()
	defer mu.Unlock()

	color, ok := levelColors[level]
	if ok {
		return color
	}

	return ansii.BLACK
}
