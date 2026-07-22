package log

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// captureSink writes to a temporary file and returns it as a sink together with
// a function that reads back what was written.
//
// A file rather than a buffer because Sink requires an *os.File, and the same
// property is under test either way: what actually reaches the output.
//
// The file is deliberately left open for the lifetime of the test binary. The
// sink registry is keyed by file-descriptor number and never evicts, so closing
// one capture file frees its descriptor for the next os.Create — and NewSink
// then returns the previous test's sink, still pointing at the file that was
// just closed. Leaking the handle keeps each test on a descriptor of its own.
// TempDir removes the files regardless.
func captureSink(t *testing.T) (*Sink, func() string) {
	t.Helper()

	path := filepath.Join(t.TempDir(), "capture.log")

	f, err := os.Create(path)
	if err != nil {
		t.Fatalf("create capture file: %v", err)
	}

	return NewSink(f, AllLevels()), func() string {
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
