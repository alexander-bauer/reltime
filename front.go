package reltime

import (
	"time"
)

func FormatRelative(format string, t time.Time) string {
	return RelTime(time.Now()).FormatRelative(format, t)
}

func RelativeDay(t time.Time) string {
	return RelTime(time.Now()).RelativeDay(t)
}

func RelativeWeekday(t time.Time) string {
	return RelTime(time.Now()).RelativeWeekday(t)
}
