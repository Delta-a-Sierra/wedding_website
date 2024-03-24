package entities

import (
	"time"
)

type Countdown struct {
	Days    int
	Hours   int
	Minutes int
}

func CreateCountDownFromDateTime(date time.Time) Countdown {
	remaining := time.Until(date)

	days := remaining.Seconds() / 86400
	remaining -= time.Duration(days) * 24 * time.Hour

	hours := remaining.Seconds() / 3600
	remaining -= time.Duration(hours) * time.Hour

	minutes := remaining.Seconds() / 60

	return Countdown{
		Days:    int(days),
		Hours:   int(hours),
		Minutes: int(minutes),
	}
}

func (c *Countdown) ZeroOutMinusValues() {
	if c.Days < 0 {
		c.Days = 0
	}
	if c.Hours < 0 {
		c.Hours = 0
	}
	if c.Minutes < 0 {
		c.Minutes = 0
	}
}
