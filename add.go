package goment

import (
	"time"
)

// Add mutates the original Goment by adding time.
func (g *Goment) Add(args ...interface{}) *Goment {
	if len(args) > 0 {
		switch val := args[0].(type) {
		case time.Duration:
			g.addDuration(val)
		case int:
			if len(args) == 2 {
				unit, ok := args[1].(string)
				if ok {
					switch unit {
					case "y", "year", "years":
						g.addYears(val)
					case "Q", "quarter", "quarters":
						g.addQuarters(val)
					case "M", "month", "months":
						g.addMonths(val)
					case "w", "week", "weeks":
						g.addWeeks(val)
					case "d", "day", "days":
						g.addDays(val)
					case "h", "hour", "hours":
						g.addHours(val)
					case "m", "minute", "minutes":
						g.addMinutes(val)
					case "s", "second", "seconds":
						g.addSeconds(val)
					case "ms", "millisecond", "milliseconds":
						g.addMilliseconds(val)
					case "ns", "nanosecond", "nanoseconds":
						g.addNanoseconds(val)
					}
				}
			}
		}
	}
	return g
}

func (g *Goment) addYears(years int) *Goment {
	g.time = g.ToTime().AddDate(years, 0, 0)
	return g
}

func (g *Goment) addQuarters(quarters int) *Goment {
	return g.addMonths(quarters * 3)
}

func (g *Goment) addMonths(months int) *Goment {
	g.time = g.ToTime().AddDate(0, months, 0)
	return g
}

func (g *Goment) addWeeks(weeks int) *Goment {
	return g.addDays(weeks * 7)
}

func (g *Goment) addDays(days int) *Goment {
	g.time = g.ToTime().AddDate(0, 0, days)
	return g
}

func (g *Goment) addHours(hours int) *Goment {
	return g.addDuration(time.Hour * time.Duration(hours))
}

func (g *Goment) addMinutes(minutes int) *Goment {
	return g.addDuration(time.Minute * time.Duration(minutes))
}

func (g *Goment) addSeconds(seconds int) *Goment {
	return g.addDuration(time.Second * time.Duration(seconds))
}

func (g *Goment) addMilliseconds(milliseconds int) *Goment {
	return g.addDuration(time.Millisecond * time.Duration(milliseconds))
}

func (g *Goment) addNanoseconds(nanoseconds int) *Goment {
	return g.addDuration(time.Nanosecond * time.Duration(nanoseconds))
}

func (g *Goment) addDuration(d time.Duration) *Goment {
	g.time = g.ToTime().Add(d)
	return g
}
