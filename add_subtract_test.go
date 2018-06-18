package goment

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAddNoArgsIgnored(t *testing.T) {
	testTime := time.Date(2017, 1, 1, 18, 0, 0, 0, chicagoLocation())

	lib, err := New(testTime)
	if assert.NoError(t, err) {
		lib.Add()
		assert.True(t, testTime.Equal(lib.ToTime()))
	}
}

func TestAddTooManyArgsIgnored(t *testing.T) {
	testTime := time.Date(2017, 1, 1, 18, 0, 0, 0, chicagoLocation())

	lib, err := New(testTime)
	if assert.NoError(t, err) {
		lib.Add(1, "year", "day")
		assert.True(t, testTime.Equal(lib.ToTime()))
	}
}

func TestAddInvalidFirstArgTypeIgnored(t *testing.T) {
	testTime := time.Date(2017, 1, 1, 18, 0, 0, 0, chicagoLocation())

	lib, err := New(testTime)
	if assert.NoError(t, err) {
		lib.Add("year")
		assert.True(t, testTime.Equal(lib.ToTime()))
	}
}

func TestAddInvalidSecondArgTypeIgnored(t *testing.T) {
	testTime := time.Date(2017, 1, 1, 18, 0, 0, 0, chicagoLocation())

	lib, err := New(testTime)
	if assert.NoError(t, err) {
		lib.Add(1, 2)
		assert.True(t, testTime.Equal(lib.ToTime()))
	}
}

func TestAddInvalidSecondArgValueIgnored(t *testing.T) {
	testTime := time.Date(2017, 1, 1, 18, 0, 0, 0, chicagoLocation())

	lib, err := New(testTime)
	if assert.NoError(t, err) {
		lib.Add(1, "custom")
		assert.True(t, testTime.Equal(lib.ToTime()))
	}
}

func TestAddDuration(t *testing.T) {
	lib, err := New("2011-01-04 14:01:01")
	if assert.NoError(t, err) {
		lib.Add(time.Duration(4) * time.Hour)
		assert.Equal(t, lib.Hour(), 18)
		lib.Add(time.Duration(25) * time.Minute)
		assert.Equal(t, lib.Minute(), 26)
	}
}

func TestAddYears(t *testing.T) {
	testTime := time.Date(2017, 1, 1, 18, 0, 0, 0, chicagoLocation())

	lib, err := New(testTime)
	if assert.NoError(t, err) {
		lib.Add(1, "year")
		assert.Equal(t, lib.Year(), 2018)
	}
}

func TestAddQuarters(t *testing.T) {
	testTime := time.Date(2014, 4, 11, 0, 0, 0, 0, chicagoLocation())

	lib, err := New(testTime)
	if assert.NoError(t, err) {
		assert.Equal(t, lib.Quarter(), 2)
		lib.Add(2, "quarters")
		assert.Equal(t, lib.Month(), 10)
	}
}

func TestAddMonths(t *testing.T) {
	testTime := time.Date(2017, 1, 1, 18, 0, 0, 0, chicagoLocation())

	lib, err := New(testTime)
	if assert.NoError(t, err) {
		lib.Add(5, "months")
		assert.Equal(t, lib.Year(), 2017)
		assert.Equal(t, lib.Month(), 6)
	}
}

func TestAddWeeks(t *testing.T) {
	testTime := time.Date(2017, 1, 1, 18, 0, 0, 0, chicagoLocation())

	lib, err := New(testTime)
	if assert.NoError(t, err) {
		lib.Add(3, "weeks")
		assert.Equal(t, lib.Year(), 2017)
		assert.Equal(t, lib.Month(), 1)
		assert.Equal(t, lib.Date(), 22)
	}
}

func TestAddDays(t *testing.T) {
	testTime := time.Date(2017, 1, 1, 18, 0, 0, 0, chicagoLocation())

	lib, err := New(testTime)
	if assert.NoError(t, err) {
		lib.Add(16, "d")
		assert.Equal(t, lib.Year(), 2017)
		assert.Equal(t, lib.Month(), 1)
		assert.Equal(t, lib.Date(), 17)
	}
}

