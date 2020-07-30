package goment

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestFromNowRelativeTime(t *testing.T) {
	assert := assert.New(t)

	testTime := time.Date(2000, 12, 15, 17, 8, 0, 0, time.UTC)
	timeNow = func() time.Time {
		return testTime
	}

	// Seconds to minutes threshold.
	lib := simpleTime(testTime)

	lib.Subtract(44, "seconds")
	assert.Equal("a few seconds ago", lib.FromNow(), "below default seconds to minutes threshold")
	assert.Equal("a few seconds", lib.FromNow(true), "below default seconds to minutes threshold without suffix")
	lib.Subtract(1, "second")
	assert.Equal("a minute ago", lib.FromNow(), "above default seconds to minutes threshold")
	assert.Equal("a minute", lib.FromNow(true), "above default seconds to minutes threshold without suffix")

	// Minutes to hours threshold.
	lib2 := simpleTime(testTime)

	lib2.Subtract(44, "minutes")
	assert.Equal("44 minutes ago", lib2.FromNow(), "below default minute to hour threshold")
	assert.Equal("44 minutes", lib2.FromNow(true), "below default minute to hour threshold without suffix")
	lib2.Subtract(1, "minutes")
	assert.Equal("an hour ago", lib2.FromNow(), "above default minute to hour threshold")
	assert.Equal("an hour", lib2.FromNow(true), "above default minute to hour threshold without suffix")

	// Hours to days threshold.
	lib3 := simpleTime(testTime)

	lib3.Subtract(21, "hours")
	assert.Equal("21 hours ago", lib3.FromNow(), "below default hours to day threshold")
	assert.Equal("21 hours", lib3.FromNow(true), "below default hours to day threshold without suffix")
	lib3.Subtract(1, "hour")
	assert.Equal("a day ago", lib3.FromNow(), "above default hours to day threshold")
	assert.Equal("a day", lib3.FromNow(true), "above default hours to day threshold without suffix")

	// Days to month threshold.
	lib4 := simpleTime(testTime)

	lib4.Subtract(25, "days")
	assert.Equal("25 days ago", lib4.FromNow(), "below default days to month threshold")
	assert.Equal("25 days", lib4.FromNow(true), "below default days to month threshold without suffix")
	lib4.Subtract(1, "day")
	assert.Equal("a month ago", lib4.FromNow(), "above default days to month threshold")
	assert.Equal("a month", lib4.FromNow(true), "above default days to month threshold without suffix")

	// Months to year threshold.
	lib5 := simpleTime(testTime)

	lib5.Subtract(10, "months")
	assert.Equal("10 months ago", lib5.FromNow(), "below default days to years threshold")
	assert.Equal("10 months", lib5.FromNow(true), "below default days to years threshold without suffix")
	lib5.Subtract(1, "month")
	assert.Equal("a year ago", lib5.FromNow(), "above default days to years threshold")
	assert.Equal("a year", lib5.FromNow(true), "above default days to years threshold without suffix")
	lib5.Subtract(2, "years")
	assert.Equal("3 years ago", lib5.FromNow(), "years threshold")
	assert.Equal("3 years", lib5.FromNow(true), "years threshold without suffix")

	// Reset timeNow.
	timeNow = time.Now
}

