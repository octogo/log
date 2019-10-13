package log

import (
	"sync"
)

var (
	regOutputs = map[string]Output{}
	outMu      = &sync.Mutex{}
)

// RegisterOutput registers the given output under the given name.
func RegisterOutput(url string, output Output) Output {
	outMu.Lock()
	defer outMu.Unlock()
	existing, exists := regOutputs[url]
	if exists {
		return existing
	}
	regOutputs[url] = output
	return regOutputs[url]
}

// GetOutput returns the output for the given name if it has been registered or
// nil if no output with that name has been registered.
func GetOutput(url string) Output {
	outMu.Lock()
	defer outMu.Unlock()
	existing, exists := regOutputs[url]
	if exists {
		return existing
	}
	return nil
}
