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
func New(args ...string) (*Goment, error) {
	switch len(args) {
	case 0:
		return fromNow()
	case 1:
		return fromISOString(args[0])
	default:
		return &Goment{}, errors.New("Invalid number of arguments")
	}
}

func fromNow() (*Goment, error) {
	return &Goment{timeNow()}, nil
}

func fromISOString(date string) (*Goment, error) {
	parsed, err := parseISOString(date)
	if err != nil {
		return &Goment{}, err
	}

	return &Goment{parsed}, nil
}
