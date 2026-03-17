package log

import "fmt"

// AnsiiAttr is an integer representing the ANSII attribute-code.
type AnsiiAttr int

// String implements fmt.Stringer
func (attr AnsiiAttr) String() string {
	return fmt.Sprintf("%d", attr)
}

// ANSII attributes
const (
	Normal    AnsiiAttr = 0
	Bold      AnsiiAttr = 1
	Underline AnsiiAttr = 4
	Blink     AnsiiAttr = 5
	Inverse   AnsiiAttr = 7
	Invisible AnsiiAttr = 8
)

// Attributes contains the list of all supported attributes.
var Attributes = []AnsiiAttr{
	Normal,
	Bold,
	Underline,
	Blink,
	Inverse,
	Invisible,
}
