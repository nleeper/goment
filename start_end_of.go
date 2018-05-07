package goment

// StartOf mutates the original Goment by setting it to the start of a unit of time.
func (g *Goment) StartOf(unit string) *Goment {
	switch unit {
	case "year":
		g.startOfYear()
	case "quarter":
		g.startOfQuarter()
	case "month":
		g.startOfMonth()
	case "week":
		g.startOfWeek()
	case "isoWeek":
		g.startOfISOWeek()
	case "day", "date":
		g.startOfDay()
	case "hour":
		g.startOfHour()
	case "minute":
		g.startOfMinute()
	case "second":
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
	return g
}

func (g *Goment) startOfISOWeek() *Goment {
	return g.SetISOWeekday(1).startOfDay()
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
func (g *Goment) EndOf(unit string) *Goment {
	switch unit {
	case "year":
		g.endOfYear()
	case "quarter":
		g.endOfQuarter()
	case "month":
		g.endOfMonth()
	case "week":
		g.endOfWeek()
	case "isoWeek":
		g.endOfISOWeek()
	case "day", "date":
		g.endOfDay()
	case "hour":
		g.endOfHour()
	case "minute":
		g.endOfMinute()
	case "second":
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
	return g
}

func (g *Goment) endOfISOWeek() *Goment {
	return g.SetISOWeekday(6).startOfDay()
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
