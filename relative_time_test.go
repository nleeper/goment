package goment

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestFromNowRelativeTime(t *testing.T) {
	var lib *Goment
	var err error

	testTime := time.Date(2000, 12, 15, 17, 8, 0, 0, time.UTC)
	timeNow = func() time.Time {
		return testTime
	}

	// Seconds to minutes threshold.
	lib, err = New(testTime)
	if assert.NoError(t, err) {
		lib.Subtract(44, "seconds")
		assert.Equal(t, "a few seconds ago", lib.FromNow(), "below default seconds to minutes threshold")
		assert.Equal(t, "a few seconds", lib.FromNow(true), "below default seconds to minutes threshold without suffix")
		lib.Subtract(1, "second")
		assert.Equal(t, "a minute ago", lib.FromNow(), "above default seconds to minutes threshold")
		assert.Equal(t, "a minute", lib.FromNow(true), "above default seconds to minutes threshold without suffix")
	}

	// Minutes to hours threshold.
	lib, err = New(testTime)
	if assert.NoError(t, err) {
		lib.Subtract(44, "minutes")
		assert.Equal(t, "44 minutes ago", lib.FromNow(), "below default minute to hour threshold")
		assert.Equal(t, "44 minutes", lib.FromNow(true), "below default minute to hour threshold without suffix")
		lib.Subtract(1, "minutes")
		assert.Equal(t, "an hour ago", lib.FromNow(), "above default minute to hour threshold")
		assert.Equal(t, "an hour", lib.FromNow(true), "above default minute to hour threshold without suffix")
	}

	// Hours to days threshold.
	lib, err = New(testTime)
	if assert.NoError(t, err) {
		lib.Subtract(21, "hours")
		assert.Equal(t, "21 hours ago", lib.FromNow(), "below default hours to day threshold")
		assert.Equal(t, "21 hours", lib.FromNow(true), "below default hours to day threshold without suffix")
		lib.Subtract(1, "hour")
		assert.Equal(t, "a day ago", lib.FromNow(), "above default hours to day threshold")
		assert.Equal(t, "a day", lib.FromNow(true), "above default hours to day threshold without suffix")
	}

	// Days to month threshold.
	lib, err = New(testTime)
	if assert.NoError(t, err) {
		lib.Subtract(25, "days")
		assert.Equal(t, "25 days ago", lib.FromNow(), "below default days to month threshold")
		assert.Equal(t, "25 days", lib.FromNow(true), "below default days to month threshold without suffix")
		lib.Subtract(1, "day")
		assert.Equal(t, "a month ago", lib.FromNow(), "above default days to month threshold")
		assert.Equal(t, "a month", lib.FromNow(true), "above default days to month threshold without suffix")
	}

	// Months to year threshold.
	lib, err = New(testTime)
	if assert.NoError(t, err) {
		lib.Subtract(10, "months")
		assert.Equal(t, "10 months ago", lib.FromNow(), "below default days to years threshold")
		assert.Equal(t, "10 months", lib.FromNow(true), "below default days to years threshold without suffix")
		lib.Subtract(1, "month")
		assert.Equal(t, "a year ago", lib.FromNow(), "above default days to years threshold")
		assert.Equal(t, "a year", lib.FromNow(true), "above default days to years threshold without suffix")
		lib.Subtract(2, "years")
		assert.Equal(t, "3 years ago", lib.FromNow(), "years threshold")
		assert.Equal(t, "3 years", lib.FromNow(true), "years threshold without suffix")
	}

	// Reset timeNow.
	timeNow = time.Now
}

