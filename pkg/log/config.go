package log

var (
	// DefaultLogFormat is defined as
	DefaultLogFormat = "{{.Date}} {{.Time}} {{.BoldColor}}{{.Logger}} {{.Level}}{{.NoColor}} {{.Color}}{{.Message}}{{.NoColor}}"

	// DefaultDebugFormat is defined as
	DefaultDebugFormat = "{{.Date}} {{.Time}}{{.Nano}} {{.BoldColor}}{{.GID}}|{{.Logger}}|{{.LID}}{{.NoColor}} {{.Color}}{{.Message}}{{.NoColor}} {{.Func}} {{.File}}:{{.Line}}"

	// LoggerName defines the name of the standard logger.
	LoggerName = "main"

	// DefaultOutputs defines the default outputs for all loggers.
	DefaultOutputs = []string{
		"file:///dev/stdout",
		"file:///dev/stderr",
	}
)
