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
	// The standard streams are passed through as themselves rather than
	// re-wrapped. os.NewFile on the same descriptor yields a second handle to
	// the same stream, and sinks are identified by handle — so wrapping would
	// build a second stdout sink beside the one every Logger already shares,
	// writing each line twice and invisible to FindSink(os.Stdout), which is how
	// AddLevel reaches it.
	case os.Stdout.Name():
		return NewSink(os.Stdout, wants)
	case os.Stderr.Name():
		return NewSink(os.Stderr, wants)
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
