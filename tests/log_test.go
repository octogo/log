package log

import (
	"testing"

	"github.com/octogo/log/v2"
)

func TestDropinReplacement(t *testing.T) {
	log.DefaultLogger.SetFormat(log.MinimalFormat)

	t.Run("Println", func(t *testing.T) {
		result, _ := captureOutput(func() {
			log.Println("Test!")
		})
		expected := "Test!\n"
		if result != expected {
			t.Errorf("Println() => %q; want %q", result, expected)
		}
	})

	t.Run("Printf", func(t *testing.T) {
		result, _ := captureOutput(func() {
			log.Printf("This is a %s", "TEST!")
		})
		expected := "This is a TEST!\n"
		if result != expected {
			t.Errorf("Printf() => %q; want %q", result, expected)
		}
	})

	t.Run("Fatal", func(t *testing.T) {
		result := ""
		didExit, code := captureOsExit(func() {
			_, result = captureOutput(func() {
				log.Fatal("Test!")
			})
		})

		if !didExit {
			t.Error("Fatal() !os.Exit(); want os.Exit()")
		}
		if code != 1 {
			t.Errorf("Exit-code = %d; want 1", code)
		}
		expected := "Test!\n"
		if result != expected {
			t.Errorf("Fatal() = %q; want %q", result, expected)
		}
	})

	t.Run("Fatalf", func(t *testing.T) {
		result := ""
		didExit, code := captureOsExit(func() {
			_, result = captureOutput(func() {
				log.Fatalf("Another %s", "TEST!")
			})
		})
		t.Logf("stdout=%q", result)

		if !didExit {
			t.Error("Fatal() !os.Exit(); want os.Exit()")
		}
		if code != 1 {
			t.Errorf("Exit-code = %d; want 1", code)
		}
		expected := "Another TEST!\n"
		if result != expected {
			t.Errorf("Fatalf() = %q; want %q", result, expected)
		}
	})
}
