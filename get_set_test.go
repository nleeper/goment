package goment

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetByUnits(t *testing.T) {
	assert := assert.New(t)

	lib := simpleString("2015-04-06 10:11:12.4567")
	assert.Equal(2015, lib.Get("y"))
	assert.Equal(2015, lib.Get("year"))
	assert.Equal(2015, lib.Get("years"))
	assert.Equal(4, lib.Get("M"))
	assert.Equal(4, lib.Get("month"))
	assert.Equal(4, lib.Get("months"))
	assert.Equal(6, lib.Get("D"))
	assert.Equal(6, lib.Get("date"))
	assert.Equal(6, lib.Get("dates"))
	assert.Equal(10, lib.Get("h"))
	assert.Equal(10, lib.Get("hour"))
	assert.Equal(10, lib.Get("hours"))
	assert.Equal(11, lib.Get("m"))
	assert.Equal(11, lib.Get("minute"))
	assert.Equal(11, lib.Get("minutes"))
	assert.Equal(12, lib.Get("s"))
	assert.Equal(12, lib.Get("second"))
	assert.Equal(12, lib.Get("seconds"))
	assert.Equal(12000, lib.Get("ms"))
	assert.Equal(12000, lib.Get("millisecond"))
	assert.Equal(12000, lib.Get("milliseconds"))
	assert.Equal(456700000, lib.Get("ns"))
	assert.Equal(456700000, lib.Get("nanosecond"))
	assert.Equal(456700000, lib.Get("nanoseconds"))
	assert.Equal(0, lib.Get("unknown"))
}

func TestGetNanosecond(t *testing.T) {
	testTime := time.Date(2017, 8, 2, 18, 0, 33, 5000, time.UTC)
	assert.Equal(t, 5000, simpleTime(testTime).Nanosecond())
}

func TestGetMillisecond(t *testing.T) {
	testTime := time.Date(2017, 8, 2, 18, 0, 33, 0, time.UTC)
	assert.Equal(t, 33000, simpleTime(testTime).Millisecond())
}

func TestGetSecond(t *testing.T) {
	testTime := time.Date(2017, 8, 2, 18, 0, 33, 0, time.UTC)
	assert.Equal(t, 33, simpleTime(testTime).Second())
}

func TestGetMinute(t *testing.T) {
	testTime := time.Date(2017, 8, 2, 18, 44, 33, 0, time.UTC)
	assert.Equal(t, 44, simpleTime(testTime).Minute())
}

func TestGetHour(t *testing.T) {
	testTime := time.Date(2017, 8, 2, 18, 0, 33, 0, chicagoLocation())
	assert.Equal(t, 18, simpleTime(testTime).Hour())
}

func TestGetDate(t *testing.T) {
	testTime := time.Date(2017, 8, 2, 18, 0, 0, 0, time.UTC)
	assert.Equal(t, 2, simpleTime(testTime).Date())
}

func TestGetDay(t *testing.T) {
	testTime := time.Date(2017, 8, 2, 18, 0, 0, 0, time.UTC)
	assert.Equal(t, 3, simpleTime(testTime).Day())
}

func TestGetWeekday(t *testing.T) {
	testTime := time.Date(2020, 7, 27, 18, 0, 0, 0, time.UTC) // Monday, July 27 2020
	assert.Equal(t, 1, simpleTime(testTime).Weekday())
}

func TestGetISOWeekday(t *testing.T) {
	assert := assert.New(t)

	testTime := time.Date(2017, 8, 6, 18, 0, 0, 0, time.UTC)

	lib := simpleTime(testTime)
	assert.Equal(0, lib.Day())
	assert.Equal(7, lib.ISOWeekday())
}

func TestDayOfYear(t *testing.T) {
	testTime := time.Date(2017, 1, 3, 18, 0, 0, 0, time.UTC)
	assert.Equal(t, 3, simpleTime(testTime).DayOfYear())
}

