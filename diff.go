package goment

import (
	"math"
	"time"
)

type diff struct {
	Start *Goment
	End   *Goment
}

// InYears returns the duration in number of years.
func (d diff) InYears() int {
	return d.monthDiff() / 12
}

// InMonths returns the duration in number of months.
func (d diff) InMonths() int {
	return d.monthDiff()
}

// InWeeks returns the duration in number of weeks.
func (d diff) InWeeks() int {
	return absFloor(math.Floor(float64(d.InDays() / 7)))
}

// InDays returns the duration in number of days.
func (d diff) InDays() int {
	return absFloor(math.Floor(float64(d.InHours()) / 24))
}

// InHours returns the duration in number of hours.
func (d diff) InHours() int {
	return absFloor(d.subtract().Hours())
}

// InMinutes returns the duration in number of minutes.
func (d diff) InMinutes() int {
	return absFloor(d.subtract().Minutes())
}

// InSeconds returns the duration in number of seconds.
func (d diff) InSeconds() int {
	return absFloor(d.subtract().Seconds())
}

func (d diff) subtract() time.Duration {
	return d.Start.ToTime().Sub(d.End.ToTime())
}

func (d diff) monthDiff() int {
	startYear, startMonth := d.Start.Year(), d.Start.Month()
	endYear, endMonth := d.End.Year(), d.End.Month()

	wholeMonthDiff := ((endYear - startYear) * 12) + int(endMonth-startMonth)
	anchor := d.Start.Clone().Add(wholeMonthDiff, "months")

	var adjust int
	var anchor2 *Goment

	if d.End.ToTime().Sub(anchor.ToTime()) < 0 {
		anchor2 = d.Start.Clone().Add(wholeMonthDiff-1, "months")
		adjust = int(d.End.ToTime().Sub(anchor.ToTime()).Seconds() / (anchor.ToTime().Sub(anchor2.ToTime())).Seconds())
	} else {
		anchor2 = d.Start.Clone().Add(wholeMonthDiff+1, "months")
		adjust = int(d.End.ToTime().Sub(anchor.ToTime()).Seconds() / (anchor2.ToTime().Sub(anchor.ToTime())).Seconds())
	}

	if wholeMonthDiff == 0 {
		return 0
	}

	return -(wholeMonthDiff + adjust)
}

func absFloor(number float64) int {
	if number < 0 {
		return int(math.Ceil(number))
	}
	return int(math.Floor(number))
}
