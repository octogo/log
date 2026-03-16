package log

import (
	"bytes"
	"io"
	"os"
	"syscall"

	"github.com/octogo/log/v2"
)

func captureOutput(f func()) (stdout string, stderr string) {
	// Create pipes for stdout and stderr
	rOut, wOut, _ := os.Pipe()
	rErr, wErr, _ := os.Pipe()

	// Duplicate original FDs
	oldOut, _ := syscall.Dup(1)
	oldErr, _ := syscall.Dup(2)

	// Redirect FDs to our pipes
	syscall.Dup2(int(wOut.Fd()), 1)
	syscall.Dup2(int(wErr.Fd()), 2)

	outCh := make(chan string)
	errCh := make(chan string)

	// Concurrently read stdout
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, rOut)
		outCh <- buf.String()
	}()

	// Concurrently read stderr
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, rErr)
		errCh <- buf.String()
	}()

	// Run the function
	f()

	// Close writers to signal EOF
	wOut.Close()
	wErr.Close()

	// Restore original FDs
	syscall.Dup2(oldOut, 1)
	syscall.Dup2(oldErr, 2)

	// Read captured output
	stdout = <-outCh
	stderr = <-errCh

	rOut.Close()
	rErr.Close()

	return stdout, stderr
}

func captureOsExit(f func()) (bool, int) {
	didExit := false
	exitCode := 0
	log.ExitHandler(func(code int) {
		didExit = true
		exitCode = code
	})
	f()
	log.ExitHandler(os.Exit)
	return didExit, exitCode
}
