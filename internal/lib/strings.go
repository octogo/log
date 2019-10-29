package lib

import (
	"errors"
	"strings"

	"github.com/octogo/log/pkg/level"
)

// StringInSlice returns true if the given string is contained in the given
// []string.
func StringInSlice(s string, slice []string) bool {
	for i := range slice {
		if s == slice[i] {
			return true
		}
	}
	return false
}

const schemaSep = "://"

// URL format the given schema and URI as URL (e.g. file:///tmp/octolog).
func URL(schema, uri string) string {
	return strings.Join([]string{schema, uri}, schemaSep)
}

// ParseURL returns the schema and URI from the given URL.
func ParseURL(url string) (schema, uri string, err error) {
	split := strings.Split(url, schemaSep)
	if len(split) < 2 {
		err = errors.New("malformed URL")
		return
	}
	schema = split[0]
	uri = split[1]
	return
}

// ParseLevels wraps level.Parse that parses more than one level.
func ParseLevels(levels ...string) []level.Level {
	if levels == nil || len(levels) == 0 {
		return nil
	}
	out := make([]level.Level, len(levels))
	for i := range levels {
		lvl, err := level.Parse(levels[i])
		if err != nil {
			panic(err)
		}
		out[i] = lvl
	}
	return out
}
