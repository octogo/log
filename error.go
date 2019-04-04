package octolog

import "os"

func internalError(err error) {
	os.Stderr.WriteString(err.Error() + "\n")
}
