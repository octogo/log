package uid

import "sync/atomic"

// UID implements a generator for unique IDs.
type UID struct {
	uid uint64
}

// Latest returns the latest unique ID.
func (uid *UID) Latest() uint64 {
	return atomic.AddUint64(&uid.uid, 0)
}

// Next returns the next unique ID and increases the internal counter for the
// next caller.
func (uid *UID) Next() uint64 {
	return atomic.AddUint64(&uid.uid, 1)
}
