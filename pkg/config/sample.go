package config

import (
	"bytes"
	"os"
	"time"
)

// SampleConfig holds the sample configuration file.
var SampleConfig = `
# rootlogger defines the name of the standard logger
# default: main
rootlogger: 'main'

# default format defines the default log-format for all
# outputs that habe no expicit log-format in the
# configuration.
#
# supported formatting labels:
# {{.Date}}   - the date formatted as 2006/01/02
# {{.Time}}   - the time formatted as 15:04:05
# {{.Milli}}  - the time formatted as .000
# {{.Nano}}   - the time formatted as .000000
# {{.PID}}    - the process' ID
# {{.PPID}}   - the partent-process' ID
# {{.GID}}    - the global message ID
# {{.LID}}    - the logger message ID
# {{.Logger}} - the name of the logger
# {{.Level}}  - the log-level of the entry
# {{.Func}}   - the name of the calling function
# {{.File}}   - the source file of the calling function
# {{.Line}}   - the line in the above source file
#
# default: '{{.Date}} {{.Time}} {{.Level}} {{.Message}}'
defaultformat: '{{.Date}} {{.Time}} {{.BoldColor}}{{.Logger}} {{.Level}}{{.NoColor}} {{.Color}}{{.Message}}{{.NoColor}}'

# defaultoutputs defines the default outputs for every logger
# that has no explicit outputs configured.
# default: [ 'file:///dev/stdout', 'file:///dev/stderr' ]
defaultoutputs:
  - 'file:///dev/stdout'
  - 'file:///dev/stderr'

# outputs defines the outputs that should automatically be
# initialize upon startup.
#
# an output is defined as:
#   {
#     url:      # the URL, i.e. file://octo.log
#     wants:    # the list of log-levels to log
#               # providing no log-levels implies 'all'
#     format:   # log-format to use, if not the global default
#   }
outputs:
  # log INFO and NOTICE to STDOUT
  - url: 'file:///dev/stdout'
    wants: [ INFO, NOTICE ]
  # log WARNING and ERROR to STDERR
  - url: 'file:///dev/stderr'
    wants: [ WARNING, ERROR ]

# loggers defines the loggers that should automatically be
# initialized upon startup.
#
# a logger is defined as:
#   {
#     name:     # a unique name
#     wants:    # list of log-levels to log
#               # providing no log-levels implies 'all'
#     outputs:  # list of output URLs to communicate with
#               # providing no URLs implies 'defaultoutputs'
#   }
loggers:
  - name: main
`

// GetSampleConfig returns a sample configuration file.
func GetSampleConfig(v string) string {
	buf := new(bytes.Buffer)
	buf.WriteString("# octolog sample configuration file\n")
	buf.WriteString("# generated: " + time.Now().Format("2006/01/02 15:04:05") + "\n")
	buf.WriteString("# version: " + v + "\n")
	buf.WriteString(SampleConfig)
	return buf.String()
}

// WriteSampleToFile writes the sample configuration to the given *os.File.
// Existing files will never be overwritten.
func WriteSampleToFile(version string, file *os.File) {
	file.WriteString(GetSampleConfig(version))
}
