package goment

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewNoDate(t *testing.T) {
	testTime := time.Date(2000, 12, 15, 17, 8, 0, 0, time.Local)
	timeNow = func() time.Time {
		return testTime
	}

	lib, err := New()
	if assert.NoError(t, err) {
		assert.True(t, testTime.Local().Equal(lib.ToTime()))
		assert.True(t, lib.ToTime().Location() == time.Local)
	}

	// Reset timeNow.
	timeNow = time.Now
}

func TestNewFromISOString(t *testing.T) {
	date := "2011-05-13"
	testTime := time.Date(2011, 5, 13, 0, 0, 0, 0, time.UTC)

	lib, err := New(date)
	if assert.NoError(t, err) {
		assert.True(t, testTime.Equal(lib.ToTime()))
		assert.True(t, lib.ToTime().Location() == time.UTC)
	}
}

func TestNewFromFormat(t *testing.T) {
	date := "05-13-2011"
	testTime := time.Date(2011, 5, 13, 0, 0, 0, 0, time.Local)

	lib, err := New(date, "MM-DD-YYYY")
	if assert.NoError(t, err) {
		assert.True(t, testTime.Equal(lib.ToTime()))
	}
}

func TestNewFromTime(t *testing.T) {
	testTime := time.Date(2015, 11, 10, 5, 30, 0, 0, time.UTC)

	lib, err := New(testTime)
	if assert.NoError(t, err) {
		assert.True(t, testTime.Equal(lib.ToTime()))
		assert.True(t, lib.ToTime().Location() == time.UTC)
	}
}

func TestNewFromDateTime(t *testing.T) {
	testGomentTime := DateTime{
		Year:  2015,
		Month: 1,
		Day:   25,
		Hour:  10,
	}

	testTime := time.Date(2015, 1, 25, 10, 0, 0, 0, time.Local)

	lib, err := New(testGomentTime)
	if assert.NoError(t, err) {
		assert.True(t, testTime.Equal(lib.ToTime()))
		assert.True(t, lib.ToTime().Location() == time.Local)
	}
}

func TestNewFromUnixNanoseconds(t *testing.T) {
	testTime := time.Now()

	lib, err := New(testTime.UnixNano())
	if assert.NoError(t, err) {
		assert.True(t, testTime.Equal(lib.ToTime()))
		assert.True(t, lib.ToTime().Location() == time.Local)
	}
}

func TestNewFromGomentReturnsClone(t *testing.T) {
	testTime := time.Date(2011, 5, 13, 0, 0, 0, 0, time.UTC)

	lib, err := New(testTime)
	if assert.NoError(t, err) {
		clone, err2 := New(lib)
		if assert.NoError(t, err2) {
			assert.True(t, clone.ToTime().Equal(lib.ToTime()))

			clone.Add(1, "d")
			assert.False(t, clone.ToTime().Equal(lib.ToTime()))
		}
	}
}

func TestNewThrowErrorFromInvalidArgumentType(t *testing.T) {
	_, err := New(1)
	assert.EqualError(t, err, "Invalid argument type")
}

func TestNewThrowErrorFromInvalidISOString(t *testing.T) {
	_, err := New("2011-05a13")
	assert.EqualError(t, err, "Not a matching ISO-8601 date")
}

func TestNewThrowErrorIfTooManyArgs(t *testing.T) {
	_, err := New("2018-01-01", "YYYY-MM-DD", "third")
	assert.EqualError(t, err, "Invalid number of arguments")
}

func TestNewThrowErrorIfInvalidArgsForFormat(t *testing.T) {
	_, err := New(1, "YYYY-MM-DD")
	assert.EqualError(t, err, "First argument must be a datetime string")

	_, err = New("2018-01-01", 2)
	assert.EqualError(t, err, "Second argument must be a format string")
}

func TestCloneReturnsClone(t *testing.T) {
	testTime := time.Date(2011, 5, 13, 0, 0, 0, 0, time.UTC)

	lib, err := New(testTime)
	if assert.NoError(t, err) {
		clone := lib.Clone()
		assert.True(t, clone.ToTime().Equal(lib.ToTime()))

		clone.Add(1, "M")
		assert.False(t, clone.ToTime().Equal(lib.ToTime()))
	}
}

func TestUnixFromSeconds(t *testing.T) {
	// time.Unix does not have microsecond info, so we must
	// compare the Unix versions of the time.
	testTime := time.Now()
	testTimeUnix := time.Unix(testTime.Unix(), 0)

	lib, err := Unix(testTime.Unix())
	if assert.NoError(t, err) {
		assert.True(t, testTimeUnix.Equal(lib.ToTime()))
		assert.True(t, lib.ToTime().Location() == time.Local)
	}
}
