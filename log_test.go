package log

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"testing"
	"time"
)

// captureSink writes to a temporary file and returns it as a sink together with
// a function that reads back what was written.
//
// A file rather than a buffer because Sink requires an *os.File, and the same
// property is under test either way: what actually reaches the output.
func captureSink(t *testing.T) (*Sink, func() string) {
	t.Helper()

	path := filepath.Join(t.TempDir(), "capture.log")

	f, err := os.Create(path)
	if err != nil {
		t.Fatalf("create capture file: %v", err)
	}

	sink := NewSink(f, AllLevels())
	t.Cleanup(func() { sink.Close() })

	return sink, func() string {
		data, err := os.ReadFile(path)
		if err != nil {
			t.Fatalf("read capture file: %v", err)
		}

		return string(data)
	}
}

// TestNewPopulatesFormatter guards a failure that leaves no trace of itself.
//
// New must set the exported Formatter field, not only the private compiled
// template. Spawn builds a child by passing that field through SetFormat, and
// SetFormat("") compiles an empty template — which is valid, and renders every
// line as the empty string. A logger whose Formatter is empty therefore works
// perfectly while silently discarding everything its children log.
func TestNewPopulatesFormatter(t *testing.T) {
	logger := New("parent", nil)

	if logger.Formatter == "" {
		t.Fatal("New left Formatter empty; every logger spawned from this one will render nothing")
	}

	if logger.Formatter != DefaultFormat {
		t.Errorf("Formatter = %q, want DefaultFormat %q", logger.Formatter, DefaultFormat)
	}
}

// TestSpawnedLoggersRender is the same bug seen from the outside: not a field
// left empty, but output that disappears.
func TestSpawnedLoggersRender(t *testing.T) {
	sink, captured := captureSink(t)

	parent := New("parent", AllLevels(), sink)
	child := parent.Spawn("child")
	grandchild := child.Spawn("grandchild")

	parent.Println("from parent")
	child.Println("from child")
	grandchild.Println("from grandchild")

	out := captured()

	for _, want := range []string{"from parent", "from child", "from grandchild"} {
		if !strings.Contains(out, want) {
			t.Errorf("%q missing from output; a spawned logger discarded it\ngot:\n%s", want, out)
		}
	}
}

// TestSpawnInheritsExplicitFormat confirms a format set on the parent still
// reaches children — the behaviour Spawn exists for, which must survive the fix
// to New.
func TestSpawnInheritsExplicitFormat(t *testing.T) {
	sink, captured := captureSink(t)

	parent := New("parent", AllLevels(), sink)
	parent.SetFormat("[{{.Logger}}] {{.Message}}")

	child := parent.Spawn("child")

	if child.Formatter != parent.Formatter {
		t.Errorf("child Formatter = %q, want the parent's %q", child.Formatter, parent.Formatter)
	}

	child.Println("hello")

	if out := captured(); !strings.Contains(out, "[parent.child] hello") {
		t.Errorf("child did not use the inherited format\ngot: %s", out)
	}
}

// TestMessagesAreNotHtmlEscaped guards the output context.
//
// Rendering with html/template escapes quotes, angle brackets and ampersands
// into entities, mangling precisely the messages that were quoted because their
// exact text mattered — a CID, an identifier or a filename logged with %q
// arrives as &#34;value&#34;. Nothing this package writes to is an HTML
// document.
func TestMessagesAreNotHtmlEscaped(t *testing.T) {
	sink, captured := captureSink(t)

	logger := New("test", AllLevels(), sink)

	logger.Printf("value is %q and it's <fine> & unescaped", "quoted")

	out := captured()

	for _, entity := range []string{"&#34;", "&#39;", "&lt;", "&gt;", "&amp;"} {
		if strings.Contains(out, entity) {
			t.Errorf("output contains HTML entity %s; messages are being escaped\ngot: %s", entity, out)
		}
	}

	if !strings.Contains(out, `value is "quoted" and it's <fine> & unescaped`) {
		t.Errorf("message did not survive verbatim\ngot: %s", out)
	}
}

