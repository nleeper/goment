package goment

import (
	"errors"
	"time"

	"github.com/nleeper/goment/locales"
)

var timeNow = time.Now

// Goment is the main class.
type Goment struct {
	time   time.Time
	locale locales.LocaleDetails
}

// DateTime is a class to define a date & time.
type DateTime struct {
	Year       int
	Month      int
	Day        int
	Hour       int
	Minute     int
	Second     int
	Nanosecond int
	Location   *time.Location
}

// New creates an instance of the Goment library.
func New(args ...interface{}) (*Goment, error) {
	loadParseReplacements()
	loadFormatReplacements()

	switch len(args) {
	case 0:
		return fromNow()
	case 1:
		switch v := args[0].(type) {
		case string:
			return fromISOString(v)
		case time.Time:
			return fromExistingTime(v)
		case int64:
			return fromUnixNanoseconds(v)
		case *Goment:
			return fromGoment(v)
		case Goment:
			return fromGoment(&v)
		case DateTime:
			return fromDateTime(v)
		default:
			return &Goment{}, errors.New("Invalid argument type")
		}
	case 2:
		if date, ok := args[0].(string); ok {
			if format, ok := args[1].(string); ok {
				return fromStringWithFormat(date, format, getGlobalLocaleDetails())
			}
			return &Goment{}, errors.New("Second argument must be a format string")
		}
		return &Goment{}, errors.New("First argument must be a datetime string")
	case 3:
		// TODO - cleanup argument parsing.
		if date, ok := args[0].(string); ok {
			if format, ok := args[1].(string); ok {
				if localeCode, ok := args[2].(string); ok {
					locale, err := loadLocale(localeCode)
					if err != nil {
						return &Goment{}, errors.New("Invalid locale code")
					}
					return fromStringWithFormat(date, format, locale)
				}
			}
			return &Goment{}, errors.New("Second argument must be a format string")
		}
		return &Goment{}, errors.New("First argument must be a datetime string")
	default:
		return &Goment{}, errors.New("Invalid number of arguments")
	}
}

// Unix creates an instance of the Goment library from the Unix timestamp (seconds since the Unix Epoch).
func Unix(unixSeconds int64) (*Goment, error) {
	t := time.Unix(unixSeconds, 0)
	return createGoment(t)
}

// Clone creates a new copy of the Goment instance.
func (g *Goment) Clone() *Goment {
	copy, _ := New()
	copy.time = g.ToTime()
	copy.locale = g.locale

	return copy
}

func fromDateTime(dt DateTime) (*Goment, error) {
	tn := timeNow()

	year := tn.Year()
	if dt.Year != 0 {
		year = dt.Year
	}

	month := tn.Month()
	if dt.Month > 0 {
		month = time.Month(dt.Month)
	}

	day := tn.Day()
	if dt.Day > 0 {
		day = dt.Day
	}

	// Default to local time if not provided.
	loc := tn.Location()
	if dt.Location != nil {
		loc = dt.Location
	}

	d := time.Date(year, month, day, dt.Hour, dt.Minute, dt.Second, dt.Nanosecond, loc)
	return fromExistingTime(d)
}

func fromGoment(g *Goment) (*Goment, error) {
	return g.Clone(), nil
}

func fromNow() (*Goment, error) {
	return fromExistingTime(timeNow())
}

func fromUnixNanoseconds(unixNano int64) (*Goment, error) {
	return fromExistingTime(time.Unix(0, unixNano))
}

func fromExistingTime(t time.Time) (*Goment, error) {
	return createGoment(t)
}

func fromStringWithFormat(date string, format string, locale locales.LocaleDetails) (*Goment, error) {
	parsed, err := parseFromFormat(date, format, locale)
	if err != nil {
		return &Goment{}, err
	}

	return createGomentWithLocale(parsed, locale)
}

func fromISOString(date string) (*Goment, error) {
	parsed, err := parseISOString(date)
	if err != nil {
		return &Goment{}, err
	}

	return createGoment(parsed)
}

func createGoment(t time.Time) (*Goment, error) {
	return createGomentWithLocale(t, getGlobalLocaleDetails())
}

func createGomentWithLocale(t time.Time, ld locales.LocaleDetails) (*Goment, error) {
	return &Goment{t, ld}, nil
}
