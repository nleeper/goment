package locales

import (
	"regexp"
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

type week struct {
	Dow int
	Doy int
}

// LocaleDetails contains the details of the loaded locale.
type LocaleDetails struct {
	Code                   string
	Weekdays               []string
	WeekdaysMin            []string
	WeekdaysShort          []string
	Months                 []string
	MonthsShort            []string
	OrdinalFunc            ordinalFunction
	MeridiemFunc           meridiemFunction
	Week                   week
	LongDateFormats        longDateFormats
	RelativeTimes          relativeTimeFormats
	Calendar               calendarFunctions
	MonthsRegex            *regexp.Regexp
	MonthsShortRegex       *regexp.Regexp
	WeekdaysRegex          *regexp.Regexp
	WeekdaysShortRegex     *regexp.Regexp
	WeekdaysMinRegex       *regexp.Regexp
	DayOfMonthOrdinalRegex *regexp.Regexp
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

	ld.LongDateFormats[key] = strings.Join(mapString(regexps.TokenRegex.FindAllString(formatUpper, -1), func(token string) string {
		switch token {
		case "MMMM", "MM", "DD", "dddd":
			return token[1:]
		default:
			return token
		}
	}), "")

	return ld.LongDateFormats[key], true
}

// GetMonthNumber returns the number for the month name.
func (ld *LocaleDetails) GetMonthNumber(month string) int {
	var idx = 1
	for _, s := range ld.Months {
		if strings.ToLower(month) == strings.ToLower(s) {
			return idx
		}
		idx = idx + 1
	}
	return -1
}

// GetMonthShortNumber returns the number for the short month name.
func (ld *LocaleDetails) GetMonthShortNumber(month string) int {
	var idx = 1
	for _, s := range ld.MonthsShort {
		if strings.ToLower(month) == strings.ToLower(s) {
			return idx
		}
		idx = idx + 1
	}
	return -1
}

// GetWeekdayNumber returns the number for the weekday name.
func (ld *LocaleDetails) GetWeekdayNumber(wd string) int {
	var idx = 0
	for _, s := range ld.Weekdays {
		if strings.ToLower(wd) == strings.ToLower(s) {
			return idx
		}
		idx = idx + 1
	}
	return -1
}

// GetWeekdayShortNumber returns the number for the short weekday name.
func (ld *LocaleDetails) GetWeekdayShortNumber(wd string) int {
	var idx = 0
	for _, s := range ld.WeekdaysShort {
		if strings.ToLower(wd) == strings.ToLower(s) {
			return idx
		}
		idx = idx + 1
	}
	return -1
}

// GetWeekdayMinNumber returns the number for the min weekday name.
func (ld *LocaleDetails) GetWeekdayMinNumber(wd string) int {
	var idx = 0
	for _, s := range ld.WeekdaysMin {
		if strings.ToLower(wd) == strings.ToLower(s) {
			return idx
		}
		idx = idx + 1
	}
	return -1
}

// GetWeekdays returns the list of weekdays for the locale. If the shifted param is true, the weekdays will
// take the DayOfWeek for the locale and shift the weekdays so the first DayOfWeek is 0,
func (ld *LocaleDetails) GetWeekdays(shifted bool) []string {
	dow := 0
	if shifted {
		dow = ld.Week.Dow
	}
	return append(ld.Weekdays[dow:7], ld.Weekdays[0:dow]...)
}

// GetWeekdaysShort returns the list of short weekdays for the locale. If the shifted param is true, the short weekdays will
// take the DayOfWeek for the locale and shift the short weekdays so the first DayOfWeek is 0,
func (ld *LocaleDetails) GetWeekdaysShort(shifted bool) []string {
	dow := 0
	if shifted {
		dow = ld.Week.Dow
	}
	return append(ld.WeekdaysShort[dow:7], ld.WeekdaysShort[0:dow]...)
}

// GetWeekdaysMin returns the list of min weekdays for the locale. If the shifted param is true, the min weekdays will
// take the DayOfWeek for the locale and shift the min weekdays so the first DayOfWeek is 0,
func (ld *LocaleDetails) GetWeekdaysMin(shifted bool) []string {
	dow := 0
	if shifted {
		dow = ld.Week.Dow
	}
	return append(ld.WeekdaysMin[dow:7], ld.WeekdaysMin[0:dow]...)
}

func mapString(vs []string, f func(string) string) []string {
	vsm := make([]string, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}

func newLocale(code string, wd []string, wds []string, wdm []string, m []string, ms []string, of ordinalFunction,
	mf meridiemFunction, wk week, ld longDateFormats, rt relativeTimeFormats, cal calendarFunctions,
	monthsRegex string, monthsShortRegex string, weekdaysRegex string, weekdaysShortRegex string, weekdaysMinRegex string, domOrdinalRegex string) LocaleDetails {
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

	// TODO - build regexs for weekdays based off arrays of weekday names.
	return LocaleDetails{
		Code:                   code,
		Weekdays:               wd,
		WeekdaysShort:          wds,
		WeekdaysMin:            wdm,
		Months:                 m,
		MonthsShort:            ms,
		OrdinalFunc:            of,
		MeridiemFunc:           mf,
		Week:                   wk,
		LongDateFormats:        ld,
		RelativeTimes:          rt,
		Calendar:               cal,
		MonthsRegex:            regexp.MustCompile(monthsRegex),
		MonthsShortRegex:       regexp.MustCompile(monthsShortRegex),
		WeekdaysRegex:          regexp.MustCompile(weekdaysRegex),
		WeekdaysShortRegex:     regexp.MustCompile(weekdaysShortRegex),
		WeekdaysMinRegex:       regexp.MustCompile(weekdaysMinRegex),
		DayOfMonthOrdinalRegex: regexp.MustCompile(domOrdinalRegex),
	}
}
