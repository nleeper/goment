package goment

import (
	"math"
	"strconv"
	"strings"
)

var thresholds = map[string]int{
	"ss": 44, // a few seconds to seconds
	"s":  45, // seconds to minute
	"m":  45, // minutes to hour
	"h":  22, // hours to day
	"d":  26, // days to month
	"M":  11, // months to year
}

var defaultRelativeTime = map[string]string{
	"future": "in %s",
	"past":   "%s ago",
	"s":      "a few seconds",
	"ss":     "%d seconds",
	"m":      "a minute",
	"mm":     "%d minutes",
	"h":      "an hour",
	"hh":     "%d hours",
	"d":      "a day",
	"dd":     "%d days",
	"M":      "a month",
	"MM":     "%d months",
	"y":      "a year",
	"yy":     "%d years",
}

var defaultCalendar = map[string]string{
	"sameDay":  "Today at %t",
	"nextDay":  "Tomorrow at %t",
	"nextWeek": "%d at %t",
	"lastDay":  "Yesterday at %t",
	"lastWeek": "Last %d at %t",
	"sameElse": "%v",
}

// ToNow returns the relative time to now to the Goment time.
func (g *Goment) ToNow(args ...interface{}) string {
	withoutSuffix := false

	if len(args) >= 1 {
		withoutSuffix = args[0].(bool)
	}

	now, err := New()
	if err != nil {
		return ""
	}

	return g.To(now, withoutSuffix)
}

// To returns the relative time from the Goment time to the supplied time.
func (g *Goment) To(args ...interface{}) string {
	var to *Goment
	var err error

	numArgs := len(args)
	if numArgs > 0 {
		switch v := args[0].(type) {
		case *Goment:
			to = v
		default:
			to, err = New(v)
			if err != nil {
				return ""
			}
		}

		withoutSuffix := false
		if numArgs > 1 {
			withoutSuffix = args[1].(bool)
		}

		return humanize(to, g, withoutSuffix)
	}
	return ""
}

// FromNow returns the relative time from now to the Goment time.
func (g *Goment) FromNow(args ...interface{}) string {
	withoutSuffix := false

	if len(args) >= 1 {
		withoutSuffix = args[0].(bool)
	}

	now, err := New()
	if err != nil {
		return ""
	}

	return g.From(now, withoutSuffix)
}

// From returns the relative time from the supplied time to the Goment time.
func (g *Goment) From(args ...interface{}) string {
	var from *Goment
	var err error

	numArgs := len(args)
	if numArgs > 0 {
		switch v := args[0].(type) {
		case *Goment:
			from = v
		default:
			from, err = New(v)
			if err != nil {
				return ""
			}
		}

		withoutSuffix := false
		if numArgs > 1 {
			withoutSuffix = args[1].(bool)
		}

		return humanize(g, from, withoutSuffix)
	}
	return ""
}

// Calendar displays time relative to a given referenceTime (defaults to now).
func (g *Goment) Calendar(args ...interface{}) string {
	var refTime *Goment
	var err error

	switch len(args) {
	case 0:
		refTime, err = New()
	default:
		switch v := args[1].(type) {
		case *Goment:
			refTime = v
		default:
			refTime, err = New(v)
		}
	}

	if err != nil {
		return ""
	}

	sod := refTime.StartOf("day")

	format := getCalendarFormat(g, sod)

	timeReplaced := strings.Replace(defaultCalendar[format], "%t", g.ToTime().Format("3:04pm"), 1)
	dayReplaced := strings.Replace(timeReplaced, "%d", g.ToTime().Weekday().String(), 1)
	dateReplaced := strings.Replace(dayReplaced, "%v", g.ToTime().Format("02/01/2006"), 1)

	return dateReplaced
}

func getCalendarFormat(g *Goment, sod *Goment) string {
	diff := g.ToTime().Sub(sod.ToTime())
	days := diff.Seconds() / 86400

	if days < -6 {
		return "sameElse"
	}
	if days < -1 {
		return "lastWeek"
	}
	if days < 0 {
		return "lastDay"
	}
	if days < 1 {
		return "sameDay"
	}
	if days < 2 {
		return "nextDay"
	}
	if days < 7 {
		return "nextWeek"
	}
	return "sameElse"
}

func humanize(to *Goment, from *Goment, withoutSuffix bool) string {
	localTo := to.Local()
	localFrom := from.Local()

	past := localTo.IsBefore(localFrom)

	diff := localFrom.ToTime().Sub(localTo.ToTime())

	seconds := roundAndAbs(diff.Seconds())
	hours := roundAndAbs(diff.Hours())
	minutes := roundAndAbs(diff.Minutes())
	days := roundAndAbs(divideSeconds(seconds, 86400))
	months := roundAndAbs(divideSeconds(seconds, 2600640))
	years := roundAndAbs(divideSeconds(seconds, 31207680))

	format := "yy"
	number := years

	if years <= 1 {
		format = "y"
	}

	if months < thresholds["M"] {
		format = "MM"
		number = months
	}

	if months <= 1 {
		format = "M"
	}

	if days < thresholds["d"] {
		format = "dd"
		number = days
	}

	if days <= 1 {
		format = "d"
	}

	if hours < thresholds["h"] {
		format = "hh"
		number = hours
	}

	if hours <= 1 {
		format = "h"
	}

	if minutes < thresholds["m"] {
		format = "mm"
		number = minutes
	}

	if minutes <= 1 {
		format = "m"
	}

	if seconds < thresholds["s"] {
		format = "ss"
		number = seconds
	}

	if seconds <= thresholds["ss"] {
		format = "s"
		number = seconds
	}

	relTime := strings.Replace(defaultRelativeTime[format], "%d", strconv.Itoa(number), 1)

	if withoutSuffix {
		return relTime
	}

	futurePast := defaultRelativeTime["future"]
	if past {
		futurePast = defaultRelativeTime["past"]
	}

	return strings.Replace(futurePast, "%s", relTime, 1)
}

func roundAndAbs(num float64) int {
	return abs(roundTime(num))
}

func divideSeconds(seconds int, divisor int) float64 {
	return float64(seconds) / float64(divisor)
}

func roundTime(input float64) int {
	var result float64

	if input < 0 {
		result = math.Ceil(input - 0.5)
	} else {
		result = math.Floor(input + 0.5)
	}

	i, _ := math.Modf(result)

	return int(i)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
