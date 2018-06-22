package goment

import "time"

type testParseable struct {
	DateTime   string
	ParsedTime time.Time
}

func chicagoLocation() *time.Location {
	location, _ := time.LoadLocation("America/Chicago")
	return location
}

func simpleNow() *Goment {
	lib, _ := New()
	return lib
}

func simple(dateTime DateTime) *Goment {
	lib, _ := New(dateTime)
	return lib
}

func simpleTime(time time.Time) *Goment {
	lib, _ := New(time)
	return lib
}

func simpleString(time string) *Goment {
	lib, _ := New(time)
	return lib
}

func simpleGoment(g *Goment) *Goment {
	lib, _ := New(g)
	return lib
}

func dstForYear(year int) *Goment {
	start := simple(DateTime{Year: year, Month: 1})
	end := simple(DateTime{Year: year + 1, Month: 1})
	current := start.Clone()

	var last *Goment

	for current.IsBefore(end) {
		last = current.Clone()
		current.Add(24, "hours")

		if last.UTCOffset() != current.UTCOffset() {
			end = current.Clone()
			current = last.Clone()
			break
		}
	}

	for current.IsBefore(end) {
		last = current.Clone()
		current.Add(1, "hour")

		if last.UTCOffset() != current.UTCOffset() {
			return last
		}
	}

	return &Goment{}
}
