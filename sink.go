package log

import (
	"os"
	"slices"
	"sync"
)

func init() {
	sinks[os.Stderr.Name()] = NewSink(os.Stderr, WantsLevels("ERROR", "WARNING"))
	sinks[os.Stdout.Name()] = NewSink(os.Stdout, WantsLevels("NOTICE", "INFO", "DEBUG"))
}

var (
	sinks  = map[string]*Sink{}
	muSink = &sync.Mutex{}
)

type Sink struct {
	f     *os.File
	wants []Level
	mu    *sync.Mutex
}

func NewSink(f *os.File, wants []Level) *Sink {
	muSink.Lock()
	defer muSink.Unlock()

	output, ok := sinks[f.Name()]
	if ok {
		return output
	}

	if wants == nil {
		wants = AllLevels()
	}

	sinks[f.Name()] = &Sink{
		f:     f,
		wants: wants,
		mu:    &sync.Mutex{},
	}

	return sinks[f.Name()]
}

func (s *Sink) Path() string {
	return s.f.Name()
}

func (s *Sink) Close() error {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.f.Close()
}

func (s *Sink) Wants(level Level) bool {
	return slices.Contains(s.wants, level)
}

func (s *Sink) Write(p []byte) (int, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.f.Write(p)
}

func (s *Sink) SetWants(wants []Level) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if wants == nil {
		wants = AllLevels()
	}
	s.wants = wants
}
