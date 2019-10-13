# Architecture

Octolog is made up of two main components, the *Logger* and the *Output*.

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
