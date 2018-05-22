package goment

import (
	"regexp"
)

var inclusivityRegex = regexp.MustCompile("^[\\[\\(]{1}[\\]\\)]{1}$")

// IsBefore will check if a Goment is before another Goment.
func (g *Goment) IsBefore(args ...interface{}) bool {
	var err error
	var input *Goment

	numArgs := len(args)
	if numArgs == 0 {
		input, err = New()
	} else {
		input, err = New(args[0])
	}

	if err != nil {
		return false
	}

	if numArgs <= 1 {
		return g.ToTime().Before(input.ToTime())
	}

	if units, ok := args[1].(string); ok {
		return g.ToTime().Before(input.StartOf(units).ToTime())
	}

	return false
}

// IsAfter will check if a Goment is after another Goment.
func (g *Goment) IsAfter(args ...interface{}) bool {
	var err error
	var input *Goment

	numArgs := len(args)
	if numArgs == 0 {
		input, err = New()
	} else {
		input, err = New(args[0])
	}

	if err != nil {
		return false
	}

	if numArgs <= 1 {
		return g.ToTime().After(input.ToTime())
	}

	if units, ok := args[1].(string); ok {
		return g.ToTime().After(input.EndOf(units).ToTime())
	}

	return false
}

// IsSame will check if a Goment is the same as another Goment.
func (g *Goment) IsSame(args ...interface{}) bool {
	numArgs := len(args)
	if numArgs > 0 {
		input, err := New(args[0])
		if err != nil {
			return false
		}

		if numArgs == 1 {
			return g.ToTime().Equal(input.ToTime())
		}

		if units, ok := args[1].(string); ok {
			return g.StartOf(units).ToTime().Equal(input.StartOf(units).ToTime())
		}
	}
	return false
}

// IsSameOrBefore will check if a Goment is before or the same as another Goment.
func (g *Goment) IsSameOrBefore(args ...interface{}) bool {
	return g.IsSame(args...) || g.IsBefore(args...)
}

// IsSameOrAfter will check if a Goment is after or the same as another Goment.
func (g *Goment) IsSameOrAfter(args ...interface{}) bool {
	return g.IsSame(args...) || g.IsAfter(args...)
}

// IsBetween will check if a Goment is between two other Goments.
func (g *Goment) IsBetween(args ...interface{}) bool {
	numArgs := len(args)
	if numArgs >= 2 {
		units := ""
		inclusivity := "()"
		fromResult, toResult := false, false

		from, err := New(args[0])
		if err != nil {
			return false
		}

		to, err := New(args[1])
		if err != nil {
			return false
		}

		if numArgs >= 3 {
			if parsedUnits, ok := args[2].(string); ok {
				units = parsedUnits
			}
		}

		if numArgs == 4 {
			if parsedInclusivity, ok := args[3].(string); ok {
				if inclusivityRegex.MatchString(parsedInclusivity) {
					inclusivity = parsedInclusivity
				}
			}
		}

		if inclusivity[0] == '(' {
			fromResult = g.IsAfter(from, units)
		} else {
			fromResult = !g.IsBefore(from, units)
		}

		if inclusivity[1] == ')' {
			toResult = g.IsBefore(to, units)
		} else {
			toResult = !g.IsAfter(to, units)
		}

		return fromResult && toResult
	}

	return false
}
