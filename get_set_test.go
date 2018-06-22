package goment

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetByUnits(t *testing.T) {
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
	testTime := time.Date(2017, 8, 2, 18, 0, 33, 0, chicagoLocation())

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
	testTime := time.Date(2005, 1, 1, 0, 0, 0, 0, chicagoLocation())
	testTime2 := time.Date(2007, 12, 30, 0, 0, 0, 0, chicagoLocation())

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
	testTime := time.Date(2005, 1, 1, 0, 0, 0, 0, chicagoLocation())
	testTime2 := time.Date(2007, 12, 30, 0, 0, 0, 0, chicagoLocation())

	lib, err := New(testTime)
	if assert.NoError(t, err) {
		assert.Equal(t, lib.ISOWeekYear(), 2004)
	}

	lib2, err2 := New(testTime2)
	if assert.NoError(t, err2) {
		assert.Equal(t, lib2.ISOWeekYear(), 2007)
	}
}

func TestSetByUnits(t *testing.T) {
	lib, err := New("2015-04-06 10:11:12.4567")
	if assert.NoError(t, err) {
		lib.Set("y", 2016)
		assert.Equal(t, lib.Year(), 2016)
		lib.Set("year", 2017)
		assert.Equal(t, lib.Year(), 2017)
		lib.Set("years", 2018)
		assert.Equal(t, lib.Year(), 2018)
		lib.Set("M", 9)
		assert.Equal(t, lib.Month(), 9)
		lib.Set("month", 10)
		assert.Equal(t, lib.Month(), 10)
		lib.Set("months", 11)
		assert.Equal(t, lib.Month(), 11)
		lib.Set("D", 8)
		assert.Equal(t, lib.Date(), 8)
		lib.Set("D", 9)
		assert.Equal(t, lib.Date(), 9)
		lib.Set("D", 10)
		assert.Equal(t, lib.Date(), 10)
		lib.Set("h", 14)
		assert.Equal(t, lib.Hour(), 14)
		lib.Set("hour", 15)
		assert.Equal(t, lib.Hour(), 15)
		lib.Set("hours", 16)
		assert.Equal(t, lib.Hour(), 16)
		lib.Set("m", 17)
		assert.Equal(t, lib.Minute(), 17)
		lib.Set("minute", 18)
		assert.Equal(t, lib.Minute(), 18)
		lib.Set("minutes", 19)
		assert.Equal(t, lib.Minute(), 19)
		lib.Set("s", 20)
		assert.Equal(t, lib.Second(), 20)
		lib.Set("second", 21)
		assert.Equal(t, lib.Second(), 21)
		lib.Set("seconds", 22)
		assert.Equal(t, lib.Second(), 22)
		lib.Set("ms", 23000)
		assert.Equal(t, lib.Millisecond(), 23000)
		lib.Set("millisecond", 23000)
		assert.Equal(t, lib.Millisecond(), 23000)
		lib.Set("milliseconds", 23000)
		assert.Equal(t, lib.Millisecond(), 23000)
		lib.Set("ns", 100000)
		assert.Equal(t, lib.Nanosecond(), 100000)
		lib.Set("nanosecond", 100001)
		assert.Equal(t, lib.Nanosecond(), 100001)
		lib.Set("nanoseconds", 100002)
		assert.Equal(t, lib.Nanosecond(), 100002)
	}
}

func TestSetUnknownUnits(t *testing.T) {
	testTime := time.Date(2011, 10, 11, 15, 20, 1, 10000, chicagoLocation())

	lib, err := New(testTime)
	if assert.NoError(t, err) {
		lib.Set("unknown", 1)
		assert.True(t, testTime.Equal(lib.ToTime()))
	}
}

func TestSetNanosecond(t *testing.T) {
	testTime := time.Date(2011, 10, 11, 15, 20, 1, 10000, chicagoLocation())

	lib, err := New(testTime)
	if assert.NoError(t, err) {
		assert.Equal(t, lib.Nanosecond(), 10000)
		lib.SetNanosecond(300)
		assert.Equal(t, lib.Nanosecond(), 300)
	}
}

func TestSetNanosecondOutOfRange(t *testing.T) {
	testTime := time.Date(2011, 10, 11, 15, 20, 1, 600000, chicagoLocation())

	lib, err := New(testTime)
	if assert.NoError(t, err) {
		assert.Equal(t, lib.Nanosecond(), 600000)
		lib.SetNanosecond(-1)
		assert.Equal(t, lib.Nanosecond(), 600000)
		lib.SetNanosecond(9999999991)
		assert.Equal(t, lib.Nanosecond(), 600000)
	}
}

func TestSetMillisecond(t *testing.T) {
	testTime := time.Date(2011, 10, 11, 15, 20, 5, 0, chicagoLocation())

	lib, err := New(testTime)
	if assert.NoError(t, err) {
		assert.Equal(t, lib.Millisecond(), 5000)
		assert.Equal(t, lib.Second(), 5)
		lib.SetMillisecond(30000)
		assert.Equal(t, lib.Millisecond(), 30000)
		assert.Equal(t, lib.Second(), 30)
	}
}

