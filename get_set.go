package goment

import (
	"math"
)

// Get is a string getter using the supplied units. Returns 0 if unsupported property.
func (g *Goment) Get(units string) int {
	switch units {
	case "y", "year", "years":
		return g.Year()
	case "M", "month", "months":
		return g.Month()
	case "D", "date", "dates":
		return g.Date()
	case "h", "hour", "hours":
		return g.Hour()
	case "m", "minute", "minutes":
		return g.Minute()
	case "s", "second", "seconds":
		return g.Second()
	case "ms", "millisecond", "milliseconds":
		return g.Millisecond()
	case "ns", "nanosecond", "nanoseconds":
		return g.Nanosecond()
	}
	return 0
}

// Nanosecond gets the nanoseconds.
func (g *Goment) Nanosecond() int {
	return g.ToTime().Nanosecond()
}

// Millisecond gets the milliseconds.
func (g *Goment) Millisecond() int {
	return g.Second() * 1000
}

// Second gets the seconds.
func (g *Goment) Second() int {
	return g.ToTime().Second()
}

// Minute gets the minutes.
func (g *Goment) Minute() int {
	return g.ToTime().Minute()
}

// Hour gets the hour.
func (g *Goment) Hour() int {
	return g.ToTime().Hour()
}

// Date gets the day of the month.
func (g *Goment) Date() int {
	return g.ToTime().Day()
}

// Day gets the day of the week (Sunday = 0...).
func (g *Goment) Day() int {
	return int(g.ToTime().Weekday())
}

// Weekday gets the day of the week according to the locale.
func (g *Goment) Weekday() int {
	return (g.Day() + 7 - g.locale.Week.Dow) % 7
}

// ISOWeekday gets the ISO day of the week with 1 being Monday and 7 being Sunday.
func (g *Goment) ISOWeekday() int {
	wd := g.Day()
	if wd == 0 {
		wd = 7
	}
	return wd
}

// DayOfYear gets the day of the year.
func (g *Goment) DayOfYear() int {
	return g.ToTime().YearDay()
}

// Week gets the week of the year according to the locale.
func (g *Goment) Week() int {
	return weekOfYear(g, g.locale.Week.Dow, g.locale.Week.Doy).week
}

// ISOWeek gets the ISO week of the year.
func (g *Goment) ISOWeek() int {
	_, week := g.ToTime().ISOWeek()
	return week
}

// Month gets the month (January = 1...).
func (g *Goment) Month() int {
	return int(g.ToTime().Month())
}

// Quarter gets the quarter (1 to 4).
func (g *Goment) Quarter() int {
	return int(math.Ceil(float64(g.Month()) / 3))
}

// Year gets the year.
func (g *Goment) Year() int {
	return g.ToTime().Year()
}

// WeekYear gets the week-year according to the locale.
func (g *Goment) WeekYear() int {
	return weekOfYear(g, g.locale.Week.Dow, g.locale.Week.Doy).year
}

// ISOWeekYear gets the ISO week-year.
func (g *Goment) ISOWeekYear() int {
	year, _ := g.ToTime().ISOWeek()
	return year
}

// WeeksInYear gets the number of weeks according to locale in the current Goment's year.
func (g *Goment) WeeksInYear() int {
	return weeksInYear(g.Year(), g.locale.Week.Dow, g.locale.Week.Doy)
}

// ISOWeeksInYear gets the number of weeks in the current Goment's year, according to ISO weeks.
func (g *Goment) ISOWeeksInYear() int {
	return weeksInYear(g.Year(), 1, 4)
}

// Set is a generic setter, accepting units as the first argument, and value as the second.
func (g *Goment) Set(units string, value int) *Goment {
	switch units {
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
func (g *Goment) SetDay(args ...interface{}) *Goment {
	if len(args) != 1 {
		return g
	}

	switch v := args[0].(type) {
	case string:
		parsed := g.locale.WeekdaysRegex.FindString(v)
		if parsed != "" {
			val := g.locale.GetWeekdayNumber(parsed)
			if val != -1 {
				return g.addDays(val - g.Day())
			}
		}
	case int:
		return g.addDays(v - g.Day())
	}

	return g
}

// SetWeekday sets the day of the week according to the locale.
func (g *Goment) SetWeekday(weekday int) *Goment {
	currWeekday := g.Weekday()
	return g.addDays(weekday - currWeekday)
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
	return g.addDays((week - g.Week()) * 7)
}

// SetISOWeek sets the ISO week of the year.
func (g *Goment) SetISOWeek(week int) *Goment {
	woy := weekOfYear(g, 1, 4).week
	return g.addDays((week - woy) * 7)
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
	return setWeekYearHelper(g, weekYear, g.Week(), g.Weekday(), g.locale.Week.Dow, g.locale.Week.Doy)
}

// SetISOWeekYear sets the ISO week-year.
func (g *Goment) SetISOWeekYear(weekYear int) *Goment {
	return setWeekYearHelper(g, weekYear, g.ISOWeek(), g.ISOWeekday(), 1, 4)
}

func setWeekYearHelper(g *Goment, weekYear int, week int, weekday int, dow int, doy int) *Goment {
	weeksTarget := weeksInYear(weekYear, dow, doy)

	if week > weeksTarget {
		week = weeksTarget
	}

	dayOfYearData := dayOfYearFromWeeks(weekYear, week, weekday, dow, doy)

	d, _ := New(DateTime{Year: dayOfYearData.year, Month: 1, Day: dayOfYearData.dayOfYear})
	d.UTC()

	g.SetYear(d.Year())
	g.SetMonth(d.Month())
	g.SetDate(d.Date())

	return g
}
