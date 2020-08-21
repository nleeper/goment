package goment

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewNoDate(t *testing.T) {
	assert := assert.New(t)

	testTime := time.Date(2000, 12, 15, 17, 8, 0, 0, time.Local)
	timeNow = func() time.Time {
		return testTime
	}

	lib := simpleNow()

	assert.Equal(simpleTime(testTime).Format(), lib.Format())
	assert.Equal(time.Local, lib.ToTime().Location())

	// Reset timeNow.
	timeNow = time.Now
}

func TestNewFromISOString(t *testing.T) {
	assert := assert.New(t)

	date := "2011-05-13"
	testTime := time.Date(2011, 5, 13, 0, 0, 0, 0, time.UTC)

	lib := simpleString(date)

	assert.Equal(simpleTime(testTime).Format(), lib.Format())
	assert.Equal(time.UTC, lib.ToTime().Location())
}

func TestNewFromFormat(t *testing.T) {
	date := "05-13-2011"
	testTime := time.Date(2011, 5, 13, 0, 0, 0, 0, time.Local)

	lib := simpleFormat(date, "MM-DD-YYYY")

	assert.Equal(t, simpleTime(testTime).Format(), lib.Format())
}

func TestNewFromFormatUsesGlobalLocale(t *testing.T) {
	assert := assert.New(t)

	date := "05-13-2011"
	testTime := time.Date(2011, 5, 13, 0, 0, 0, 0, time.Local)

	SetLocale("es")

	lib := simpleFormat(date, "MM-DD-YYYY")

	assert.Equal(simpleTime(testTime).Format(), lib.Format())
	assert.Equal("es", lib.Locale())

	SetLocale("en")
}

func TestNewFromFormatWithLocale(t *testing.T) {
	assert := assert.New(t)

	date := "05-13-2011"
	testTime := time.Date(2011, 5, 13, 0, 0, 0, 0, time.Local)

	lib := simpleFormatLocale(date, "MM-DD-YYYY", "fr")

	assert.Equal(simpleTime(testTime).Format(), lib.Format())
	assert.Equal("fr", lib.Locale())
}

func TestNewFromTime(t *testing.T) {
	assert := assert.New(t)

	testTime := time.Date(2015, 11, 10, 5, 30, 0, 0, time.UTC)

	lib := simpleTime(testTime)

	assert.Equal(simpleTime(testTime).Format(), lib.Format())
	assert.Equal(time.UTC, lib.ToTime().Location())
}

func TestNewFromDateTime(t *testing.T) {
	assert := assert.New(t)

	testGomentTime := DateTime{
		Year:   2015,
		Month:  1,
		Day:    25,
		Hour:   10,
		Minute: 30,
	}

	testTime := time.Date(2015, 1, 25, 10, 30, 0, 0, time.Local)

	lib := simple(testGomentTime)

	assert.Equal(simpleTime(testTime).Format(), lib.Format())
	assert.Equal(time.Local, lib.ToTime().Location())
}

func TestNewFromDateTimeOnlyYear(t *testing.T) {
	assert := assert.New(t)

	testGomentTime := DateTime{
		Year: 2015,
	}

	now := simpleNow()

	testTime := time.Date(2015, time.Month(now.Month()), now.Date(), 0, 0, 0, 0, time.Local)

	lib := simple(testGomentTime)

	assert.Equal(simpleTime(testTime).Format(), lib.Format())
	assert.Equal(time.Local, lib.ToTime().Location())
}

func TestNewFromDateTimeYearMonth(t *testing.T) {
	assert := assert.New(t)

	testGomentTime := DateTime{
		Year:  2015,
		Month: 2,
	}

	now := simpleNow()

	testTime := time.Date(2015, 2, now.Date(), 0, 0, 0, 0, time.Local)

	lib := simple(testGomentTime)

	assert.Equal(simpleTime(testTime).Format(), lib.Format())
	assert.Equal(time.Local, lib.ToTime().Location())
}

func TestNewFromDateTimeYearMonthDay(t *testing.T) {
	assert := assert.New(t)

	testGomentTime := DateTime{
		Year:  2015,
		Month: 2,
		Day:   3,
	}

	testTime := time.Date(2015, 2, 3, 0, 0, 0, 0, time.Local)

	lib := simple(testGomentTime)

	assert.Equal(simpleTime(testTime).Format(), lib.Format())
	assert.Equal(time.Local, lib.ToTime().Location())
}

func TestNewFromUnixNanoseconds(t *testing.T) {
	assert := assert.New(t)

	testTime := time.Now()

	lib := simpleUnixNano(testTime.UnixNano())

	assert.Equal(simpleTime(testTime).Format(), lib.Format())
	assert.Equal(time.Local, lib.ToTime().Location())
}

func TestNewFromGomentReturnsClone(t *testing.T) {
	assert := assert.New(t)

	testTime := time.Date(2011, 5, 13, 0, 0, 0, 0, time.UTC)

	lib := simpleTime(testTime)

	clone := simpleGoment(lib)

	assert.Equal(clone.Format(), lib.Format())

	clone.Add(1, "d")
	assert.NotEqual(clone.Format(), lib.Format())
}

func TestNewFromGomentReturnsCloneWithSetLocale(t *testing.T) {
	assert := assert.New(t)

	testTime := time.Date(2011, 5, 13, 0, 0, 0, 0, time.UTC)

	lib := simpleTime(testTime)
	lib.SetLocale("es")

	clone := simpleGoment(lib)

	assert.Equal(simpleTime(testTime).Format(), lib.Format())
	assert.Equal(lib.Locale(), clone.Locale())
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
	_, err := New("2018-01-01", "YYYY-MM-DD", "fr", "stuff")
	assert.EqualError(t, err, "Invalid number of arguments")
}

func TestNewThrowErrorIfInvalidArgsForFormat(t *testing.T) {
	assert := assert.New(t)

	_, err := New(1, "YYYY-MM-DD")
	assert.EqualError(err, "First argument must be a datetime string")

	_, err = New("2018-01-01", 2)
	assert.EqualError(err, "Second argument must be a format string")
}

func TestCloneReturnsClone(t *testing.T) {
	assert := assert.New(t)

	testTime := time.Date(2011, 5, 13, 0, 0, 0, 0, time.UTC)

	lib := simpleTime(testTime)

	clone := lib.Clone()
	assert.Equal(clone.Format(), lib.Format())

	clone.Add(1, "M")
	assert.NotEqual(clone.Format(), lib.Format())
}

func TestCloneReturnsCloneWithLocale(t *testing.T) {
	assert := assert.New(t)

	testTime := time.Date(2011, 5, 13, 0, 0, 0, 0, time.UTC)

	lib := simpleTime(testTime)
	lib.SetLocale("es")

	clone := lib.Clone()

	assert.Equal(simpleTime(testTime).Format(), lib.Format())
	assert.Equal(lib.Locale(), clone.Locale())
}

func TestUnixFromSeconds(t *testing.T) {
	assert := assert.New(t)

	// time.Unix does not have microsecond info, so we must
	// compare the Unix versions of the time.
	testTime := time.Now()
	testTimeUnix := time.Unix(testTime.Unix(), 0)

	lib := simpleUnix(testTime.Unix())

	assert.Equal(simpleTime(testTimeUnix).Format(), lib.Format())
	assert.Equal(time.Local, lib.ToTime().Location())
}