func TestGetWeek(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(1, simple(DateTime{Year: 2012, Month: 1, Day: 1}).Week(), "Jan 1 2012 should be week 1")
	assert.Equal(1, simple(DateTime{Year: 2012, Month: 1, Day: 7}).Week(), "Jan 7 2012 should be week 1")
	assert.Equal(2, simple(DateTime{Year: 2012, Month: 1, Day: 8}).Week(), "Jan 8 2012 should be week 2")
	assert.Equal(2, simple(DateTime{Year: 2012, Month: 1, Day: 14}).Week(), "Jan 14 2012 should be week 2")
	assert.Equal(3, simple(DateTime{Year: 2012, Month: 1, Day: 15}).Week(), "Jan 15 2012 should be week 3")
	assert.Equal(1, simple(DateTime{Year: 2006, Month: 12, Day: 31}).Week(), "Dec 31 2006 should be week 1")
	assert.Equal(1, simple(DateTime{Year: 2007, Month: 1, Day: 1}).Week(), "Jan 1 2007 should be week 1")
	assert.Equal(1, simple(DateTime{Year: 2007, Month: 1, Day: 6}).Week(), "Jan 6 2007 should be week 1")
	assert.Equal(2, simple(DateTime{Year: 2007, Month: 1, Day: 7}).Week(), "Jan 7 2007 should be week 2")
	assert.Equal(2, simple(DateTime{Year: 2007, Month: 1, Day: 13}).Week(), "Jan 13 2007 should be week 2")
	assert.Equal(3, simple(DateTime{Year: 2007, Month: 1, Day: 14}).Week(), "Jan 14 2007 should be week 3")
	assert.Equal(52, simple(DateTime{Year: 2007, Month: 12, Day: 29}).Week(), "Dec 29 2007 should be week 52")
	assert.Equal(1, simple(DateTime{Year: 2008, Month: 1, Day: 1}).Week(), "Jan 1 2008 should be week 1")
	assert.Equal(1, simple(DateTime{Year: 2008, Month: 1, Day: 5}).Week(), "Jan 5 2008 should be week 1")
	assert.Equal(2, simple(DateTime{Year: 2008, Month: 1, Day: 6}).Week(), "Jan 6 2008 should be week 2")
	assert.Equal(2, simple(DateTime{Year: 2008, Month: 1, Day: 12}).Week(), "Jan 12 2008 should be week 2")
	assert.Equal(3, simple(DateTime{Year: 2008, Month: 1, Day: 13}).Week(), "Jan 13 2008 should be week 3")
	assert.Equal(1, simple(DateTime{Year: 2002, Month: 12, Day: 29}).Week(), "Dec 29 2002 should be week 1")
	assert.Equal(1, simple(DateTime{Year: 2003, Month: 1, Day: 1}).Week(), "Jan 1 2003 should be week 1")
	assert.Equal(1, simple(DateTime{Year: 2003, Month: 1, Day: 4}).Week(), "Jan 4 2003 should be week 1")
	assert.Equal(2, simple(DateTime{Year: 2003, Month: 1, Day: 5}).Week(), "Jan 5 2003 should be week 2")
	assert.Equal(2, simple(DateTime{Year: 2003, Month: 1, Day: 11}).Week(), "Jan 11 2003 should be week 2")
	assert.Equal(3, simple(DateTime{Year: 2003, Month: 1, Day: 12}).Week(), "Jan 12 2003 should be week 3")
	assert.Equal(1, simple(DateTime{Year: 2008, Month: 12, Day: 28}).Week(), "Dec 28 2008 should be week 1")
	assert.Equal(1, simple(DateTime{Year: 2009, Month: 1, Day: 1}).Week(), "Jan 1 2009 should be week 1")
	assert.Equal(1, simple(DateTime{Year: 2009, Month: 1, Day: 3}).Week(), "Jan 3 2009 should be week 1")
	assert.Equal(2, simple(DateTime{Year: 2009, Month: 1, Day: 4}).Week(), "Jan 4 2009 should be week 2")
	assert.Equal(2, simple(DateTime{Year: 2009, Month: 1, Day: 10}).Week(), "Jan 10 2009 should be week 2")
	assert.Equal(3, simple(DateTime{Year: 2009, Month: 1, Day: 11}).Week(), "Jan 11 2009 should be week 3")
	assert.Equal(1, simple(DateTime{Year: 2009, Month: 12, Day: 27}).Week(), "Dec 27 2009 should be week 1")
	assert.Equal(1, simple(DateTime{Year: 2010, Month: 1, Day: 1}).Week(), "Jan 1 2010 should be week 1")
	assert.Equal(1, simple(DateTime{Year: 2010, Month: 1, Day: 2}).Week(), "Jan 2 2010 should be week 1")
	assert.Equal(2, simple(DateTime{Year: 2010, Month: 1, Day: 3}).Week(), "Jan 3 2010 should be week 2")
	assert.Equal(2, simple(DateTime{Year: 2010, Month: 1, Day: 9}).Week(), "Jan 9 2010 should be week 2")
	assert.Equal(3, simple(DateTime{Year: 2010, Month: 1, Day: 10}).Week(), "Jan 10 2010 should be week 3")
	assert.Equal(1, simple(DateTime{Year: 2010, Month: 12, Day: 26}).Week(), "Dec 26 2010 should be week 1")
	assert.Equal(1, simple(DateTime{Year: 2011, Month: 1, Day: 1}).Week(), "Jan 1 2011 should be week 1")
	assert.Equal(2, simple(DateTime{Year: 2011, Month: 1, Day: 2}).Week(), "Jan 2 2011 should be week 2")
	assert.Equal(2, simple(DateTime{Year: 2011, Month: 1, Day: 8}).Week(), "Jan 8 2011 should be week 2")
	assert.Equal(3, simple(DateTime{Year: 2011, Month: 1, Day: 9}).Week(), "Jan 9 2011 should be week 3")
}

func TestGetISOWeek(t *testing.T) {
	assert := assert.New(t)

	testTime := time.Date(2005, 1, 1, 0, 0, 0, 0, chicagoLocation())
	testTime2 := time.Date(2007, 12, 30, 0, 0, 0, 0, chicagoLocation())

	assert.Equal(53, simpleTime(testTime).ISOWeek())
	assert.Equal(52, simpleTime(testTime2).ISOWeek())
}

