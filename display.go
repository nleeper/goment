package goment

import (
	"time"
)

// Diff returns the difference between two Goments as an integer.
func (g *Goment) Diff(args ...interface{}) int {
	numArgs := len(args)
	if numArgs > 0 {
		units := ""

		input, err := New(args[0])
		if err != nil {
			return 0
		}

		if numArgs > 1 {
			if parsedUnits, ok := args[1].(string); ok {
				units = parsedUnits
			}
		}

		d := diff{
			Start: g,
			End:   input,
		}

		switch units {
		case "y", "year", "years":
			return d.InYears()
		case "M", "month", "months":
			return d.InMonths()
		case "w", "week", "weeks":
			return d.InWeeks()
		case "d", "day", "days":
			return d.InDays()
		case "h", "hour", "hours":
			return d.InHours()
		case "m", "minute", "minutes":
			return d.InMinutes()
		default:
			return d.InSeconds()
		}
	}
	return 0
}

// DaysInMonth returns the number of days in the set month.
func (g *Goment) DaysInMonth() int {
	return daysInMonth(g.Month(), g.Year())
}

// ToTime returns the time.Time object that is wrapped by Goment.
func (g *Goment) ToTime() time.Time {
	return g.time
}

// ToUnix returns the Unix timestamp (the number of seconds since the Unix Epoch).
func (g *Goment) ToUnix() int64 {
	return g.ToTime().Unix()
}

// ToArray returns an array that mirrors the parameters from time.Date().
func (g *Goment) ToArray() []int {
	return []int{g.Year(), g.Month(), g.Date(), g.Hour(), g.Minute(), g.Second(), g.Nanosecond()}
}

// ToDateTime returns a DateTime struct.
func (g *Goment) ToDateTime() DateTime {
	return DateTime{
		Year:       g.Year(),
		Month:      g.Month(),
		Day:        g.Date(),
		Hour:       g.Hour(),
		Minute:     g.Minute(),
		Second:     g.Second(),
		Nanosecond: g.Nanosecond(),
		Location:   g.ToTime().Location(),
	}
}

// ToString returns a string representation of the Goment time.
func (g *Goment) ToString() string {
	return g.ToTime().String()
}

// ToISOString returns a ISO8601 standard representation of the Goment time.
func (g *Goment) ToISOString() string {
	return g.ToTime().Format("2006-01-02T15:04:05.999Z07:00")
}

func daysInMonth(month, year int) int {
	return time.Date(year, time.Month(month+1), 0, 0, 0, 0, 0, time.UTC).Day()
}
