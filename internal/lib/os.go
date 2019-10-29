package lib

import "os"

// OpenFile returns a *os.File for the given path.
func OpenFile(path string) *os.File {
	file, err := os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}
	return file
}
