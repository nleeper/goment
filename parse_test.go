package goment

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestISOParsing(t *testing.T) {
	parseable := []testParseable{
		testParseable{"2010-01-01", time.Date(2010, 1, 1, 0, 0, 0, 0, time.UTC)},
		testParseable{"20090406", time.Date(2009, 4, 6, 0, 0, 0, 0, time.UTC)},
		testParseable{"2011-02", time.Date(2011, 2, 1, 0, 0, 0, 0, time.UTC)},
		testParseable{"2013-02-08 09", time.Date(2013, 2, 8, 9, 0, 0, 0, time.UTC)},
		testParseable{"2013-02-08T09", time.Date(2013, 2, 8, 9, 0, 0, 0, time.UTC)},
		testParseable{"2013-02-08 09:30", time.Date(2013, 2, 8, 9, 30, 0, 0, time.UTC)},
		testParseable{"2013-02-08 09:30:26", time.Date(2013, 2, 8, 9, 30, 26, 0, time.UTC)},
		testParseable{"2013-02-08 09:30:26.123", time.Date(2013, 2, 8, 9, 30, 26, calculateNanoseconds(123), time.UTC)},
		testParseable{"20130208 0930", time.Date(2013, 2, 8, 9, 30, 0, 0, time.UTC)},
		testParseable{"20130208T0930", time.Date(2013, 2, 8, 9, 30, 0, 0, time.UTC)},
		testParseable{"20130208 093026", time.Date(2013, 2, 8, 9, 30, 26, 0, time.UTC)},
		testParseable{"20130208T093026", time.Date(2013, 2, 8, 9, 30, 26, 0, time.UTC)},
		testParseable{"20130208 093026.123", time.Date(2013, 2, 8, 9, 30, 26, calculateNanoseconds(123), time.UTC)},
		testParseable{"20130208T093026.123", time.Date(2013, 2, 8, 9, 30, 26, calculateNanoseconds(123), time.UTC)},
		testParseable{"2013-02-08 09:30:26.123-0600", time.Date(2013, 2, 8, 9, 30, 26, calculateNanoseconds(123), chicagoLocation())},
		testParseable{"2013-02-08 09+0700", time.Date(2013, 2, 8, 9, 0, 0, 0, getLocation("Antarctica/Davis"))},
		testParseable{"2013-02-08T09:30:26.123Z", time.Date(2013, 2, 8, 9, 30, 26, 123, time.UTC)},
		testParseable{"2013-02-08T09:30:26Z", time.Date(2013, 2, 8, 9, 30, 26, 0, time.UTC)},
		// testParseable{"2013-02-08 09+07:00", time.Date(2013, 2, 8, 9, 0, 0, 0, getLocation("America/Chicago"))}, // Need to support : in tz
		// testParseable{"2013-02-08 09:30:26,123", time.Date(2013, 2, 8, 9, 30, 26, 123*1000*1000, chicagoLocation())}, comma in date string not supported by Go
		// testParseable{"+002010-01-01", time.Date(2010, 1, 1, 0, 0, 0, 0, chicagoLocation())},
		// testParseable{"-002010-01-01", time.Date(2010, 1, 1, 0, 0, 0, 0, chicagoLocation())},
	}

	for _, p := range parseable {
		parsed, _ := parseISOString(p.dateTime)
		assert.Equal(t, simpleTime(p.parsedTime).Format(), simpleTime(parsed).Format(), fmt.Sprintf("%s not equal to %s", parsed, p.parsedTime))
	}
}

