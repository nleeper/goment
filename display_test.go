package goment

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDiff(t *testing.T) {
	nowLib, err := New()
	if assert.NoError(t, err) {
		oneHourLib, oneHourErr := New(nowLib)
		if assert.NoError(t, oneHourErr) {
			oneHourLib.Add(1, "hours")
		}

		assert.Equal(t, 60*60, oneHourLib.Diff(nowLib), "1 hour from now = 3600")
	}
}

func TestDiffAfter(t *testing.T) {
	nowLib, err := New()
	if assert.NoError(t, err) {
		oneHourLib, oneHourErr := New(nowLib)
		if assert.NoError(t, oneHourErr) {
			oneHourLib.Add(1, "hours")
		}

		assert.Equal(t, -1*60*60, nowLib.Diff(oneHourLib), "1 hour from now = 3600")
	}
}

func TestDiffAfterWithUnits(t *testing.T) {
	lib, err := New(DateTime{
		Year:  2010,
		Month: 1,
	})

	if assert.NoError(t, err) {
		assert.Equal(t, -1, lib.Diff(DateTime{Year: 2011, Month: 1}, "years"), "year diff")
		assert.Equal(t, -2, lib.Diff(DateTime{Year: 2010, Month: 3}, "months"), "month diff")
		assert.Equal(t, 0, lib.Diff(DateTime{Year: 2010, Month: 1, Day: 6}, "weeks"), "week diff")
		assert.Equal(t, -1, lib.Diff(DateTime{Year: 2010, Month: 1, Day: 7}, "weeks"), "week diff")
		assert.Equal(t, -2, lib.Diff(DateTime{Year: 2010, Month: 1, Day: 20}, "weeks"), "week diff")
		assert.Equal(t, -3, lib.Diff(DateTime{Year: 2010, Month: 1, Day: 21}, "weeks"), "week diff")
		assert.Equal(t, -3, lib.Diff(DateTime{Year: 2010, Month: 1, Day: 3}, "days"), "day diff")
		assert.Equal(t, -4, lib.Diff(DateTime{Year: 2010, Month: 1, Day: 0, Hour: 4}, "hours"), "hour diff")
		assert.Equal(t, -5, lib.Diff(DateTime{Year: 2010, Month: 1, Day: 0, Hour: 0, Minute: 5}, "minutes"), "minute diff")
		assert.Equal(t, -6, lib.Diff(DateTime{Year: 2010, Month: 1, Day: 0, Hour: 0, Minute: 0, Second: 6}, "seconds"), "second diff")
	}
}

func TestDiffBeforeWithUnits(t *testing.T) {
	assert.Equal(t, 1, simple(DateTime{Year: 2010, Month: 1}).Diff(DateTime{Year: 2009, Month: 1}, "years"), "year diff")
	assert.Equal(t, 2, simple(DateTime{Year: 2010, Month: 3}).Diff(DateTime{Year: 2010, Month: 1}, "months"), "month diff")
	assert.Equal(t, 3, simple(DateTime{Year: 2010, Month: 1, Day: 3}).Diff(DateTime{Year: 2010, Month: 1, Day: 0}, "days"), "day diff")
	assert.Equal(t, 0, simple(DateTime{Year: 2010, Month: 1, Day: 6}).Diff(DateTime{Year: 2010, Month: 1}, "weeks"), "week diff")
	assert.Equal(t, 1, simple(DateTime{Year: 2010, Month: 1, Day: 7}).Diff(DateTime{Year: 2010, Month: 1}, "weeks"), "week diff")
	assert.Equal(t, 2, simple(DateTime{Year: 2010, Month: 1, Day: 20}).Diff(DateTime{Year: 2010, Month: 1}, "weeks"), "week diff")
	assert.Equal(t, 3, simple(DateTime{Year: 2010, Month: 1, Day: 21}).Diff(DateTime{Year: 2010, Month: 1}, "weeks"), "week diff")
	assert.Equal(t, 4, simple(DateTime{Year: 2010, Month: 1, Day: 0, Hour: 4}).Diff(DateTime{Year: 2010, Month: 1}, "hours"), "hour diff")
	assert.Equal(t, 5, simple(DateTime{Year: 2010, Month: 1, Day: 0, Hour: 0, Minute: 5}).Diff(DateTime{Year: 2010, Month: 1}, "minutes"), "minute diff")
	assert.Equal(t, 6, simple(DateTime{Year: 2010, Month: 1, Day: 0, Hour: 0, Minute: 0, Second: 6}).Diff(DateTime{Year: 2010, Month: 1}, "seconds"), "second diff")
}

func TestDiffMonth(t *testing.T) {
	assert.Equal(t, -1, simple(DateTime{Year: 2011, Month: 1, Day: 31}).Diff(DateTime{Year: 2011, Month: 3, Day: 1}, "months"), "month diff")
}

func TestDiffOverflow(t *testing.T) {
	assert.Equal(t, 12, simple(DateTime{Year: 2011, Month: 1}).Diff(DateTime{Year: 2010, Month: 1}, "months"), "month diff")
	assert.Equal(t, 24, simple(DateTime{Year: 2010, Month: 1, Day: 1}).Diff(DateTime{Year: 2010, Month: 1}, "hours"), "hour diff")
	assert.Equal(t, 120, simple(DateTime{Year: 2010, Month: 1, Hour: 2}).Diff(DateTime{Year: 2010, Month: 1}, "minutes"), "minute diff")
	assert.Equal(t, 240, simple(DateTime{Year: 2010, Month: 1, Minute: 4}).Diff(DateTime{Year: 2010, Month: 1}), "second diff")
}