func TestSetMillisecondOutOfRange(t *testing.T) {
	testTime := time.Date(2011, 10, 11, 15, 20, 55, 0, chicagoLocation())

	lib, err := New(testTime)
	if assert.NoError(t, err) {
		assert.Equal(t, lib.Millisecond(), 55000)
		assert.Equal(t, lib.Second(), 55)
		lib.SetMillisecond(-1)
		assert.Equal(t, lib.Millisecond(), 55000)
		assert.Equal(t, lib.Second(), 55)
		lib.SetMillisecond(600001)
		assert.Equal(t, lib.Millisecond(), 55000)
		assert.Equal(t, lib.Second(), 55)
	}
}

func TestSetSecond(t *testing.T) {
	testTime := time.Date(2011, 10, 11, 15, 20, 25, 0, chicagoLocation())

	lib, err := New(testTime)
	if assert.NoError(t, err) {
		assert.Equal(t, lib.Second(), 25)
		lib.SetSecond(42)
		assert.Equal(t, lib.Second(), 42)
	}
}

func TestSetSecondOutOfRange(t *testing.T) {
	testTime := time.Date(2011, 10, 11, 15, 20, 41, 0, chicagoLocation())

	lib, err := New(testTime)
	if assert.NoError(t, err) {
		assert.Equal(t, lib.Second(), 41)
		lib.SetSecond(-1)
		assert.Equal(t, lib.Second(), 41)
		lib.SetSecond(60)
		assert.Equal(t, lib.Second(), 41)
	}
}

func TestSetMinute(t *testing.T) {
	testTime := time.Date(2011, 10, 11, 15, 20, 0, 0, chicagoLocation())

	lib, err := New(testTime)
	if assert.NoError(t, err) {
		assert.Equal(t, lib.Minute(), 20)
		lib.SetMinute(11)
		assert.Equal(t, lib.Minute(), 11)
	}
}

func TestSetMinuteOutOfRange(t *testing.T) {
	testTime := time.Date(2011, 10, 11, 15, 20, 0, 0, chicagoLocation())

	lib, err := New(testTime)
	if assert.NoError(t, err) {
		assert.Equal(t, lib.Minute(), 20)
		lib.SetMinute(-1)
		assert.Equal(t, lib.Minute(), 20)
		lib.SetMinute(60)
		assert.Equal(t, lib.Minute(), 20)
	}
}

func TestSetHour(t *testing.T) {
	testTime := time.Date(2011, 10, 11, 15, 0, 0, 0, chicagoLocation())

	lib, err := New(testTime)
	if assert.NoError(t, err) {
		assert.Equal(t, lib.Hour(), 15)
		lib.SetHour(3)
		assert.Equal(t, lib.Hour(), 3)
	}
}

func TestSetHourOutOfRange(t *testing.T) {
	testTime := time.Date(2011, 10, 11, 23, 0, 0, 0, chicagoLocation())

	lib, err := New(testTime)
	if assert.NoError(t, err) {
		assert.Equal(t, lib.Hour(), 23)
		lib.SetHour(-1)
		assert.Equal(t, lib.Hour(), 23)
		lib.SetHour(24)
		assert.Equal(t, lib.Hour(), 23)
	}
}

func TestSetDate(t *testing.T) {
	lib, err := New("2001-02-13")
	if assert.NoError(t, err) {
		assert.Equal(t, lib.Date(), 13)
		lib.SetDate(25)
		assert.Equal(t, lib.Date(), 25)
	}

	lib, err = New("2001-06-25")
	if assert.NoError(t, err) {
		assert.Equal(t, lib.Date(), 25)
		lib.SetDate(31)
		assert.Equal(t, lib.Date(), 30)
	}

	lib, err = New("2001-01-10")
	if assert.NoError(t, err) {
		assert.Equal(t, lib.Date(), 10)
		lib.SetDate(31)
		assert.Equal(t, lib.Date(), 31)
	}
}

func TestSetDateOutOfRange(t *testing.T) {
	lib, err := New("2016-09-04")
	if assert.NoError(t, err) {
		lib.SetDate(-1)
		assert.Equal(t, lib.Date(), 4)
		lib.SetDate(35)
		assert.Equal(t, lib.Date(), 4)
	}
}

func TestSetDatePastEndOfMonth(t *testing.T) {
	lib, err := New("2001-02-13")
	if assert.NoError(t, err) {
		lib.SetDate(29)
		assert.Equal(t, lib.Date(), 28)
	}

	lib, err = New("2001-06-25")
	if assert.NoError(t, err) {
		lib.SetDate(31)
		assert.Equal(t, lib.Date(), 30)
	}
}

func TestSetDateHandlesLeapyear(t *testing.T) {
	lib, err := New("2000-02-13")
	if assert.NoError(t, err) {
		lib.SetDate(29)
		assert.Equal(t, lib.Date(), 29)
	}
}