func TestFormatParsing(t *testing.T) {
	assert := assert.New(t)

	formats := map[string][]string{
		"YYYY-Q":                    []string{"2014-4"},
		"MM-DD-YYYY":                []string{"12-02-1999"},
		"DD-MM-YYYY":                []string{"12-02-1999"},
		"DD/MM/YYYY":                []string{"12/02/1999"},
		"DD_MM_YYYY":                []string{"12_02_1999"},
		"DD:MM:YYYY":                []string{"12:02:1999"},
		"D-M-YY":                    []string{"2-2-99"},
		"Y":                         []string{"-0025"},
		"YY":                        []string{"99"},
		"DDD-YYYY":                  []string{"300-1999"},
		"YYYY-DDD":                  []string{"1999-300"},
		"YYYY MM Do":                []string{"2014 01 3rd", "2015 11 21st", "2016 05 16th"},
		"MMM":                       []string{"Apr"},
		"MMMM":                      []string{"December"},
		"YYYY MMMM":                 []string{"2018 October"},
		"D":                         []string{"3", "27"},
		"DD":                        []string{"04", "23"},
		"DDD":                       []string{"7", "300"},
		"DDDD":                      []string{"008", "211", "312"},
		"h":                         []string{"4"},
		"H":                         []string{"1", "10", "23"},
		"DD-MM-YYYY h:m:s":          []string{"12-02-1999 2:45:10"},
		"DD-MM-YYYY h:m:s a":        []string{"12-02-1999 2:45:10 am", "12-02-1999 2:45:10 pm"},
		"h:mm a":                    []string{"12:00 pm", "12:30 pm", "12:00 am", "12:30 am"},
		"HH:mm":                     []string{"12:00"},
		"kk:mm":                     []string{"12:00"},
		"YYYY-MM-DDTHH:mm:ss":       []string{"2011-11-11T11:11:11"},
		"MM-DD-YYYY [M]":            []string{"12-02-1999 M"},
		"ddd MMM DD HH:mm:ss YYYY":  []string{"Wed Apr 08 22:52:51 2009"},
		"dddd MMM DD HH:mm:ss YYYY": []string{"Saturday Apr 11 22:52:51 2009"},
		"HH:mm:ss":                  []string{"12:00:00", "12:30:11", "00:00:00"},
		"kk:mm:ss":                  []string{"12:00:10", "12:30:42", "24:00:00", "09:00:00"},
		"YYYY-MM-DD HH:mm:ss ZZ":    []string{"2000-05-15 17:08:00 -0700"},
		"YYYY-MM-DD HH:mm Z":        []string{"2010-10-20 04:30 +00:00"},
		"e":                         []string{"0", "5"},
		"E":                         []string{"1", "7"},
		// "HH:mm:ss.S": []string{"00:30:00.1"},
		// "HH:mm:ss SS":               "00:30:00 12",
		// "HH:mm:ss SSS":              "00:30:00 123",
		// "HH:mm:ss S":                "00:30:00 7",
		// "HH:mm:ss SS":               "00:30:00 78",
		// "HH:mm:ss SSS":              "00:30:00 789",
		// "kk:mm:ss S":   "24:30:00 1",
		// "kk:mm:ss SS":  "24:30:00 12",
		// "kk:mm:ss SSS": "24:30:00 123",
		// "kk:mm:ss S":   "24:30:00 7",
		// "kk:mm:ss SS":  "24:30:00 78",
		// "kk:mm:ss SSS": "24:30:00 789",
		"X":    []string{"1234567890"},
		"H Z":  []string{"6 -06:00"},
		"H ZZ": []string{"5 -0700"},
		"LT":   []string{"12:30 AM"},
		"LTS":  []string{"12:30:29 AM"},
		"L":    []string{"09/02/1999"},
		"l":    []string{"9/2/1999"},
		"LL":   []string{"September 2, 1999"},
		"ll":   []string{"Sep 2, 1999"},
		"LLL":  []string{"September 2, 1999 12:30 AM"},
		"lll":  []string{"Sep 2, 1999 12:30 AM"},
		"LLLL": []string{"Thursday, September 2, 1999 12:30 AM"},
		"llll": []string{"Thu, Sep 2, 1999 12:30 AM"},
	}

	for format, dates := range formats {
		for _, date := range dates {
			assert.Equal(date, simpleFormat(date, format).Format(format), fmt.Sprintf("%v: %v", format, date))
		}
	}
}

func TestFormatNoSeparators(t *testing.T) {
	assert := assert.New(t)

	formats := map[string]string{
		"MMDDYYYY":  "12021999",
		"DDMMYYYY":  "12021999",
		"YYYYMMDD":  "19991202",
		"DDMMMYYYY": "10Sep2001",
	}

	for format, date := range formats {
		assert.Equal(date, simpleFormat(date, format).Format(format), fmt.Sprintf("%v: %v", format, date))
	}
}

func TestHourFormatParsing(t *testing.T) {
	lib := simpleFormat("23", "h")
	assert.Equal(t, "11", lib.Format("h"), "h: 23")
}