func TestAddHours(t *testing.T) {
	testTime := time.Date(2017, 1, 1, 15, 0, 0, 0, chicagoLocation())

	lib, err := New(testTime)
	if assert.NoError(t, err) {
		lib.Add(5, "hours")
		assert.Equal(t, lib.Hour(), 20)
	}
}

func TestAddMinutes(t *testing.T) {
	testTime := time.Date(2017, 1, 1, 15, 30, 0, 0, chicagoLocation())

	lib, err := New(testTime)
	if assert.NoError(t, err) {
		lib.Add(15, "m")
		assert.Equal(t, lib.Minute(), 45)
	}
}

func TestAddSeconds(t *testing.T) {
	testTime := time.Date(2017, 1, 1, 15, 0, 0, 0, chicagoLocation())

	lib, err := New(testTime)
	if assert.NoError(t, err) {
		lib.Add(55, "seconds")
		assert.Equal(t, lib.Second(), 55)
	}
}

func TestAddMilliseconds(t *testing.T) {
	testTime := time.Date(2017, 1, 1, 15, 0, 0, 0, chicagoLocation())

	lib, err := New(testTime)
	if assert.NoError(t, err) {
		lib.Add(6000, "ms")
		assert.Equal(t, lib.Second(), 6)
		assert.Equal(t, lib.Millisecond(), 6000)
	}
}

func TestAddNanoseconds(t *testing.T) {
	testTime := time.Date(2017, 1, 1, 15, 0, 0, 0, chicagoLocation())

	lib, err := New(testTime)
	if assert.NoError(t, err) {
		lib.Add(600, "nanoseconds")
		assert.Equal(t, lib.Nanosecond(), 600)
	}
}

func TestSubtractNoArgsIgnored(t *testing.T) {
	testTime := time.Date(2017, 1, 1, 18, 0, 0, 0, chicagoLocation())

	lib, err := New(testTime)
	if assert.NoError(t, err) {
		lib.Subtract()
		assert.True(t, testTime.Equal(lib.ToTime()))
	}
}

func TestSubtractTooManyArgsIgnored(t *testing.T) {
	testTime := time.Date(2017, 1, 1, 18, 0, 0, 0, chicagoLocation())

	lib, err := New(testTime)
	if assert.NoError(t, err) {
		lib.Subtract(1, "year", "day")
		assert.True(t, testTime.Equal(lib.ToTime()))
	}
}

func TestSubtractInvalidFirstArgTypeIgnored(t *testing.T) {
	testTime := time.Date(2017, 1, 1, 18, 0, 0, 0, chicagoLocation())

	lib, err := New(testTime)
	if assert.NoError(t, err) {
		lib.Subtract("year")
		assert.True(t, testTime.Equal(lib.ToTime()))
	}
}

func TestSubtractInvalidSecondArgTypeIgnored(t *testing.T) {
	testTime := time.Date(2017, 1, 1, 18, 0, 0, 0, chicagoLocation())

	lib, err := New(testTime)
	if assert.NoError(t, err) {
		lib.Subtract(1, 2)
		assert.True(t, testTime.Equal(lib.ToTime()))
	}
}

func TestSubtractInvalidSecondArgValueIgnored(t *testing.T) {
	testTime := time.Date(2017, 1, 1, 18, 0, 0, 0, chicagoLocation())

	lib, err := New(testTime)
	if assert.NoError(t, err) {
		lib.Subtract(1, "custom")
		assert.True(t, testTime.Equal(lib.ToTime()))
	}
}

func TestSubtractDuration(t *testing.T) {
	lib, err := New("2011-01-04 14:30:01")
	if assert.NoError(t, err) {
		lib.Subtract(time.Duration(4) * time.Hour)
		assert.Equal(t, lib.Hour(), 10)
		lib.Subtract(time.Duration(25) * time.Minute)
		assert.Equal(t, lib.Minute(), 5)
	}
}

