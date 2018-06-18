package goment

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestLocalUsesLocalTime(t *testing.T) {
	testTime := time.Date(2011, 5, 13, 14, 0, 0, 0, time.UTC)

	lib, err := New(testTime)
	if assert.NoError(t, err) {
		assert.Equal(t, lib.ToTime().Location(), time.UTC)
		lib.Local()
		assert.Equal(t, lib.ToTime().Location(), time.Local)
	}
}

func TestUTCUsesUTCTime(t *testing.T) {
	testTime := time.Date(2011, 5, 13, 14, 0, 0, 0, chicagoLocation())

	lib, err := New(testTime)
	if assert.NoError(t, err) {
		assert.Equal(t, lib.ToTime().Location(), chicagoLocation())
		lib.UTC()
		assert.Equal(t, lib.ToTime().Location(), time.UTC)
	}
}

func TestUTCOffsetForLocal(t *testing.T) {
	testTime := time.Date(2011, 5, 13, 14, 0, 0, 0, chicagoLocation())

	lib, err := New(testTime)
	if assert.NoError(t, err) {
		_, o := testTime.Zone()
		assert.Equal(t, lib.UTCOffset(), o/60)
	}
}

func TestUTCOffsetAfterSet(t *testing.T) {
	testTime := time.Date(2011, 5, 13, 14, 0, 0, 0, time.UTC)

	lib, err := New(testTime)
	if assert.NoError(t, err) {
		assert.Equal(t, lib.ToTime().Location(), time.UTC)
		lib.SetUTCOffset(-120)
		assert.Equal(t, lib.Hour(), 12)
		assert.Equal(t, lib.UTCOffset(), -120)
	}
}

func TestSetUTCOffsetInHours(t *testing.T) {
	testTime := time.Date(2011, 5, 13, 14, 0, 0, 0, time.UTC)

	lib, err := New(testTime)
	if assert.NoError(t, err) {
		assert.Equal(t, lib.ToTime().Location(), time.UTC)
		lib.SetUTCOffset(5)
		assert.Equal(t, lib.Hour(), 19)
	}
}

func TestSetUTCOffsetInMinutes(t *testing.T) {
	testTime := time.Date(2011, 5, 13, 14, 0, 0, 0, time.UTC)

	lib, err := New(testTime)
	if assert.NoError(t, err) {
		assert.Equal(t, lib.ToTime().Location(), time.UTC)
		lib.SetUTCOffset(-120)
		assert.Equal(t, lib.Hour(), 12)
	}
}
