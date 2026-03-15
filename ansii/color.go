package ansii

import "fmt"

// Color is an integer representing an ANSII color-code.
type Color int

func (c Color) String() string {
	return fmt.Sprintf("%d", c)
}

const (
	BLACK    Color = 30
	RED      Color = 31
	GREEN    Color = 32
	YELLOW   Color = 33
	BLUE     Color = 34
	MAGENTA  Color = 35
	CYAN     Color = 36
	WHITE    Color = 37
	bgOffset Color = 10 // for conveniently converting to background colors
)

var (
	Colors = []Color{
		BLACK,
		RED,
		GREEN,
		YELLOW,
		BLUE,
		MAGENTA,
		CYAN,
		WHITE,
	}
)
