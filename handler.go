package log

import (
	"os"
	"sync/atomic"
)

var (
	nextGID    uint64
	backends   []Backend
	mainLogger *Logger

	chClose    = make(chan struct{})
	chBackends = make(chan []Backend)
	chStatus   = make(chan string)
	chLog      = make(chan Entry)
)

func init() {
	mainLogger = NewLogger("main")
	backends = DefaultBackends()
	go func() {
		for {
			select {
			case <-chClose:
				close(chClose)
				close(chBackends)
				close(chStatus)
				close(chLog)
				return

			case chStatus <- "TODO: STATUS STRING":
			case chBackends <- backends:
			case b := <-chBackends:
				backends = b

			case e := <-chLog:
				log(e)
			}
		}
	}()
}

// NextGID returns the next GID.
func NextGID() uint64 {
	return atomic.AddUint64(&nextGID, 1)
}

// Close closes all internal channels and drains all remaining log entries.
func Close() {
	chClose <- struct{}{}
	for e := range chLog {
		log(e)
	}
}

// Log is the primary logging function.
func Log(e Entry) {
	chLog <- e
}

func log(e Entry) {
	for _, backend := range backends {
		if backend.Wants(e) {
			backend.Log(e)
		}
	}
}

// SetBackends sets the given backends into context, replacing all previously
// configured backends.
func SetBackends(b ...Backend) {
	chBackends <- b
}

// AddBackends adds the given backends into context.
// All loggers send all log messages to all Backends, the Backends decide for
// themselves if they actually want a message.
func AddBackends(b ...Backend) {
	chBackends <- append(<-chBackends, b...)
}

// Printf logs the given interfaces as formatted INFO.
func Printf(format string, v ...interface{}) {
	mainLogger.Infof(format, v...)
}

// Println logs the given interfaces as INFO.
func Println(v ...interface{}) {
	mainLogger.Info(v...)
}

// Error logs the given interface as ERROR.
func Error(v interface{}) {
	mainLogger.Error(v)
}

// Errorf logs the given interface as formatted ERROR.
func Errorf(format string, v interface{}) {
	mainLogger.Errorf(format, v)
}

// Fatal logs the given interfaces as ERROR and exists and exists non-zero.
func Fatal(v ...interface{}) {
	mainLogger.Error(v...)
	Close()
	os.Exit(1)
}

// Fatalf logs the given interaces as formatted ERROR and exists non-zero.
func Fatalf(format string, v ...interface{}) {
	mainLogger.Errorf(format, v...)
	Close()
	os.Exit(1)
}
