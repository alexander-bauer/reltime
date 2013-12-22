package reltime

import (
	"testing"
	"time"
)

const format = "2006-01-02 15:04:05.999999999"

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
