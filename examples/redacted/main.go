package main

import (
	"fmt"

	"github.com/octogo/log"
)

// SecretCredentials demonstrates the use of log.Redactor.
type SecretCredentials struct {
	Username, Password string
}

// String satisfies fmt.Stringer and will print out the password in
// cleartext.
func (sc SecretCredentials) String() string {
	return fmt.Sprintf("%s:%s", sc.Username, sc.Password)
}

// Redacted satisfies log.Redactor and therefor ensures that the
// Password attribute will never appear in the logs anywhere.
func (sc SecretCredentials) Redacted() string {
	return fmt.Sprintf("%s:%s", sc.Username, "********")
}

func main() {
	secret := SecretCredentials{
		Username: "fooser",
		Password: "t0ps3cr3t",
	}

	log.Init()
	logger := log.New("myapp", nil)
	logger.Noticef(
		"Note how the password gets redacted in the logs: %s",
		secret,
	)
}