func TestGetISOWeekYearsWith53Weeks(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(53, simple(DateTime{Year: 2004, Month: 12, Day: 31}).ISOWeek(), "Dec 31 2004 should be iso week 53")
	assert.Equal(53, simple(DateTime{Year: 2009, Month: 12, Day: 31}).ISOWeek(), "Dec 31 2009 should be iso week 53")
	assert.Equal(53, simple(DateTime{Year: 2015, Month: 12, Day: 31}).ISOWeek(), "Dec 31 2015 should be iso week 53")
	assert.Equal(53, simple(DateTime{Year: 2020, Month: 12, Day: 31}).ISOWeek(), "Dec 31 2020 should be iso week 53")
	assert.Equal(53, simple(DateTime{Year: 2026, Month: 12, Day: 31}).ISOWeek(), "Dec 31 2026 should be iso week 53")
	assert.Equal(53, simple(DateTime{Year: 2032, Month: 12, Day: 31}).ISOWeek(), "Dec 31 2032 should be iso week 53")
	assert.Equal(53, simple(DateTime{Year: 2037, Month: 12, Day: 31}).ISOWeek(), "Dec 31 2037 should be iso week 53")
	assert.Equal(53, simple(DateTime{Year: 2043, Month: 12, Day: 31}).ISOWeek(), "Dec 31 2043 should be iso week 53")
	assert.Equal(53, simple(DateTime{Year: 2048, Month: 12, Day: 31}).ISOWeek(), "Dec 31 2048 should be iso week 53")
	assert.Equal(53, simple(DateTime{Year: 2054, Month: 12, Day: 31}).ISOWeek(), "Dec 31 2054 should be iso week 53")
	assert.Equal(53, simple(DateTime{Year: 2060, Month: 12, Day: 31}).ISOWeek(), "Dec 31 2060 should be iso week 53")
	assert.Equal(53, simple(DateTime{Year: 2065, Month: 12, Day: 31}).ISOWeek(), "Dec 31 2065 should be iso week 53")
	assert.Equal(53, simple(DateTime{Year: 2071, Month: 12, Day: 31}).ISOWeek(), "Dec 31 2071 should be iso week 53")
	assert.Equal(53, simple(DateTime{Year: 2076, Month: 12, Day: 31}).ISOWeek(), "Dec 31 2076 should be iso week 53")
	assert.Equal(53, simple(DateTime{Year: 2082, Month: 12, Day: 31}).ISOWeek(), "Dec 31 2082 should be iso week 53")
	assert.Equal(53, simple(DateTime{Year: 2088, Month: 12, Day: 31}).ISOWeek(), "Dec 31 2088 should be iso week 53")
	assert.Equal(53, simple(DateTime{Year: 2093, Month: 12, Day: 31}).ISOWeek(), "Dec 31 2093 should be iso week 53")
	assert.Equal(53, simple(DateTime{Year: 2099, Month: 12, Day: 31}).ISOWeek(), "Dec 31 2099 should be iso week 53")
	assert.Equal(53, simple(DateTime{Year: 2105, Month: 12, Day: 31}).ISOWeek(), "Dec 31 2105 should be iso week 53")
	assert.Equal(53, simple(DateTime{Year: 2111, Month: 12, Day: 31}).ISOWeek(), "Dec 31 2111 should be iso week 53")
	assert.Equal(53, simple(DateTime{Year: 2116, Month: 12, Day: 31}).ISOWeek(), "Dec 31 2116 should be iso week 53")
	assert.Equal(53, simple(DateTime{Year: 2122, Month: 12, Day: 31}).ISOWeek(), "Dec 31 2122 should be iso week 53")
	assert.Equal(53, simple(DateTime{Year: 2128, Month: 12, Day: 31}).ISOWeek(), "Dec 31 2128 should be iso week 53")
	assert.Equal(53, simple(DateTime{Year: 2133, Month: 12, Day: 31}).ISOWeek(), "Dec 31 2133 should be iso week 53")
	assert.Equal(53, simple(DateTime{Year: 2139, Month: 12, Day: 31}).ISOWeek(), "Dec 31 2139 should be iso week 53")
	assert.Equal(53, simple(DateTime{Year: 2144, Month: 12, Day: 31}).ISOWeek(), "Dec 31 2144 should be iso week 53")
	assert.Equal(53, simple(DateTime{Year: 2150, Month: 12, Day: 31}).ISOWeek(), "Dec 31 2150 should be iso week 53")
	assert.Equal(53, simple(DateTime{Year: 2156, Month: 12, Day: 31}).ISOWeek(), "Dec 31 2156 should be iso week 53")
	assert.Equal(53, simple(DateTime{Year: 2161, Month: 12, Day: 31}).ISOWeek(), "Dec 31 2161 should be iso week 53")
	assert.Equal(53, simple(DateTime{Year: 2167, Month: 12, Day: 31}).ISOWeek(), "Dec 31 2167 should be iso week 53")
	assert.Equal(53, simple(DateTime{Year: 2172, Month: 12, Day: 31}).ISOWeek(), "Dec 31 2172 should be iso week 53")
	assert.Equal(53, simple(DateTime{Year: 2178, Month: 12, Day: 31}).ISOWeek(), "Dec 31 2178 should be iso week 53")
	assert.Equal(53, simple(DateTime{Year: 2184, Month: 12, Day: 31}).ISOWeek(), "Dec 31 2184 should be iso week 53")
	assert.Equal(53, simple(DateTime{Year: 2189, Month: 12, Day: 31}).ISOWeek(), "Dec 31 2189 should be iso week 53")
	assert.Equal(53, simple(DateTime{Year: 2195, Month: 12, Day: 31}).ISOWeek(), "Dec 31 2195 should be iso week 53")
	assert.Equal(53, simple(DateTime{Year: 2201, Month: 12, Day: 31}).ISOWeek(), "Dec 31 2201 should be iso week 53")
	assert.Equal(53, simple(DateTime{Year: 2207, Month: 12, Day: 31}).ISOWeek(), "Dec 31 2207 should be iso week 53")
	assert.Equal(53, simple(DateTime{Year: 2212, Month: 12, Day: 31}).ISOWeek(), "Dec 31 2212 should be iso week 53")
	assert.Equal(53, simple(DateTime{Year: 2218, Month: 12, Day: 31}).ISOWeek(), "Dec 31 2218 should be iso week 53")
	assert.Equal(53, simple(DateTime{Year: 2224, Month: 12, Day: 31}).ISOWeek(), "Dec 31 2224 should be iso week 53")
	assert.Equal(53, simple(DateTime{Year: 2229, Month: 12, Day: 31}).ISOWeek(), "Dec 31 2229 should be iso week 53")
	assert.Equal(53, simple(DateTime{Year: 2235, Month: 12, Day: 31}).ISOWeek(), "Dec 31 2235 should be iso week 53")
	assert.Equal(53, simple(DateTime{Year: 2240, Month: 12, Day: 31}).ISOWeek(), "Dec 31 2240 should be iso week 53")
	assert.Equal(53, simple(DateTime{Year: 2246, Month: 12, Day: 31}).ISOWeek(), "Dec 31 2246 should be iso week 53")
	assert.Equal(53, simple(DateTime{Year: 2252, Month: 12, Day: 31}).ISOWeek(), "Dec 31 2252 should be iso week 53")
	assert.Equal(53, simple(DateTime{Year: 2257, Month: 12, Day: 31}).ISOWeek(), "Dec 31 2257 should be iso week 53")
	assert.Equal(53, simple(DateTime{Year: 2263, Month: 12, Day: 31}).ISOWeek(), "Dec 31 2263 should be iso week 53")
	assert.Equal(53, simple(DateTime{Year: 2268, Month: 12, Day: 31}).ISOWeek(), "Dec 31 2268 should be iso week 53")
	assert.Equal(53, simple(DateTime{Year: 2274, Month: 12, Day: 31}).ISOWeek(), "Dec 31 2274 should be iso week 53")
	assert.Equal(53, simple(DateTime{Year: 2280, Month: 12, Day: 31}).ISOWeek(), "Dec 31 2280 should be iso week 53")
	assert.Equal(53, simple(DateTime{Year: 2285, Month: 12, Day: 31}).ISOWeek(), "Dec 31 2285 should be iso week 53")
	assert.Equal(53, simple(DateTime{Year: 2291, Month: 12, Day: 31}).ISOWeek(), "Dec 31 2291 should be iso week 53")
	assert.Equal(53, simple(DateTime{Year: 2296, Month: 12, Day: 31}).ISOWeek(), "Dec 31 2296 should be iso week 53")
	assert.Equal(53, simple(DateTime{Year: 2303, Month: 12, Day: 31}).ISOWeek(), "Dec 31 2303 should be iso week 53")
	assert.Equal(53, simple(DateTime{Year: 2308, Month: 12, Day: 31}).ISOWeek(), "Dec 31 2308 should be iso week 53")
	assert.Equal(53, simple(DateTime{Year: 2314, Month: 12, Day: 31}).ISOWeek(), "Dec 31 2314 should be iso week 53")
	assert.Equal(53, simple(DateTime{Year: 2320, Month: 12, Day: 31}).ISOWeek(), "Dec 31 2320 should be iso week 53")
	assert.Equal(53, simple(DateTime{Year: 2325, Month: 12, Day: 31}).ISOWeek(), "Dec 31 2325 should be iso week 53")
	assert.Equal(53, simple(DateTime{Year: 2331, Month: 12, Day: 31}).ISOWeek(), "Dec 31 2331 should be iso week 53")
	assert.Equal(53, simple(DateTime{Year: 2336, Month: 12, Day: 31}).ISOWeek(), "Dec 31 2336 should be iso week 53")
	assert.Equal(53, simple(DateTime{Year: 2342, Month: 12, Day: 31}).ISOWeek(), "Dec 31 2342 should be iso week 53")
	assert.Equal(53, simple(DateTime{Year: 2348, Month: 12, Day: 31}).ISOWeek(), "Dec 31 2348 should be iso week 53")
	assert.Equal(53, simple(DateTime{Year: 2353, Month: 12, Day: 31}).ISOWeek(), "Dec 31 2353 should be iso week 53")
	assert.Equal(53, simple(DateTime{Year: 2359, Month: 12, Day: 31}).ISOWeek(), "Dec 31 2359 should be iso week 53")
	assert.Equal(53, simple(DateTime{Year: 2364, Month: 12, Day: 31}).ISOWeek(), "Dec 31 2364 should be iso week 53")
	assert.Equal(53, simple(DateTime{Year: 2370, Month: 12, Day: 31}).ISOWeek(), "Dec 31 2370 should be iso week 53")
	assert.Equal(53, simple(DateTime{Year: 2376, Month: 12, Day: 31}).ISOWeek(), "Dec 31 2376 should be iso week 53")
	assert.Equal(53, simple(DateTime{Year: 2381, Month: 12, Day: 31}).ISOWeek(), "Dec 31 2381 should be iso week 53")
	assert.Equal(53, simple(DateTime{Year: 2387, Month: 12, Day: 31}).ISOWeek(), "Dec 31 2387 should be iso week 53")
	assert.Equal(53, simple(DateTime{Year: 2392, Month: 12, Day: 31}).ISOWeek(), "Dec 31 2392 should be iso week 53")
	assert.Equal(53, simple(DateTime{Year: 2398, Month: 12, Day: 31}).ISOWeek(), "Dec 31 2398 should be iso week 53")
}

