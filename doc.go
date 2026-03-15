// Package log is a drop-in replacement for the builtin log package.
// It features:
//
//   - concurrent logging across multiple goroutines
//   - logging in and filtering by log-levels
//   - custom formatting of log-messages
//   - individual ANSII colors per log-level
//
// The DefaultLogger comes with a slightly more verbose formatter than the
// builtin log package. It also automatically sends all errors and warnings to
// the STDERR sink instead of STDOUT.
//
// Everything can be customized via code.
package log
