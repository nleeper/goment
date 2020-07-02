package locales

import (
	"strings"

	"github.com/nleeper/goment/regexps"
)

type ordinalFunction func(int, string) string

type meridiemFunction func(int, int, bool) string

type longDateFormats map[string]string

// LocaleDetails contains the details of the loaded locale.
type LocaleDetails struct {
	Code            string
	Weekdays        []string
	WeekdaysMin     []string
	WeekdaysShort   []string
	Months          []string
	MonthsShort     []string
	OrdinalFunc     ordinalFunction
	MeridiemFunc    meridiemFunction
	FirstDayOfWeek  int
	LongDateFormats longDateFormats
}

// LongDateFormat returns the format for the matching long date token.
func (ld *LocaleDetails) LongDateFormat(key string) (string, bool) {
	format, formatOk := ld.LongDateFormats[key]
	formatUpper, formatUpperOk := ld.LongDateFormats[strings.ToUpper(key)]

	// If we didnt't find the key or its upper, return false.
	if !formatOk && !formatUpperOk {
		return "", false
	}

	if formatOk || !formatUpperOk {
		return format, true
	}

	ld.LongDateFormats[key] = strings.Join(Map(regexps.TokenRegex.FindAllString(formatUpper, -1), func(token string) string {
		switch token {
		case "MMMM", "MM", "DD", "dddd":
			return token[1:]
		default:
			return token
		}
	}), "")

	return ld.LongDateFormats[key], true
}

// Map will iterate over string slice and call function on each item.
func Map(vs []string, f func(string) string) []string {
	vsm := make([]string, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}

// NewLocale loads new LocaleDetails.
func NewLocale(code string, wd []string, wds []string, wdm []string, m []string, ms []string, of ordinalFunction, mf meridiemFunction, dow int, ld longDateFormats) LocaleDetails {
	if mf == nil {
		mf = func(hours int, minutes int, isLower bool) string {
			m := ""
			if hours > 11 {
				m = "pm"
			} else {
				m = "am"
			}
			if !isLower {
				m = strings.ToUpper(m)
			}
			return m
		}
	}

	return LocaleDetails{
		Code:            code,
		Weekdays:        wd,
		WeekdaysShort:   wds,
		WeekdaysMin:     wdm,
		Months:          m,
		MonthsShort:     ms,
		OrdinalFunc:     of,
		MeridiemFunc:    mf,
		FirstDayOfWeek:  dow,
		LongDateFormats: ld,
	}
}
