package octolog

var (
	// DefaultLoggerName defines the default name for all loggers.
	DefaultLoggerName = "octolog"

	// DefaultLogFormat defines a basic default format.
	DefaultLogFormat = "{{.Date}} {{.TimeExact}} [{{.PID}}]:{{.GID}} {{.Logger}}:{{.LID}} {{.Level}} ▶ {{.Body}}"

	defaultInternalLogFormat = "{{.Body}}"
)
