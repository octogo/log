package log

import (
	"testing"

	"github.com/octogo/log/v2"
	"github.com/octogo/log/v2/ansii"
)

func TestLevels(t *testing.T) {
	tests := []struct {
		name     string
		lvl      log.Level
		expected string
	}{
		{"ERROR", log.ERROR, "ERROR"},
		{"WARNING", log.WARNING, "WARNING"},
		{"NOTICE", log.NOTICE, "NOTICE"},
		{"INFO", log.INFO, "INFO"},
		{"DEBUG", log.DEBUG, "DEBUG"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := string(tt.lvl)
			if result != tt.expected {
				t.Errorf("string(%s) = %s; want = %s", tt.lvl, result, tt.expected)
			}
		})
	}
}

func TestColors(t *testing.T) {
	tests := []struct {
		name     string
		lvl      log.Level
		expected ansii.Color
	}{
		{"ERROR", log.ERROR, ansii.RED},
		{"WARNING", log.WARNING, ansii.YELLOW},
		{"NOTICE", log.NOTICE, ansii.GREEN},
		{"INFO", log.INFO, ansii.WHITE},
		{"DEBUG", log.DEBUG, ansii.CYAN},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := log.ColorOf(tt.lvl)
			if result != tt.expected {
				t.Errorf("LevelColors[%s] = %s; want %s", tt.lvl, result, tt.expected)
			}
		})
	}
}

func TestCustom(t *testing.T) {
	CUSTOM := log.Level("CUSTOM")

	t.Run("Create custom", func(t *testing.T) {
		log.AddLevel(string(CUSTOM), ansii.MAGENTA)
		levels := log.AllLevels()
		if len(levels) != 6 {
			t.Errorf("All() = %v; want %v", levels, append(levels, CUSTOM))
		}
	})

	t.Run("Get color", func(t *testing.T) {
		color := log.ColorOf(CUSTOM)
		if color != ansii.MAGENTA {
			t.Errorf("Color(CUSTOM) = %s; want %s", color, ansii.MAGENTA)
		}
	})

	t.Run("Change color", func(t *testing.T) {
		log.ChangeColor(CUSTOM, ansii.WHITE)
		color := log.ColorOf(CUSTOM)
		if color != ansii.WHITE {
			t.Errorf("Color(CUSTOM) = %s; want %s", color, ansii.WHITE)
		}
	})
}