func TestToNowRelativeTime(t *testing.T) {
	var lib *Goment
	var err error

	testTime := time.Date(2000, 12, 15, 17, 8, 0, 0, time.UTC)
	timeNow = func() time.Time {
		return testTime
	}

	// Seconds to minutes threshold.
	lib, err = New(testTime)
	if assert.NoError(t, err) {
		lib.Subtract(44, "seconds")
		assert.Equal(t, "in a few seconds", lib.ToNow(), "below default seconds to minutes threshold")
		assert.Equal(t, "a few seconds", lib.ToNow(true), "below default seconds to minutes threshold without suffix")
		lib.Subtract(1, "second")
		assert.Equal(t, "in a minute", lib.ToNow(), "above default seconds to minutes threshold")
		assert.Equal(t, "a minute", lib.ToNow(true), "above default seconds to minutes threshold without suffix")
	}

	// Minutes to hours threshold.
	lib, err = New(testTime)
	if assert.NoError(t, err) {
		lib.Subtract(44, "minutes")
		assert.Equal(t, "in 44 minutes", lib.ToNow(), "below default minute to hour threshold")
		assert.Equal(t, "44 minutes", lib.ToNow(true), "below default minute to hour threshold without suffix")
		lib.Subtract(1, "minutes")
		assert.Equal(t, "in an hour", lib.ToNow(), "above default minute to hour threshold")
		assert.Equal(t, "an hour", lib.ToNow(true), "above default minute to hour threshold without suffix")
	}

	// Hours to days threshold.
	lib, err = New(testTime)
	if assert.NoError(t, err) {
		lib.Subtract(21, "hours")
		assert.Equal(t, "in 21 hours", lib.ToNow(), "below default hours to day threshold")
		assert.Equal(t, "21 hours", lib.ToNow(true), "below default hours to day threshold without suffix")
		lib.Subtract(1, "hour")
		assert.Equal(t, "in a day", lib.ToNow(), "above default hours to day threshold")
		assert.Equal(t, "a day", lib.ToNow(true), "above default hours to day threshold without suffix")
	}

	// Days to month threshold.
	lib, err = New(testTime)
	if assert.NoError(t, err) {
		lib.Subtract(25, "days")
		assert.Equal(t, "in 25 days", lib.ToNow(), "below default days to month threshold")
		assert.Equal(t, "25 days", lib.ToNow(true), "below default days to month threshold without suffix")
		lib.Subtract(1, "day")
		assert.Equal(t, "in a month", lib.ToNow(), "above default days to month threshold")
		assert.Equal(t, "a month", lib.ToNow(true), "above default days to month threshold without suffix")
	}

	// Months to year threshold.
	lib, err = New(testTime)
	if assert.NoError(t, err) {
		lib.Subtract(10, "months")
		assert.Equal(t, "in 10 months", lib.ToNow(), "below default days to years threshold")
		assert.Equal(t, "10 months", lib.ToNow(true), "below default days to years threshold without suffix")
		lib.Subtract(1, "month")
		assert.Equal(t, "in a year", lib.ToNow(), "above default days to years threshold")
		assert.Equal(t, "a year", lib.ToNow(true), "above default days to years threshold without suffix")
		lib.Subtract(2, "years")
		assert.Equal(t, "in 3 years", lib.ToNow(), "years threshold")
		assert.Equal(t, "3 years", lib.ToNow(true), "years threshold without suffix")
	}

	// Reset timeNow.
	timeNow = time.Now
}

