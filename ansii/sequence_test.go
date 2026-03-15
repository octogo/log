package ansii

import "testing"

func TestColors(t *testing.T) {
	tests := []struct {
		name     string
		color    Color
		expected string
	}{
		{"Black", BLACK, "\033[0;30m"},
		{"Red", RED, "\033[0;31m"},
		{"Green", GREEN, "\033[0;32m"},
		{"Yellow", YELLOW, "\033[0;33m"},
		{"Blue", BLUE, "\033[0;34m"},
		{"Magenta", MAGENTA, "\033[0;35m"},
		{"Cyan", CYAN, "\033[0;36m"},
		{"White", WHITE, "\033[0;37m"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Sequence(Normal, tt.color)
			if result != tt.expected {
				t.Errorf("Sequence(%s) = %s; want %q", tt.color, result, tt.expected)
			}
		})
	}
}

func TestBGColors(t *testing.T) {
	tests := []struct {
		name     string
		color    Color
		expected string
	}{
		{"Black", BLACK, "\033[0;30;40m"},
		{"Red", RED, "\033[0;30;41m"},
		{"Green", GREEN, "\033[0;30;42m"},
		{"Yellow", YELLOW, "\033[0;30;43m"},
		{"Blue", BLUE, "\033[0;30;44m"},
		{"Magenta", MAGENTA, "\033[0;30;45m"},
		{"Cyan", CYAN, "\033[0;30;46m"},
		{"White", WHITE, "\033[0;30;47m"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Sequence(Normal, BLACK, tt.color)
			if result != tt.expected {
				t.Errorf("Sequence(%s) = %q; want %q", tt.color, result, tt.expected)
			}
		})
	}
}

func TestReset(t *testing.T) {
	expected := "\033[0m"
	result := Reset()
	if result != expected {
		t.Errorf("Reset() = %q; want %q", result, expected)
	}
}

func TestAttributes(t *testing.T) {
	tests := []struct {
		name     string
		attr     Attribute
		expected string
	}{
		{"Normal", Normal, "\033[0m"},
		{"Bold", Bold, "\033[1m"},
		{"Underline", Underline, "\033[4m"},
		{"Blink", Blink, "\033[5m"},
		{"Inverse", Inverse, "\033[7m"},
		{"Invisible", Invisible, "\033[8m"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Sequence(tt.attr)
			if result != tt.expected {
				t.Errorf("Sequence(%s) = %q; want %q", tt.attr, result, tt.expected)
			}
		})
	}
}
