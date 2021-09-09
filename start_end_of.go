package goment

// StartOf mutates the original Goment by setting it to the start of a unit of time.
func (g *Goment) StartOf(units string) *Goment {
	switch units {
	case "y", "year", "years":
		g.startOfYear()
	case "Q", "quarter", "quarters":
		g.startOfQuarter()
	case "M", "month", "months":
		g.startOfMonth()
	case "w", "week", "weeks":
		g.startOfWeek()
	case "W", "isoWeek", "isoWeeks":
		g.startOfISOWeek()
	case "d", "day", "days", "date":
		g.startOfDay()
	case "h", "hour", "hours":
		g.startOfHour()
	case "m", "minute", "minutes":
		g.startOfMinute()
	case "s", "second", "seconds":
		g.startOfSecond()
	}
	return g
}

func (g *Goment) startOfYear() *Goment {
	return g.SetMonth(1).startOfMonth()
}

func (g *Goment) startOfQuarter() *Goment {
	firstMonthOfQuarter := (g.Quarter() * 3) - 2
	return g.SetMonth(firstMonthOfQuarter).startOfMonth()
}

func (g *Goment) startOfMonth() *Goment {
	return g.SetDate(1).startOfDay()
}

func (g *Goment) startOfWeek() *Goment {
	return g.SetWeekday(0).startOfDay()
}

func (g *Goment) startOfISOWeek() *Goment {
	return g.SetDate(g.Date() - (g.ISOWeekday() - 1)).StartOf("day")
}

func (g *Goment) startOfDay() *Goment {
	return g.SetHour(0).startOfHour()
}

func (g *Goment) startOfHour() *Goment {
	return g.SetMinute(0).startOfMinute()
}

func (g *Goment) startOfMinute() *Goment {
	return g.SetSecond(0).startOfSecond()
}

func (g *Goment) startOfSecond() *Goment {
	return g.SetNanosecond(0)
}

// EndOf mutates the original Goment by setting it to the end of a unit of time.
func (g *Goment) EndOf(units string) *Goment {
	switch units {
	case "y", "year", "years":
		g.endOfYear()
	case "Q", "quarter", "quarters":
		g.endOfQuarter()
	case "M", "month", "months":
		g.endOfMonth()
	case "w", "week", "weeks":
		g.endOfWeek()
	case "W", "isoWeek", "isoWeeks":
		g.endOfISOWeek()
	case "d", "day", "days", "date":
		g.endOfDay()
	case "h", "hour", "hours":
		g.endOfHour()
	case "m", "minute", "minutes":
		g.endOfMinute()
	case "s", "second", "seconds":
		g.endOfSecond()
	}
	return g
}

func (g *Goment) endOfYear() *Goment {
	return g.SetMonth(12).endOfMonth()
}

func (g *Goment) endOfQuarter() *Goment {
	lastMonthOfQuarter := g.Quarter() * 3
	return g.SetMonth(lastMonthOfQuarter).endOfMonth()
}

func (g *Goment) endOfMonth() *Goment {
	return g.SetDate(g.DaysInMonth()).endOfDay()
}

func (g *Goment) endOfWeek() *Goment {
	return g.startOfWeek().addDays(6).endOfDay()
}

func (g *Goment) endOfISOWeek() *Goment {
	return g.startOfISOWeek().addDays(6).endOfDay()
}

func (g *Goment) endOfDay() *Goment {
	return g.SetHour(23).endOfHour()
}

func (g *Goment) endOfHour() *Goment {
	return g.SetMinute(59).endOfMinute()
}

func (g *Goment) endOfMinute() *Goment {
	return g.SetSecond(59).endOfSecond()
}

func (g *Goment) endOfSecond() *Goment {
	return g.SetNanosecond(999999999)
}
