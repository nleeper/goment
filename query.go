package goment

import (
	"time"
)

// IsGoment will check if a variable is a Goment object.
func IsGoment(obj interface{}) bool {
	_, ok := obj.(*Goment)
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
