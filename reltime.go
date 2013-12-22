// Package reltime implements human-readable stringification for
// relative times.
//
// Copyleft (C) 2013 Alexander Bauer
package reltime

import (
	"time"
)

type RelTime time.Time

// RelativeString stringifies the given time relative to the
// RelTime. For example, it might stringify to "Tomorrow" or
// "Yesterday". If it is not possible to stringify to a specific
// relative day, it will return a blank string.
func (rt RelTime) RelativeDay(t time.Time) string {
	switch {
	case rt.Today(t):
		return "Today"
	case rt.Tomorrow(t):
		return "Tomorrow"
	case rt.Yesterday(t):
		return "Yesterday"
	default:
		return ""
	}
}
