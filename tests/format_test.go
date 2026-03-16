package log

import (
	"testing"
	"time"

	"github.com/octogo/log/v2"
)

func TestFormat(t *testing.T) {
	DefaultFormatExpected := "{{.Date}} {{.Time}} {{.Color}}{{.Logger}}{{.NoColor}} [{{.BoldColor}}{{.Level}}{{.NoColor}}] {{.Message}}"
	DebugFormatExpected := "{{.Date}} {{.Time}}{{.Nano}} {{.BoldColor}}{{.Logger}}{{.NoColor}} {{.Color}}{{.Message}}{{.NoColor}} [{{.File}}:{{.Line}} ({{.Caller}})]"
	MinimalFormatExpected := "{{.Message}}"

	t.Run("Test DefaultFormat", func(t *testing.T) {
		if log.DefaultFormat != DefaultFormatExpected {
			t.Errorf("DefaultFormat = %s; want %s", log.DefaultFormat, DefaultFormatExpected)
		}
	})

	t.Run("DebugFormat", func(t *testing.T) {
		if log.DebugFormat != DebugFormatExpected {
			t.Errorf("DebugFormat = %s; want %s", log.DebugFormat, DebugFormatExpected)
		}
	})

	t.Run("MinimalFormat", func(t *testing.T) {
		if log.MinimalFormat != MinimalFormatExpected {
			t.Errorf("MinimalFormat = %s; want %s", log.MinimalFormat, MinimalFormatExpected)
		}
	})

	t.Run("NewFormatter", func(t *testing.T) {
		l := log.New("testing", log.WantsAllLevels(), log.DebugFormat, log.DefaultSinks()...)
		fmt := log.NewFormatter(log.DebugFormat)
		now := time.Now()
		m := log.Message{
			Timestamp: now,
			Logger:    l,
			Level:     log.DEBUG,
			Msg:       "TEST!",
		}

		result := fmt.Format(m, false)
		expected := log.NewFormatter(DebugFormatExpected).Format(m, false)

		if result != expected {
			t.Errorf("New() = %v; want %v", result, expected)
		}
	})
}