func TestGetMonth(t *testing.T) {
	testTime := time.Date(2017, 6, 3, 18, 0, 0, 0, time.UTC)
	assert.Equal(t, 6, simpleTime(testTime).Month())
}

func TestGetQuarter(t *testing.T) {
	assert := assert.New(t)

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
		assert.Equal(q, simpleTime(testTime).Quarter())
	}
}

func TestGetYear(t *testing.T) {
	testTime := time.Date(2017, 1, 3, 18, 0, 0, 0, time.UTC)
	assert.Equal(t, 2017, simpleTime(testTime).Year())
}

func TestGetWeekYear(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(2005, simple(DateTime{Year: 2005, Month: 1, Day: 1}).WeekYear())
	assert.Equal(2005, simple(DateTime{Year: 2005, Month: 1, Day: 2}).WeekYear())
	assert.Equal(2005, simple(DateTime{Year: 2005, Month: 1, Day: 3}).WeekYear())
	assert.Equal(2005, simple(DateTime{Year: 2005, Month: 12, Day: 31}).WeekYear())
	assert.Equal(2006, simple(DateTime{Year: 2006, Month: 1, Day: 1}).WeekYear())
	assert.Equal(2006, simple(DateTime{Year: 2006, Month: 1, Day: 2}).WeekYear())
	assert.Equal(2007, simple(DateTime{Year: 2007, Month: 1, Day: 1}).WeekYear())
	assert.Equal(2008, simple(DateTime{Year: 2007, Month: 12, Day: 30}).WeekYear())
	assert.Equal(2008, simple(DateTime{Year: 2007, Month: 12, Day: 31}).WeekYear())
	assert.Equal(2008, simple(DateTime{Year: 2008, Month: 1, Day: 1}).WeekYear())
	assert.Equal(2009, simple(DateTime{Year: 2008, Month: 12, Day: 28}).WeekYear())
	assert.Equal(2009, simple(DateTime{Year: 2008, Month: 12, Day: 29}).WeekYear())
	assert.Equal(2009, simple(DateTime{Year: 2008, Month: 12, Day: 30}).WeekYear())
	assert.Equal(2009, simple(DateTime{Year: 2008, Month: 12, Day: 31}).WeekYear())
	assert.Equal(2009, simple(DateTime{Year: 2009, Month: 1, Day: 1}).WeekYear())
	assert.Equal(2010, simple(DateTime{Year: 2010, Month: 1, Day: 1}).WeekYear())
	assert.Equal(2010, simple(DateTime{Year: 2010, Month: 1, Day: 2}).WeekYear())
	assert.Equal(2010, simple(DateTime{Year: 2010, Month: 1, Day: 3}).WeekYear())
	assert.Equal(2010, simple(DateTime{Year: 2010, Month: 1, Day: 4}).WeekYear())
}

