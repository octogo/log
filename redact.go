package log

import "github.com/octogo/log/pkg/log"

// Redactor is defined as any type with a Redact() string function.
// All values logged will be checked if they satisfy the Redactor
// interface and if they do, the output of their Redact() function
// will be logged instead of their string representation.
type Redactor interface {
	log.Redactor
}