func TestDiffBetweenUTCAndLocal(t *testing.T) {
	if simple(DateTime{Year: 2012, Month: 1}).UTCOffset() == simple(DateTime{Year: 2011, Month: 1}).UTCOffset() {
		// Russia's utc offset on 1st of Jan 2012 vs 2011 is different
		assert.Equal(t, 1, simple(DateTime{Year: 2012, Month: 1}).Diff(simple(DateTime{Year: 2011, Month: 1}), "years"), "year diff")
	}

	assert.Equal(t, 2, simple(DateTime{Year: 2010, Month: 3, Day: 2}).UTC().Diff(DateTime{Year: 2010, Month: 1, Day: 2}, "months"), "month diff")
	assert.Equal(t, 3, simple(DateTime{Year: 2010, Month: 1, Day: 3}).UTC().Diff(DateTime{Year: 2010, Month: 1}, "days"), "day diff")
	assert.Equal(t, 3, simple(DateTime{Year: 2010, Month: 1, Day: 21}).UTC().Diff(DateTime{Year: 2010, Month: 1}, "weeks"), "week diff")
	assert.Equal(t, 4, simple(DateTime{Year: 2010, Month: 1, Hour: 4}).UTC().Diff(DateTime{Year: 2010, Month: 1}, "hours"), "hour diff")
	assert.Equal(t, 5, simple(DateTime{Year: 2010, Month: 1, Minute: 5}).UTC().Diff(DateTime{Year: 2010, Month: 1}, "minutes"), "minute diff")
	assert.Equal(t, 6, simple(DateTime{Year: 2010, Month: 1, Second: 6}).UTC().Diff(DateTime{Year: 2010, Month: 1}, "seconds"), "second diff")
}

func TestDiffFloored(t *testing.T) {
	assert.Equal(t, 0, simple(DateTime{Year: 2010, Month: 1, Hour: 23}).Diff(DateTime{Year: 2010, Month: 1}, "day"), "23 hours = 0 days")
	assert.Equal(t, 0, simple(DateTime{Year: 2010, Month: 1, Hour: 23, Minute: 59}).Diff(DateTime{Year: 2010, Month: 1}, "day"), "23:59 hours = 0 days")
	assert.Equal(t, 1, simple(DateTime{Year: 2010, Month: 1, Hour: 24}).Diff(DateTime{Year: 2010, Month: 1}, "day"), "24 hours = 1 day")
	assert.Equal(t, 0, simple(DateTime{Year: 2010, Month: 1, Day: 1}).Diff(DateTime{Year: 2011, Month: 1}, "year"), "year rounded down")
	assert.Equal(t, 0, simple(DateTime{Year: 2011, Month: 1}).Diff(DateTime{Year: 2010, Month: 1, Day: 1}, "year"), "year rounded down")
	assert.Equal(t, -1, simple(DateTime{Year: 2010, Month: 1}).Diff(DateTime{Year: 2011, Month: 1}, "year"), "year rounded down")
	assert.Equal(t, 1, simple(DateTime{Year: 2011, Month: 1}).Diff(DateTime{Year: 2010, Month: 1}, "year"), "year rounded down")
}

func TestDaysInMonth(t *testing.T) {
	lib, err := New("2000-01-25")
	if assert.NoError(t, err) {
		assert.Equal(t, lib.DaysInMonth(), 31)
	}

	lib, err = New("2000-02-05")
	if assert.NoError(t, err) {
		assert.Equal(t, lib.DaysInMonth(), 29)
	}

	lib, err = New("2001-02-05")
	if assert.NoError(t, err) {
		assert.Equal(t, lib.DaysInMonth(), 28)
	}

	lib, err = New("2000-06-12")
	if assert.NoError(t, err) {
		assert.Equal(t, lib.DaysInMonth(), 30)
	}

	lib, err = New("2000-12-25")
	if assert.NoError(t, err) {
		assert.Equal(t, lib.DaysInMonth(), 31)
	}
}

func TestToUnix(t *testing.T) {
	testTime := time.Now()

	lib, err := New(testTime)
	if assert.NoError(t, err) {
		assert.Equal(t, testTime.Unix(), lib.ToUnix())
	}
}

func TestToArray(t *testing.T) {
	testTime := time.Now()

	lib, err := New(testTime)
	if assert.NoError(t, err) {
		assert.Equal(t, []int{testTime.Year(), int(testTime.Month()), testTime.Day(), testTime.Hour(), testTime.Minute(), testTime.Second(), testTime.Nanosecond()}, lib.ToArray())
	}
}

func TestToDateTime(t *testing.T) {
	testTime := time.Now()

	lib, err := New(testTime)
	if assert.NoError(t, err) {
		assert.Equal(t, DateTime{testTime.Year(), int(testTime.Month()), testTime.Day(), testTime.Hour(), testTime.Minute(), testTime.Second(), testTime.Nanosecond(), time.Local}, lib.ToDateTime())
	}
}

func TestToDateTimeWithUTC(t *testing.T) {
	testTime := time.Now().UTC()

	lib, err := New(testTime)
	if assert.NoError(t, err) {
		assert.Equal(t, DateTime{testTime.Year(), int(testTime.Month()), testTime.Day(), testTime.Hour(), testTime.Minute(), testTime.Second(), testTime.Nanosecond(), time.UTC}, lib.ToDateTime())
	}
}

func TestToString(t *testing.T) {
	testTime := time.Now()

	lib, err := New(testTime)
	if assert.NoError(t, err) {
		assert.Equal(t, testTime.String(), lib.ToString())
	}
}

func TestToISOString(t *testing.T) {
	lib, err := New(time.Date(2016, 4, 12, 19, 46, 47, 286000000, time.UTC))
	if assert.NoError(t, err) {
		assert.Equal(t, "2016-04-12T19:46:47.286Z", lib.ToISOString())
	}
}