func TestToNowRelativeTime(t *testing.T) {
	assert := assert.New(t)

	testTime := time.Date(2000, 12, 15, 17, 8, 0, 0, time.UTC)
	timeNow = func() time.Time {
		return testTime
	}

	// Seconds to minutes threshold.
	lib := simpleTime(testTime)

	lib.Subtract(44, "seconds")
	assert.Equal("in a few seconds", lib.ToNow(), "below default seconds to minutes threshold")
	assert.Equal("a few seconds", lib.ToNow(true), "below default seconds to minutes threshold without suffix")
	lib.Subtract(1, "second")
	assert.Equal("in a minute", lib.ToNow(), "above default seconds to minutes threshold")
	assert.Equal("a minute", lib.ToNow(true), "above default seconds to minutes threshold without suffix")

	// Minutes to hours threshold.
	lib2 := simpleTime(testTime)

	lib2.Subtract(44, "minutes")
	assert.Equal("in 44 minutes", lib2.ToNow(), "below default minute to hour threshold")
	assert.Equal("44 minutes", lib2.ToNow(true), "below default minute to hour threshold without suffix")
	lib2.Subtract(1, "minutes")
	assert.Equal("in an hour", lib2.ToNow(), "above default minute to hour threshold")
	assert.Equal("an hour", lib2.ToNow(true), "above default minute to hour threshold without suffix")

	// Hours to days threshold.
	lib3 := simpleTime(testTime)

	lib3.Subtract(21, "hours")
	assert.Equal("in 21 hours", lib3.ToNow(), "below default hours to day threshold")
	assert.Equal("21 hours", lib3.ToNow(true), "below default hours to day threshold without suffix")
	lib3.Subtract(1, "hour")
	assert.Equal("in a day", lib3.ToNow(), "above default hours to day threshold")
	assert.Equal("a day", lib3.ToNow(true), "above default hours to day threshold without suffix")

	// Days to month threshold.
	lib4 := simpleTime(testTime)

	lib4.Subtract(25, "days")
	assert.Equal("in 25 days", lib4.ToNow(), "below default days to month threshold")
	assert.Equal("25 days", lib4.ToNow(true), "below default days to month threshold without suffix")
	lib4.Subtract(1, "day")
	assert.Equal("in a month", lib4.ToNow(), "above default days to month threshold")
	assert.Equal("a month", lib4.ToNow(true), "above default days to month threshold without suffix")

	// Months to year threshold.
	lib5 := simpleTime(testTime)

	lib5.Subtract(10, "months")
	assert.Equal("in 10 months", lib5.ToNow(), "below default days to years threshold")
	assert.Equal("10 months", lib5.ToNow(true), "below default days to years threshold without suffix")
	lib5.Subtract(1, "month")
	assert.Equal("in a year", lib5.ToNow(), "above default days to years threshold")
	assert.Equal("a year", lib5.ToNow(true), "above default days to years threshold without suffix")
	lib5.Subtract(2, "years")
	assert.Equal("in 3 years", lib5.ToNow(), "years threshold")
	assert.Equal("3 years", lib5.ToNow(true), "years threshold without suffix")

	// Reset timeNow.
	timeNow = time.Now
}

