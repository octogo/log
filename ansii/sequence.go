package ansii

import (
	"fmt"
	"strings"
)

const seqPrefix = "\033"

func Sequence(attr Attribute, colors ...Color) string {
	colorFields := []string{}

loop:
	for _, c := range colors {
		switch len(colorFields) {
		case 0:
			colorFields = append(colorFields, c.String())
		case 1:
			colorFields = append(colorFields, (c + bgOffset).String())
		case 2:
			break loop
		}
	}

	delim := ""
	if len(colors) > 0 {
		delim = ";"
	}

	return fmt.Sprintf(
		"%s[%s%s%sm",
		seqPrefix,
		attr.String(),
		delim,
		strings.Join(colorFields, ";"),
	)
}

func Reset() string {
	return Sequence(Normal)
}

func Wrap(s string, attr Attribute, colors ...Color) string {
	return fmt.Sprintf(
		"%s%s%s",
		Sequence(attr, colors...),
		s,
		Reset(),
	)
}
