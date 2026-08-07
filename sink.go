package log

import (
	"os"
	"slices"
	"sync"

	"golang.org/x/term"
)

// sinks is the process-wide registry of open sinks, keyed by the file handle
// they write to.
//
// It exists so that every Logger shares one Sink per destination: DefaultSinks
// is called by each New, and AddWants on the stdout sink (see AddLevel) has to
// affect every logger that writes there. Two NewSink calls for the same handle
// must therefore return the same Sink.
//
// The key is the *os.File, not its descriptor number. A descriptor number is
// unique only among descriptors that are currently open — close a file and the
// kernel hands the same number to the next one opened. Keying on it meant that
// after a log rotation, NewSink returned the sink belonging to the file that had
// just been closed, and every subsequent line went to a closed descriptor. The
// handle is a stable identity for as long as anything can still write to it.
var (
	sinks  = map[*os.File]*Sink{}
	sinkMu = sync.Mutex{}
)

type Sink struct {
	path string
	f    *os.File

	// isTerm records whether this sink writes to a terminal, and therefore
	// whether its output should carry ANSI colour.
	//
	// It is resolved once, when the sink is created. term.IsTerminal is an
	// ioctl, and asking it per line meant a syscall on every log statement to
	// answer a question that cannot change: an open handle does not become a
	// terminal, or stop being one, while it is open.
	isTerm bool

	wants []Level
	mu    *sync.Mutex
}

// NewSink returns the Sink for a file, creating it on first use.
//
// A handle already registered yields its existing Sink, wants and all — the
// caller's argument is ignored in that case, because the registered Sink is
// shared and silently rewriting its levels would reconfigure every logger
// already using it. Use SetWants or AddWants to change them deliberately.
func NewSink(f *os.File, wants []Level) *Sink {
	// Resolved before sinkMu is taken, deliberately. AllLevels acquires lvlMu,
	// and AddLevel acquires lvlMu and then reaches into the sink registry — so
	// taking them in the other order here would complete a lock cycle and let
	// two goroutines deadlock. lvlMu is always the outer lock.
	if wants == nil {
		wants = AllLevels()
	}

	isTerm := term.IsTerminal(int(f.Fd()))

	sinkMu.Lock()
	defer sinkMu.Unlock()

	if existing, found := sinks[f]; found {
		// don't f.Close()!
		// the handle must stay alive.
		return existing
	}

	sinks[f] = &Sink{
		path:   f.Name(),
		f:      f,
		isTerm: isTerm,
		wants:  wants,
		mu:     &sync.Mutex{},
	}

	return sinks[f]
}

func (s *Sink) Path() string {
	return s.path
}

// Close closes the underlying file and removes the Sink from the registry, so
// that a later NewSink for a new handle builds a new Sink rather than finding
// this one.
//
// Closing through the Sink is what makes log rotation work. A caller that closes
// the *os.File directly leaves the entry behind: harmless in that it can only
// ever be returned for that same handle, but the Sink stays reachable and its
// writes will fail.
//
// The stdout/stderr guard compares descriptor numbers rather than handles, and
// deliberately so — this one is a question about the process, not about
// identity. Any handle sitting on descriptor 1 or 2 *is* the process's standard
// output or error, whichever pointer happens to wrap it, and closing it would
// take the stream away from everything else in the program.
func (s *Sink) Close() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.f.Fd() == os.Stdout.Fd() || s.f.Fd() == os.Stderr.Fd() {
		return nil
	}

	sinkMu.Lock()
	delete(sinks, s.f)
	sinkMu.Unlock()

	return s.f.Close()
}

func (s *Sink) Wants(level Level) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	return slices.Contains(s.wants, level)
}

func (s *Sink) Log(msg Message, fmt Formatter) (int, error) {
	line := fmt.Format(msg, s.isTerm)
	return s.Write([]byte(line))
}

// IsTerminal reports whether this sink writes to a terminal, as determined when
// it was created.
func (s *Sink) IsTerminal() bool {
	return s.isTerm
}

func (s *Sink) Write(p []byte) (int, error) {
	// os.File.Write is already concurrency-safe
	return s.f.Write(p)
}

func (s *Sink) SetWants(wants []Level) {
	// Resolved before s.mu, for the reason given in NewSink: AllLevels takes
	// lvlMu, and AddLevel holds lvlMu while calling AddWants, which takes s.mu.
	// Acquiring them in both orders is a deadlock waiting for the right timing.
	if wants == nil {
		wants = AllLevels()
	}

	s.mu.Lock()
	defer s.mu.Unlock()

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

// IsFile reports whether this Sink writes to the given file.
//
// Handles are compared, not descriptor numbers, for the reason given on the
// registry: a recycled descriptor number would make this report true for a file
// this Sink has never written to.
func (s *Sink) IsFile(f *os.File) bool {
	return s.f == f
}

// FindSink returns the registered Sink for a file, or nil if there is none.
func FindSink(f *os.File) *Sink {
	sinkMu.Lock()
	defer sinkMu.Unlock()

	return sinks[f]
}