func TestGetISOWeekYear(t *testing.T) {
	assert := assert.New(t)

	testTime := time.Date(2005, 1, 1, 0, 0, 0, 0, chicagoLocation())
	testTime2 := time.Date(2007, 12, 30, 0, 0, 0, 0, chicagoLocation())

	assert.Equal(2004, simpleTime(testTime).ISOWeekYear())
	assert.Equal(2007, simpleTime(testTime2).ISOWeekYear())
}

func TestGetWeeksInYear(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(52, simple(DateTime{Year: 2004}).WeeksInYear(), "2004 has 52 weeks")
	assert.Equal(53, simple(DateTime{Year: 2005}).WeeksInYear(), "2005 has 53 weeks")
	assert.Equal(52, simple(DateTime{Year: 2006}).WeeksInYear(), "2006 has 52 weeks")
	assert.Equal(52, simple(DateTime{Year: 2007}).WeeksInYear(), "2007 has 52 weeks")
	assert.Equal(52, simple(DateTime{Year: 2008}).WeeksInYear(), "2008 has 52 weeks")
	assert.Equal(52, simple(DateTime{Year: 2009}).WeeksInYear(), "2009 has 52 weeks")
	assert.Equal(52, simple(DateTime{Year: 2010}).WeeksInYear(), "2010 has 52 weeks")
	assert.Equal(53, simple(DateTime{Year: 2011}).WeeksInYear(), "2011 has 53 weeks")
	assert.Equal(52, simple(DateTime{Year: 2012}).WeeksInYear(), "2012 has 52 weeks")
	assert.Equal(52, simple(DateTime{Year: 2013}).WeeksInYear(), "2013 has 52 weeks")
	assert.Equal(52, simple(DateTime{Year: 2014}).WeeksInYear(), "2014 has 52 weeks")
	assert.Equal(52, simple(DateTime{Year: 2015}).WeeksInYear(), "2015 has 52 weeks")
}

func TestGetISOWeeksInYear(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(52, simple(DateTime{Year: 2005}).ISOWeeksInYear(), "ISO year 2005 has 52 iso weeks")
	assert.Equal(52, simple(DateTime{Year: 2006}).ISOWeeksInYear(), "ISO year 2006 has 52 iso weeks")
	assert.Equal(53, simple(DateTime{Year: 2009}).ISOWeeksInYear(), "ISO year 2009 has 53 iso weeks")
	assert.Equal(52, simple(DateTime{Year: 2010}).ISOWeeksInYear(), "ISO year 2010 has 52 iso weeks")
}

