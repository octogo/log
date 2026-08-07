package log

import (
	"fmt"
	"strings"
)

const seqPrefix = "\033"

func AnsiiSequence(attr AnsiiAttr, colors ...AnsiiColor) string {
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

func ansiiReset() string {
	return AnsiiSequence(Normal)
}
