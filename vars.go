package log

var (
	// DefaultLogFormat defines a very basic default log format.
	DefaultLogFormat = "{{.Msg}}"

	// DebugLogFormat defines a more verbose default format.
	DebugLogFormat = "{{.Date}} {{.TimeExact}} [{{.PID}}]:{{.GID}} {{.Logger}}:{{.LID}} |{{.Level}}| {{.Msg}}"
)