func TestFromRelativeTime(t *testing.T) {
	testTime := time.Date(2007, 1, 28, 0, 0, 0, 0, chicagoLocation())
	lib := simpleTime(testTime)

	assert.Equal(t, "a few seconds ago", lib.From(simpleTime(testTime).Add(44, "s")), "44 seconds = a few seconds ago")
	assert.Equal(t, "a few seconds", lib.From(simpleTime(testTime).Add(44, "s"), true), "44 seconds = a few seconds")
	assert.Equal(t, "a minute ago", lib.From(simpleTime(testTime).Add(45, "s")), "45 seconds = a minute ago")
	assert.Equal(t, "a minute", lib.From(simpleTime(testTime).Add(45, "s"), true), "45 seconds = a minute")
	assert.Equal(t, "a minute ago", lib.From(simpleTime(testTime).Add(89, "s")), "89 seconds = a minute ago")
	assert.Equal(t, "a minute", lib.From(simpleTime(testTime).Add(89, "s"), true), "89 seconds = a minute")
	assert.Equal(t, "2 minutes ago", lib.From(simpleTime(testTime).Add(90, "s")), "90 seconds = 2 minutes ago")
	assert.Equal(t, "2 minutes", lib.From(simpleTime(testTime).Add(90, "s"), true), "89 seconds = 2 minutes")
	assert.Equal(t, "44 minutes ago", lib.From(simpleTime(testTime).Add(44, "m")), "44 minutes = 44 minutes ago")
	assert.Equal(t, "44 minutes", lib.From(simpleTime(testTime).Add(44, "m"), true), "44 minutes = 44 minutes")
	assert.Equal(t, "an hour ago", lib.From(simpleTime(testTime).Add(45, "m")), "45 minutes = an hour ago")
	assert.Equal(t, "an hour", lib.From(simpleTime(testTime).Add(45, "m"), true), "45 minutes = an hour")
	assert.Equal(t, "an hour ago", lib.From(simpleTime(testTime).Add(89, "m")), "89 minutes = an hour ago")
	assert.Equal(t, "an hour", lib.From(simpleTime(testTime).Add(89, "m"), true), "89 minutes = an hour")
	assert.Equal(t, "2 hours ago", lib.From(simpleTime(testTime).Add(90, "m")), "90 minutes = 2 hours ago")
	assert.Equal(t, "2 hours", lib.From(simpleTime(testTime).Add(90, "m"), true), "90 minutes = 2 hours")
	assert.Equal(t, "5 hours ago", lib.From(simpleTime(testTime).Add(5, "h")), "5 hours = 5 hours ago")
	assert.Equal(t, "5 hours", lib.From(simpleTime(testTime).Add(5, "h"), true), "5 hours = 5 hours")
	assert.Equal(t, "21 hours ago", lib.From(simpleTime(testTime).Add(21, "h")), "21 hours = 21 hours ago")
	assert.Equal(t, "21 hours", lib.From(simpleTime(testTime).Add(21, "h"), true), "21 hours = 21 hours")
	assert.Equal(t, "a day ago", lib.From(simpleTime(testTime).Add(22, "h")), "22 hours = a day ago")
	assert.Equal(t, "a day", lib.From(simpleTime(testTime).Add(22, "h"), true), "22 hours = a day")
	assert.Equal(t, "a day ago", lib.From(simpleTime(testTime).Add(35, "h")), "35 hours = a day ago")
	assert.Equal(t, "a day", lib.From(simpleTime(testTime).Add(35, "h"), true), "35 hours = a day")
	assert.Equal(t, "2 days ago", lib.From(simpleTime(testTime).Add(36, "h")), "36 hours = 2 days ago")
	assert.Equal(t, "2 days", lib.From(simpleTime(testTime).Add(36, "h"), true), "36 hours = 2 days")
	assert.Equal(t, "a day ago", lib.From(simpleTime(testTime).Add(1, "d")), "1 day = a day ago")
	assert.Equal(t, "a day", lib.From(simpleTime(testTime).Add(1, "d"), true), "1 day = a day")
	assert.Equal(t, "5 days ago", lib.From(simpleTime(testTime).Add(5, "d")), "5 days = 5 days ago")
	assert.Equal(t, "5 days", lib.From(simpleTime(testTime).Add(5, "d"), true), "5 days = 5 days")
	assert.Equal(t, "25 days ago", lib.From(simpleTime(testTime).Add(25, "d")), "25 days = 25 days ago")
	assert.Equal(t, "25 days", lib.From(simpleTime(testTime).Add(25, "d"), true), "25 days = 25 days")
	assert.Equal(t, "a month ago", lib.From(simpleTime(testTime).Add(26, "d")), "26 days = a month ago")
	assert.Equal(t, "a month", lib.From(simpleTime(testTime).Add(26, "d"), true), "26 days = a month")
	assert.Equal(t, "a month ago", lib.From(simpleTime(testTime).Add(30, "d")), "30 days = a month ago")
	assert.Equal(t, "a month", lib.From(simpleTime(testTime).Add(30, "d"), true), "30 days = a month")
	assert.Equal(t, "a month ago", lib.From(simpleTime(testTime).Add(43, "d")), "43 days = a month ago")
	assert.Equal(t, "a month", lib.From(simpleTime(testTime).Add(43, "d"), true), "43 days = a month")
	assert.Equal(t, "2 months ago", lib.From(simpleTime(testTime).Add(46, "d")), "46 days = 2 months ago")
	assert.Equal(t, "2 months", lib.From(simpleTime(testTime).Add(46, "d"), true), "46 days = 2 months")
	assert.Equal(t, "2 months ago", lib.From(simpleTime(testTime).Add(75, "d")), "75 days = 2 months ago")
	assert.Equal(t, "2 months", lib.From(simpleTime(testTime).Add(75, "d"), true), "75 days = 2 months")
	assert.Equal(t, "3 months ago", lib.From(simpleTime(testTime).Add(76, "d")), "76 days = 3 months ago")
	assert.Equal(t, "3 months", lib.From(simpleTime(testTime).Add(76, "d"), true), "76 days = 3 months")
	assert.Equal(t, "a month ago", lib.From(simpleTime(testTime).Add(1, "M")), "1 month = a month ago")
	assert.Equal(t, "a month", lib.From(simpleTime(testTime).Add(1, "M"), true), "1 month = a month")
	assert.Equal(t, "5 months ago", lib.From(simpleTime(testTime).Add(5, "M")), "5 months = 5 months ago")
	assert.Equal(t, "5 months", lib.From(simpleTime(testTime).Add(5, "M"), true), "5 months = 5 months")
	assert.Equal(t, "a year ago", lib.From(simpleTime(testTime).Add(345, "d")), "345 days = a year ago")
	assert.Equal(t, "a year", lib.From(simpleTime(testTime).Add(345, "d"), true), "345 days = a year")
	assert.Equal(t, "2 years ago", lib.From(simpleTime(testTime).Add(548, "d")), "548 days = 2 years ago")
	assert.Equal(t, "2 years", lib.From(simpleTime(testTime).Add(548, "d"), true), "548 days = 2 years")
	assert.Equal(t, "a year ago", lib.From(simpleTime(testTime).Add(1, "y")), "1 year = a year ago")
	assert.Equal(t, "a year", lib.From(simpleTime(testTime).Add(1, "y"), true), "1 year = a year")
	assert.Equal(t, "5 years ago", lib.From(simpleTime(testTime).Add(5, "y")), "5 years = 5 years ago")
	assert.Equal(t, "5 years", lib.From(simpleTime(testTime).Add(5, "y"), true), "5 years = 5 years")
}

