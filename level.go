package log

import (
	"strings"

	"github.com/pkg/errors"
)

// Level serves as pseudo-type for refering to log-levels.
type Level uint8

// LevelSet serves as pseudo-type for refering to a set of log-levels.
type LevelSet map[Level]struct{}

// LevelSlice serves as pseudo-type for refering to a slice of log-levels.
type LevelSlice []Level

// Log Levels
const (
	ERROR Level = iota
	WARNING
	ALERT
	NOTICE
	INFO
	DEBUG
)

var levelNames = []string{
	"ERROR",
	"WARNING",
	"ALERT",
	"NOTICE",
	"INFO",
	"DEBUG",
}

// ErrInvalidLogLevel defines an error indicating the use of an invalid log-level.
var ErrInvalidLogLevel = errors.New("octolog: invalid log-level")

// String returns the log-level as type string.
func (level Level) String() string {
	return levelNames[level]
}

// Int returns the log-level as type int.
func (level Level) Int() int {
	return int(level)
}

// AllLevelSlice returns a LevelSlice containing all levels.
func AllLevelSlice() LevelSlice {
	return LevelSlice{
		ERROR,
		WARNING,
		ALERT,
		NOTICE,
		INFO,
		DEBUG,
	}
}

// AllLevelSet returns a LevelSet containing all levels.
func AllLevelSet() LevelSet {
	levels := make(LevelSet)
	for _, level := range AllLevelSlice() {
		levels[level] = struct{}{}
	}
	return levels
}

// LevelByName returns the log-level corresponging to the given log-level name.
// Takes a log-level as string (case-insensitive).
func LevelByName(level string) (Level, error) {
	for i, name := range levelNames {
		if strings.EqualFold(name, level) {
			return Level(i), nil
		}
	}
	return ERROR, ErrInvalidLogLevel
}

func levelSliceToSet(s LevelSlice) LevelSet {
	levels := make(LevelSet)
	for _, level := range s {
		levels[level] = struct{}{}
	}
	return levels
}

func levelSetToSlice(m LevelSet) LevelSlice {
	levels := make(LevelSlice, 0, len(m))
	for level := range m {
		levels = append(levels, level)
	}
	return levels
}