// TestSinkIsNotReusedAcrossFiles is the log-rotation bug.
//
// Descriptor numbers are unique only among descriptors currently open. Close a
// file and the kernel hands the same number to the next one opened, so a
// registry keyed on that number returned the *previous* file's Sink — and every
// line written afterwards went to a closed descriptor. Rotation is exactly this
// sequence, and the failure is silent until someone checks the writes.
func TestSinkIsNotReusedAcrossFiles(t *testing.T) {
	dir := t.TempDir()

	oldPath := filepath.Join(dir, "old.log")
	newPath := filepath.Join(dir, "new.log")

	old, err := os.Create(oldPath)
	if err != nil {
		t.Fatalf("create old: %v", err)
	}

	oldSink := NewSink(old, AllLevels())
	oldFd := old.Fd()

	if err := oldSink.Close(); err != nil {
		t.Fatalf("close old sink: %v", err)
	}

	// The freed descriptor number is normally handed straight back.
	fresh, err := os.Create(newPath)
	if err != nil {
		t.Fatalf("create new: %v", err)
	}
	defer fresh.Close()

	newSink := NewSink(fresh, AllLevels())

	if newSink == oldSink {
		t.Fatal("NewSink returned the closed file's sink for a different file")
	}

	if !newSink.IsFile(fresh) {
		t.Error("the new sink does not report the file it was created for")
	}

	if newSink.IsFile(old) {
		t.Error("the new sink claims to write to the closed file")
	}

	logger := New("rotate", AllLevels(), newSink)
	logger.Println("after rotation")

	data, err := os.ReadFile(newPath)
	if err != nil {
		t.Fatalf("read new: %v", err)
	}

	if !strings.Contains(string(data), "after rotation") {
		t.Errorf("line did not reach the rotated file\ngot: %q", data)
	}

	// Worth stating plainly: this test only proves anything when the descriptor
	// number really was recycled, which is the ordinary case but not guaranteed.
	if fresh.Fd() != oldFd {
		t.Logf("descriptor not recycled (%d -> %d); the collision case went untested this run", oldFd, fresh.Fd())
	}
}

// TestCloseUnregistersSink covers the eviction that makes the above work.
func TestCloseUnregistersSink(t *testing.T) {
	f, err := os.Create(filepath.Join(t.TempDir(), "closing.log"))
	if err != nil {
		t.Fatalf("create: %v", err)
	}

	sink := NewSink(f, AllLevels())

	if FindSink(f) != sink {
		t.Fatal("a new sink is not findable by its file")
	}

	if err := sink.Close(); err != nil {
		t.Fatalf("close: %v", err)
	}

	if got := FindSink(f); got != nil {
		t.Error("a closed sink is still registered; its file handle can never be reused")
	}
}

// TestSinksAreSharedPerHandle is the property the registry exists for: every
// Logger must write to the same stdout Sink, so that adjusting its levels — as
// AddLevel does — affects all of them rather than one.
func TestSinksAreSharedPerHandle(t *testing.T) {
	if NewSink(os.Stdout, nil) != NewSink(os.Stdout, nil) {
		t.Error("two NewSink calls for os.Stdout returned different sinks")
	}

	if FindSink(os.Stdout) != NewSink(os.Stdout, nil) {
		t.Error("FindSink and NewSink disagree about the stdout sink")
	}

	if NewSink(os.Stdout, nil) == NewSink(os.Stderr, nil) {
		t.Error("stdout and stderr share a sink")
	}
}

// TestCloseLeavesStandardStreamsOpen guards the one place a descriptor number is
// the right question: any handle on descriptor 1 or 2 is the process's own
// output, and closing it would take the stream from the rest of the program.
func TestCloseLeavesStandardStreamsOpen(t *testing.T) {
	for _, f := range []*os.File{os.Stdout, os.Stderr} {
		if err := NewSink(f, nil).Close(); err != nil {
			t.Errorf("closing the %s sink returned %v", f.Name(), err)
		}

		if _, err := f.Write(nil); err != nil {
			t.Fatalf("%s was closed: %v", f.Name(), err)
		}
	}
}