func TestSubtractYears(t *testing.T) {
	testTime := time.Date(2017, 1, 1, 18, 0, 0, 0, chicagoLocation())

	lib, err := New(testTime)
	if assert.NoError(t, err) {
		lib.Subtract(1, "year")
		assert.Equal(t, lib.Year(), 2016)
	}
}

func TestSubtractQuarters(t *testing.T) {
	testTime := time.Date(2014, 8, 11, 0, 0, 0, 0, chicagoLocation())

	lib, err := New(testTime)
	if assert.NoError(t, err) {
		assert.Equal(t, lib.Quarter(), 3)
		lib.Subtract(2, "quarters")
		assert.Equal(t, lib.Month(), 2)
	}
}

func TestSubtractMonths(t *testing.T) {
	testTime := time.Date(2017, 11, 1, 18, 0, 0, 0, chicagoLocation())

	lib, err := New(testTime)
	if assert.NoError(t, err) {
		lib.Subtract(5, "months")
		assert.Equal(t, lib.Year(), 2017)
		assert.Equal(t, lib.Month(), 6)
	}
}

func TestSubtractWeeks(t *testing.T) {
	testTime := time.Date(2017, 1, 22, 18, 0, 0, 0, chicagoLocation())

	lib, err := New(testTime)
	if assert.NoError(t, err) {
		lib.Subtract(3, "weeks")
		assert.Equal(t, lib.Year(), 2017)
		assert.Equal(t, lib.Month(), 1)
		assert.Equal(t, lib.Date(), 1)
	}
}

func TestSubtractDays(t *testing.T) {
	testTime := time.Date(2017, 1, 17, 18, 0, 0, 0, chicagoLocation())

	lib, err := New(testTime)
	if assert.NoError(t, err) {
		lib.Subtract(16, "d")
		assert.Equal(t, lib.Year(), 2017)
		assert.Equal(t, lib.Month(), 1)
		assert.Equal(t, lib.Date(), 1)
	}
}

func TestSubtractHours(t *testing.T) {
	testTime := time.Date(2017, 1, 1, 15, 0, 0, 0, chicagoLocation())

	lib, err := New(testTime)
	if assert.NoError(t, err) {
		lib.Subtract(5, "hours")
		assert.Equal(t, lib.Hour(), 10)
	}
}

func TestSubtractMinutes(t *testing.T) {
	testTime := time.Date(2017, 1, 1, 15, 30, 0, 0, chicagoLocation())

	lib, err := New(testTime)
	if assert.NoError(t, err) {
		lib.Subtract(15, "m")
		assert.Equal(t, lib.Minute(), 15)
	}
}

func TestSubtractSeconds(t *testing.T) {
	testTime := time.Date(2017, 1, 1, 15, 0, 30, 0, chicagoLocation())

	lib, err := New(testTime)
	if assert.NoError(t, err) {
		lib.Subtract(25, "seconds")
		assert.Equal(t, lib.Second(), 5)
	}
}

func TestSubtractMilliseconds(t *testing.T) {
	testTime := time.Date(2017, 1, 1, 15, 0, 30, 0, chicagoLocation())

	lib, err := New(testTime)
	if assert.NoError(t, err) {
		lib.Subtract(6000, "ms")
		assert.Equal(t, lib.Second(), 24)
		assert.Equal(t, lib.Millisecond(), 24000)
	}
}

func TestSubtractNanoseconds(t *testing.T) {
	testTime := time.Date(2017, 1, 1, 15, 0, 0, 1000, chicagoLocation())

	lib, err := New(testTime)
	if assert.NoError(t, err) {
		lib.Subtract(600, "nanoseconds")
		assert.Equal(t, lib.Nanosecond(), 400)
	}
}
