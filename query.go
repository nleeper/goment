package goment

import (
	"math"
	"time"
)

type weekYear struct {
	week int
	year int
}

type dayOfYear struct {
	year      int
	dayOfYear int
}

// IsGoment will check if a variable is a Goment object.
func IsGoment(obj interface{}) bool {
	_, ok := obj.(*Goment)
	if !ok {
		_, ok2 := obj.(Goment)
		return ok2
	}
	return ok
}

// IsTime will check if a variable is a time.Time object.
func IsTime(obj interface{}) bool {
	_, ok := obj.(time.Time)
	return ok
}

// IsDST checks if the current Goment is in daylight saving time.
func (g *Goment) IsDST() bool {
	utcOffset := g.UTCOffset()
	janOffset := g.Clone().SetMonth(1).UTCOffset()
	juneOffset := g.Clone().SetMonth(6).UTCOffset()

	return utcOffset > janOffset || utcOffset > juneOffset
}

// IsLeapYear returns true if that year is a leap year, and false if it is not.
func (g *Goment) IsLeapYear() bool {
	return daysInYear(g.Year()) == 366
}

func daysInYear(year int) int {
	return time.Date(year+1, 1, 0, 0, 0, 0, 0, time.UTC).YearDay()
}

func weekOfYear(g *Goment, dow int, doy int) weekYear {
	weekOffset := firstWeekOffset(g.Year(), dow, doy)
	week := int(math.Floor(float64(g.DayOfYear()-weekOffset-1)/float64(7))) + 1

	resWeek := 0
	resYear := 0

	if week < 1 {
		resYear = g.Year() - 1
		resWeek = week + weeksInYear(resYear, dow, doy)
	} else if week > weeksInYear(g.Year(), dow, doy) {
		resWeek = week - weeksInYear(g.Year(), dow, doy)
		resYear = g.Year() + 1
	} else {
		resYear = g.Year()
		resWeek = week
	}

	return weekYear{
		week: resWeek,
		year: resYear,
	}
}

func weeksInYear(year int, dow int, doy int) int {
	weekOffset := firstWeekOffset(year, dow, doy)
	weekOffsetNext := firstWeekOffset(year+1, dow, doy)

	return (daysInYear(year) - weekOffset + weekOffsetNext) / 7
}

func firstWeekOffset(year int, dow int, doy int) int {
	fwd := 7 + dow - doy

	d, _ := New(DateTime{Year: year, Month: 1, Day: fwd})
	fwdlw := (7 + d.UTC().Day() - dow) % 7

	return -fwdlw + fwd - 1
}

func dayOfYearFromWeeks(year int, week int, weekday int, dow int, doy int) dayOfYear {
	localWeekday := (7 + weekday - dow) % 7
	weekOffset := firstWeekOffset(year, dow, doy)
	currDayOfYear := 1 + 7*(week-1) + localWeekday + weekOffset

	resYear := 0
	resDayOfYear := 0

	if currDayOfYear <= 0 {
		resYear = year - 1
		resDayOfYear = daysInYear(resYear) + currDayOfYear
	} else if currDayOfYear > daysInYear(year) {
		resYear = year + 1
		resDayOfYear = currDayOfYear - daysInYear(year)
	} else {
		resYear = year
		resDayOfYear = currDayOfYear
	}

	return dayOfYear{
		year:      resYear,
		dayOfYear: resDayOfYear,
	}
}
