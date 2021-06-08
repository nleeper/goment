package goment

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestIsGoment(t *testing.T) {
	assert := assert.New(t)

	testTime := time.Now()

	lib := simpleTime(testTime)

	assert.True(IsGoment(lib))
	assert.True(IsGoment(*lib))
	assert.False(IsGoment("hi!"))
	assert.False(IsGoment(testTime))
}

func TestIsTime(t *testing.T) {
	assert := assert.New(t)

	testTime := time.Now()

	lib := simpleTime(testTime)

	assert.True(IsTime(testTime))
	assert.False(IsTime("hi!"))
	assert.False(IsTime(lib))
}

func TestIsDST(t *testing.T) {
	assert := assert.New(t)

	location, _ := time.LoadLocation("America/Chicago")

	date := time.Date(2011, 3, 12, 0, 0, 0, 0, location)

	lib := simpleTime(date)
	assert.False(lib.IsDST())

	date = time.Date(2011, 3, 14, 0, 0, 0, 0, location)

	lib2 := simpleTime(date)
	assert.True(lib2.IsDST())

	lib3 := simpleTime(time.Date(2017, 5, 3, 0, 0, 0, 0, location))
	assert.False(lib3.SetMonth(1).IsDST())
	assert.True(lib3.SetMonth(7).IsDST())
	assert.False(lib3.SetMonth(12).IsDST())
}

func TestIsLeapYear(t *testing.T) {
	assert := assert.New(t)

	lib := simpleString("2000-06-06")
	assert.True(lib.IsLeapYear())

	lib2 := simpleString("2010-05-07")
	assert.False(lib2.IsLeapYear())
}
