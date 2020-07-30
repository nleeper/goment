package goment

import (
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestFormats(t *testing.T) {
	assert := assert.New(t)

	formats := map[string]string{
		"dddd, MMMM Do YYYY, h:mm:ss a": "Sunday, February 14th 2010, 3:25:50 pm",
		"ddd, hA":                       "Sun, 3PM",
		"M Mo MM MMMM MMM":              "2 2nd 02 February Feb",
		"YYYYYY YYYYY YYYY YY Y":        "+002010 02010 2010 10 2010",
		"D Do DD":                       "14 14th 14",
		"d do dddd ddd dd":              "0 0th Sunday Sun Su",
		"DDD DDDo DDDD":                 "45 45th 045",
		"e":                             "0",
		"E":                             "7",
		"w wo ww":                       "8 8th 08",
		"W Wo WW":                       "6 6th 06",
		"Q Qo":                          "1 1st",
		"h hh":                          "3 03",
		"H HH":                          "15 15",
		"k kk":                          "16 16",
		"m mm":                          "25 25",
		"s ss":                          "50 50",
		"a A":                           "pm PM",
		"z zz zzzz":                     "CST CST Central Standard Time",
		"Z ZZ":                          "-06:00 -0600",
		"[the] DDDo [day of the year]":  "the 45th day of the year",
		"[the] DDDo [day of the year after January]": "the 45th day of the year after January",
		"X":        "1266182750",
		"x":        "1266182750125",
		"LT":       "3:25 PM",
		"LTS":      "3:25:50 PM",
		"L":        "02/14/2010",
		"LL":       "February 14, 2010",
		"LLL":      "February 14, 2010 3:25 PM",
		"LLLL":     "Sunday, February 14, 2010 3:25 PM",
		"l":        "2/14/2010",
		"l [test]": "2/14/2010 test",
		"ll":       "Feb 14, 2010",
		"lll":      "Feb 14, 2010 3:25 PM",
		"llll":     "Sun, Feb 14, 2010 3:25 PM",
	}

	lib := simpleTime(time.Date(2010, 2, 14, 15, 25, 50, 125000000, chicagoLocation()))

	for p, r := range formats {
		assert.Equal(r, lib.Format(p), r)
	}
}

func TestDefaultFormat(t *testing.T) {
	lib := simpleTime(time.Date(2010, 2, 14, 15, 25, 50, 125000000, chicagoLocation()))
	assert.Equal(t, "2010-02-14T15:25:50-06:00", lib.Format(), "default format")
}

func TestLongYears(t *testing.T) {
	assert := assert.New(t)

	assert.Equal("+000002", simpleNow().UTC().SetYear(2).Format("YYYYYY"))
	assert.Equal("+002012", simpleNow().UTC().SetYear(2012).Format("YYYYYY"))
	assert.Equal("+020123", simpleNow().UTC().SetYear(20123).Format("YYYYYY"))

	assert.Equal("-000001", simpleNow().UTC().SetYear(-1).Format("YYYYYY"))
	assert.Equal("-002012", simpleNow().UTC().SetYear(-2012).Format("YYYYYY"))
	assert.Equal("-020123", simpleNow().UTC().SetYear(-20123).Format("YYYYYY"))
}

func TestISOWeekYear(t *testing.T) {
	assert := assert.New(t)

	cases := map[string]string{
		"2005-01-02": "2004-53",
		"2005-12-31": "2005-52",
		"2007-01-01": "2007-01",
		"2007-12-30": "2007-52",
		"2007-12-31": "2008-01",
		"2008-01-01": "2008-01",
		"2008-12-28": "2008-52",
		"2008-12-29": "2009-01",
		"2008-12-30": "2009-01",
		"2008-12-31": "2009-01",
		"2009-01-01": "2009-01",
		"2009-12-31": "2009-53",
		"2010-01-01": "2009-53",
		"2010-01-02": "2009-53",
		"2010-01-03": "2009-53",
		"404-12-31":  "0404-53",
		"405-12-31":  "0405-52",
	}

	for date, iso := range cases {
		isoWeekYear := strings.Split(iso, "-")[0]
		assert.Equal("0"+isoWeekYear, simpleFormat(date, "YYYY-MM-DD").Format("GGGGG"))
		assert.Equal(isoWeekYear, simpleFormat(date, "YYYY-MM-DD").Format("GGGG"))
		assert.Equal(isoWeekYear[2:4], simpleFormat(date, "YYYY-MM-DD").Format("GG"))
	}
}
