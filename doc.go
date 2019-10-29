// Package log is a drop-in replacement for the builtin log package.
// It has support for colors, logging in and filtering by log-levels, as well
// as support for concurrent use across multiple goroutines.
//
// The standard Logger routes all logs to the STDOUT and STDERR outputs.
// By default, the STDOUT output will only log log-levels INFO and NOTICE,
// while the STDERR output will ony log log-levels WARNING and ERROR.
//
// All defaults can be overwritten in code or optionally a plain-text YAML file.
package log
