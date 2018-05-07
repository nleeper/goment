package goment

import (
	"time"
)

// Local will set the Goment to use local time.
func (g *Goment) Local() *Goment {
	g.time = g.ToTime().Local()
	return g
}

// UTC will set the Goment to use UTC time.
func (g *Goment) UTC() *Goment {
	g.time = g.ToTime().UTC()
	return g
}

// UTCOffset get the UTC offset in minutes.
func (g *Goment) UTCOffset() int {
	_, o := g.ToTime().Zone()
	return o / 60
}

// SetUTCOffset sets the UTC offset in minutes. If the offset is less than 16 and greater than -16, the value is treated as hours.
func (g *Goment) SetUTCOffset(offset int) *Goment {
	multiplier := 60
	if offset < 16 && offset > -16 {
		multiplier = 3600
	}

	loc := time.FixedZone("Offset", offset*multiplier)
	g.time = g.ToTime().In(loc)

	return g
}
