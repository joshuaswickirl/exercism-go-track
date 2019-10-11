package clock

import (
	"fmt"
	"math"
)

// Clock has hours and minutes
type Clock struct {
	hours   int
	minutes int
}

// New creates a new clock
func New(hours, minutes int) Clock {
	c := Clock{hours, minutes}
	c.Adjust()
	return c
}

// Add seconds to the clock
func (c Clock) Add(minutes int) Clock {
	c.minutes += minutes
	c.Adjust()
	return c
}

// Subtract seconds from the clock
func (c Clock) Subtract(minutes int) Clock {
	c.minutes -= minutes
	c.Adjust()
	return c
}

// String returns the formatted time
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hours, c.minutes)
}

// Adjust corrects the time to a 24 hour clock
func (c *Clock) Adjust() {
	for {
		var changes bool
		if c.minutes >= 60 {
			c.hours += c.minutes / 60
			c.minutes = c.minutes % 60
			changes = true
		}
		if c.minutes < 0 {
			c.hours -= int(math.Abs(float64(c.minutes/60))) + 1
			c.minutes = 60 - int(math.Abs(float64(c.minutes)))%60
			changes = true
		}
		if c.hours >= 24 {
			c.hours = c.hours % 24
			changes = true
		}
		if c.hours < 0 {
			c.hours = 24 - int(math.Abs(float64(c.hours%24)))
			changes = true
		}
		if !changes {
			break
		}
	}
}
