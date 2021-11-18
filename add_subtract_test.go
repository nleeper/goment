package goment

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAddNoArgsIgnored(t *testing.T) {
	testTime := time.Date(2017, 1, 1, 18, 0, 0, 0, chicagoLocation())

	lib := simpleTime(testTime)

	lib.Add()
	assert.Equal(t, testTime, lib.ToTime())
}

func TestAddTooManyArgsIgnored(t *testing.T) {
	testTime := time.Date(2017, 1, 1, 18, 0, 0, 0, chicagoLocation())

	lib := simpleTime(testTime)

	lib.Add(1, "year", "day")
	assert.Equal(t, testTime, lib.ToTime())
}

func TestAddInvalidFirstArgTypeIgnored(t *testing.T) {
	testTime := time.Date(2017, 1, 1, 18, 0, 0, 0, chicagoLocation())

	lib := simpleTime(testTime)

	lib.Add("year")
	assert.Equal(t, testTime, lib.ToTime(), fmt.Sprintf(`Date arguments must be equal, got %q, expected %q`, testTime, lib.ToTime()))
}

func TestAddInvalidSecondArgTypeIgnored(t *testing.T) {
	testTime := time.Date(2017, 1, 1, 18, 0, 0, 0, chicagoLocation())

	lib := simpleTime(testTime)

	lib.Add(1, 2)
	assert.Equal(t, testTime, lib.ToTime())
}

func TestAddInvalidSecondArgValueIgnored(t *testing.T) {
	testTime := time.Date(2017, 1, 1, 18, 0, 0, 0, chicagoLocation())

	lib := simpleTime(testTime)

	lib.Add(1, "custom")
	assert.Equal(t, testTime, lib.ToTime())
}

func TestAddDuration(t *testing.T) {
	assert := assert.New(t)

	lib := simpleString("2011-01-04 14:01:01")

	lib.Add(time.Duration(4) * time.Hour)
	assert.Equal(18, lib.Hour())

	lib.Add(time.Duration(25) * time.Minute)
	assert.Equal(26, lib.Minute())
}

func TestAddYears(t *testing.T) {
	testTime := time.Date(2017, 1, 1, 18, 0, 0, 0, chicagoLocation())

	lib := simpleTime(testTime)

	lib.Add(1, "year")
	assert.Equal(t, 2018, lib.Year())
}

func TestAddQuarters(t *testing.T) {
	assert := assert.New(t)

	testTime := time.Date(2014, 4, 11, 0, 0, 0, 0, chicagoLocation())

	lib := simpleTime(testTime)
	assert.Equal(2, lib.Quarter())
	lib.Add(2, "quarters")

	assert.Equal(10, lib.Month())
	assert.Equal(4, lib.Quarter())
}

func TestAddMonths(t *testing.T) {
	assert := assert.New(t)

	testTime := time.Date(2017, 1, 1, 18, 0, 0, 0, chicagoLocation())

	lib := simpleTime(testTime)
	lib.Add(5, "months")

	assert.Equal(2017, lib.Year())
	assert.Equal(6, lib.Month())
}

func TestAddWeeks(t *testing.T) {
	assert := assert.New(t)

	testTime := time.Date(2017, 1, 1, 18, 0, 0, 0, chicagoLocation())

	lib := simpleTime(testTime)
	lib.Add(3, "weeks")

	assert.Equal(2017, lib.Year())
	assert.Equal(1, lib.Month())
	assert.Equal(22, lib.Date())
}

func TestAddDays(t *testing.T) {
	assert := assert.New(t)

	testTime := time.Date(2017, 1, 1, 18, 0, 0, 0, chicagoLocation())

	lib := simpleTime(testTime)
	lib.Add(16, "d")

	assert.Equal(2017, lib.Year())
	assert.Equal(1, lib.Month())
	assert.Equal(17, lib.Date())
}

func TestAddHours(t *testing.T) {
	testTime := time.Date(2017, 1, 1, 15, 0, 0, 0, chicagoLocation())

	lib := simpleTime(testTime)
	lib.Add(5, "hours")

	assert.Equal(t, 20, lib.Hour())
}

func TestAddMinutes(t *testing.T) {
	testTime := time.Date(2017, 1, 1, 15, 30, 0, 0, chicagoLocation())

	lib := simpleTime(testTime)
	lib.Add(15, "m")

	assert.Equal(t, 45, lib.Minute())
}

func TestAddSeconds(t *testing.T) {
	testTime := time.Date(2017, 1, 1, 15, 0, 0, 0, chicagoLocation())

	lib := simpleTime(testTime)
	lib.Add(55, "seconds")

	assert.Equal(t, 55, lib.Second())
}

func TestAddMilliseconds(t *testing.T) {
	assert := assert.New(t)

	testTime := time.Date(2017, 1, 1, 15, 0, 0, 0, chicagoLocation())

	lib := simpleTime(testTime)
	lib.Add(6000, "ms")

	assert.Equal(6, lib.Second())
	assert.Equal(6000, lib.Millisecond())
}

func TestAddNanoseconds(t *testing.T) {
	testTime := time.Date(2017, 1, 1, 15, 0, 0, 0, chicagoLocation())

	lib := simpleTime(testTime)
	lib.Add(600, "nanoseconds")

	assert.Equal(t, 600, lib.Nanosecond())
}

