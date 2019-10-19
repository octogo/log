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
# {{.Date}}       - date formatted as 2006/01/02
# {{.Time}}       - time formatted as 15:04:05
# {{.Milli}}      - time formatted as .000
# {{.Nano}}       - time formatted as .000000
# {{.PID}}        - process' ID
# {{.PPID}}       - partent-process' ID
# {{.GID}}        - global message ID
# {{.LID}}        - logger message ID
# {{.Logger}}     - name of the logger
# {{.Level}}      - log-level of the entry
# {{.Func}}       - name of the calling function
# {{.File}}       - source file of the calling function
# {{.Line}}       - line in the above source file
#
# supported colorize labels:
# {{.Color}}      - activates coloring
# {{.BoldColor}}  - activates bold coloring
# {{.NoColor}}    - deactivates coloring
#
# default: '{{.Date}} {{.Time}} {{.Level}} {{.Message}}'
defaultformat: '{{.Date}} {{.Time}} {{.BoldColor}}{{.Logger}} {{.Level}}{{.NoColor}} {{.Color}}{{.Message}}{{.NoColor}}'

# defaultoutputs defines the default outputs for every logger
# that has no explicit outputs configured.
# default: [ 'file:///dev/stdout', 'file:///dev/stderr' ]
defaultoutputs:
  - 'file:///dev/stdout'
  - 'file:///dev/stderr'

# levels configures the available log-levels and their corresponding colors.
# You can specify any ANSII color literal, such as
#   - black
#   - red
#   - green
#   - yellow
#   - blue
#   - magenta
#   - cyan
#   - white
# Any other value will be handled like a literal ANSII escape sequence, such as
#   - 30;41   # (black text on red background)
#   - 5;31    # blinking red text
# Note: the default levels [ ERROR, WARNING, NOTICE, INFO, DEBUG ] will always be
# registered before all custom levels. Therefor, their colors can be overridden,
# but their order can not.
levels:
  - name: ERROR
    color: red
  - name: WARNING
    color: yellow
  - name: NOTICE
    color: green
  - name: INFO
    color: white
  - name: DEBUG
    color: cyan

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
