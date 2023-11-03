package timeutil

import (
	"time"
)

type Timer struct {
	seconds time.Duration
	minutes time.Duration
	hours   time.Duration
}

// NewTimer creates a new timer with the specified duration in seconds.
func NewTimer(seconds, minutes, hours time.Duration) *Timer {
	return &Timer{
		seconds: seconds,
		minutes: minutes,
		hours:   hours,
	}
}

// Start starts the timer and waits for the specified duration.
func (t *Timer) Start() {
	<-time.After(t.seconds + t.minutes + t.hours)
}
