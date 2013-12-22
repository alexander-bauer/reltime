package reltime

import (
	"github.com/jinzhu/now"
	"time"
)

// Between checks whether the first time is between the next two
// times, t2 being inclusive and t3 being exclusive.
func Between(t1, t2, t3 time.Time) bool {
	// We want to check whether t1 comes after t2, inclusive, and
	// before t1, exclusive. Because both the time.Before and
	// time.After functions are exclusive, we use !Before to check if
	// t1 comes after t2, inclusive.
	return !t1.Before(t2) && t1.Before(t3)
}

// Today returns whether the given time is on the same day as the RelTime.
func (rt RelTime) Today(t time.Time) bool {
	// Compare the dates and see if they're the same.
	ryear, rmonth, rday := time.Time(rt).Date()
	tyear, tmonth, tday := t.Date()
	return ryear == tyear && rmonth == tmonth && rday == tday
}

// Tomorrow returns whether the given time is on the day after the
// RelTime.
func (rt RelTime) Tomorrow(t time.Time) bool {
	// Use the end of day function to determine if the given time is
	// during the next day.
	end := now.New(time.Time(rt)).EndOfDay()
	return Between(t, end, end.AddDate(0, 0, 1))
}

// Yesterday returns whether the given time is on the day before the
// RelTime.
func (rt RelTime) Yesterday(t time.Time) bool {
	// Use the beginning of day function to determine if the given
	// time is during the previous day.
	beginning := now.New(time.Time(rt)).BeginningOfDay()
	return Between(t, beginning.AddDate(0, 0, -1), beginning)
}

// ThisWeek returns whether the given time is within the 7-day period
// beginning at the start of today and ending 24 * 7 hours later.
func (rt RelTime) ThisWeek(t time.Time) bool {
	beginning := now.New(time.Time(rt)).BeginningOfDay()
	return Between(t, beginning, beginning.AddDate(0, 0, 7))
}

// NextWeek returns whether the given time is within the 7-day period
// beginning at the start of today + 24 * 7 hours, and ending 24 * 7
// hours later.
func (rt RelTime) NextWeek(t time.Time) bool {
	beginning := now.New(time.Time(rt)).BeginningOfDay().AddDate(0, 0, 7)
	return Between(t, beginning, beginning.AddDate(0, 0, 7))
}

// LastWeek returns whether the given time is within the 7-day period
// ending at the start of today and beginning 24 * 7 hours earlier.
func (rt RelTime) LastWeek(t time.Time) bool {
	end := now.New(time.Time(rt)).BeginningOfDay()
	return Between(t, end.AddDate(0, 0, -7), end)
}
