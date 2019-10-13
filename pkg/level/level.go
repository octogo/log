package level

import (
	"strings"
	"sync"
)

// Level is defined as
type Level uint8

// Log levels are
const (
	ERROR Level = iota
	WARNING
	NOTICE
	INFO
	DEBUG
)

var registeredLevels = map[string]Level{
	"ERROR":   ERROR,
	"WARNING": WARNING,
	"NOTICE":  NOTICE,
	"INFO":    INFO,
	"DEBUG":   DEBUG,
}

var mu = &sync.Mutex{}

func (l Level) String() string {
	mu.Lock()
	defer mu.Unlock()
	for k := range registeredLevels {
		if registeredLevels[k] == l {
			return k
		}
	}
	panic("level not registered")
}

// Register registers a new log-level under the given name.
func Register(name string) Level {
	mu.Lock()
	defer mu.Unlock()
	name = strings.ToUpper(name)
	for k := range registeredLevels {
		if k == name {
			return registeredLevels[k]
		}
	}
	registeredLevels[name] = Level(len(registeredLevels))
	return registeredLevels[name]
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
	return Level(0), errLevelInvalid
}
