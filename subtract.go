package goment

import "time"

// Subtract mutates the original Goment by subtracting time.
func (g *Goment) Subtract(args ...interface{}) *Goment {
	if len(args) > 0 {
		switch val := args[0].(type) {
		case time.Duration:
			g.subtractDuration(val)
		case int:
			if len(args) == 2 {
				unit, ok := args[1].(string)
				if ok {
					switch unit {
					case "y", "year", "years":
						g.subtractYears(val)
					case "Q", "quarter", "quarters":
						g.subtractQuarters(val)
					case "M", "month", "months":
						g.subtractMonths(val)
					case "w", "week", "weeks":
						g.subtractWeeks(val)
					case "d", "day", "days":
						g.subtractDays(val)
					case "h", "hour", "hours":
						g.subtractHours(val)
					case "m", "minute", "minutes":
						g.subtractMinutes(val)
					case "s", "second", "seconds":
						g.substractSeconds(val)
					case "ms", "millisecond", "milliseconds":
						g.subtractMilliseconds(val)
					case "ns", "nanosecond", "nanoseconds":
						g.subtractNanoseconds(val)
					}
				}
			}
		}
	}
	return g
}

func (g *Goment) subtractYears(years int) *Goment {
	return g.addYears(years * -1)
}

func (g *Goment) subtractQuarters(quarters int) *Goment {
	return g.addQuarters(quarters * -1)
}

func (g *Goment) subtractMonths(months int) *Goment {
	return g.addMonths(months * -1)
}

func (g *Goment) subtractWeeks(weeks int) *Goment {
	return g.addWeeks(weeks * -1)
}

func (g *Goment) subtractDays(days int) *Goment {
	return g.addDays(days * -1)
}

func (g *Goment) subtractHours(hours int) *Goment {
	return g.addHours(hours * -1)
}

func (g *Goment) subtractMinutes(minutes int) *Goment {
	return g.addMinutes(minutes * -1)
}

func (g *Goment) substractSeconds(seconds int) *Goment {
	return g.addSeconds(seconds * -1)
}

func (g *Goment) subtractMilliseconds(milliseconds int) *Goment {
	return g.addMilliseconds(milliseconds * -1)
}

func (g *Goment) subtractNanoseconds(nanoseconds int) *Goment {
	return g.addNanoseconds(nanoseconds * -1)
}

func (g *Goment) subtractDuration(d time.Duration) *Goment {
	return g.addDuration(d * -1)
}
