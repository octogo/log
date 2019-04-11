package log

import (
	"fmt"
	"os"

	"golang.org/x/crypto/ssh/terminal"
	"golang.org/x/sys/unix"
)

type color int

// Colors supported by this package.
const (
	ColorBlack color = iota + 30
	ColorRed
	ColorGreen
	ColorYellow
	ColorBlue
	ColorMagenta
	ColorCyan
	ColorWhite
)

var (
	// Colors contains all ANSI color escape sequences, ordered by log level criticality.
	Colors = []string{
		ERROR:   colorSeq(ColorRed),
		WARNING: colorSeq(ColorYellow),
		ALERT:   colorSeq(ColorMagenta),
		NOTICE:  colorSeq(ColorGreen),
		DEBUG:   colorSeq(ColorCyan),
	}

	// BoldColors contais all bold ANSI color escape sequences, ordered by log level criticality.
	BoldColors = []string{
		ERROR:   colorSeqBold(ColorRed),
		WARNING: colorSeqBold(ColorYellow),
		ALERT:   colorSeqBold(ColorMagenta),
		NOTICE:  colorSeqBold(ColorGreen),
		DEBUG:   colorSeqBold(ColorCyan),
	}

	// SmartTTY enables automatically disabling all colors if stdout is not a TTY.
	SmartTTY    = true
	stdoutIsTTY = terminal.IsTerminal(int(os.Stdout.Fd()))
)

func colorSeq(color color) string {
	return fmt.Sprintf("\033[%dm", int(color))
}

func colorSeqBold(color color) string {
	return fmt.Sprintf("\033[%d;1m", int(color))
}

func colorSeqReset() string {
	return fmt.Sprint("\033[0m")
}

func isTerminal(file *os.File) bool {
	_, err := unix.IoctlGetTermios(int(file.Fd()), unix.TCGETS)
	return err == nil
}
