package goment

import (
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
		testParseable{"2013-02-08 09:30:26.123", time.Date(2013, 2, 8, 9, 30, 26, 123*1000*1000, time.UTC)},
		testParseable{"20130208 0930", time.Date(2013, 2, 8, 9, 30, 0, 0, time.UTC)},
		testParseable{"20130208T0930", time.Date(2013, 2, 8, 9, 30, 0, 0, time.UTC)},
		testParseable{"20130208 093026", time.Date(2013, 2, 8, 9, 30, 26, 0, time.UTC)},
		testParseable{"20130208T093026", time.Date(2013, 2, 8, 9, 30, 26, 0, time.UTC)},
		testParseable{"20130208 093026.123", time.Date(2013, 2, 8, 9, 30, 26, 123*1000*1000, time.UTC)},
		testParseable{"20130208T093026.123", time.Date(2013, 2, 8, 9, 30, 26, 123*1000*1000, time.UTC)},
		// testParseable{"2013-02-08 09:30:26,123", time.Date(2013, 2, 8, 9, 30, 26, 123*1000*1000, time.UTC)}, comma in date string not supported by Go
		// testParseable{"+002010-01-01", time.Date(2010, 1, 1, 0, 0, 0, 0, time.UTC)},
		// testParseable{"-002010-01-01", time.Date(2010, 1, 1, 0, 0, 0, 0, time.UTC)},
	}

	for _, p := range parseable {
		assertParsed(t, p.DateTime, p.ParsedTime)
	}
}

func assertParsed(t *testing.T, date string, testTime time.Time) {
	parsed, err := parseISOString(date)
	if assert.NoError(t, err) {
		assert.Equal(t, parsed, testTime)
	}
}
