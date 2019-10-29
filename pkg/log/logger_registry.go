package log

import "sync"

var (
	regLoggers = map[string]*Logger{}
	logMu      = &sync.Mutex{}
)

// RegisterLogger registers the given Logger under the given name.
func RegisterLogger(name string, logger *Logger) *Logger {
	if name == "" {
		name = LoggerName
	}
	logMu.Lock()
	defer logMu.Unlock()
	existing, exists := regLoggers[name]
	if exists {
		return existing
	}
	regLoggers[name] = logger
	return regLoggers[name]
}
