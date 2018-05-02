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
		assert.Equal(t, lib.DateTime, testTime)
	}
}

func TestNewFromISOString(t *testing.T) {
	date := "2011-05-13"
	testTime := time.Date(2011, 5, 13, 0, 0, 0, 0, time.UTC)

	lib, err := New(date)
	if assert.NoError(t, err) {
		assert.Equal(t, lib.DateTime, testTime)
	}
}

func TestNewFromTime(t *testing.T) {
	testTime := time.Date(2015, 11, 10, 5, 30, 0, 0, time.UTC)

	lib, err := New(testTime)
	if assert.NoError(t, err) {
		assert.Equal(t, lib.DateTime, testTime)
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
