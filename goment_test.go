package goment

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewNoDate(t *testing.T) {
	testTime := time.Date(2000, 12, 15, 17, 8, 0, 0, time.UTC)
	timeNow = func() time.Time {
		return testTime
	}

	lib, err := New()
	if assert.NoError(t, err) {
		assert.True(t, testTime.Local().Equal(lib.DateTime))
		assert.True(t, lib.DateTime.Location() == time.Local)
	}
}

func TestNewFromISOString(t *testing.T) {
	date := "2011-05-13"
	testTime := time.Date(2011, 5, 13, 0, 0, 0, 0, time.UTC)

	lib, err := New(date)
	if assert.NoError(t, err) {
		assert.True(t, testTime.Equal(lib.DateTime))
		assert.True(t, lib.DateTime.Location() == time.UTC)
	}
}

func TestNewFromTime(t *testing.T) {
	testTime := time.Date(2015, 11, 10, 5, 30, 0, 0, time.UTC)

	lib, err := New(testTime)
	if assert.NoError(t, err) {
		assert.True(t, testTime.Equal(lib.DateTime))
		assert.True(t, lib.DateTime.Location() == time.Local)
	}
}

func TestNewFromUnixMilliseconds(t *testing.T) {
	testTime := time.Now()

	lib, err := New(testTime.UnixNano())
	if assert.NoError(t, err) {
		assert.True(t, testTime.Equal(lib.DateTime))
		assert.True(t, lib.DateTime.Location() == time.Local)
	}
}

func TestNewFromTimeConvertUTCToLocal(t *testing.T) {
	testNow := time.Now()

	lib, err := New(testNow.UTC())
	if assert.NoError(t, err) {
		assert.Equal(t, lib.DateTime, testNow.Local())
		assert.True(t, lib.DateTime.Location() == time.Local)
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
	_, err := New("2018-01-01", "YYYY-MM-DD")
	assert.EqualError(t, err, "Invalid number of arguments")
}

func TestUnixFromSeconds(t *testing.T) {
	// time.Unix does not have microsecond info, so we must
	// compare the Unix versions of the time.
	testTime := time.Now()
	testTimeUnix := time.Unix(testTime.Unix(), 0)

	lib, err := Unix(testTime.Unix())
	if assert.NoError(t, err) {
		assert.True(t, testTimeUnix.Equal(lib.DateTime))
		assert.True(t, lib.DateTime.Location() == time.Local)
	}
}
