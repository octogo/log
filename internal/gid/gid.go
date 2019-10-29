package gid

import "github.com/octogo/log/pkg/uid"

var gid = &uid.UID{}

// Latest returns the latest ID.
func Latest() uint64 {
	return gid.Latest()
}

// Next returns the next GID and increases the counter for the next caller.
func Next() uint64 {
	return gid.Next()
}
