package octolog

import (
	"bytes"
	"fmt"
	"html/template"
	"sync"
	"sync/atomic"
)

// Router serves as consumer interface.
type Router struct {
	nextGID  uint64
	backends []Backend

	chStatus      chan chan string
	chLog         chan Entry
	chAddBackends chan []Backend
	chSetBackends chan []Backend
	chClose       chan int
	waitGroup     sync.WaitGroup
}

// New returns a *Router.
func New() *Router {
	router := &Router{
		backends: []Backend{},

		chStatus:      make(chan chan string),
		chLog:         make(chan Entry),
		chAddBackends: make(chan []Backend),
		chSetBackends: make(chan []Backend),
		chClose:       make(chan int),
	}

	go router.run()
	router.SetBackends(DefaultBackends()...)
	return router
}

func (router *Router) run() {
	router.waitGroup.Add(1)
	defer router.waitGroup.Done()

	for {
		select {
		case <-router.chClose:
			close(router.chLog)
			close(router.chAddBackends)
			close(router.chSetBackends)
			close(router.chStatus)
			close(router.chClose)
			return

		case out := <-router.chStatus:
			out <- router.status()

		case entry := <-router.chLog:
			router.Log(entry)

		case backends := <-router.chAddBackends:
			router.backends = append(router.backends, backends...)

		case backends := <-router.chSetBackends:
			router.backends = backends
		}
	}
}

// Close closes the router and all its channels, gracefully.
func (router *Router) Close() {
	router.chClose <- 1
	router.waitGroup.Wait()

	for entry := range router.chLog {
		router.Log(entry)
	}
}

// Drain is an alias for Close.
func (router *Router) Drain() {
	router.Close()
}

// Status returns the status string of the router.
func (router *Router) Status() string {
	responseChan := make(chan string)
	router.chStatus <- responseChan
	return <-responseChan
}

func (router *Router) status() string {
	vars := map[string]string{
		"GID":      fmt.Sprintf("%d", router.nextGID),
		"Backends": fmt.Sprintf("%d", len(router.backends)),
	}

	format := "octolog status: GID={{.GID}} Backends={{.Backends}}"
	tmpl, err := template.New("octolog/status").Parse(format)
	if err != nil {
		panic(err)
	}
	buf := new(bytes.Buffer)
	err = tmpl.Execute(buf, vars)
	if err != nil {
		panic(err)
	}
	return buf.String()
}

// NextGID returns the next GID.
func (router *Router) NextGID() uint64 {
	return atomic.AddUint64(&router.nextGID, 1)
}

// Log routes the given entry to all its backends.
func (router *Router) Log(entry Entry) {
	for _, backend := range router.backends {
		if backend.Wants(entry) {
			backend.Log(entry)
		}
	}
}

// AddBackends adds the given backends to this router.
func (router *Router) AddBackends(backends ...Backend) {
	router.chAddBackends <- backends
}

// SetBackends sets the given backends for this router, removing all previously
// add backends.
func (router *Router) SetBackends(backends ...Backend) {
	router.chSetBackends <- backends
}

// NewLogger returns a new logging interface for this router.
func (router *Router) NewLogger(name string) Logger {
	return &l{
		name:   name,
		router: router,
	}
}