func TestFromRelativeTime(t *testing.T) {
	assert := assert.New(t)

	testTime := time.Date(2007, 1, 28, 0, 0, 0, 0, chicagoLocation())
	lib := simpleTime(testTime)

	assert.Equal("a few seconds ago", lib.From(simpleTime(testTime).Add(44, "s")), "44 seconds = a few seconds ago")
	assert.Equal("a few seconds", lib.From(simpleTime(testTime).Add(44, "s"), true), "44 seconds = a few seconds")
	assert.Equal("a minute ago", lib.From(simpleTime(testTime).Add(45, "s")), "45 seconds = a minute ago")
	assert.Equal("a minute", lib.From(simpleTime(testTime).Add(45, "s"), true), "45 seconds = a minute")
	assert.Equal("a minute ago", lib.From(simpleTime(testTime).Add(89, "s")), "89 seconds = a minute ago")
	assert.Equal("a minute", lib.From(simpleTime(testTime).Add(89, "s"), true), "89 seconds = a minute")
	assert.Equal("2 minutes ago", lib.From(simpleTime(testTime).Add(90, "s")), "90 seconds = 2 minutes ago")
	assert.Equal("2 minutes", lib.From(simpleTime(testTime).Add(90, "s"), true), "89 seconds = 2 minutes")
	assert.Equal("44 minutes ago", lib.From(simpleTime(testTime).Add(44, "m")), "44 minutes = 44 minutes ago")
	assert.Equal("44 minutes", lib.From(simpleTime(testTime).Add(44, "m"), true), "44 minutes = 44 minutes")
	assert.Equal("an hour ago", lib.From(simpleTime(testTime).Add(45, "m")), "45 minutes = an hour ago")
	assert.Equal("an hour", lib.From(simpleTime(testTime).Add(45, "m"), true), "45 minutes = an hour")
	assert.Equal("an hour ago", lib.From(simpleTime(testTime).Add(89, "m")), "89 minutes = an hour ago")
	assert.Equal("an hour", lib.From(simpleTime(testTime).Add(89, "m"), true), "89 minutes = an hour")
	assert.Equal("2 hours ago", lib.From(simpleTime(testTime).Add(90, "m")), "90 minutes = 2 hours ago")
	assert.Equal("2 hours", lib.From(simpleTime(testTime).Add(90, "m"), true), "90 minutes = 2 hours")
	assert.Equal("5 hours ago", lib.From(simpleTime(testTime).Add(5, "h")), "5 hours = 5 hours ago")
	assert.Equal("5 hours", lib.From(simpleTime(testTime).Add(5, "h"), true), "5 hours = 5 hours")
	assert.Equal("21 hours ago", lib.From(simpleTime(testTime).Add(21, "h")), "21 hours = 21 hours ago")
	assert.Equal("21 hours", lib.From(simpleTime(testTime).Add(21, "h"), true), "21 hours = 21 hours")
	assert.Equal("a day ago", lib.From(simpleTime(testTime).Add(22, "h")), "22 hours = a day ago")
	assert.Equal("a day", lib.From(simpleTime(testTime).Add(22, "h"), true), "22 hours = a day")
	assert.Equal("a day ago", lib.From(simpleTime(testTime).Add(35, "h")), "35 hours = a day ago")
	assert.Equal("a day", lib.From(simpleTime(testTime).Add(35, "h"), true), "35 hours = a day")
	assert.Equal("2 days ago", lib.From(simpleTime(testTime).Add(36, "h")), "36 hours = 2 days ago")
	assert.Equal("2 days", lib.From(simpleTime(testTime).Add(36, "h"), true), "36 hours = 2 days")
	assert.Equal("a day ago", lib.From(simpleTime(testTime).Add(1, "d")), "1 day = a day ago")
	assert.Equal("a day", lib.From(simpleTime(testTime).Add(1, "d"), true), "1 day = a day")
	assert.Equal("5 days ago", lib.From(simpleTime(testTime).Add(5, "d")), "5 days = 5 days ago")
	assert.Equal("5 days", lib.From(simpleTime(testTime).Add(5, "d"), true), "5 days = 5 days")
	assert.Equal("25 days ago", lib.From(simpleTime(testTime).Add(25, "d")), "25 days = 25 days ago")
	assert.Equal("25 days", lib.From(simpleTime(testTime).Add(25, "d"), true), "25 days = 25 days")
	assert.Equal("a month ago", lib.From(simpleTime(testTime).Add(26, "d")), "26 days = a month ago")
	assert.Equal("a month", lib.From(simpleTime(testTime).Add(26, "d"), true), "26 days = a month")
	assert.Equal("a month ago", lib.From(simpleTime(testTime).Add(30, "d")), "30 days = a month ago")
	assert.Equal("a month", lib.From(simpleTime(testTime).Add(30, "d"), true), "30 days = a month")
	assert.Equal("a month ago", lib.From(simpleTime(testTime).Add(43, "d")), "43 days = a month ago")
	assert.Equal("a month", lib.From(simpleTime(testTime).Add(43, "d"), true), "43 days = a month")
	assert.Equal("2 months ago", lib.From(simpleTime(testTime).Add(46, "d")), "46 days = 2 months ago")
	assert.Equal("2 months", lib.From(simpleTime(testTime).Add(46, "d"), true), "46 days = 2 months")
	assert.Equal("2 months ago", lib.From(simpleTime(testTime).Add(75, "d")), "75 days = 2 months ago")
	assert.Equal("2 months", lib.From(simpleTime(testTime).Add(75, "d"), true), "75 days = 2 months")
	assert.Equal("3 months ago", lib.From(simpleTime(testTime).Add(76, "d")), "76 days = 3 months ago")
	assert.Equal("3 months", lib.From(simpleTime(testTime).Add(76, "d"), true), "76 days = 3 months")
	assert.Equal("a month ago", lib.From(simpleTime(testTime).Add(1, "M")), "1 month = a month ago")
	assert.Equal("a month", lib.From(simpleTime(testTime).Add(1, "M"), true), "1 month = a month")
	assert.Equal("5 months ago", lib.From(simpleTime(testTime).Add(5, "M")), "5 months = 5 months ago")
	assert.Equal("5 months", lib.From(simpleTime(testTime).Add(5, "M"), true), "5 months = 5 months")
	assert.Equal("a year ago", lib.From(simpleTime(testTime).Add(345, "d")), "345 days = a year ago")
	assert.Equal("a year", lib.From(simpleTime(testTime).Add(345, "d"), true), "345 days = a year")
	assert.Equal("2 years ago", lib.From(simpleTime(testTime).Add(548, "d")), "548 days = 2 years ago")
	assert.Equal("2 years", lib.From(simpleTime(testTime).Add(548, "d"), true), "548 days = 2 years")
	assert.Equal("a year ago", lib.From(simpleTime(testTime).Add(1, "y")), "1 year = a year ago")
	assert.Equal("a year", lib.From(simpleTime(testTime).Add(1, "y"), true), "1 year = a year")
	assert.Equal("5 years ago", lib.From(simpleTime(testTime).Add(5, "y")), "5 years = 5 years ago")
	assert.Equal("5 years", lib.From(simpleTime(testTime).Add(5, "y"), true), "5 years = 5 years")
}

