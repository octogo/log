package log

import (
	"os"
	"strings"
)

// WantsLevels converts the given list of []string and returns them as
// []Level.
func WantsLevels(levels ...string) []Level {
	out := []Level{}
	for _, lvl := range levels {
		out = append(out, Level(strings.ToUpper(lvl)))
	}
	return out
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
		sink, ok := sinks[os.Stdout.Name()]
		if ok {
			return sink
		}
		return NewSink(os.Stdout, wants)
	case os.Stderr.Name():
		sink, ok := sinks[os.Stderr.Name()]
		if ok {
			return sink
		}
		return NewSink(os.Stderr, wants)
	default:
		f, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 0640)
		if err != nil {
			panic(err)
		}
		return NewSink(f, wants)
	}
}

// DefaultSinks returns a []*Sink of the default sinks for STDOUT and STDERR.
func DefaultSinks() []*Sink {
	return []*Sink{
		NewSink(os.Stdout, WantsLevels("DEBUG", "INFO", "NOTICE")),
		NewSink(os.Stderr, WantsLevels("ERROR", "WARNING")),
	}
}
