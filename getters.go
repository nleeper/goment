package goment

import (
	"math"
	"time"
)

// Get is a string getter using the unit. Returns 0 if unsupported property.
func (g *Goment) Get(unit string) int {
	switch unit {
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
	return 0
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
	return 0
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
	return 0
}

// ISOWeekYear gets the ISO week-year.
func (g *Goment) ISOWeekYear() int {
	year, _ := g.ToTime().ISOWeek()
	return year
}

// WeeksInYear gets the number of weeks according to locale in the current moment's year.
func (g *Goment) WeeksInYear() int {
	return 0
}

// ISOWeeksInYear gets the number of weeks in the current moment's year, according to ISO weeks.
func (g *Goment) ISOWeeksInYear() int {
	return 0
}

// Format functions.

// DaysInMonth returns the number of days in the set month.
func (g *Goment) DaysInMonth() int {
	return daysInMonth(g.Month(), g.Year())
}

func daysInMonth(month, year int) int {
	return time.Date(year, time.Month(month+1), 0, 0, 0, 0, 0, time.UTC).Day()
}

// Query functions

// IsLeapYear returns true if that year is a leap year, and false if it is not.
func (g *Goment) IsLeapYear() bool {
	last := time.Date(g.Year()+1, 1, 0, 0, 0, 0, 0, time.UTC).YearDay()
	return last == 366
}
