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
			return fromTime(v)
		default:
			return &Goment{}, errors.New("Invalid argument type")
		}
	default:
		return &Goment{}, errors.New("Invalid number of arguments")
	}
}

func fromNow() (*Goment, error) {
	return fromTime(timeNow())
}

func fromTime(time time.Time) (*Goment, error) {
	return &Goment{time}, nil
}

func fromISOString(date string) (*Goment, error) {
	parsed, err := parseISOString(date)
	if err != nil {
		return &Goment{}, err
	}

	return fromTime(parsed)
}
