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
	lib, err := New(Time{
		Year:  2011,
		Month: 3,
		Day:   12,
	})

	if assert.NoError(t, err) {
		assert.False(t, lib.IsDST())
	}

	lib, err = New(Time{
		Year:  2011,
		Month: 3,
		Day:   14,
	})

	if assert.NoError(t, err) {
		assert.True(t, lib.IsDST())
	}

	lib, err = New()
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
