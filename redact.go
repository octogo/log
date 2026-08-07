package log

// All entries logged will be checked if they satisfy the Redactor interface and
// if they do, the output of their Redact() function will be logged instead of
// their fmt.Stringer.
type Redacter interface {
	Redacted() string
}