func TestSetByUnits(t *testing.T) {
	assert := assert.New(t)

	lib := simpleString("2015-04-06 10:11:12.4567")

	assert.Equal(2016, lib.Set("y", 2016).Year())
	assert.Equal(2017, lib.Set("year", 2017).Year())
	assert.Equal(2018, lib.Set("years", 2018).Year())
	assert.Equal(9, lib.Set("M", 9).Month())
	assert.Equal(10, lib.Set("month", 10).Month())
	assert.Equal(11, lib.Set("months", 11).Month())
	assert.Equal(8, lib.Set("D", 8).Date())
	assert.Equal(9, lib.Set("D", 9).Date())
	assert.Equal(10, lib.Set("D", 10).Date())
	assert.Equal(14, lib.Set("h", 14).Hour())
	assert.Equal(15, lib.Set("hour", 15).Hour())
	assert.Equal(16, lib.Set("hours", 16).Hour())
	assert.Equal(17, lib.Set("m", 17).Minute())
	assert.Equal(18, lib.Set("minute", 18).Minute())
	assert.Equal(19, lib.Set("minutes", 19).Minute())
	assert.Equal(20, lib.Set("s", 20).Second())
	assert.Equal(21, lib.Set("second", 21).Second())
	assert.Equal(22, lib.Set("seconds", 22).Second())
	assert.Equal(23000, lib.Set("ms", 23000).Millisecond())
	assert.Equal(24000, lib.Set("millisecond", 24000).Millisecond())
	assert.Equal(25000, lib.Set("milliseconds", 25000).Millisecond())
	assert.Equal(100000, lib.Set("ns", 100000).Nanosecond())
	assert.Equal(100001, lib.Set("nanosecond", 100001).Nanosecond())
	assert.Equal(100002, lib.Set("nanoseconds", 100002).Nanosecond())
}

func TestSetUnknownUnits(t *testing.T) {
	testTime := time.Date(2011, 10, 11, 15, 20, 1, 10000, chicagoLocation())

	lib := simpleTime(testTime)
	lib.Set("unknown", 1)
	assert.Equal(t, simpleTime(testTime).Format(), lib.Format())
}

func TestSetNanosecond(t *testing.T) {
	assert := assert.New(t)

	testTime := time.Date(2011, 10, 11, 15, 20, 1, 10000, chicagoLocation())

	lib := simpleTime(testTime)
	assert.Equal(10000, lib.Nanosecond())

	lib.SetNanosecond(300)
	assert.Equal(300, lib.Nanosecond())
}

func TestSetNanosecondOutOfRange(t *testing.T) {
	assert := assert.New(t)

	testTime := time.Date(2011, 10, 11, 15, 20, 1, 600000, chicagoLocation())

	lib := simpleTime(testTime)
	assert.Equal(600000, lib.Nanosecond())

	lib.SetNanosecond(-1)
	assert.Equal(600000, lib.Nanosecond())

	lib.SetNanosecond(9999999991)
	assert.Equal(600000, lib.Nanosecond())
}

func TestSetMillisecond(t *testing.T) {
	assert := assert.New(t)

	testTime := time.Date(2011, 10, 11, 15, 20, 5, 0, chicagoLocation())

	lib := simpleTime(testTime)
	assert.Equal(5000, lib.Millisecond())
	assert.Equal(5, lib.Second())

	lib.SetMillisecond(30000)
	assert.Equal(30000, lib.Millisecond())
	assert.Equal(30, lib.Second())
}

func TestSetMillisecondOutOfRange(t *testing.T) {
	assert := assert.New(t)

	testTime := time.Date(2011, 10, 11, 15, 20, 55, 0, chicagoLocation())

	lib := simpleTime(testTime)
	assert.Equal(55000, lib.Millisecond())
	assert.Equal(55, lib.Second())

	lib.SetMillisecond(-1)
	assert.Equal(55000, lib.Millisecond())
	assert.Equal(55, lib.Second())

	lib.SetMillisecond(600001)
	assert.Equal(55000, lib.Millisecond())
	assert.Equal(55, lib.Second())
}

func TestSetSecond(t *testing.T) {
	assert := assert.New(t)

	testTime := time.Date(2011, 10, 11, 15, 20, 25, 0, chicagoLocation())

	lib := simpleTime(testTime)
	assert.Equal(25, lib.Second())

	lib.SetSecond(42)
	assert.Equal(42, lib.Second())
}

func TestSetSecondOutOfRange(t *testing.T) {
	assert := assert.New(t)

	testTime := time.Date(2011, 10, 11, 15, 20, 41, 0, chicagoLocation())

	lib := simpleTime(testTime)
	assert.Equal(41, lib.Second())

	lib.SetSecond(-1)
	assert.Equal(41, lib.Second())

	lib.SetSecond(60)
	assert.Equal(41, lib.Second())
}

func TestSetMinute(t *testing.T) {
	assert := assert.New(t)

	testTime := time.Date(2011, 10, 11, 15, 20, 0, 0, chicagoLocation())

	lib := simpleTime(testTime)
	assert.Equal(20, lib.Minute())

	lib.SetMinute(11)
	assert.Equal(11, lib.Minute())
}

func TestSetMinuteOutOfRange(t *testing.T) {
	assert := assert.New(t)

	testTime := time.Date(2011, 10, 11, 15, 20, 0, 0, chicagoLocation())

	lib := simpleTime(testTime)
	assert.Equal(20, lib.Minute())

	lib.SetMinute(-1)
	assert.Equal(20, lib.Minute())

	lib.SetMinute(60)
	assert.Equal(20, lib.Minute())
}

func TestSetHour(t *testing.T) {
	assert := assert.New(t)

	testTime := time.Date(2011, 10, 11, 15, 0, 0, 0, chicagoLocation())

	lib := simpleTime(testTime)
	assert.Equal(15, lib.Hour())

	lib.SetHour(3)
	assert.Equal(3, lib.Hour())
}

