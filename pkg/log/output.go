package log

import "github.com/octogo/log/pkg/level"

// Output is defined as
type Output interface {
	Type() string
	URI() string
	URL() string
	Log(Entry) (int, error)
	SetFormat(string)
	SetWants([]level.Level)
}
