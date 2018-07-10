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
		// testParseable{"2013-02-08 09+07:00", time.Date(2013, 2, 8, 9, 0, 0, 0, getLocation("America/Chicago"))}, // Need to support : in tz
		// testParseable{"2013-02-08 09:30:26,123", time.Date(2013, 2, 8, 9, 30, 26, 123*1000*1000, chicagoLocation())}, comma in date string not supported by Go
		// testParseable{"+002010-01-01", time.Date(2010, 1, 1, 0, 0, 0, 0, chicagoLocation())},
		// testParseable{"-002010-01-01", time.Date(2010, 1, 1, 0, 0, 0, 0, chicagoLocation())},
	}

	for _, p := range parseable {
		parsed, err := parseISOString(p.DateTime)
		if assert.NoError(t, err) {
			assert.True(t, p.ParsedTime.Equal(parsed), fmt.Sprintf("%s not equal to %s", parsed, p.ParsedTime))
		}
	}
}

func TestFormatParsing(t *testing.T) {
	formats := map[string][]string{
		"YYYY-Q":              []string{"2014-4"},
		"MM-DD-YYYY":          []string{"12-02-1999"},
		"DD-MM-YYYY":          []string{"12-02-1999"},
		"DD/MM/YYYY":          []string{"12/02/1999"},
		"DD_MM_YYYY":          []string{"12_02_1999"},
		"DD:MM:YYYY":          []string{"12:02:1999"},
		"D-M-YY":              []string{"2-2-99"},
		"Y":                   []string{"-25"},
		"YY":                  []string{"99"},
		"DDD-YYYY":            []string{"300-1999"},
		"YYYY-DDD":            []string{"1999-300"},
		"YYYY MM Do":          []string{"2014 01 3rd", "2015 11 21st", "2016 05 16th"},
		"MMM":                 []string{"Apr"},
		"MMMM":                []string{"December"},
		"YYYY MMMM":           []string{"2018 October"},
		"D":                   []string{"3", "27"},
		"DD":                  []string{"04", "23"},
		"DDD":                 []string{"7", "300"},
		"DDDD":                []string{"008", "211", "312"},
		"h":                   []string{"4"},
		"H":                   []string{"1", "10", "23"},
		"DD-MM-YYYY h:m:s":    []string{"12-02-1999 2:45:10"},
		"DD-MM-YYYY h:m:s a":  []string{"12-02-1999 2:45:10 am", "12-02-1999 2:45:10 pm"},
		"h:mm a":              []string{"12:00 pm", "12:30 pm", "12:00 am", "12:30 am"},
		"HH:mm":               []string{"12:00"},
		"kk:mm":               []string{"12:00"},
		"YYYY-MM-DDTHH:mm:ss": []string{"2011-11-11T11:11:11"},
		"MM-DD-YYYY [M]":      []string{"12-02-1999 M"},
		// "ddd MMM DD HH:mm:ss YYYY":  []string{"Wed Apr 08 22:52:51 2009"},
		// "dddd MMM DD HH:mm:ss YYYY": []string{"Saturday Apr 11 22:52:51 2009"},
		"HH:mm:ss":               []string{"12:00:00", "12:30:00", "00:00:00"},
		"YYYY-MM-DD HH:mm:ss ZZ": []string{"2000-05-15 17:08:00 -0700"},
		"YYYY-MM-DD HH:mm Z":     []string{"2010-10-20 04:30 +00:00"},
		// "e":                         []string{"0", "5"},
		// "E": []string{"1", "7"},
		// "HH:mm:ss.S": []string{"00:30:00.1"},
		// "HH:mm:ss SS":               "00:30:00 12",
		// "HH:mm:ss SSS":              "00:30:00 123",
		// "HH:mm:ss S":                "00:30:00 7",
		// "HH:mm:ss SS":               "00:30:00 78",
		// "HH:mm:ss SSS":              "00:30:00 789",
		"kk:mm:ss": []string{"12:00:00", "12:30:00", "24:00:00", "09:00:00"},
		// "kk:mm:ss S":   "24:30:00 1",
		// "kk:mm:ss SS":  "24:30:00 12",
		// "kk:mm:ss SSS": "24:30:00 123",
		// "kk:mm:ss S":   "24:30:00 7",
		// "kk:mm:ss SS":  "24:30:00 78",
		// "kk:mm:ss SSS": "24:30:00 789",
		"X": []string{"1234567890"},
		// "H Z":  []string{"6 -06:00"},
		// "H ZZ": []string{"5 -0700"},
		// "LT":   []string{"12:30 AM"},
		// "LTS":  []string{"12:30:29 AM"},
		// "L":    []string{"09/02/1999"},
		// "l":    []string{"9/2/1999"},
		// "LL":   []string{"September 2, 1999"},
		// "ll":   []string{"Sep 2, 1999"},
		// "LLL":  []string{"September 2, 1999 12:30 AM"},
		// "lll":  []string{"Sep 2, 1999 12:30 AM"},
		// "LLLL": []string{"Thursday, September 2, 1999 12:30 AM"},
		// "llll": []string{"Thu, Sep 2, 1999 12:30 AM"},
	}

	for format, dates := range formats {
		for _, date := range dates {
			lib, _ := New(date, format)
			assert.Equal(t, date, lib.Format(format), fmt.Sprintf("%v: %v", format, date))
		}
	}
}

func TestHourFormatParsing(t *testing.T) {
	lib, _ := New("23", "h")
	assert.Equal(t, "11", lib.Format("h"), "h: 23")
}

func getLocation(locationName string) *time.Location {
	location, _ := time.LoadLocation(locationName)
	return location
}

func calculateNanoseconds(ms int) int {
	return ms * 1000 * 1000
}
