package log

import (
	"os"
	"strings"
)

// Wants converts the given list of []string and returns them as
// []Level.
func Wants(levels ...string) []Level {
	out := []Level{}
	for _, lvl := range levels {
		out = append(out, Level(strings.ToUpper(lvl)))
	}
	return out
}

// WantsLevels is merely a convenience wrapper for making consumer code look
// nicer.
func WantsLevels(levels ...Level) []Level {
	return levels
}

// WantsAllLevels returns a []Level containing all log-levels.
func WantsAllLevels() []Level {
	return AllLevels()
}

// WithSink returns a *Sink for the given path.
// If the sink for that path already exists, the existing *Sink is returned,
// while the sink's Wants slice is left untouched.
func WithSink(path string, wants []Level) *Sink {
	if wants == nil {
		wants = AllLevels()
	}

	switch path {
	case os.Stdout.Name():
		f := os.NewFile(os.Stdout.Fd(), os.Stdout.Name())
		return NewSink(f, wants)
	case os.Stderr.Name():
		f := os.NewFile(os.Stderr.Fd(), os.Stderr.Name())
		return NewSink(f, wants)
	default:
		f, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0640)
		if err != nil {
			panic(err)
		}
		return NewSink(f, wants)
	}
}

// DefaultSinks returns a []*Sink of the default sinks for STDOUT and STDERR.
func DefaultSinks() []*Sink {
	return []*Sink{
		NewSink(os.Stdout, Wants("INFO")),
		NewSink(os.Stderr, Wants("DEBUG", "NOTICE", "ERROR", "WARNING")),
	}
}
