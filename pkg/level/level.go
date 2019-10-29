package level

import (
	"strings"
	"sync"

	"github.com/octogo/log/pkg/color"
)

// Level is defined as
type Level uint8

// Color returns the ANSII escape sequence for the color of this log-level.
func (lvl Level) Color() color.Sequence {
	mu.Lock()
	defer mu.Unlock()
	return registeredColorSequences[lvl]
}

// Log levels are
const (
	ERROR Level = iota
	WARNING
	NOTICE
	INFO
	DEBUG
)

var (
	registeredLevels = map[string]Level{
		"ERROR":   ERROR,
		"WARNING": WARNING,
		"NOTICE":  NOTICE,
		"INFO":    INFO,
		"DEBUG":   DEBUG,
	}
	registeredColorSequences = map[Level]color.Sequence{
		ERROR:   color.New(color.NormalDisplay, color.Red),
		WARNING: color.New(color.NormalDisplay, color.Yellow),
		NOTICE:  color.New(color.NormalDisplay, color.Green),
		INFO:    color.New(color.NormalDisplay, color.White),
		DEBUG:   color.New(color.NormalDisplay, color.Cyan),
	}
)

var mu = &sync.Mutex{}

// String implements fmt.Stringer
func (lvl Level) String() string {
	mu.Lock()
	defer mu.Unlock()
	for k := range registeredLevels {
		if registeredLevels[k] == lvl {
			return k
		}
	}
	panic("level not registered")
}

// Register registers a new log-level under the given name.
func Register(name string, colSeq color.Sequence) (Level, bool, error) {
	mu.Lock()
	defer mu.Unlock()
	name = strings.ToUpper(name)
	for k := range registeredLevels {
		if k == name {
			registeredColorSequences[registeredLevels[k]] = colSeq
			return registeredLevels[k], false, nil
		}
	}
	registeredLevels[name] = Level(len(registeredLevels))
	registeredColorSequences[registeredLevels[name]] = colSeq
	return registeredLevels[name], true, nil
}

// Levels returns a []Level of all registered levels.
func Levels() []Level {
	mu.Lock()
	defer mu.Unlock()
	levels := make([]Level, len(registeredLevels))
	var i int
	for k := range registeredLevels {
		levels[i] = registeredLevels[k]
		i++
	}
	return levels
}

// Colors returns a []color.Sequence of all registered colors.
func Colors() []color.Sequence {
	mu.Lock()
	defer mu.Unlock()
	colors := make([]color.Sequence, len(registeredLevels))
	var i int
	for lvl := range registeredLevels {
		colors[i] = registeredColorSequences[registeredLevels[lvl]]
		i++
	}
	return colors
}

// IsValid returns true if the given level is registered.
func IsValid(lvl Level) bool {
	mu.Lock()
	defer mu.Unlock()
	for k := range registeredLevels {
		if registeredLevels[k] == lvl {
			return true
		}
	}
	return false
}

// IsValidName returns true if the fiven level is registered.
func IsValidName(name string) bool {
	mu.Lock()
	defer mu.Unlock()
	name = strings.ToUpper(name)
	for k := range registeredLevels {
		if k == name {
			return true
		}
	}
	return false
}

// Parse returns the Level from parsing the given level-name.
func Parse(name string) (Level, error) {
	mu.Lock()
	defer mu.Unlock()
	name = strings.ToUpper(name)
	for k := range registeredLevels {
		if k == name {
			return registeredLevels[k], nil
		}
	}
	return Level(0), errLevelUndefined
}