func TestDefaultToCurrentDate(t *testing.T) {
	assert := assert.New(t)

	now := simpleNow()

	assert.Equal(now.Clone().SetHour(12).SetMinute(13).SetSecond(14).Format("YYYY-MM-DD hh:mm:ss"), simpleFormat("12:13:14", "hh:mm:ss").Format("YYYY-MM-DD hh:mm:ss"))
	assert.Equal(now.Clone().SetDate(5).Format("YYYY-MM-DD"), simpleFormat("05", "DD").Format("YYYY-MM-DD"))
	assert.Equal(now.Clone().SetMonth(5).SetDate(1).Format("YYYY-MM-DD"), simpleFormat("05", "MM").Format("YYYY-MM-DD"))
	assert.Equal(now.Clone().SetYear(1996).SetMonth(1).SetDate(1).Format("YYYY-MM-DD"), simpleFormat("1996", "YYYY").Format("YYYY-MM-DD"))
}

func TestParseTwoDigitYear(t *testing.T) {
	assert := assert.New(t)

	assert.Equal("09/02/1999", simpleFormat("9/2/99", "D/M/YYYY").Format("DD/MM/YYYY"))
	assert.Equal("09/02/1999", simpleFormat("9/2/1999", "D/M/YYYY").Format("DD/MM/YYYY"))
	assert.Equal("09/02/2068", simpleFormat("9/2/68", "D/M/YYYY").Format("DD/MM/YYYY"))
	assert.Equal("09/02/1969", simpleFormat("9/2/69", "D/M/YYYY").Format("DD/MM/YYYY"))

	assert.Equal("2067", simpleFormat("67", "YY").Format("YYYY"))
	assert.Equal("2068", simpleFormat("68", "YY").Format("YYYY"))
	assert.Equal("1969", simpleFormat("69", "YY").Format("YYYY"))
	assert.Equal("1970", simpleFormat("70", "YY").Format("YYYY"))
}

func TestParseLongYear(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(-270000, simpleFormat("-270000-01-01", "YYYYY-MM-DD").Year())
	assert.Equal(270000, simpleFormat("270000-01-01", "YYYYY-MM-DD").Year())
	assert.Equal(270000, simpleFormat("+270000-01-01", "YYYYY-MM-DD").Year())

	assert.Equal(-1000, simpleFormat("-1000-01-01", "YYYYY-MM-DD").Year())
}

func TestDayOfWeekParsing(t *testing.T) {
	assert := assert.New(t)

	outputFormat := "YYYY MM DD"

	lib := simpleNow()
	lib.SetWeekday(0)

	expected := lib.Format(outputFormat)

	formats := map[string]string{
		"e":    "0",
		"ddd":  "Sun",
		"dddd": "Sunday",
	}

	for format, date := range formats {
		assert.Equal(expected, simpleFormat(date, format).Format(outputFormat))
	}
}

func TestWeekYearParsing(t *testing.T) {
	assert := assert.New(t)

	outputFormat := "YYYY-MM-DD"

	assert.Equal("2006-12-31", simpleFormat("2007-01", "gggg-ww").Format(outputFormat))
	assert.Equal("2007-12-30", simpleFormat("2008-01", "gggg-ww").Format(outputFormat))
	assert.Equal("2002-12-29", simpleFormat("2003-01", "gggg-ww").Format(outputFormat))
	assert.Equal("2008-12-28", simpleFormat("2009-01", "gggg-ww").Format(outputFormat))
	assert.Equal("2009-12-27", simpleFormat("2010-01", "gggg-ww").Format(outputFormat))
	assert.Equal("2010-12-26", simpleFormat("2011-01", "gggg-ww").Format(outputFormat))
	assert.Equal("2012-01-01", simpleFormat("2012-01", "gggg-ww").Format(outputFormat))
}

