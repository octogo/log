// Package color provides an interface for logging text in ANSII colors.
package color

import (
	"fmt"
	"strings"
)

var escapeSeq = "\033"

// Sequence is defined as
type Sequence interface {
	SetAttribute(Attribute) error
	SetColors([]Color) error
	IsValid() bool

	fmt.Stringer
}

type sequence struct {
	Literal   string
	Attribute Attribute
	Colors    []Color
}

func (seq sequence) String() string {
	if seq.Literal != "" {
		return strings.Join([]string{
			escapeSeq,
			"[",
			seq.Literal,
			"m",
		}, "")
	}
	out := escapeSeq
	out += "["
	out += seq.Attribute.String()
	if seq.Colors != nil && len(seq.Colors) > 0 {
		out += ";"
	}
	colors := make([]string, len(seq.Colors))
	for i := range seq.Colors {
		colors[i] = seq.Colors[i].String()
	}
	out += strings.Join(colors, ";")
	out += "m"
	return out
}

func (seq *sequence) SetAttribute(attr Attribute) error {
	var attrValid bool
	for i := range Attributes {
		if Attributes[i] == attr {
			attrValid = true
			break
		}
	}
	if !attrValid {
		return fmt.Errorf("unsupported ANSII attribute: %d", attr)
	}
	seq.Attribute = attr
	return nil
}

func (seq *sequence) SetColors(colors []Color) error {
	for i := range colors {
		colorValid := false
		for j := range Colors {
			if Colors[j] == colors[i] {
				colorValid = true
				break
			}
		}
		if !colorValid {
			return fmt.Errorf("invalid ANSII color code: %d", colors[i])
		}
	}
	seq.Colors = colors
	return nil
}

// IsValid returns true if this ANSII sequence is valid in terms of its
// configured attributes.
func (seq sequence) IsValid() bool {
	attrValid := false
	for i := range Attributes {
		if Attributes[i] == seq.Attribute {
			attrValid = true
			break
		}
	}
	if !attrValid {
		return false
	}
	return true
}

// Attribute is defined as a type of int.
type Attribute int

func (attr Attribute) String() string {
	return fmt.Sprintf("%d", attr)
}

// Attributes
const (
	NormalDisplay Attribute = 0
	Bold          Attribute = 1
	Underline     Attribute = 4
	Blink         Attribute = 5
	ReverseVideo  Attribute = 7
	Invisible     Attribute = 8
)

// Attributes contains the list of all supported attributes.
var Attributes = []Attribute{
	NormalDisplay,
	Bold,
	Underline,
	Blink,
	ReverseVideo,
	Invisible,
}

// Color is defined as a type of int.
type Color int

func (c Color) String() string {
	return fmt.Sprintf("%d", c)
}

// Colors
const (
	// foreground colors
	Black   Color = 30
	Red     Color = 31
	Green   Color = 32
	Yellow  Color = 33
	Blue    Color = 34
	Magenta Color = 35
	Cyan    Color = 36
	White   Color = 37
	// background colors
	BGBlack   Color = 40
	BGRed     Color = 41
	BGGreen   Color = 42
	BGYellow  Color = 43
	BGBlue    Color = 44
	BGMagenta Color = 45
	BGCyan    Color = 46
	BGWhite   Color = 47
)

// Colors contains the list of all supported FGColors and BGColors
var (
	FGColors = []Color{
		Black,
		Red,
		Green,
		Yellow,
		Blue,
		Magenta,
		Cyan,
		White,
	}
	BGColors = []Color{
		BGBlack,
		BGRed,
		BGGreen,
		BGYellow,
		BGBlue,
		BGMagenta,
		BGCyan,
		BGWhite,
	}
	Colors = append(FGColors, BGColors...)
)

// New returns an ANSII escape sequence based on the given attributes and
// colors.
func New(attr Attribute, colors ...Color) Sequence {
	return &sequence{
		Attribute: attr,
		Colors:    colors,
	}
}

// NewLiteral returns a literal ANSII escape sequence based on the given
// string.
func NewLiteral(literal string) Sequence {
	return &sequence{
		Literal: literal,
	}
}
