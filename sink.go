package log

import (
	"os"
	"slices"
	"sync"

	"golang.org/x/term"
)

var (
	sinks  = map[uintptr]*Sink{}
	sinkMu = sync.Mutex{}
)

type Sink struct {
	path  string
	f     *os.File
	wants []Level
	mu    *sync.Mutex
}

func NewSink(f *os.File, wants []Level) *Sink {
	sinkMu.Lock()
	defer sinkMu.Unlock()

	if existing, found := sinks[f.Fd()]; found {
		// don't f.Close()!
		// the FD must stay alive.
		return existing
	}

	if wants == nil {
		wants = AllLevels()
	}

	sinks[f.Fd()] = &Sink{
		path:  f.Name(),
		f:     f,
		wants: wants,
		mu:    &sync.Mutex{},
	}

	return sinks[f.Fd()]
}

func (s *Sink) Path() string {
	return s.path
}

func (s *Sink) Close() error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.f.Fd() == os.Stdout.Fd() || s.f.Fd() == os.Stderr.Fd() {
		return nil
	}
	return s.f.Close()
}

func (s *Sink) Wants(level Level) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	return slices.Contains(s.wants, level)
}

func (s *Sink) Log(msg Message, fmt Formatter) (int, error) {
	line := fmt.Format(msg, term.IsTerminal(int(s.f.Fd())))
	return s.Write([]byte(line))
}

func (s *Sink) Write(p []byte) (int, error) {
	// os.File.Write is already concurrency-safe
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

func (s *Sink) AddWants(levels ...Level) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for _, lvl := range levels {
		if slices.Contains(s.wants, lvl) {
			continue
		}
		s.wants = append(s.wants, lvl)
	}
}

func (s *Sink) IsFile(f *os.File) bool {
	return f.Fd() == s.f.Fd()
}

func FindSink(f *os.File) *Sink {
	sinkMu.Lock()
	defer sinkMu.Unlock()

	for _, sink := range sinks {
		if sink.f.Fd() == f.Fd() {
			return sink
		}
	}

	return nil
}
