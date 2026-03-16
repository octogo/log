package log

import (
	"os"
	"testing"

	"github.com/octogo/log/v2"
)

func TestLogger(t *testing.T) {
	t.Run("Initialized", func(t *testing.T) {
		if log.DefaultLogger == nil {
			t.Errorf("DefaultLogger = nil; want !nil")
		}
	})

	t.Run("Name", func(t *testing.T) {
		result := log.DefaultLogger.Name
		expected := "main"
		if result != expected {
			t.Errorf("DefaultLogger.Name = %s; want %s", result, expected)
		}
	})

	t.Run("Setname", func(t *testing.T) {
		log.DefaultLogger.SetName("testing")
		result := log.DefaultLogger.Name
		expected := "testing"
		if result != expected {
			t.Errorf("SetName() -> %s; want %s", result, expected)
		}
	})

	t.Run("SetFormat", func(t *testing.T) {
		log.DefaultLogger.SetFormat("testing")
		result := log.DefaultLogger.Formatter
		expected := "testing"
		if result != expected {
			t.Errorf("SetFormat -> %s; want %s", result, expected)
		}
	})

	t.Run("SetWants", func(t *testing.T) {
		tests := []struct {
			name     string
			level    log.Level
			expected []log.Level
		}{
			{"DEBUG", log.DEBUG, []log.Level{log.DEBUG}},
			{"INFO", log.INFO, []log.Level{log.INFO}},
			{"NOTICE", log.NOTICE, []log.Level{log.NOTICE}},
			{"WARNING", log.WARNING, []log.Level{log.WARNING}},
			{"ERROR", log.ERROR, []log.Level{log.ERROR}},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				log.DefaultLogger.SetWants([]log.Level{tt.level})
				result := log.DefaultLogger.Wants
				if len(result) != 1 {
					t.Errorf("len(SetWants(%s)) = %d; want 1", tt.level, len(result))
				}
				if result[0] != tt.expected[0] {
					t.Errorf("SetWants(%s) = %s; want %s", tt.level, result, tt.expected)
				}
			})
		}
	})

	t.Run("AddSinks", func(t *testing.T) {
		t.Run("Before testing AddSink", func(t *testing.T) {
			result := len(log.DefaultLogger.Sinks)
			expected := 2
			if result != expected {
				t.Errorf("len(AddSinks()) = %d; want %d", result, expected)
			}
		})

		expected := log.WithSink("test.log", log.WantsAllLevels())
		log.DefaultLogger.AddSinks(expected)
		result := log.DefaultLogger.Sinks[len(log.DefaultLogger.Sinks)-1]
		if result != expected {
			t.Errorf("AddSink() -> %v; want %v", result, expected)
		}

		if _, err := os.Stat("test.log"); os.IsNotExist(err) {
			t.Error("os.Stat(test.log) os.NotExists; want !os.NotExists")
		}
	})

	t.Run("SetSinks", func(t *testing.T) {
		log.DefaultLogger.SetSinks(
			append(
				log.DefaultSinks(),
				log.WithSink("test.log", log.WantsAllLevels()),
				log.WithSink("test.err", log.WantsAllLevels()),
			)...,
		)

		t.Run("Should have 4 sinks", func(t *testing.T) {
			result := len(log.DefaultLogger.Sinks)
			expected := 4
			if result != expected {
				t.Errorf("len(DefaultLogger.Sinks) = %d; want %d", result, expected)
			}
		})

		if _, err := os.Stat("test.err"); os.IsNotExist(err) {
			t.Error("os.Stat(test.err) => os.NotExists; want !os.NotExists")
		}
	})
}
