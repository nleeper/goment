package goment

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewNoDateCallsNow(t *testing.T) {
	testTime := time.Date(2000, 12, 15, 17, 8, 00, 0, time.UTC)
	timeNow = func() time.Time {
		return testTime
	}

	lib, err := New()
	if assert.NoError(t, err) {
		assert.Equal(t, lib.DateTime, testTime)
	}
}
