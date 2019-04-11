package log

// Backend defines the interface for a logging backend.
type Backend interface {
	SetFormat(string)      // sets the format template string
	Format() string        // returns the format template string
	Levels() []Level       // returns the log levels this backend will log
	SetLevels(...Level)    // sets the log levels
	AddLevels(...Level)    // adds log levels
	RemoveLevels(...Level) // removes log levels for this backend
	Wants(Entry) bool      // Check if handler is interested in Entry
	Log(Entry)             // Logs the entry
}

// BaseBackend implements abstract Backend.
type BaseBackend struct {
	format string
	levels LevelSet
}

// Format returns the format string of this backend.
func (baseBackend BaseBackend) Format() string {
	return baseBackend.format
}

// SetFormat sets the format string of this backend to the given format.
func (baseBackend *BaseBackend) SetFormat(format string) {
	baseBackend.format = format
}

// Levels returns the log-levels configured for this backend.
func (baseBackend BaseBackend) Levels() []Level {
	keys := make([]Level, 0, len(baseBackend.levels))
	for k := range baseBackend.levels {
		keys = append(keys, k)
	}
	return keys
}

// SetLevels sets the log-levels for this backend.
func (baseBackend *BaseBackend) SetLevels(levels ...Level) {
	newLevels := make(map[Level]struct{})
	for _, level := range levels {
		newLevels[level] = struct{}{}
	}
	baseBackend.levels = newLevels
}

// AddLevels adds the given log-levels to this backend.
func (baseBackend *BaseBackend) AddLevels(levels ...Level) {
	for _, level := range levels {
		baseBackend.levels[level] = struct{}{}
	}
}

// RemoveLevels will remove the given log levels from this backend.
func (baseBackend *BaseBackend) RemoveLevels(levels ...Level) {
	for _, level := range levels {
		if _, ok := baseBackend.levels[level]; ok == true {
			delete(baseBackend.levels, level)
		}
	}
}

// Wants returns true if the given Entry is logged by this backend.
func (baseBackend BaseBackend) Wants(entry Entry) bool {
	return baseBackend.wants(entry)
}

func (baseBackend BaseBackend) wants(entry Entry) bool {
	_, ok := baseBackend.levels[entry.Level()]
	return ok == true
}

// FormattedLogLine returns the formatted log line.
func (baseBackend BaseBackend) FormattedLogLine(entry Entry) string {
	return FormattedEntry(entry, baseBackend.format)
}
