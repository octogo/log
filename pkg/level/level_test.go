package level

import (
	"testing"

	"github.com/octogo/log/pkg/color"
)

func TestBuiltinLevels(t *testing.T) {
	if ERROR != Level(0) {
		t.Errorf("expected %v, got %v", 0, ERROR)
	}
	if WARNING != Level(1) {
		t.Errorf("expected %v, got %v", 1, WARNING)
	}
	if NOTICE != Level(2) {
		t.Errorf("expected %v, got %v", 2, NOTICE)
	}
	if INFO != Level(3) {
		t.Errorf("expected %v, got %v", 3, INFO)
	}
	if DEBUG != Level(4) {
		t.Errorf("expected %v, got %v", 4, DEBUG)
	}
}

func TestRegister(t *testing.T) {
	if lvl, _, _ := Register("ERROR", color.New(color.NormalDisplay, color.Red)); lvl != ERROR {
		t.Errorf("expected %v, got %v", ERROR, lvl)
	}
	if lvl, _, _ := Register("debug", color.New(color.NormalDisplay, color.Yellow)); lvl != DEBUG {
		t.Errorf("ecxpected %v, got %v", DEBUG, lvl)
	}
	if lvl, _, _ := Register("custom", color.New(color.NormalDisplay, color.Green)); int(lvl) != 5 {
		t.Errorf("expected %v, got %v", 5, lvl)
	}
}

func TestLevels(t *testing.T) {
	if len(Levels()) != len(registeredLevels) {
		t.Errorf("expected %v, got %v", len(registeredLevels), len(Levels()))
	}
}

func TestIsValid(t *testing.T) {
	if IsValid(Level(99)) {
		t.Errorf("expected %v, got %v", false, true)
	}
	testLevel, _, _ := Register("test_1", color.New(color.NormalDisplay, color.Magenta))
	if !IsValid(testLevel) {
		t.Errorf("expected %v, got %v", true, false)
	}
}

func TestIsValidName(t *testing.T) {
	if IsValidName("test_2") {
		t.Errorf("expected %v, got %v", false, true)
	}
	Register("test_2", color.New(color.NormalDisplay, color.Magenta))
	if !IsValidName("test_2") {
		t.Errorf("expected %v, got %v", true, false)
	}
}

func TestParse(t *testing.T) {
	var (
		level Level
		err   error
	)
	level, err = Parse("eRrOr")
	if err != nil {
		t.Errorf("expected %v, got %v", nil, err)
	}
	if level != ERROR {
		t.Errorf("expected %v, got %v", ERROR, level)
	}
}
