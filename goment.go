package goment

import (
	"time"
)

var timeNow = time.Now

// Goment is the main class.
type Goment struct {
	DateTime time.Time
}

// New creates an instance of the Goment library.
func New() (*Goment, error) {
	return newFromNow(), nil
}

func newFromNow() *Goment {
	return &Goment{
		timeNow(),
	}
}
