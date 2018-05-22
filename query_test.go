package goment

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestIsGoment(t *testing.T) {
	testTime := time.Now()

	lib, err := New(testTime)
	if assert.NoError(t, err) {
		assert.True(t, IsGoment(lib))
		assert.False(t, IsGoment("hi!"))
		assert.False(t, IsGoment(testTime))
	}
}

func TestIsTime(t *testing.T) {
	testTime := time.Now()

	lib, err := New(testTime)
	if assert.NoError(t, err) {
		assert.True(t, IsTime(testTime))
		assert.False(t, IsTime("hi!"))
		assert.False(t, IsTime(lib))
	}
}

func TestIsDST(t *testing.T) {
	location, _ := time.LoadLocation("America/Chicago")

	date := time.Date(2011, 3, 12, 0, 0, 0, 0, location)

	lib, err := New(date)
	if assert.NoError(t, err) {
		assert.False(t, lib.IsDST())
	}

	date = time.Date(2011, 3, 14, 0, 0, 0, 0, location)

	lib, err = New(date)
	if assert.NoError(t, err) {
		assert.True(t, lib.IsDST())
	}

	lib, err = New(time.Date(2017, 5, 3, 0, 0, 0, 0, location))
	if assert.NoError(t, err) {
		assert.False(t, lib.SetMonth(1).IsDST())
		assert.True(t, lib.SetMonth(7).IsDST())
		assert.False(t, lib.SetMonth(12).IsDST())
	}
}

func TestIsLeapYear(t *testing.T) {
	lib, err := New("2000-06-06")
	if assert.NoError(t, err) {
		assert.True(t, lib.IsLeapYear())
	}

	lib, err = New("2010-05-07")
	if assert.NoError(t, err) {
		assert.False(t, lib.IsLeapYear())
	}
}
