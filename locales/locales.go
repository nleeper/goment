package locales

import (
	"strconv"
	"strings"

	"github.com/nleeper/goment/regexps"
)

type ordinalFunction func(int, string) string

type meridiemFunction func(int, int, bool) string

type calendarFunction func(int, int) string

type longDateFormats map[string]string

type relativeTimeFormats map[string]string

type calendarFunctions map[string]calendarFunction

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
	RelativeTimes   relativeTimeFormats
	Calendar        calendarFunctions
}

// RelativeTime returns the relative time for the period.
func (ld *LocaleDetails) RelativeTime(format string, number int, withoutSuffix bool, past bool) string {
	relTime := strings.Replace(ld.RelativeTimes[format], "%d", strconv.Itoa(number), 1)

	if withoutSuffix {
		return relTime
	}

	futurePast := ld.RelativeTimes["future"]
	if past {
		futurePast = ld.RelativeTimes["past"]
	}

	return strings.Replace(futurePast, "%s", relTime, 1)
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
func NewLocale(code string, wd []string, wds []string, wdm []string, m []string, ms []string, of ordinalFunction, mf meridiemFunction, dow int, ld longDateFormats, rt relativeTimeFormats, cal calendarFunctions) LocaleDetails {
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
		RelativeTimes:   rt,
		Calendar:        cal,
	}
}