func TestToRelativeTime(t *testing.T) {
	assert := assert.New(t)

	testTime := time.Date(2007, 1, 28, 0, 0, 0, 0, chicagoLocation())
	lib := simpleTime(testTime)

	assert.Equal("in a few seconds", lib.To(simpleTime(testTime).Add(44, "s")), "44 seconds = in a few seconds")
	assert.Equal("a few seconds", lib.To(simpleTime(testTime).Add(44, "s"), true), "44 seconds = a few seconds")
	assert.Equal("in a minute", lib.To(simpleTime(testTime).Add(45, "s")), "45 seconds = a minute ago")
	assert.Equal("a minute", lib.To(simpleTime(testTime).Add(45, "s"), true), "45 seconds = a minute")
	assert.Equal("in a minute", lib.To(simpleTime(testTime).Add(89, "s")), "89 seconds = a minute ago")
	assert.Equal("a minute", lib.To(simpleTime(testTime).Add(89, "s"), true), "89 seconds = a minute")
	assert.Equal("in 2 minutes", lib.To(simpleTime(testTime).Add(90, "s")), "90 seconds = 2 minutes ago")
	assert.Equal("2 minutes", lib.To(simpleTime(testTime).Add(90, "s"), true), "89 seconds = 2 minutes")
	assert.Equal("in 44 minutes", lib.To(simpleTime(testTime).Add(44, "m")), "44 minutes = 44 minutes ago")
	assert.Equal("44 minutes", lib.To(simpleTime(testTime).Add(44, "m"), true), "44 minutes = 44 minutes")
	assert.Equal("in an hour", lib.To(simpleTime(testTime).Add(45, "m")), "45 minutes = an hour ago")
	assert.Equal("an hour", lib.To(simpleTime(testTime).Add(45, "m"), true), "45 minutes = an hour")
	assert.Equal("in an hour", lib.To(simpleTime(testTime).Add(89, "m")), "89 minutes = an hour ago")
	assert.Equal("an hour", lib.To(simpleTime(testTime).Add(89, "m"), true), "89 minutes = an hour")
	assert.Equal("in 2 hours", lib.To(simpleTime(testTime).Add(90, "m")), "90 minutes = 2 hours ago")
	assert.Equal("2 hours", lib.To(simpleTime(testTime).Add(90, "m"), true), "90 minutes = 2 hours")
	assert.Equal("in 5 hours", lib.To(simpleTime(testTime).Add(5, "h")), "5 hours = 5 hours ago")
	assert.Equal("5 hours", lib.To(simpleTime(testTime).Add(5, "h"), true), "5 hours = 5 hours")
	assert.Equal("in 21 hours", lib.To(simpleTime(testTime).Add(21, "h")), "21 hours = 21 hours ago")
	assert.Equal("21 hours", lib.To(simpleTime(testTime).Add(21, "h"), true), "21 hours = 21 hours")
	assert.Equal("in a day", lib.To(simpleTime(testTime).Add(22, "h")), "22 hours = a day ago")
	assert.Equal("a day", lib.To(simpleTime(testTime).Add(22, "h"), true), "22 hours = a day")
	assert.Equal("in a day", lib.To(simpleTime(testTime).Add(35, "h")), "35 hours = a day ago")
	assert.Equal("a day", lib.To(simpleTime(testTime).Add(35, "h"), true), "35 hours = a day")
	assert.Equal("in 2 days", lib.To(simpleTime(testTime).Add(36, "h")), "36 hours = 2 days ago")
	assert.Equal("2 days", lib.To(simpleTime(testTime).Add(36, "h"), true), "36 hours = 2 days")
	assert.Equal("in a day", lib.To(simpleTime(testTime).Add(1, "d")), "1 day = a day ago")
	assert.Equal("a day", lib.To(simpleTime(testTime).Add(1, "d"), true), "1 day = a day")
	assert.Equal("in 5 days", lib.To(simpleTime(testTime).Add(5, "d")), "5 days = 5 days ago")
	assert.Equal("5 days", lib.To(simpleTime(testTime).Add(5, "d"), true), "5 days = 5 days")
	assert.Equal("in 25 days", lib.To(simpleTime(testTime).Add(25, "d")), "25 days = 25 days ago")
	assert.Equal("25 days", lib.To(simpleTime(testTime).Add(25, "d"), true), "25 days = 25 days")
	assert.Equal("in a month", lib.To(simpleTime(testTime).Add(26, "d")), "26 days = a month ago")
	assert.Equal("a month", lib.To(simpleTime(testTime).Add(26, "d"), true), "26 days = a month")
	assert.Equal("in a month", lib.To(simpleTime(testTime).Add(30, "d")), "30 days = a month ago")
	assert.Equal("a month", lib.To(simpleTime(testTime).Add(30, "d"), true), "30 days = a month")
	assert.Equal("in a month", lib.To(simpleTime(testTime).Add(43, "d")), "43 days = a month ago")
	assert.Equal("a month", lib.To(simpleTime(testTime).Add(43, "d"), true), "43 days = a month")
	assert.Equal("in 2 months", lib.To(simpleTime(testTime).Add(46, "d")), "46 days = 2 months ago")
	assert.Equal("2 months", lib.To(simpleTime(testTime).Add(46, "d"), true), "46 days = 2 months")
	assert.Equal("in 2 months", lib.To(simpleTime(testTime).Add(75, "d")), "75 days = 2 months ago")
	assert.Equal("2 months", lib.To(simpleTime(testTime).Add(75, "d"), true), "75 days = 2 months")
	assert.Equal("in 3 months", lib.To(simpleTime(testTime).Add(76, "d")), "76 days = 3 months ago")
	assert.Equal("3 months", lib.To(simpleTime(testTime).Add(76, "d"), true), "76 days = 3 months")
	assert.Equal("in a month", lib.To(simpleTime(testTime).Add(1, "M")), "1 month = a month ago")
	assert.Equal("a month", lib.To(simpleTime(testTime).Add(1, "M"), true), "1 month = a month")
	assert.Equal("in 5 months", lib.To(simpleTime(testTime).Add(5, "M")), "5 months = 5 months ago")
	assert.Equal("5 months", lib.To(simpleTime(testTime).Add(5, "M"), true), "5 months = 5 months")
	assert.Equal("in a year", lib.To(simpleTime(testTime).Add(345, "d")), "345 days = a year ago")
	assert.Equal("a year", lib.To(simpleTime(testTime).Add(345, "d"), true), "345 days = a year")
	assert.Equal("in 2 years", lib.To(simpleTime(testTime).Add(548, "d")), "548 days = 2 years ago")
	assert.Equal("2 years", lib.To(simpleTime(testTime).Add(548, "d"), true), "548 days = 2 years")
	assert.Equal("in a year", lib.To(simpleTime(testTime).Add(1, "y")), "1 year = a year ago")
	assert.Equal("a year", lib.To(simpleTime(testTime).Add(1, "y"), true), "1 year = a year")
	assert.Equal("in 5 years", lib.To(simpleTime(testTime).Add(5, "y")), "5 years = 5 years ago")
	assert.Equal("5 years", lib.To(simpleTime(testTime).Add(5, "y"), true), "5 years = 5 years")
}

