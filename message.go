package log

import (
	"time"
)

// Message is a container for a log-message.
type Message struct {
	Timestamp time.Time
	Logger    *Logger
	Level     Level
	Msg       string
	Caller    string
	File      string
	Line      int
}
