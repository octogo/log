package log

var redactRune = rune('*')

// Redactor defines the interface for a redactor.
// Redactors will be redacted (replaced with ***) before being logged.
type Redactor interface {
	Redacted() string
}

// Redacted returns a string of redactRunes of length n.
func Redacted(n uint) (out string) {
	if n == 0 {
		n = 3
	}
	for i := 0; i < int(n); i++ {
		out = out + string(redactRune)
	}
	return out
}
