// Package color provides routines for cooring ASCII test with ANSII escape
// sequences. See: https://en.wikipedia.org/wiki/ANSI_escape_code
//
// A color, in the sense of this package, is an integer which directly
// translates to the corresponding ANSII sequence numer.
//
// Some colors are pre-defined, such as BLACK(30), RED(31), GREEN(32),
// YELOW(33), BLUE(34), MAGENTA(35), CYAN(36) and WHITE(37).
//
// Easily change the mapping between log-level and color like:
//
//		Colors[level.ERROR] = colors.Red
//
// Implements you own colors easily:
//
// 		var MyColor int = "\u001b[48;5;"
//		Colors[level.ERROR] = MyColor
//
package color

import (
	"fmt"
	"os"

	"github.com/octogo/log/pkg/level"
	"golang.org/x/sys/unix"
)

// Color is defined as a type of int.
type Color int

// Colors supported by this package.
const (
	Black Color = iota + 30
	Red
	Green
	Yellow
	Blue
	Magenta
	Cyan
	White
)

// ResetSeq return a string resetting all ANSI colors.
func ResetSeq() string {
	return fmt.Sprint("\033[0m")
}

// Seq returns the ANSII color escape sequence for the given color.
func Seq(Color Color) string {
	return fmt.Sprintf("\033[%dm", int(Color))
}

// SeqBold returns the ANSII color escape sequence for the given BOLD color.
func SeqBold(Color Color) string {
	return fmt.Sprintf("\033[%d;1m", int(Color))
}

func isTerminal(file *os.File) bool {
	_, err := unix.IoctlGetTermios(int(file.Fd()), unix.TCGETS)
	return err == nil
}

// Colors contains all ANSI Color escape sequences.
var Colors = map[level.Level]Color{
	level.ERROR:   Red,
	level.WARNING: Yellow,
	level.NOTICE:  Green,
	level.INFO:    White,
	level.DEBUG:   Cyan,
}

// Colorize takes a level.Level{} and an interface{} and then wraps the string
// representation of the interface in the color configured for the given
// level.Level{}.
func Colorize(l level.Level, v interface{}) string {
	return Seq(Colors[l]) + fmt.Sprintf("%s", v) + ResetSeq()
}

// ColorizeBold takes a level.Level and an interface{} and then wraps the
// string representation of the interface{} in the bold color configured for
// the given level.Level.
func ColorizeBold(l level.Level, v interface{}) string {
	return SeqBold(Colors[l]) + fmt.Sprintf("%s", v) + ResetSeq()
}
