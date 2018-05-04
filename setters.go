package goment

// Set is a generic setter, accepting unit as the first argument, and value as the second.
func (g *Goment) Set(unit string, value int) *Goment {
	switch unit {
	case "y", "year", "years":
		return g.SetYear(value)
	case "M", "month", "months":
		return g.SetMonth(value)
	case "D", "date", "dates":
		return g.SetDate(value)
	case "h", "hour", "hours":
		return g.SetHour(value)
	case "m", "minute", "minutes":
		return g.SetMinute(value)
	case "s", "second", "seconds":
		return g.SetSecond(value)
	case "ms", "millisecond", "milliseconds":
		return g.SetMillisecond(value)
	case "ns", "nanosecond", "nanoseconds":
		return g.SetNanosecond(value)
	}
	return g
}

// SetNanosecond sets the nanoseconds.
func (g *Goment) SetNanosecond(nanoseconds int) *Goment {
	if nanoseconds >= 0 && nanoseconds <= 999999999 {
		return g.addNanoseconds(nanoseconds - g.Nanosecond())
	}
	return g
}

// SetMillisecond sets the milliseconds.
func (g *Goment) SetMillisecond(milliseconds int) *Goment {
	if milliseconds >= 0 && milliseconds <= 59000 {
		return g.addMilliseconds(milliseconds - g.Millisecond())
	}
	return g
}

// SetSecond sets the seconds.
func (g *Goment) SetSecond(seconds int) *Goment {
	if seconds >= 0 && seconds <= 59 {
		return g.addSeconds(seconds - g.Second())
	}
	return g
}

// SetMinute sets the minutes.
func (g *Goment) SetMinute(minutes int) *Goment {
	if minutes >= 0 && minutes <= 59 {
		return g.addMinutes(minutes - g.Minute())
	}
	return g
}

// SetHour sets the hour.
func (g *Goment) SetHour(hours int) *Goment {
	if hours >= 0 && hours <= 23 {
		return g.addHours(hours - g.Hour())
	}
	return g
}

// SetDate sets the day of the month. If the date passed in is greater than the number of days in the month,
// then the day is set to the last day of the month.
func (g *Goment) SetDate(date int) *Goment {
	if date >= 1 && date <= 31 {
		daysInMonth := g.DaysInMonth()
		if date >= daysInMonth {
			date = daysInMonth
		}
		return g.addDays(date - g.Date())
	}
	return g
}

// SetDay sets the day of the week (Sunday = 0...).
func (g *Goment) SetDay(day int) *Goment {
	if day >= 0 && day <= 6 {
		return g.addDays(day - g.Day())
	}
	return g
}

// SetWeekday sets the day of the week according to the locale.
func (g *Goment) SetWeekday(weekday int) *Goment {
	return g
}

// SetISOWeekday sets the ISO day of the week with 1 being Monday and 7 being Sunday.
func (g *Goment) SetISOWeekday(weekday int) *Goment {
	if weekday >= 1 && weekday <= 7 {
		if weekday == 7 {
			weekday = 0
		}
		return g.SetDay(weekday)
	}
	return g
}

// SetDayOfYear sets the day of the year. For non-leap years, 366 is treated as 365.
func (g *Goment) SetDayOfYear(doy int) *Goment {
	if doy >= 1 && doy <= 366 {
		if !g.IsLeapYear() && doy == 366 {
			doy = 365
		}
		return g.addDays(doy - g.DayOfYear())
	}
	return g
}

// SetWeek sets the week of the year according to the locale.
func (g *Goment) SetWeek(week int) *Goment {
	return g
}

// SetISOWeek sets the ISO week of the year.
func (g *Goment) SetISOWeek(week int) *Goment {
	return g
}

// SetMonth sets the month (January = 1...). If new month has less days than current month,
// the date is pinned to the end of the target month.
func (g *Goment) SetMonth(month int) *Goment {
	if month >= 1 && month <= 12 {
		currentDate := g.Date()
		newDaysInMonth := daysInMonth(month, g.Year())
		if currentDate > newDaysInMonth {
			g.SetDate(newDaysInMonth)
		}
		return g.addMonths(month - g.Month())
	}
	return g
}

// SetQuarter sets the quarter (1 to 4).
func (g *Goment) SetQuarter(quarter int) *Goment {
	if quarter >= 1 && quarter <= 4 {
		return g.addQuarters(quarter - g.Quarter())
	}
	return g
}

// SetYear sets the year.
func (g *Goment) SetYear(year int) *Goment {
	return g.addYears(year - g.Year())
}

// SetWeekYear sets the week-year according to the locale.
func (g *Goment) SetWeekYear(weekYear int) *Goment {
	return g
}

// SetISOWeekYear sets the ISO week-year.
func (g *Goment) SetISOWeekYear(weekYear int) *Goment {
	return g
}
