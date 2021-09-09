package goment

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestStartOfYear(t *testing.T) {
	start := time.Date(2012, 1, 1, 0, 0, 0, 0, time.UTC)

	lib := simpleString("2012-05-09 14:17:12")

	lib.StartOf("year")
	assert.Equal(t, simpleTime(start).Format(), lib.Format())
}

func TestStartOfQuarter(t *testing.T) {
	start := time.Date(2012, 7, 1, 0, 0, 0, 0, time.UTC)

	lib := simpleString("2012-09-15 15:10:55")

	lib.StartOf("quarter")
	assert.Equal(t, simpleTime(start).Format(), lib.Format())
}

func TestStartOfMonth(t *testing.T) {
	start := time.Date(2012, 2, 1, 0, 0, 0, 0, time.UTC)

	lib := simpleString("2012-02-28 15:10:55")

	lib.StartOf("month")
	assert.Equal(t, simpleTime(start).Format(), lib.Format())
}

func TestStartOfWeek(t *testing.T) {
	assert.Equal(t, "Sunday", simpleNow().StartOf("week").Format("dddd"))
}

func TestStartOfISOWeek(t *testing.T) {
	start := time.Date(2021, 4, 5, 0, 0, 0, 0, time.UTC)

	lib := simpleString("2021-04-11 13:45:12")

	lib.StartOf("isoWeek")
	assert.Equal(t, simpleTime(start).Format(), lib.Format())
}

func TestStartOfDay(t *testing.T) {
	start := time.Date(2012, 2, 28, 0, 0, 0, 0, time.UTC)

	lib := simpleString("2012-02-28 15:10:55")

	lib.StartOf("day")
	assert.Equal(t, simpleTime(start).Format(), lib.Format())
}

func TestStartOfDate(t *testing.T) {
	start := time.Date(2012, 2, 15, 0, 0, 0, 0, time.UTC)

	lib := simpleString("2012-02-15 15:10:55")

	lib.StartOf("date")
	assert.Equal(t, simpleTime(start).Format(), lib.Format())
}

func TestStartOfHour(t *testing.T) {
	start := time.Date(2012, 2, 15, 15, 0, 0, 0, time.UTC)

	lib := simpleString("2012-02-15 15:10:55.123456")

	lib.StartOf("hour")
	assert.Equal(t, simpleTime(start).Format(), lib.Format())
}

func TestStartOfMinute(t *testing.T) {
	start := time.Date(2012, 2, 15, 15, 10, 0, 0, time.UTC)

	lib := simpleString("2012-02-15 15:10:55.123456")

	lib.StartOf("minute")
	assert.Equal(t, simpleTime(start).Format(), lib.Format())
}

func TestStartOfSecond(t *testing.T) {
	start := time.Date(2012, 2, 15, 15, 10, 55, 0, time.UTC)

	lib := simpleString("2012-02-15 15:10:55.123456")

	lib.StartOf("second")
	assert.Equal(t, simpleTime(start).Format(), lib.Format())
}

func TestEndOfYear(t *testing.T) {
	end := time.Date(2012, 12, 31, 23, 59, 59, 999999999, time.UTC)

	lib := simpleString("2012-05-09 14:17:12")

	lib.EndOf("year")
	assert.Equal(t, simpleTime(end).Format(), lib.Format())
}

func TestEndOfQuarter(t *testing.T) {
	end := time.Date(2012, 9, 30, 23, 59, 59, 999999999, time.UTC)

	lib := simpleString("2012-07-15 15:10:55")

	lib.EndOf("quarter")
	assert.Equal(t, simpleTime(end).Format(), lib.Format())
}

func TestEndOfMonth(t *testing.T) {
	end := time.Date(2011, 2, 28, 23, 59, 59, 999999999, time.UTC)

	lib := simpleString("2011-02-01 15:10:55")

	lib.EndOf("month")
	assert.Equal(t, simpleTime(end).Format(), lib.Format())
}

func TestEndOfWeek(t *testing.T) {
	assert.Equal(t, "Saturday", simpleNow().EndOf("week").Format("dddd"))

	end := time.Date(2021, 9, 11, 23, 59, 59, 999999999, time.UTC)
	lib := simpleString("2021-09-09 13:45:12")
	lib.EndOf("week")

	assert.Equal(t, simpleTime(end).Format(), lib.Format())
}

func TestEndOfISOWeek(t *testing.T) {
	assert.Equal(t, "Sunday", simpleNow().EndOf("isoWeek").Format("dddd"))

	end := time.Date(2021, 9, 12, 23, 59, 59, 999999999, time.UTC)

	lib := simpleString("2021-09-09 13:45:12")
	lib.EndOf("isoWeek")

	assert.Equal(t, simpleTime(end).Format(), lib.Format())
}

func TestEndOfDay(t *testing.T) {
	end := time.Date(2012, 2, 28, 23, 59, 59, 999999999, time.UTC)

	lib := simpleString("2012-02-28 15:10:55")

	lib.EndOf("day")
	assert.Equal(t, simpleTime(end).Format(), lib.Format())
}

func TestEndOfDate(t *testing.T) {
	end := time.Date(2012, 2, 15, 23, 59, 59, 999999999, time.UTC)

	lib := simpleString("2012-02-15 15:10:55")

	lib.EndOf("date")
	assert.Equal(t, simpleTime(end).Format(), lib.Format())
}

func TestEndOfHour(t *testing.T) {
	end := time.Date(2012, 2, 15, 15, 59, 59, 999999999, time.UTC)

	lib := simpleString("2012-02-15 15:10:55.123456")

	lib.EndOf("hour")
	assert.Equal(t, simpleTime(end).Format(), lib.Format())
}

func TestEndOfMinute(t *testing.T) {
	end := time.Date(2012, 2, 15, 15, 10, 59, 999999999, time.UTC)

	lib := simpleString("2012-02-15 15:10:55.123456")

	lib.EndOf("minute")
	assert.Equal(t, simpleTime(end).Format(), lib.Format())
}

func TestEndOfSecond(t *testing.T) {
	end := time.Date(2012, 2, 15, 15, 10, 55, 999999999, time.UTC)

	lib := simpleString("2012-02-15 15:10:55.123456")

	lib.EndOf("second")
	assert.Equal(t, simpleTime(end).Format(), lib.Format())
}
