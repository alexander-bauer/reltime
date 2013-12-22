package reltime

import (
	"time"
)

func FormatRelative(format, fallback string, t time.Time) string {
	return RelTime(time.Now()).FormatRelative(format, fallback, t)
}

func RelativeDay(t time.Time) string {
	return RelTime(time.Now()).RelativeDay(t)
}

func RelativeWeekday(t time.Time) string {
	return RelTime(time.Now()).RelativeWeekday(t)
}
