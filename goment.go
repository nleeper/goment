package goment

import (
	"time"

	"github.com/pkg/errors"
)

var timeNow = time.Now

// Goment is the main class.
type Goment struct {
	DateTime time.Time
}

// New creates an instance of the Goment library.
func New(args ...interface{}) (*Goment, error) {
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
		default:
			return &Goment{}, errors.New("Invalid argument type")
		}
	default:
		return &Goment{}, errors.New("Invalid number of arguments")
	}
}

// Unix creates an instance of the Goment library from the Unix timestamp (seconds since the Unix Epoch).
func Unix(unixSeconds int64) (*Goment, error) {
	t := time.Unix(unixSeconds, 0)
	return createGoment(toLocalTime(t))
}

func fromNow() (*Goment, error) {
	// Convert current time to Local.
	now := timeNow()
	return createGoment(toLocalTime(now))
}

func fromUnixNanoseconds(unixNano int64) (*Goment, error) {
	// Convert time to Local.
	t := time.Unix(0, unixNano)
	return createGoment(toLocalTime(t))
}

func fromExistingTime(t time.Time) (*Goment, error) {
	// Convert existing time to Local.
	return createGoment(toLocalTime(t))
}

func fromISOString(date string) (*Goment, error) {
	// Use whatever tz is parsed for now, need to figure out if we should convert to Local.
	parsed, err := parseISOString(date)
	if err != nil {
		return &Goment{}, err
	}

	return createGoment(parsed)
}

func toLocalTime(t time.Time) time.Time {
	return t.Local()
}

func createGoment(t time.Time) (*Goment, error) {
	return &Goment{t}, nil
}
