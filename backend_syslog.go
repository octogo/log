package octolog

import "log/syslog"

// SyslogBackend defines a backend logging to syslog.
type SyslogBackend struct {
	BaseBackend
	writer *syslog.Writer
}

// NewSyslogBackend returns a newly initialized SyslogBackend.
func NewSyslogBackend(format string, levels LevelSlice, priority syslog.Priority, tag string) (*SyslogBackend, error) {
	writer, err := syslog.New(priority, tag)
	if err != nil {
		return nil, err
	}

	return &SyslogBackend{
		BaseBackend: BaseBackend{
			format: format,
			levels: levelSliceToSet(levels),
		},
		writer: writer,
	}, nil
}

// Log takes a Record and logs it to syslog.
func (syslogBackend SyslogBackend) Log(entry Entry) {
	line := syslogBackend.FormattedLogLine(entry)
	syslogBackend.writer.Write([]byte(line))
}
