package reltime

import (
	"testing"
	"time"
)

const format = "2006-01-02 15:04:05.999999999"

func TestFormatRelative(t *testing.T) {
	// rt will be the reference time.
	rt := RelTime(time.Date(2013, 12, 21, 11, 12, 13, 123456789, time.UTC))

	if rt.FormatRelative("%s at 15:04", format,
		time.Date(2013, 12, 21, 18, 0, 0, 0, time.UTC)) != "Today at 18:00" {
		t.Errorf("RelativeDay")
	}
	if rt.FormatRelative("%s", format, time.Date(2013, 12, 31, 0, 0, 0, 0,
		time.UTC)) != "Next Tuesday" {
		t.Errorf("RelativeWeekday")
	}
}

func TestRelativeDay(t *testing.T) {
	// rt will be the reference time.
	rt := RelTime(time.Date(2013, 12, 21, 11, 12, 13, 123456789, time.UTC))

	if rt.RelativeDay(time.Date(2013, 12, 21, 0, 0, 0, 0, time.UTC)) !=
		"Today" {
		t.Errorf("Today")
	}
	if rt.RelativeDay(time.Date(2013, 12, 22, 0, 0, 0, 0, time.UTC)) !=
		"Tomorrow" {
		t.Errorf("Tomorrow")
	}
	if rt.RelativeDay(time.Date(2013, 12, 20, 0, 0, 0, 0, time.UTC)) !=
		"Yesterday" {
		t.Logf(rt.RelativeDay(time.Date(2013, 12, 20, 0, 0, 0, 0, time.UTC)))
		t.Errorf("Yesterday")
	}
}

func TestRelativeWeekday(t *testing.T) {
	// rt will be the reference time.
	rt := RelTime(time.Date(2013, 12, 21, 11, 12, 13, 123456789, time.UTC))

	if rt.RelativeWeekday(time.Date(2013, 12, 23, 0, 0, 0, 0, time.UTC)) !=
		"Monday" {
		t.Errorf("This week")
	}
	if rt.RelativeWeekday(time.Date(2013, 12, 31, 0, 0, 0, 0, time.UTC)) !=
		"Next Tuesday" {
		t.Errorf("Next Week")
	}
	if rt.RelativeWeekday(time.Date(2013, 12, 20, 0, 0, 0, 0, time.UTC)) !=
		"Last Friday" {
		t.Errorf("Last Week")
	}
}
