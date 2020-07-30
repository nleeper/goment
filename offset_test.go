package goment

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestLocalUsesLocalTime(t *testing.T) {
	assert := assert.New(t)

	testTime := time.Date(2011, 5, 13, 14, 0, 0, 0, time.UTC)

	lib := simpleTime(testTime)

	assert.Equal(time.UTC, lib.ToTime().Location())

	lib.Local()
	assert.Equal(time.Local, lib.ToTime().Location())
}

func TestUTCUsesUTCTime(t *testing.T) {
	assert := assert.New(t)

	testTime := time.Date(2011, 5, 13, 14, 0, 0, 0, chicagoLocation())

	lib := simpleTime(testTime)

	assert.Equal(chicagoLocation(), lib.ToTime().Location())

	lib.UTC()
	assert.Equal(time.UTC, lib.ToTime().Location())
}

func TestUTCOffsetForLocal(t *testing.T) {
	testTime := time.Date(2011, 5, 13, 14, 0, 0, 0, chicagoLocation())

	lib := simpleTime(testTime)

	_, o := testTime.Zone()
	assert.Equal(t, o/60, lib.UTCOffset())
}

func TestUTCOffsetAfterSet(t *testing.T) {
	assert := assert.New(t)

	testTime := time.Date(2011, 5, 13, 14, 0, 0, 0, time.UTC)

	lib := simpleTime(testTime)

	assert.Equal(time.UTC, lib.ToTime().Location())

	lib.SetUTCOffset(-120)
	assert.Equal(12, lib.Hour())
	assert.Equal(-120, lib.UTCOffset())
}

func TestSetUTCOffsetInHours(t *testing.T) {
	assert := assert.New(t)

	testTime := time.Date(2011, 5, 13, 14, 0, 0, 0, time.UTC)

	lib := simpleTime(testTime)

	assert.Equal(time.UTC, lib.ToTime().Location())

	lib.SetUTCOffset(5)
	assert.Equal(19, lib.Hour())
}

func TestSetUTCOffsetInMinutes(t *testing.T) {
	assert := assert.New(t)

	testTime := time.Date(2011, 5, 13, 14, 0, 0, 0, time.UTC)

	lib := simpleTime(testTime)

	assert.Equal(time.UTC, lib.ToTime().Location())

	lib.SetUTCOffset(-120)
	assert.Equal(12, lib.Hour())
}
