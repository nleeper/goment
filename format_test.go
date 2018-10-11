package goment

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestFormats(t *testing.T) {
	formats := map[string]string{
		"dddd, MMMM Do YYYY, h:mm:ss a": "Sunday, February 14th 2010, 3:25:50 pm",
		"ddd, hA":                       "Sun, 3PM",
		"M Mo MM MMMM MMM":              "2 2nd 02 February Feb",
		"YYYY YY Y":                     "2010 10 2010",
		"D Do DD":                       "14 14th 14",
		"d do dddd ddd dd":              "0 0th Sunday Sun Su",
		"DDD DDDo DDDD":                 "45 45th 045",
		// "w wo ww":                       "8 8th 08",
		"W Wo WW": "6 6th 06",
		"h hh":    "3 03",
		"H HH":    "15 15",
		"k kk":    "16 16",
		"m mm":    "25 25",
		"s ss":    "50 50",
		"a A":     "pm PM",
		"z zz":    "CST CST",
		"Z ZZ":    "-06:00 -0600",
		"[the] DDDo [day of the year]": "the 45th day of the year",
		"LTS":  "3:25:50 PM",
		"L":    "02/14/2010",
		"LL":   "February 14, 2010",
		"LLL":  "February 14, 2010 3:25 PM",
		"LLLL": "Sunday, February 14, 2010 3:25 PM",
		"l":    "2/14/2010",
		"ll":   "Feb 14, 2010",
		"lll":  "Feb 14, 2010 3:25 PM",
		"llll": "Sun, Feb 14, 2010 3:25 PM",
	}

	lib := simpleTime(time.Date(2010, 2, 14, 15, 25, 50, 125000000, chicagoLocation()))

	for p, r := range formats {
		assert.Equal(t, r, lib.Format(p), r)
	}
}

func TestDefaultFormat(t *testing.T) {
	lib := simpleTime(time.Date(2010, 2, 14, 15, 25, 50, 125000000, chicagoLocation()))
	assert.Equal(t, "2010-02-14T15:25:50-06:00", lib.Format(), "default format")
}
