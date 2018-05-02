package goment

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type testParseable struct {
	DateTime   string
	ParsedTime time.Time
}

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
		testParseable{"2013-02-08 09:30:26.123-0600", time.Date(2013, 2, 8, 9, 30, 26, calculateNanoseconds(123), getLocation("America/Chicago"))},
		testParseable{"2013-02-08 09+0700", time.Date(2013, 2, 8, 9, 0, 0, 0, getLocation("Antarctica/Davis"))},
		// testParseable{"2013-02-08 09+07:00", time.Date(2013, 2, 8, 9, 0, 0, 0, getLocation("America/Chicago"))}, // Need to support : in tz
		// testParseable{"2013-02-08 09:30:26,123", time.Date(2013, 2, 8, 9, 30, 26, 123*1000*1000, time.Local)}, comma in date string not supported by Go
		// testParseable{"+002010-01-01", time.Date(2010, 1, 1, 0, 0, 0, 0, time.Local)},
		// testParseable{"-002010-01-01", time.Date(2010, 1, 1, 0, 0, 0, 0, time.Local)},
	}

	for _, p := range parseable {
		assertParsed(t, p.DateTime, p.ParsedTime)
	}
}

func assertParsed(t *testing.T, date string, testTime time.Time) {
	parsed, err := parseISOString(date)
	if assert.NoError(t, err) {
		assert.True(t, testTime.Equal(parsed), fmt.Sprintf("%s not equal to %s", parsed, testTime))
	}
}

func getLocation(locationName string) *time.Location {
	location, _ := time.LoadLocation(locationName)
	return location
}

func calculateNanoseconds(ms int) int {
	return ms * 1000 * 1000
}