// TestSinkTerminalStateIsResolvedOnce covers both halves of the colour
// decision: it is answered from the sink rather than by asking the OS per line,
// and a file is correctly not a terminal, so no ANSI escapes reach a log file.
func TestSinkTerminalStateIsResolvedOnce(t *testing.T) {
	sink, captured := captureSink(t)

	if sink.IsTerminal() {
		t.Error("a regular file was detected as a terminal")
	}

	logger := New("colour", AllLevels(), sink)
	logger.Error("no colour here")

	if out := captured(); strings.Contains(out, "\x1b[") {
		t.Errorf("ANSI escape written to a file\ngot: %q", out)
	}
}

// TestAddLevelWithoutStdoutSink covers a nil dereference in the one function a
// program calls while configuring how it reports things — the worst place to
// crash, because nothing is set up yet to say why.
//
// The stdout sink is normally registered during package initialisation, so the
// registry is emptied here to reach the branch at all.
func TestAddLevelWithoutStdoutSink(t *testing.T) {
	sinkMu.Lock()
	saved, hadSink := sinks[os.Stdout]
	delete(sinks, os.Stdout)
	sinkMu.Unlock()

	t.Cleanup(func() {
		sinkMu.Lock()
		if hadSink {
			sinks[os.Stdout] = saved
		}
		sinkMu.Unlock()

		lvlMu.Lock()
		delete(levelColors, Level("TESTLEVEL"))
		lvlMu.Unlock()
	})

	// Must not panic.
	if got := AddLevel("TESTLEVEL", CYAN); got != Level("TESTLEVEL") {
		t.Errorf("AddLevel returned %q, want TESTLEVEL", got)
	}
}

// TestLevelAndSinkLocksDoNotDeadlock guards a lock-order inversion.
//
// AllLevels takes lvlMu, and NewSink and SetWants both called it while already
// holding sinkMu or a sink's own mutex. AddLevel goes the other way: it holds
// lvlMu and then reaches into the sink registry. Two goroutines taking those
// pairs in opposite orders wedge permanently — and it is timing-dependent, so it
// surfaces in production rather than in development.
//
// Against the unfixed code this test hangs rather than failing, which the
// timeout converts into a failure.
// deadlockIterations is tuned by experiment: the inversion needs both goroutines
// to interleave inside a window of a few instructions, and at 40 iterations the
// unfixed code passed comfortably. At this count it hangs every time.
const deadlockIterations = 2000

func TestLevelAndSinkLocksDoNotDeadlock(t *testing.T) {
	t.Cleanup(func() {
		lvlMu.Lock()
		for i := range deadlockIterations {
			delete(levelColors, Level(fmt.Sprintf("RACE%d", i)))
		}
		lvlMu.Unlock()
	})

	done := make(chan struct{})

	go func() {
		defer close(done)

		var wg sync.WaitGroup
		wg.Add(2)

		go func() {
			defer wg.Done()
			for i := range deadlockIterations {
				AddLevel(fmt.Sprintf("RACE%d", i), CYAN)
			}
		}()

		go func() {
			defer wg.Done()
			for range deadlockIterations {
				// wants == nil is the path that reached AllLevels while a sink
				// lock was held.
				NewSink(os.Stdout, nil).SetWants(nil)
			}
		}()

		wg.Wait()
	}()

	select {
	case <-done:
	case <-time.After(10 * time.Second):
		t.Fatal("deadlock: lvlMu and the sink locks are being acquired in both orders")
	}
}

// TestWithSinkReusesStandardStreams confirms that naming a standard stream by
// path yields the same shared sink as passing the handle.
//
// WithSink used to re-wrap the descriptor with os.NewFile, producing a second
// handle to the same stream. Since sinks are identified by handle, that would
// now register a duplicate stdout sink: every line written twice, and invisible
// to the FindSink(os.Stdout) lookup that AddLevel uses to reach it.
func TestWithSinkReusesStandardStreams(t *testing.T) {
	if got := WithSink(os.Stdout.Name(), nil); got != NewSink(os.Stdout, nil) {
		t.Error("WithSink by path returned a different sink than the stdout handle")
	}

	if got := WithSink(os.Stderr.Name(), nil); got != NewSink(os.Stderr, nil) {
		t.Error("WithSink by path returned a different sink than the stderr handle")
	}
}
