// Package reltime implements human-readable stringification for
// relative times.
//
// Copyleft (C) 2013 Alexander Bauer
package reltime

import (
	"fmt"
	"time"
)

type RelTime time.Time

// FormatRelative tries to format the given time relative to the
// RelTime as best it can, making use of available "Relative*"
// functions, and defaults to using time.Format with the given
// fallback string if it cannot. If it does succeed, it tries to
// substitute the produced stringification into the given format
// string.
//
// For example, if myTime is a time which occurs next Tuesday at 8:05:
//
//     FormatRelative("%s at 15:04", "2006-01-02 15:04", myTime)
//     // Next Tuesday at 08:05
//
func (rt RelTime) FormatRelative(format, fallback string, t time.Time) string {
	// Find the relative representation, if applicable.
	representation := rt.tryRelatives(t)

	// If there is no relative representation, format it with the
	// fallback.
	if representation == "" {
		return t.Format(fallback)
	}

	// Otherwise, format using the main format string, then use
	// Sprintf to substitute in the representation.
	return fmt.Sprintf(t.Format(format), representation)
}

func (rt RelTime) tryRelatives(t time.Time) string {
	var output string

	// Try to pin it down to a "Today", "Tomorrow", or "Yesterday".
	if output = rt.RelativeDay(t); output != "" {
		return output
	} else if output = rt.RelativeWeekday(t); output != "" {
		return output
	} else {
		return ""
	}
}

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

// RelativeWeekday stringifies the given time to a weekday name
// relative to the RelTime. For example, it might stringify to
// "Monday" or "Last Wednesday". If it is not possible to stringify to
// a specific relative weekday, it will return a blank string.
func (rt RelTime) RelativeWeekday(t time.Time) string {
	switch {
	case rt.ThisWeek(t):
		return t.Weekday().String()
	case rt.NextWeek(t):
		return "Next " + t.Weekday().String()
	case rt.LastWeek(t):
		return "Last " + t.Weekday().String()
	default:
		return ""
	}
}