func TestSetHourOutOfRange(t *testing.T) {
	assert := assert.New(t)

	testTime := time.Date(2011, 10, 11, 23, 0, 0, 0, chicagoLocation())

	lib := simpleTime(testTime)
	assert.Equal(23, lib.Hour())

	lib.SetHour(-1)
	assert.Equal(23, lib.Hour())

	lib.SetHour(24)
	assert.Equal(23, lib.Hour())
}

func TestSetDate(t *testing.T) {
	assert := assert.New(t)

	lib := simpleString("2001-02-13")
	assert.Equal(13, lib.Date())

	lib.SetDate(25)
	assert.Equal(25, lib.Date())

	lib2 := simpleString("2001-06-25")
	assert.Equal(25, lib2.Date())

	lib2.SetDate(31)
	assert.Equal(30, lib2.Date())

	lib3 := simpleString("2001-01-10")
	assert.Equal(10, lib3.Date())

	lib3.SetDate(31)
	assert.Equal(31, lib3.Date())
}

func TestSetDateOutOfRange(t *testing.T) {
	assert := assert.New(t)

	lib := simpleString("2016-09-04")
	lib.SetDate(-1)
	assert.Equal(4, lib.Date())

	lib.SetDate(35)
	assert.Equal(4, lib.Date())
}

func TestSetDatePastEndOfMonth(t *testing.T) {
	assert := assert.New(t)

	lib := simpleString("2001-02-13")
	lib.SetDate(29)
	assert.Equal(28, lib.Date())

	lib2 := simpleString("2001-06-25")
	lib2.SetDate(31)
	assert.Equal(30, lib2.Date())
}

func TestSetDateHandlesLeapyear(t *testing.T) {
	assert := assert.New(t)

	lib := simpleString("2000-02-13")
	assert.Equal(13, lib.Date())
	assert.Equal(29, lib.SetDate(29).Date())
}

func TestSetDay(t *testing.T) {
	assert := assert.New(t)

	lib := simpleString("2016-09-04")
	assert.Equal(0, lib.Day()) // Sunday

	lib.SetDay(3) // Wednesday
	assert.Equal(7, lib.Date())

	lib.SetDay(6) // Saturday
	assert.Equal(10, lib.Date())
}

func TestSetDayInvalid(t *testing.T) {
	assert := assert.New(t)

	lib := simpleString("2016-09-04")
	assert.Equal(0, lib.Day()) // Sunday

	lib.SetDay(true)
	assert.Equal(4, lib.Date())

	lib.SetDay(6, "test")
	assert.Equal(4, lib.Date())
}

func TestSetDayByName(t *testing.T) {
	assert := assert.New(t)

	lib := simpleString("2016-09-04")
	assert.Equal(0, lib.Day()) // Sunday

	lib.SetDay("Wednesday") // Wednesday
	assert.Equal(7, lib.Date())

	lib.SetDay("Saturday") // Saturday
	assert.Equal(10, lib.Date())
}

func TestSetDayRangeExceeded(t *testing.T) {
	assert := assert.New(t)

	lib := simpleString("2020-07-26")
	assert.Equal(26, lib.Date())

	lib.SetDay(-7)
	assert.Equal(19, lib.Date())

	lib.SetDay(9)
	assert.Equal(28, lib.Date())
}

func TestSetWeekday(t *testing.T) {
	assert := assert.New(t)

	lib := simpleString("2016-09-04")
	assert.Equal(0, lib.Day()) // Sunday

	lib.SetWeekday(3) // Wednesday
	assert.Equal(7, lib.Date())

	lib.SetWeekday(6) // Saturday
	assert.Equal(10, lib.Date())
}

func TestSetWeekdayRangeExceeded(t *testing.T) {
	assert := assert.New(t)

	lib := simpleString("2020-07-26")
	assert.Equal(26, lib.Date())

	lib.SetWeekday(-7)
	assert.Equal(19, lib.Date())

	lib.SetWeekday(9)
	assert.Equal(28, lib.Date())
}

func TestSetISOWeekday(t *testing.T) {
	assert := assert.New(t)

	lib := simpleString("2016-09-05")
	assert.Equal(1, lib.ISOWeekday()) // Monday

	lib.SetISOWeekday(3) // Wednesday
	assert.Equal(7, lib.Date())

	lib.SetISOWeekday(7) // Sunday
	assert.Equal(4, lib.Date())
}

func TestSetISOWeekdayOutOfRange(t *testing.T) {
	assert := assert.New(t)

	lib := simpleString("2016-09-04")

	lib.SetISOWeekday(-1)
	assert.Equal(4, lib.Date())

	lib.SetISOWeekday(8)
	assert.Equal(4, lib.Date())
}

func TestSetDayOfYear(t *testing.T) {
	assert := assert.New(t)

	lib := simpleString("2000-01-01") // Leap year
	assert.Equal(1, lib.DayOfYear())

	lib.SetDayOfYear(100)
	assert.Equal(100, lib.DayOfYear())

	lib.SetDayOfYear(366)
	assert.Equal(366, lib.DayOfYear())

	lib2 := simpleString("2001-01-01") // Non-leap year
	assert.Equal(1, lib2.DayOfYear())

	lib2.SetDayOfYear(100)
	assert.Equal(100, lib2.DayOfYear())

	lib2.SetDayOfYear(366)
	assert.Equal(365, lib2.DayOfYear())
}