func TestSetDay(t *testing.T) {
	lib, err := New("2016-09-04")
	if assert.NoError(t, err) {
		assert.Equal(t, lib.Day(), 0) // Sunday
		lib.SetDay(3)                 // Wednesday
		assert.Equal(t, lib.Date(), 7)
		lib.SetDay(6) // Saturday
		assert.Equal(t, lib.Date(), 10)
	}
}

func TestSetDayOutOfRange(t *testing.T) {
	lib, err := New("2016-09-04")
	if assert.NoError(t, err) {
		lib.SetDay(-1)
		assert.Equal(t, lib.Date(), 4)
		lib.SetDay(7)
		assert.Equal(t, lib.Date(), 4)
	}
}

func TestSetISOWeekday(t *testing.T) {
	lib, err := New("2016-09-05")
	if assert.NoError(t, err) {
		assert.Equal(t, lib.ISOWeekday(), 1) // Monday
		lib.SetISOWeekday(3)                 // Wednesday
		assert.Equal(t, lib.Date(), 7)
		lib.SetISOWeekday(7) // Sunday
		assert.Equal(t, lib.Date(), 4)
	}
}

func TestSetISOWeekdayOutOfRange(t *testing.T) {
	lib, err := New("2016-09-04")
	if assert.NoError(t, err) {
		lib.SetISOWeekday(-1)
		assert.Equal(t, lib.Date(), 4)
		lib.SetISOWeekday(8)
		assert.Equal(t, lib.Date(), 4)
	}
}

func TestSetDayOfYear(t *testing.T) {
	lib, err := New("2000-01-01") // Leap year
	if assert.NoError(t, err) {
		assert.Equal(t, lib.DayOfYear(), 1)
		lib.SetDayOfYear(100)
		assert.Equal(t, lib.DayOfYear(), 100)
		lib.SetDayOfYear(366)
		assert.Equal(t, lib.DayOfYear(), 366)
	}

	lib, err = New("2001-01-01") // Non-leap year
	if assert.NoError(t, err) {
		assert.Equal(t, lib.DayOfYear(), 1)
		lib.SetDayOfYear(100)
		assert.Equal(t, lib.DayOfYear(), 100)
		lib.SetDayOfYear(366)
		assert.Equal(t, lib.DayOfYear(), 365)
	}
}

func TestSetDayOfYearOutOfRange(t *testing.T) {
	lib, err := New("2016-09-04")
	if assert.NoError(t, err) {
		lib.SetDayOfYear(0)
		assert.Equal(t, lib.DayOfYear(), 248)
		lib.SetDayOfYear(368)
		assert.Equal(t, lib.DayOfYear(), 248)
	}
}

func TestSetMonth(t *testing.T) {
	lib, err := New("2004-04-15")
	if assert.NoError(t, err) {
		lib.SetMonth(6)
		assert.Equal(t, lib.Month(), 6)
		assert.Equal(t, lib.Date(), 15)
	}
}

func TestSetMonthOutOfRange(t *testing.T) {
	lib, err := New("2004-04-15")
	if assert.NoError(t, err) {
		lib.SetMonth(0)
		assert.Equal(t, lib.Month(), 4)
		lib.SetMonth(15)
		assert.Equal(t, lib.Month(), 4)
	}
}

func TestSetMonthWithLessDays(t *testing.T) {
	lib, err := New("2001-01-31")
	if assert.NoError(t, err) {
		assert.Equal(t, lib.Date(), 31)
		lib.SetMonth(2)
		assert.Equal(t, lib.Month(), 2)
		assert.Equal(t, lib.Date(), 28)
	}
}

func TestSetQuarter(t *testing.T) {
	lib, err := New("2001-02-01")
	if assert.NoError(t, err) {
		lib.SetQuarter(1)
		assert.Equal(t, lib.Month(), 2)
		assert.Equal(t, lib.Quarter(), 1)
		lib.SetQuarter(2)
		assert.Equal(t, lib.Month(), 5)
		assert.Equal(t, lib.Quarter(), 2)
		lib.SetQuarter(3)
		assert.Equal(t, lib.Month(), 8)
		assert.Equal(t, lib.Quarter(), 3)
		lib.SetQuarter(4)
		assert.Equal(t, lib.Month(), 11)
		assert.Equal(t, lib.Quarter(), 4)
	}

	lib, err = New("2001-12-01")
	if assert.NoError(t, err) {
		lib.SetQuarter(3)
		assert.Equal(t, lib.Month(), 9)
		assert.Equal(t, lib.Quarter(), 3)
	}
}

func TestSetQuarterOutOfRange(t *testing.T) {
	lib, err := New("2001-12-01")
	if assert.NoError(t, err) {
		lib.SetQuarter(0)
		assert.Equal(t, lib.Month(), 12)
		assert.Equal(t, lib.Quarter(), 4)
		lib.SetQuarter(6)
		assert.Equal(t, lib.Month(), 12)
		assert.Equal(t, lib.Quarter(), 4)
	}
}

func TestSetYear(t *testing.T) {
	lib, err := New("1985-05-05")
	if assert.NoError(t, err) {
		lib.SetYear(2015)
		assert.Equal(t, lib.Year(), 2015)
	}
}
