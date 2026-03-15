package ansii

import "fmt"

// Attribute is an integer representing the ANSII attribute-code.
type Attribute int

// String implements fmt.Stringer
func (attr Attribute) String() string {
	return fmt.Sprintf("%d", attr)
}

// ANSII attributes
const (
	Normal    Attribute = 0
	Bold      Attribute = 1
	Underline Attribute = 4
	Blink     Attribute = 5
	Inverse   Attribute = 7
	Invisible Attribute = 8
)

// Attributes contains the list of all supported attributes.
var Attributes = []Attribute{
	Normal,
	Bold,
	Underline,
	Blink,
	Inverse,
	Invisible,
}