func TestSetDayOfYearOutOfRange(t *testing.T) {
	assert := assert.New(t)
	lib := simpleString("2016-09-04")

	lib.SetDayOfYear(0)
	assert.Equal(248, lib.DayOfYear())

	lib.SetDayOfYear(368)
	assert.Equal(248, lib.DayOfYear())
}

func TestSetWeek(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(30, simple(DateTime{Year: 2012, Month: 1, Day: 1}).SetWeek(30).Week(), "Setting Jan 1 2012 to week 30 should work")
	assert.Equal(30, simple(DateTime{Year: 2012, Month: 1, Day: 7}).SetWeek(30).Week(), "Setting Jan 7 2012 to week 30 should work")
	assert.Equal(30, simple(DateTime{Year: 2012, Month: 1, Day: 8}).SetWeek(30).Week(), "Setting Jan 8 2012 to week 30 should work")
	assert.Equal(30, simple(DateTime{Year: 2012, Month: 1, Day: 14}).SetWeek(30).Week(), "Setting Jan 14 2012 to week 30 should work")
	assert.Equal(30, simple(DateTime{Year: 2012, Month: 1, Day: 15}).SetWeek(30).Week(), "Setting Jan 15 2012 to week 30 should work")
}

func TestSetISOWeek(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(9, simple(DateTime{Year: 2012, Month: 1, Day: 1}).SetISOWeek(1).DayOfYear(), "Setting Jan 1 2012 to week 1 should be day of year 9")
	assert.Equal(2011, simple(DateTime{Year: 2012, Month: 1, Day: 1}).SetISOWeek(1).Year(), "Setting Jan 1 2012 to week 1 should be year 2011")
	assert.Equal(2, simple(DateTime{Year: 2012, Month: 1, Day: 2}).SetISOWeek(1).DayOfYear(), "Setting Jan 2 2012 to week 1 should be day of year 2")
	assert.Equal(8, simple(DateTime{Year: 2012, Month: 1, Day: 8}).SetISOWeek(1).DayOfYear(), "Setting Jan 8 2012 to week 1 should be day of year 8")
	assert.Equal(2, simple(DateTime{Year: 2012, Month: 1, Day: 9}).SetISOWeek(1).DayOfYear(), "Setting Jan 9 2012 to week 1 should be day of year 2")
	assert.Equal(8, simple(DateTime{Year: 2012, Month: 1, Day: 15}).SetISOWeek(1).DayOfYear(), "Setting Jan 15 2012 to week 1 should be day of year 8")
}

func TestSetMonth(t *testing.T) {
	assert := assert.New(t)

	lib := simpleString("2004-04-15")

	lib.SetMonth(6)
	assert.Equal(6, lib.Month())
	assert.Equal(15, lib.Date())
}

func TestSetMonthOutOfRange(t *testing.T) {
	assert := assert.New(t)

	lib := simpleString("2004-04-15")

	lib.SetMonth(0)
	assert.Equal(4, lib.Month())

	lib.SetMonth(15)
	assert.Equal(4, lib.Month())
}

func TestSetMonthWithLessDays(t *testing.T) {
	assert := assert.New(t)

	lib := simpleString("2001-01-31")
	assert.Equal(31, lib.Date())

	lib.SetMonth(2)
	assert.Equal(2, lib.Month())
	assert.Equal(28, lib.Date())
}

func TestSetQuarter(t *testing.T) {
	assert := assert.New(t)

	lib := simpleString("2001-02-01")

	lib.SetQuarter(1)
	assert.Equal(2, lib.Month())
	assert.Equal(1, lib.Quarter())

	lib.SetQuarter(2)
	assert.Equal(5, lib.Month())
	assert.Equal(2, lib.Quarter())

	lib.SetQuarter(3)
	assert.Equal(8, lib.Month())
	assert.Equal(3, lib.Quarter())

	lib.SetQuarter(4)
	assert.Equal(11, lib.Month())
	assert.Equal(4, lib.Quarter())

	lib2 := simpleString("2001-12-01")

	lib2.SetQuarter(3)
	assert.Equal(9, lib2.Month())
	assert.Equal(3, lib2.Quarter())
}

func TestSetQuarterOutOfRange(t *testing.T) {
	assert := assert.New(t)
	lib := simpleString("2001-12-01")

	lib.SetQuarter(0)
	assert.Equal(12, lib.Month())
	assert.Equal(4, lib.Quarter())

	lib.SetQuarter(6)
	assert.Equal(12, lib.Month())
	assert.Equal(4, lib.Quarter())
}

func TestSetYear(t *testing.T) {
	lib := simpleString("1985-05-05")
	assert.Equal(t, 2015, lib.SetYear(2015).Year())
}

func TestSetWeekYear(t *testing.T) {
	assert := assert.New(t)

	assert.Equal("2010-12-31", simpleString("2010-01-01").SetWeekYear(2011).Format("YYYY-MM-DD"))
	assert.Equal("2001-06-10 05:06:07", simpleString("2002-06-09 05:06:07").SetWeekYear(2001).Format("YYYY-MM-DD HH:MM:ss"))
}

func TestSetISOWeekYear(t *testing.T) {
	assert := assert.New(t)

	for year := 2000; year <= 2020; year++ {
		assert.Equal(year, simpleString("2012-12-31T00:00:00.000Z").UTC().SetISOWeekYear(year).ISOWeekYear(), fmt.Sprintf("setting iso-week-year to %d", year))
	}
}
