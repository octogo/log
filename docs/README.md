# Architecture

Octolog is made up of the following components.

## Logger

- serves as interface for using octolog in other packages
- has wrapper-methods for logging in several log-levels
- can be configured to log only pre-defined log-levels
- initializes entries and passes them to all its configured outputs
- initializing two loggers with the same name will return the same logger

## Output

- writes log entries
- can be configured to log only pre-defined log-levels
- string-formats the entry before logging it
- initializing two outputs with the same URL will return the same outout

## Entry

- data-container for the log messag and its meta-data
- created by the Logger and passed to the Outputs
- string-templated against the log-format string by the Outputs

## Level

- form of classification of the severity or criticality of an entry
- simple integers; where `0 == ERROR`, the most severe or criticial log-level
- levels have colors assigned to them
- five built-in log-levels are always configured:
  - 0: ERROR (red)
  - 1: WARNING (yellow)
  - 2: NOTICE (green)
  - 3: INFO (white)
  - 4: DEBUG (cyan)
- custom log-levels can be registeredand and will carry the values 5 and up
- the colors of pre-defined and custom levels can be changed

## Color

- helper for injecting ANSII escape sequences into strings
- colors are configurable, custom ANSII sequences can be registered for use

## Config

- interface for configuration of package and its run-time behaviours
- basically a nested struct with lots of publicly exposed string values
- can be populated in code during run-time
- can be populated by a configuration file during initialization phase
- configuration file overrides all values in initialization phase
- run-time changes after initialization phase override the configuration file
