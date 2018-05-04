package goment

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetByUnit(t *testing.T) {
	lib, err := New("2015-04-06 10:11:12.4567")
	if assert.NoError(t, err) {
		assert.Equal(t, lib.Get("y"), 2015)
		assert.Equal(t, lib.Get("year"), 2015)
		assert.Equal(t, lib.Get("years"), 2015)
		assert.Equal(t, lib.Get("M"), 4)
		assert.Equal(t, lib.Get("month"), 4)
		assert.Equal(t, lib.Get("months"), 4)
		assert.Equal(t, lib.Get("D"), 6)
		assert.Equal(t, lib.Get("date"), 6)
		assert.Equal(t, lib.Get("dates"), 6)
		assert.Equal(t, lib.Get("h"), 10)
		assert.Equal(t, lib.Get("hour"), 10)
		assert.Equal(t, lib.Get("hours"), 10)
		assert.Equal(t, lib.Get("m"), 11)
		assert.Equal(t, lib.Get("minute"), 11)
		assert.Equal(t, lib.Get("minutes"), 11)
		assert.Equal(t, lib.Get("s"), 12)
		assert.Equal(t, lib.Get("second"), 12)
		assert.Equal(t, lib.Get("seconds"), 12)
		assert.Equal(t, lib.Get("ms"), 12000)
		assert.Equal(t, lib.Get("millisecond"), 12000)
		assert.Equal(t, lib.Get("milliseconds"), 12000)
		assert.Equal(t, lib.Get("ns"), 456700000)
		assert.Equal(t, lib.Get("nanosecond"), 456700000)
		assert.Equal(t, lib.Get("nanoseconds"), 456700000)
		assert.Equal(t, lib.Get("unknown"), 0)
	}
}

func TestGetNanosecond(t *testing.T) {
	testTime := time.Date(2017, 8, 2, 18, 0, 33, 5000, time.UTC)

	lib, err := New(testTime)
	if assert.NoError(t, err) {
		assert.Equal(t, lib.Nanosecond(), 5000)
	}
}

func TestGetMillisecond(t *testing.T) {
	testTime := time.Date(2017, 8, 2, 18, 0, 33, 0, time.UTC)

	lib, err := New(testTime)
	if assert.NoError(t, err) {
		assert.Equal(t, lib.Millisecond(), 33000)
	}
}

func TestGetSecond(t *testing.T) {
	testTime := time.Date(2017, 8, 2, 18, 0, 33, 0, time.UTC)

	lib, err := New(testTime)
	if assert.NoError(t, err) {
		assert.Equal(t, lib.Second(), 33)
	}
}

func TestGetMinute(t *testing.T) {
	testTime := time.Date(2017, 8, 2, 18, 44, 33, 0, time.UTC)

	lib, err := New(testTime)
	if assert.NoError(t, err) {
		assert.Equal(t, lib.Minute(), 44)
	}
}

func TestGetHour(t *testing.T) {
	testTime := time.Date(2017, 8, 2, 18, 0, 33, 0, time.Local)

	lib, err := New(testTime)
	if assert.NoError(t, err) {
		assert.Equal(t, lib.Hour(), 18)
	}
}

func TestGetDate(t *testing.T) {
	testTime := time.Date(2017, 8, 2, 18, 0, 0, 0, time.UTC)

	lib, err := New(testTime)
	if assert.NoError(t, err) {
		assert.Equal(t, lib.Date(), 2)
	}
}

func TestGetDay(t *testing.T) {
	testTime := time.Date(2017, 8, 2, 18, 0, 0, 0, time.UTC)

	lib, err := New(testTime)
	if assert.NoError(t, err) {
		assert.Equal(t, lib.Day(), 3)
	}
}

func TestGetISOWeekday(t *testing.T) {
	testTime := time.Date(2017, 8, 6, 18, 0, 0, 0, time.UTC)

	lib, err := New(testTime)
	if assert.NoError(t, err) {
		assert.Equal(t, lib.Day(), 0)
		assert.Equal(t, lib.ISOWeekday(), 7)
	}
}

func TestDayOfYear(t *testing.T) {
	testTime := time.Date(2017, 1, 3, 18, 0, 0, 0, time.UTC)

	lib, err := New(testTime)
	if assert.NoError(t, err) {
		assert.Equal(t, lib.DayOfYear(), 3)
	}
}

func TestISOWeek(t *testing.T) {
	testTime := time.Date(2005, 1, 1, 0, 0, 0, 0, time.Local)
	testTime2 := time.Date(2007, 12, 30, 0, 0, 0, 0, time.Local)

	lib, err := New(testTime)
	if assert.NoError(t, err) {
		assert.Equal(t, lib.ISOWeek(), 53)
	}

	lib2, err2 := New(testTime2)
	if assert.NoError(t, err2) {
		assert.Equal(t, lib2.ISOWeek(), 52)
	}
}

func TestMonth(t *testing.T) {
	testTime := time.Date(2017, 6, 3, 18, 0, 0, 0, time.UTC)

	lib, err := New(testTime)
	if assert.NoError(t, err) {
		assert.Equal(t, lib.Month(), 6)
	}
}

func TestQuarter(t *testing.T) {
	monthQuarters := map[time.Month]int{
		time.January:   1,
		time.February:  1,
		time.March:     1,
		time.April:     2,
		time.May:       2,
		time.June:      2,
		time.July:      3,
		time.August:    3,
		time.September: 3,
		time.October:   4,
		time.November:  4,
		time.December:  4,
	}

	for m, q := range monthQuarters {
		testTime := time.Date(2017, m, 1, 18, 0, 0, 0, time.UTC)
		lib, err := New(testTime)
		if assert.NoError(t, err) {
			assert.Equal(t, lib.Quarter(), q)
		}
	}
}

func TestYear(t *testing.T) {
	testTime := time.Date(2017, 1, 3, 18, 0, 0, 0, time.UTC)

	lib, err := New(testTime)
	if assert.NoError(t, err) {
		assert.Equal(t, lib.Year(), 2017)
	}
}

func TestISOWeekYear(t *testing.T) {
	testTime := time.Date(2005, 1, 1, 0, 0, 0, 0, time.Local)
	testTime2 := time.Date(2007, 12, 30, 0, 0, 0, 0, time.Local)

	lib, err := New(testTime)
	if assert.NoError(t, err) {
		assert.Equal(t, lib.ISOWeekYear(), 2004)
	}

	lib2, err2 := New(testTime2)
	if assert.NoError(t, err2) {
		assert.Equal(t, lib2.ISOWeekYear(), 2007)
	}
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

func TestIsLeapYear(t *testing.T) {
	lib, err := New("2000-06-06")
	if assert.NoError(t, err) {
		assert.True(t, lib.IsLeapYear())
	}

	lib, err = New("2010-05-07")
	if assert.NoError(t, err) {
		assert.False(t, lib.IsLeapYear())
	}
}
