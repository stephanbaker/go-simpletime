// Package simpletime is an extension of the go standard library time
// package which provides a number of convenience functions for retrieving
// new times based on the provided reference time
package simpletime

import (
	"time"
)

// SimpleDuration struct embeds time.Duration
type SimpleDuration struct {
	time.Duration
}

// NewSimpleDuration returns a pointer to a SimpleDuration struct
func NewSimpleDuration(d time.Duration) *SimpleDuration {
	return &SimpleDuration{d}
}

// Compare test if the reference duration is equal to the supplied
// duration by testing the equality of nanonseconds
func (d *SimpleDuration) Compare(e *SimpleDuration) bool {
	if d == nil || e == nil {
		return false
	}

	if d.Duration.Nanoseconds() == e.Duration.Nanoseconds() {
		return true
	}

	return false
}

// Days returns the time elapsed since t in days
func (d *SimpleDuration) Days() float64 {
	return d.Hours() / 24
}