func TestWeekWeekdayParsing(t *testing.T) {
	assert := assert.New(t)

	format := "YYYY MM DD"

	currentWeekOfYear := simpleNow().Week()

	expectedDate2012 := simple(DateTime{Year: 2012, Month: 1, Day: 1}).SetDay(0).Add(currentWeekOfYear-1, "weeks").Format(format)
	expectedDate1999 := simple(DateTime{Year: 1999, Month: 1, Day: 1}).SetDay(0).Add(currentWeekOfYear-1, "weeks").Format(format)

	assert.Equal(expectedDate2012, simpleFormat("12", "gg").Format(format), "week-year two digits")
	assert.Equal(expectedDate2012, simpleFormat("2012", "gggg").Format(format), "week-year four digits")
	assert.Equal(expectedDate1999, simpleFormat("99", "gg").Format(format), "week-year two digits previous year")
	assert.Equal(expectedDate1999, simpleFormat("1999", "gggg").Format(format), "week-year four digits previous year")

	assert.Equal("1999 01 04", simpleFormat("99", "GG").Format(format), "iso week-year two digits")
	assert.Equal("1999 01 04", simpleFormat("1999", "GGGG").Format(format), "iso week-year four digits")

	assert.Equal("2012 12 31", simpleFormat("13", "GG").Format(format), "iso week-year two digits previous year")
	assert.Equal("2012 12 31", simpleFormat("2013", "GGGG").Format(format), "iso week-year four digits previous year")

	// Year + week parsing
	assert.Equal("1999 09 05", simpleFormat("1999 37", "gggg w").Format(format), "week")
	assert.Equal("1999 09 05", simpleFormat("1999 37", "gggg ww").Format(format), "week double")
	assert.Equal("1999 09 13", simpleFormat("1999 37", "GGGG W").Format(format), "iso week")
	assert.Equal("1999 09 13", simpleFormat("1999 37", "GGGG WW").Format(format), "iso week double")

	assert.Equal("1999 09 16", simpleFormat("1999 37 4", "GGGG WW E").Format(format), "iso day")
	assert.Equal("1999 09 16", simpleFormat("1999 37 04", "GGGG WW E").Format(format), "iso day wide")

	assert.Equal("1999 09 09", simpleFormat("1999 37 4", "gggg ww e").Format(format), "day")
	assert.Equal("1999 09 09", simpleFormat("1999 37 04", "gggg ww e").Format(format), "day wide")

	// Year + week + day parsing
	assert.Equal("1999 09 09", simpleFormat("1999 37 4", "gggg ww d").Format(format), "d")
	assert.Equal("1999 09 09", simpleFormat("1999 37 Th", "gggg ww dd").Format(format), "dd")
	assert.Equal("1999 09 09", simpleFormat("1999 37 Thu", "gggg ww ddd").Format(format), "ddd")
	assert.Equal("1999 09 09", simpleFormat("1999 37 Thursday", "gggg ww dddd").Format(format), "dddd")

	// Lower-order only parsing
	assert.Equal(22, simpleFormat("22", "ww").Week(), "week sets the week by itself")
	assert.Equal(simpleNow().WeekYear(), simpleFormat("22", "ww").WeekYear(), "week keeps this year")
	assert.Equal(2012, simpleFormat("2012 22", "YYYY ww").WeekYear(), "week keeps parsed year")

	assert.Equal(22, simpleFormat("22", "WW").ISOWeek(), "iso week sets the week by itself")
	assert.Equal(simpleNow().ISOWeekYear(), simpleFormat("22", "ww").ISOWeekYear(), "iso week keeps this year")
	assert.Equal(2012, simpleFormat("2012 22", "YYYY WW").WeekYear(), "iso week keeps parsed year")

	// Order
	assert.Equal("2013 01 12", simpleFormat("6 2013 2", "e gggg w").Format(format), "order doesn't matter")
	assert.Equal("2013 01 12", simpleFormat("6 2013 2", "E GGGG W").Format(format), "iso order doesn't matter")

	// Can parse other fields with weeks
	assert.Equal("1999 09 16 03:30", simpleFormat("1999-W37-4 3:30", "GGGG-[W]WW-E HH:mm").Format("YYYY MM DD HH:mm"), "parsing weeks and hours")
}

func TestParseWeekdayMismatch(t *testing.T) {
	assert := assert.New(t)

	_, err := New("Wed 08-10-2017", "ddd MM-DD-YYYY") // 8-10-2017 is a Thursday
	assert.EqualError(err, "There is a mismatch between parsed weekday and expected weekday")

	_, err2 := New("Thu 08-10-2017", "ddd MM-DD-YYYY") // 8-10-2017 is a Thursday
	assert.NoError(err2)
}

func TestISO8601Timestamp(t *testing.T) {
	assert := assert.New(t)

	outputFormat := "YYYY-MM-DDTHH:mm:ssZ"

	assert.Equal("2020-09-01T20:46:07+00:00", simpleFormat("2020-09-01T20:46:07Z", "YYYY-MM-DDTHH:mm:ssZ").Format(outputFormat))
}

func getLocation(locationName string) *time.Location {
	location, _ := time.LoadLocation(locationName)
	return location
}

func calculateNanoseconds(ms int) int {
	return ms * 1000 * 1000
}
