package log

import (
	"time"
)

// Message is a container for a log-message.
type Message struct {
	timestamp time.Time
	logger    *Logger
	level     Level
	msg       string
	caller    string
	file      string
	line      int
}
