package log

import "os"

// ExitFunc defines the signature of an exit function, such as os.Exit().
// It is needed for testing.
type ExitFunc func(int)

// exit holds a reference to os.Exit.
// This is needed for tests to continue wihtout actually calling os.Exit.
var exit ExitFunc = os.Exit

func ExitHandler(f func(code int)) {
	exit = f
}
