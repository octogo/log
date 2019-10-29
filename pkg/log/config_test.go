package log

import (
	"testing"
)

func TestConfigDefaults(t *testing.T) {
	assertEquals := func(a, b string) {
		if a != b {
			t.Errorf("expected %v, got %v", b, a)
		}
	}
	assertEquals(
		DefaultLogFormat,
		"{{.Date}} {{.Time}} {{.BoldColor}}{{.Logger}} {{.Level}}{{.NoColor}} {{.Color}}{{.Message}}{{.NoColor}}",
	)
	assertEquals(
		DefaultDebugFormat,
		"{{.Date}} {{.Time}}{{.Nano}} {{.BoldColor}}{{.GID}}|{{.Logger}}|{{.LID}}{{.NoColor}} {{.Color}}{{.Message}}{{.NoColor}} {{.Func}} {{.File}}:{{.Line}}",
	)
	assertEquals(
		LoggerName,
		"main",
	)
	assertEquals(
		DefaultOutputs[0],
		"file:///dev/stdout",
	)
	assertEquals(
		DefaultOutputs[1],
		"file:///dev/stderr",
	)
}
