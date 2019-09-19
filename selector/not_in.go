package selector

import (
	"fmt"
	"strings"
)

var (
	_ Selector = (*NotIn)(nil)
)

// NotIn returns if a key does not match a set of values.
type NotIn struct {
	Key    string
	Values []string
}

// Matches returns the selector result.
func (ni NotIn) Matches(labels Labels) bool {
	if value, hasValue := labels[ni.Key]; hasValue {
		for _, iv := range ni.Values {
			if iv == value {
				// the key does not equal any of the values
				return false
			}
		}
	}
	// the value doesn't exist.
	return true
}

// String returns a string representation of the selector.
func (ni NotIn) String() string {
	return fmt.Sprintf("%s notin (%s)", ni.Key, strings.Join(ni.Values, ", "))
}
