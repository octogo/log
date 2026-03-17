package log

import "fmt"

// AnsiiColor is an integer representing an ANSII color-code.
type AnsiiColor int

func (c AnsiiColor) String() string {
	return fmt.Sprintf("%d", c)
}

const (
	BLACK    AnsiiColor = 30
	RED      AnsiiColor = 31
	GREEN    AnsiiColor = 32
	YELLOW   AnsiiColor = 33
	BLUE     AnsiiColor = 34
	MAGENTA  AnsiiColor = 35
	CYAN     AnsiiColor = 36
	WHITE    AnsiiColor = 37
	bgOffset AnsiiColor = 10 // for conveniently converting to background colors
)

var (
	Colors = []AnsiiColor{
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