func TestSubtractNoArgsIgnored(t *testing.T) {
	testTime := time.Date(2017, 1, 1, 18, 0, 0, 0, chicagoLocation())

	lib := simpleTime(testTime)
	lib.Subtract()

	assert.Equal(t, testTime, lib.ToTime())
}

func TestSubtractTooManyArgsIgnored(t *testing.T) {
	testTime := time.Date(2017, 1, 1, 18, 0, 0, 0, chicagoLocation())

	lib := simpleTime(testTime)
	lib.Subtract(1, "year", "day")

	assert.Equal(t, testTime, lib.ToTime())
}

func TestSubtractInvalidFirstArgTypeIgnored(t *testing.T) {
	testTime := time.Date(2017, 1, 1, 18, 0, 0, 0, chicagoLocation())

	lib := simpleTime(testTime)
	lib.Subtract("year")

	assert.Equal(t, testTime, lib.ToTime())
}

func TestSubtractInvalidSecondArgTypeIgnored(t *testing.T) {
	testTime := time.Date(2017, 1, 1, 18, 0, 0, 0, chicagoLocation())

	lib := simpleTime(testTime)
	lib.Subtract(1, 2)

	assert.Equal(t, testTime, lib.ToTime())
}

func TestSubtractInvalidSecondArgValueIgnored(t *testing.T) {
	testTime := time.Date(2017, 1, 1, 18, 0, 0, 0, chicagoLocation())

	lib := simpleTime(testTime)
	lib.Subtract(1, "custom")

	assert.Equal(t, testTime, lib.ToTime())
}

func TestSubtractDuration(t *testing.T) {
	assert := assert.New(t)

	lib := simpleString("2011-01-04 14:30:01")
	lib.Subtract(time.Duration(4) * time.Hour)

	assert.Equal(10, lib.Hour())

	lib.Subtract(time.Duration(25) * time.Minute)

	assert.Equal(5, lib.Minute())
}

func TestSubtractYears(t *testing.T) {
	testTime := time.Date(2017, 1, 1, 18, 0, 0, 0, chicagoLocation())

	lib := simpleTime(testTime)
	lib.Subtract(1, "year")

	assert.Equal(t, 2016, lib.Year())
}

func TestSubtractQuarters(t *testing.T) {
	assert := assert.New(t)

	testTime := time.Date(2014, 8, 11, 0, 0, 0, 0, chicagoLocation())

	lib := simpleTime(testTime)
	assert.Equal(3, lib.Quarter())
	lib.Subtract(2, "quarters")

	assert.Equal(2, lib.Month())
	assert.Equal(1, lib.Quarter())
}

func TestSubtractMonths(t *testing.T) {
	assert := assert.New(t)

	testTime := time.Date(2017, 11, 1, 18, 0, 0, 0, chicagoLocation())

	lib := simpleTime(testTime)
	lib.Subtract(5, "months")

	assert.Equal(2017, lib.Year())
	assert.Equal(6, lib.Month())
}

func TestSubtractWeeks(t *testing.T) {
	assert := assert.New(t)

	testTime := time.Date(2017, 1, 22, 18, 0, 0, 0, chicagoLocation())

	lib := simpleTime(testTime)
	lib.Subtract(3, "weeks")

	assert.Equal(2017, lib.Year())
	assert.Equal(1, lib.Month())
	assert.Equal(1, lib.Date())
}

func TestSubtractDays(t *testing.T) {
	assert := assert.New(t)

	testTime := time.Date(2017, 1, 17, 18, 0, 0, 0, chicagoLocation())

	lib := simpleTime(testTime)
	lib.Subtract(16, "d")

	assert.Equal(2017, lib.Year())
	assert.Equal(1, lib.Month())
	assert.Equal(1, lib.Date())
}

func TestSubtractHours(t *testing.T) {
	testTime := time.Date(2017, 1, 1, 15, 0, 0, 0, chicagoLocation())

	lib := simpleTime(testTime)
	lib.Subtract(5, "hours")

	assert.Equal(t, 10, lib.Hour())
}

func TestSubtractMinutes(t *testing.T) {
	testTime := time.Date(2017, 1, 1, 15, 30, 0, 0, chicagoLocation())

	lib := simpleTime(testTime)
	lib.Subtract(15, "m")

	assert.Equal(t, 15, lib.Minute())
}

func TestSubtractSeconds(t *testing.T) {
	testTime := time.Date(2017, 1, 1, 15, 0, 30, 0, chicagoLocation())

	lib := simpleTime(testTime)
	lib.Subtract(25, "seconds")

	assert.Equal(t, 5, lib.Second())
}

func TestSubtractMilliseconds(t *testing.T) {
	assert := assert.New(t)

	testTime := time.Date(2017, 1, 1, 15, 0, 30, 0, chicagoLocation())

	lib := simpleTime(testTime)
	lib.Subtract(6000, "ms")

	assert.Equal(24, lib.Second())
	assert.Equal(24000, lib.Millisecond())
}

func TestSubtractNanoseconds(t *testing.T) {
	testTime := time.Date(2017, 1, 1, 15, 0, 0, 1000, chicagoLocation())

	lib := simpleTime(testTime)
	lib.Subtract(600, "nanoseconds")

	assert.Equal(t, 400, lib.Nanosecond())
}
