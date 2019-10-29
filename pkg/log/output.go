package log

import "github.com/octogo/log/pkg/level"

// Output is defined as
type Output interface {
	Type() string           // i.e.: file://
	URI() string            // i.e.: debug.log
	URL() string            // string(Type() + URI())
	Log(Entry) (int, error) // Logs the given entry
	SetFormat(string)       // sets the log-format for this output to the given string
	SetWants([]level.Level) // sets the whitelisted log-levels (nil implies 'all')
}
