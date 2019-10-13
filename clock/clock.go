package clock

import (
	"fmt"
)

// Clock is the time in minutes
type Clock int

var minutesInDay = 1440

// New creates a new clock
func New(hours, minutes int) Clock {
	c := (hours*60 + minutes) % minutesInDay
	if c < 0 {
		c += minutesInDay
	}
	return Clock(c)
}

// Add seconds to the clock
func (c Clock) Add(minutes int) Clock {
	return New(0, int(c)+minutes)
}

// Subtract seconds from the clock
func (c Clock) Subtract(minutes int) Clock {
	return New(0, int(c)-minutes)
}

// String returns the formatted time
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c/60, c%60)
}