func TestToRelativeTime(t *testing.T) {
	testTime := time.Date(2007, 1, 28, 0, 0, 0, 0, chicagoLocation())
	lib := simpleTime(testTime)

	assert.Equal(t, "in a few seconds", lib.To(simpleTime(testTime).Add(44, "s")), "44 seconds = in a few seconds")
	assert.Equal(t, "a few seconds", lib.To(simpleTime(testTime).Add(44, "s"), true), "44 seconds = a few seconds")
	assert.Equal(t, "in a minute", lib.To(simpleTime(testTime).Add(45, "s")), "45 seconds = a minute ago")
	assert.Equal(t, "a minute", lib.To(simpleTime(testTime).Add(45, "s"), true), "45 seconds = a minute")
	assert.Equal(t, "in a minute", lib.To(simpleTime(testTime).Add(89, "s")), "89 seconds = a minute ago")
	assert.Equal(t, "a minute", lib.To(simpleTime(testTime).Add(89, "s"), true), "89 seconds = a minute")
	assert.Equal(t, "in 2 minutes", lib.To(simpleTime(testTime).Add(90, "s")), "90 seconds = 2 minutes ago")
	assert.Equal(t, "2 minutes", lib.To(simpleTime(testTime).Add(90, "s"), true), "89 seconds = 2 minutes")
	assert.Equal(t, "in 44 minutes", lib.To(simpleTime(testTime).Add(44, "m")), "44 minutes = 44 minutes ago")
	assert.Equal(t, "44 minutes", lib.To(simpleTime(testTime).Add(44, "m"), true), "44 minutes = 44 minutes")
	assert.Equal(t, "in an hour", lib.To(simpleTime(testTime).Add(45, "m")), "45 minutes = an hour ago")
	assert.Equal(t, "an hour", lib.To(simpleTime(testTime).Add(45, "m"), true), "45 minutes = an hour")
	assert.Equal(t, "in an hour", lib.To(simpleTime(testTime).Add(89, "m")), "89 minutes = an hour ago")
	assert.Equal(t, "an hour", lib.To(simpleTime(testTime).Add(89, "m"), true), "89 minutes = an hour")
	assert.Equal(t, "in 2 hours", lib.To(simpleTime(testTime).Add(90, "m")), "90 minutes = 2 hours ago")
	assert.Equal(t, "2 hours", lib.To(simpleTime(testTime).Add(90, "m"), true), "90 minutes = 2 hours")
	assert.Equal(t, "in 5 hours", lib.To(simpleTime(testTime).Add(5, "h")), "5 hours = 5 hours ago")
	assert.Equal(t, "5 hours", lib.To(simpleTime(testTime).Add(5, "h"), true), "5 hours = 5 hours")
	assert.Equal(t, "in 21 hours", lib.To(simpleTime(testTime).Add(21, "h")), "21 hours = 21 hours ago")
	assert.Equal(t, "21 hours", lib.To(simpleTime(testTime).Add(21, "h"), true), "21 hours = 21 hours")
	assert.Equal(t, "in a day", lib.To(simpleTime(testTime).Add(22, "h")), "22 hours = a day ago")
	assert.Equal(t, "a day", lib.To(simpleTime(testTime).Add(22, "h"), true), "22 hours = a day")
	assert.Equal(t, "in a day", lib.To(simpleTime(testTime).Add(35, "h")), "35 hours = a day ago")
	assert.Equal(t, "a day", lib.To(simpleTime(testTime).Add(35, "h"), true), "35 hours = a day")
	assert.Equal(t, "in 2 days", lib.To(simpleTime(testTime).Add(36, "h")), "36 hours = 2 days ago")
	assert.Equal(t, "2 days", lib.To(simpleTime(testTime).Add(36, "h"), true), "36 hours = 2 days")
	assert.Equal(t, "in a day", lib.To(simpleTime(testTime).Add(1, "d")), "1 day = a day ago")
	assert.Equal(t, "a day", lib.To(simpleTime(testTime).Add(1, "d"), true), "1 day = a day")
	assert.Equal(t, "in 5 days", lib.To(simpleTime(testTime).Add(5, "d")), "5 days = 5 days ago")
	assert.Equal(t, "5 days", lib.To(simpleTime(testTime).Add(5, "d"), true), "5 days = 5 days")
	assert.Equal(t, "in 25 days", lib.To(simpleTime(testTime).Add(25, "d")), "25 days = 25 days ago")
	assert.Equal(t, "25 days", lib.To(simpleTime(testTime).Add(25, "d"), true), "25 days = 25 days")
	assert.Equal(t, "in a month", lib.To(simpleTime(testTime).Add(26, "d")), "26 days = a month ago")
	assert.Equal(t, "a month", lib.To(simpleTime(testTime).Add(26, "d"), true), "26 days = a month")
	assert.Equal(t, "in a month", lib.To(simpleTime(testTime).Add(30, "d")), "30 days = a month ago")
	assert.Equal(t, "a month", lib.To(simpleTime(testTime).Add(30, "d"), true), "30 days = a month")
	assert.Equal(t, "in a month", lib.To(simpleTime(testTime).Add(43, "d")), "43 days = a month ago")
	assert.Equal(t, "a month", lib.To(simpleTime(testTime).Add(43, "d"), true), "43 days = a month")
	assert.Equal(t, "in 2 months", lib.To(simpleTime(testTime).Add(46, "d")), "46 days = 2 months ago")
	assert.Equal(t, "2 months", lib.To(simpleTime(testTime).Add(46, "d"), true), "46 days = 2 months")
	assert.Equal(t, "in 2 months", lib.To(simpleTime(testTime).Add(75, "d")), "75 days = 2 months ago")
	assert.Equal(t, "2 months", lib.To(simpleTime(testTime).Add(75, "d"), true), "75 days = 2 months")
	assert.Equal(t, "in 3 months", lib.To(simpleTime(testTime).Add(76, "d")), "76 days = 3 months ago")
	assert.Equal(t, "3 months", lib.To(simpleTime(testTime).Add(76, "d"), true), "76 days = 3 months")
	assert.Equal(t, "in a month", lib.To(simpleTime(testTime).Add(1, "M")), "1 month = a month ago")
	assert.Equal(t, "a month", lib.To(simpleTime(testTime).Add(1, "M"), true), "1 month = a month")
	assert.Equal(t, "in 5 months", lib.To(simpleTime(testTime).Add(5, "M")), "5 months = 5 months ago")
	assert.Equal(t, "5 months", lib.To(simpleTime(testTime).Add(5, "M"), true), "5 months = 5 months")
	assert.Equal(t, "in a year", lib.To(simpleTime(testTime).Add(345, "d")), "345 days = a year ago")
	assert.Equal(t, "a year", lib.To(simpleTime(testTime).Add(345, "d"), true), "345 days = a year")
	assert.Equal(t, "in 2 years", lib.To(simpleTime(testTime).Add(548, "d")), "548 days = 2 years ago")
	assert.Equal(t, "2 years", lib.To(simpleTime(testTime).Add(548, "d"), true), "548 days = 2 years")
	assert.Equal(t, "in a year", lib.To(simpleTime(testTime).Add(1, "y")), "1 year = a year ago")
	assert.Equal(t, "a year", lib.To(simpleTime(testTime).Add(1, "y"), true), "1 year = a year")
	assert.Equal(t, "in 5 years", lib.To(simpleTime(testTime).Add(5, "y")), "5 years = 5 years ago")
	assert.Equal(t, "5 years", lib.To(simpleTime(testTime).Add(5, "y"), true), "5 years = 5 years")
}

