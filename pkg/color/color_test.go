package color

import (
	"fmt"
	"testing"

	"github.com/octogo/log/pkg/level"
)

const (
	testColorBlack Color = iota + 30
	testColorRed
	testColorGreen
	testColorYellow
	_
	_
	testColorCyan
	testColorWhite
)

var (
	testColors = map[Color]Color{
		Black:  testColorBlack,
		Red:    testColorRed,
		Green:  testColorGreen,
		Yellow: testColorYellow,
		Cyan:   testColorCyan,
		White:  testColorWhite,
	}
	testLevels = map[level.Level]Color{
		level.ERROR:   testColorRed,
		level.WARNING: testColorYellow,
		level.NOTICE:  testColorGreen,
		level.INFO:    testColorWhite,
		level.DEBUG:   testColorCyan,
	}
)

func TestColors(t *testing.T) {
	for c, tc := range testColors {
		if c != tc {
			t.Errorf("expected %v, got %v", tc, c)
		}
	}
}

func TestColorSeq(t *testing.T) {
	for c, tc := range testColors {
		if Seq(c) != fmt.Sprintf("\033[%dm", tc) {
			t.Errorf("expected %v, got %v", c, tc)
		}
	}
}

func TestColorSeqBold(t *testing.T) {
	for c, tc := range testColors {
		if colorSeqBold(c) != fmt.Sprintf("\033[%d;1m", c) {
			t.Errorf("expected %v, got %v", c, tc)
		}
	}
}

func TestColorSeqReset(t *testing.T) {
	expected := fmt.Sprintf("\033[0m")
	got := ResetSeq()
	if got != expected {
		t.Errorf("expected %v, got %v", expected, got)
	}
}

func TestLevelColors(t *testing.T) {
	for l, c := range testLevels {
		if Colors[l] != Seq(c) {
			t.Errorf("expected %v, got %v", c, Colors[l])
		}
	}
}

func TestLevelColorsBold(t *testing.T) {
	for l, c := range testLevels {
		if BoldColors[l] != colorSeqBold(c) {
			t.Errorf("expected: %v, got %v", c, BoldColors[l])
		}
	}
}
