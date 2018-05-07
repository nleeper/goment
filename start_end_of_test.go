package goment

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestStartOfYear(t *testing.T) {
	start := time.Date(2012, 1, 1, 0, 0, 0, 0, time.UTC)

	lib, err := New("2012-05-09 14:17:12")
	if assert.NoError(t, err) {
		lib.StartOf("year")
		assert.True(t, start.Equal(lib.ToTime()))
	}
}

func TestStartOfQuarter(t *testing.T) {
	start := time.Date(2012, 7, 1, 0, 0, 0, 0, time.UTC)

	lib, err := New("2012-09-15 15:10:55")
	if assert.NoError(t, err) {
		lib.StartOf("quarter")
		assert.True(t, start.Equal(lib.ToTime()))
	}
}

func TestStartOfMonth(t *testing.T) {
	start := time.Date(2012, 2, 1, 0, 0, 0, 0, time.UTC)

	lib, err := New("2012-02-28 15:10:55")
	if assert.NoError(t, err) {
		lib.StartOf("month")
		assert.True(t, start.Equal(lib.ToTime()))
	}
}

func TestStartOfISOWeek(t *testing.T) {
	start := time.Date(2017, 9, 4, 0, 0, 0, 0, time.UTC)

	lib, err := New("2017-09-07 13:45:12")
	if assert.NoError(t, err) {
		lib.StartOf("isoWeek")
		assert.True(t, start.Equal(lib.ToTime()))
	}
}

func TestStartOfDay(t *testing.T) {
	start := time.Date(2012, 2, 28, 0, 0, 0, 0, time.UTC)

	lib, err := New("2012-02-28 15:10:55")
	if assert.NoError(t, err) {
		lib.StartOf("day")
		assert.True(t, start.Equal(lib.ToTime()))
	}
}

func TestStartOfDate(t *testing.T) {
	start := time.Date(2012, 2, 15, 0, 0, 0, 0, time.UTC)

	lib, err := New("2012-02-15 15:10:55")
	if assert.NoError(t, err) {
		lib.StartOf("date")
		assert.True(t, start.Equal(lib.ToTime()))
	}
}

func TestStartOfHour(t *testing.T) {
	start := time.Date(2012, 2, 15, 15, 0, 0, 0, time.UTC)

	lib, err := New("2012-02-15 15:10:55.123456")
	if assert.NoError(t, err) {
		lib.StartOf("hour")
		assert.True(t, start.Equal(lib.ToTime()))
	}
}

func TestStartOfMinute(t *testing.T) {
	start := time.Date(2012, 2, 15, 15, 10, 0, 0, time.UTC)

	lib, err := New("2012-02-15 15:10:55.123456")
	if assert.NoError(t, err) {
		lib.StartOf("minute")
		assert.True(t, start.Equal(lib.ToTime()))
	}
}

func TestStartOfSecond(t *testing.T) {
	start := time.Date(2012, 2, 15, 15, 10, 55, 0, time.UTC)

	lib, err := New("2012-02-15 15:10:55.123456")
	if assert.NoError(t, err) {
		lib.StartOf("second")
		assert.True(t, start.Equal(lib.ToTime()))
	}
}

func TestEndOfYear(t *testing.T) {
	end := time.Date(2012, 12, 31, 23, 59, 59, 999999999, time.UTC)

	lib, err := New("2012-05-09 14:17:12")
	if assert.NoError(t, err) {
		lib.EndOf("year")
		assert.True(t, end.Equal(lib.ToTime()))
	}
}

func TestEndOfQuarter(t *testing.T) {
	end := time.Date(2012, 9, 30, 23, 59, 59, 999999999, time.UTC)

	lib, err := New("2012-07-15 15:10:55")
	if assert.NoError(t, err) {
		lib.EndOf("quarter")
		assert.True(t, end.Equal(lib.ToTime()))
	}
}

func TestEndOfMonth(t *testing.T) {
	end := time.Date(2011, 2, 28, 23, 59, 59, 999999999, time.UTC)

	lib, err := New("2011-02-01 15:10:55")
	if assert.NoError(t, err) {
		lib.EndOf("month")
		assert.True(t, end.Equal(lib.ToTime()))
	}
}

func TestEndOfISOWeek(t *testing.T) {
	end := time.Date(2017, 9, 9, 0, 0, 0, 0, time.UTC)

	lib, err := New("2017-09-03 13:45:12")
	if assert.NoError(t, err) {
		lib.EndOf("isoWeek")
		assert.True(t, end.Equal(lib.ToTime()))
	}
}

func TestEndOfDay(t *testing.T) {
	end := time.Date(2012, 2, 28, 23, 59, 59, 999999999, time.UTC)

	lib, err := New("2012-02-28 15:10:55")
	if assert.NoError(t, err) {
		lib.EndOf("day")
		assert.True(t, end.Equal(lib.ToTime()))
	}
}

func TestEndOfDate(t *testing.T) {
	end := time.Date(2012, 2, 15, 23, 59, 59, 999999999, time.UTC)

	lib, err := New("2012-02-15 15:10:55")
	if assert.NoError(t, err) {
		lib.EndOf("date")
		assert.True(t, end.Equal(lib.ToTime()))
	}
}

func TestEndOfHour(t *testing.T) {
	end := time.Date(2012, 2, 15, 15, 59, 59, 999999999, time.UTC)

	lib, err := New("2012-02-15 15:10:55.123456")
	if assert.NoError(t, err) {
		lib.EndOf("hour")
		assert.True(t, end.Equal(lib.ToTime()))
	}
}

func TestEndOfMinute(t *testing.T) {
	start := time.Date(2012, 2, 15, 15, 10, 59, 999999999, time.UTC)

	lib, err := New("2012-02-15 15:10:55.123456")
	if assert.NoError(t, err) {
		lib.EndOf("minute")
		assert.True(t, start.Equal(lib.ToTime()))
	}
}

func TestEndOfSecond(t *testing.T) {
	end := time.Date(2012, 2, 15, 15, 10, 55, 999999999, time.UTC)

	lib, err := New("2012-02-15 15:10:55.123456")
	if assert.NoError(t, err) {
		lib.EndOf("second")
		assert.True(t, end.Equal(lib.ToTime()))
	}
}