func TestCalendarDay(t *testing.T) {
	testTime := time.Date(2000, 12, 15, 12, 0, 0, 0, time.UTC)
	timeNow = func() time.Time {
		return testTime
	}

	refTime := simpleTime(testTime)

	assert.Equal(t, "Today at 12:00pm", simpleGoment(refTime).Calendar(), "today at the same time")
	assert.Equal(t, "Today at 12:25pm", simpleGoment(refTime).Add(25, "m").Calendar(), "now plus 25 min")
	assert.Equal(t, "Today at 1:00pm", simpleGoment(refTime).Add(1, "h").Calendar(), "now plus 1 hour")
	assert.Equal(t, "Tomorrow at 12:00pm", simpleGoment(refTime).Add(1, "d").Calendar(), "tomorrow at the same time")
	assert.Equal(t, "Today at 11:00am", simpleGoment(refTime).Subtract(1, "h").Calendar(), "now minus 1 hour")
	assert.Equal(t, "Yesterday at 12:00pm", simpleGoment(refTime).Subtract(1, "d").Calendar(), "yesterday at the same time")

	refTime = simpleTime(testTime)
	assert.Equal(t, "Sunday at 12:00pm", refTime.Add(2, "d").Calendar(), "Today + 2 days current time")
	refTime.StartOf("day")
	assert.Equal(t, "Sunday at 12:00am", refTime.Calendar(), "Today + 2 days beginning of day")
	refTime.EndOf("day")
	assert.Equal(t, "Sunday at 11:59pm", refTime.Calendar(), "Today + 2 days end of day")

	refTime = simpleTime(testTime)
	assert.Equal(t, "Last Wednesday at 12:00pm", refTime.Subtract(2, "d").Calendar(), "Today - 2 days current time")
	refTime.StartOf("day")
	assert.Equal(t, "Last Wednesday at 12:00am", refTime.Calendar(), "Today - 2 days beginning of day")
	refTime.EndOf("day")
	assert.Equal(t, "Last Wednesday at 11:59pm", refTime.Calendar(), "Today - 2 days end of day")

	weeksAgo := simpleTime(testTime).Subtract(1, "w")
	weeksFromNow := simpleTime(testTime).Add(1, "w")

	assert.Equal(t, "08/12/2000", weeksAgo.Calendar())
	assert.Equal(t, "22/12/2000", weeksFromNow.Calendar())

	weeksAgo = simpleTime(testTime).Subtract(2, "w")
	weeksFromNow = simpleTime(testTime).Add(2, "w")

	assert.Equal(t, "01/12/2000", weeksAgo.Calendar())
	assert.Equal(t, "29/12/2000", weeksFromNow.Calendar())

	// Reset timeNow.
	timeNow = time.Now
}