func TestCalendarDay(t *testing.T) {
	assert := assert.New(t)

	testTime := time.Date(2000, 12, 15, 12, 0, 0, 0, time.UTC)
	timeNow = func() time.Time {
		return testTime
	}

	refTime := simpleTime(testTime)

	assert.Equal("Today at 12:00 PM", simpleGoment(refTime).Calendar(), "today at the same time")
	assert.Equal("Today at 12:25 PM", simpleGoment(refTime).Add(25, "m").Calendar(), "now plus 25 min")
	assert.Equal("Today at 1:00 PM", simpleGoment(refTime).Add(1, "h").Calendar(), "now plus 1 hour")
	assert.Equal("Tomorrow at 12:00 PM", simpleGoment(refTime).Add(1, "d").Calendar(), "tomorrow at the same time")
	assert.Equal("Today at 11:00 AM", simpleGoment(refTime).Subtract(1, "h").Calendar(), "now minus 1 hour")
	assert.Equal("Yesterday at 12:00 PM", simpleGoment(refTime).Subtract(1, "d").Calendar(), "yesterday at the same time")

	refTime2 := simpleTime(testTime)
	assert.Equal("Sunday at 12:00 PM", refTime2.Add(2, "d").Calendar(), "Today + 2 days current time")
	refTime2.StartOf("day")
	assert.Equal("Sunday at 12:00 AM", refTime2.Calendar(), "Today + 2 days beginning of day")
	refTime2.EndOf("day")
	assert.Equal("Sunday at 11:59 PM", refTime2.Calendar(), "Today + 2 days end of day")

	refTime3 := simpleTime(testTime)
	assert.Equal("Last Wednesday at 12:00 PM", refTime3.Subtract(2, "d").Calendar(), "Today - 2 days current time")
	refTime3.StartOf("day")
	assert.Equal("Last Wednesday at 12:00 AM", refTime3.Calendar(), "Today - 2 days beginning of day")
	refTime3.EndOf("day")
	assert.Equal("Last Wednesday at 11:59 PM", refTime3.Calendar(), "Today - 2 days end of day")

	weeksAgo := simpleTime(testTime).Subtract(1, "w")
	weeksFromNow := simpleTime(testTime).Add(1, "w")

	assert.Equal("12/08/2000", weeksAgo.Calendar())
	assert.Equal("12/22/2000", weeksFromNow.Calendar())

	weeksAgo = simpleTime(testTime).Subtract(2, "w")
	weeksFromNow = simpleTime(testTime).Add(2, "w")

	assert.Equal("12/01/2000", weeksAgo.Calendar())
	assert.Equal("12/29/2000", weeksFromNow.Calendar())

	// Reset timeNow.
	timeNow = time.Now
}
