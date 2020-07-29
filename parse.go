package goment

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/nleeper/goment/locales"
	"github.com/nleeper/goment/regexps"
)

var yearIdx = 0
var monthIdx = 1
var dateIdx = 2
var hourIdx = 3
var minuteIdx = 4
var secondIdx = 5
var nanosecondIdx = 6
var weekIdx = 7
var weekdayIdx = 8

var isoDates = []isoDateFormat{
	// createISOFormat("+002006-01-02", `\+\d{6}-\d\d-\d\d`),
	// createISOFormat("-002006-01-02", `-\d{6}-\d\d-\d\d`),
	createISODateFormat("2006-01-02", `\d{4}-\d\d-\d\d`),
	createISODateFormat("2006-01", `\d{4}-\d\d`, false),
	createISODateFormat("20060102", `\d{8}`),
}

var isoTimes = []isoTimeFormat{
	createISOTimeFormat("15:04:05.9999", `\d\d:\d\d:\d\d\.\d+`),
	// createISOTimeFormat("15:04:05,9999", `\d\d:\d\d:\d\d,\d+`), // TODO - will need to replace the comma with a period in date string
	createISOTimeFormat("15:04:05", `\d\d:\d\d:\d\d`),
	createISOTimeFormat("15:04", `\d\d:\d\d`),
	createISOTimeFormat("150405.9999", `\d\d\d\d\d\d\.\d+`),
	// createISOTimeFormat("150405,9999", `\d\d\d\d\d\d,\d+`), // TODO - replace comma
	createISOTimeFormat("150405", `\d\d\d\d\d\d`),
	createISOTimeFormat("1504", `\d\d\d\d`),
	createISOTimeFormat("15", `\d\d`),
}

// var isoDates = [
//     ['YYYYYY-MM-DD', /[+-]\d{6}-\d\d-\d\d/],
//     ['YYYY-MM-DD', /\d{4}-\d\d-\d\d/],
//     ['GGGG-[W]WW-E', /\d{4}-W\d\d-\d/],
//     ['GGGG-[W]WW', /\d{4}-W\d\d/, false],
//     ['YYYY-DDD', /\d{4}-\d{3}/],
//     ['YYYY-MM', /\d{4}-\d\d/, false],
//     ['YYYYYYMMDD', /[+-]\d{10}/],
//     ['YYYYMMDD', /\d{8}/],
//     // YYYYMM is NOT allowed by the standard
//     ['GGGG[W]WWE', /\d{4}W\d{3}/],
//     ['GGGG[W]WW', /\d{4}W\d{2}/, false],
//     ['YYYYDDD', /\d{7}/]
// ]

// var isoTimes = [
//     ['HH:mm:ss.SSSS', /\d\d:\d\d:\d\d\.\d+/],
//     ['HH:mm:ss,SSSS', /\d\d:\d\d:\d\d,\d+/],
//     ['HH:mm:ss', /\d\d:\d\d:\d\d/],
//     ['HH:mm', /\d\d:\d\d/],
//     ['HHmmss.SSSS', /\d\d\d\d\d\d\.\d+/],
//     ['HHmmss,SSSS', /\d\d\d\d\d\d,\d+/],
//     ['HHmmss', /\d\d\d\d\d\d/],
//     ['HHmm', /\d\d\d\d/],
//     ['HH', /\d\d/]
// ]

type configFunc func(string, *parseConfig, locales.LocaleDetails)
type findTokenFunc func(string, locales.LocaleDetails) (string, string)

type isoTimeFormat struct {
	format string
	regex  *regexp.Regexp
}

type isoDateFormat struct {
	format    string
	regex     *regexp.Regexp
	allowTime bool
}

type parseConfig struct {
	isUTC             bool
	tzMinutes         int
	dayOfYear         int
	overflowDayOfYear bool
	meridiem          string
	week              map[string]int
	parsedArray       map[int]int
	date              *Goment
}

type parseReplacement struct {
	configFunction    configFunc
	findTokenFunction findTokenFunc
}

var parseReplacements = map[string]parseReplacement{}

func loadParseReplacements() {
	if len(parseReplacements) > 0 {
		return
	}

	// 	// // "gggg": createParseReplacement("gggg", "isoWeekYear", regexps.MatchOneToFour, nil),
	// 	// // "gg":   createParseReplacement("gg", "isoWeekYear", regexps.MatchOneToTwo, nil),
	// 	// // "w":  createParseReplacement("w", "weekOfYear", regexps.MatchOneToTwo, nil),
	// 	// // "ww": createParseReplacement("ww", "weekOfYear", regexps.MatchOneToTwo, nil),
	// 	// // "e":    createParseReplacement("e", "dayOfWeek", regexps.MatchOneToTwo, nil),
	// 	// // "GGGG": createParseReplacement("GGGG", "isoWeekyear", regexps.MatchOneToFour, nil),
	// 	// // "GG":   createParseReplacement("GG", "isoWeekYear", regexps.MatchOneToTwo, nil),
	// 	// // "W": createParseReplacement("W", "isoWeekOfYear", regexps.MatchOneToTwo, nil),
	// 	// // "WW": createParseReplacement("WW", "isoWeekOfYear", regexps.MatchOneToTwo, nil),
	// 	// "E":  createParseReplacement("E", "isoWeekday", regexps.MatchOneToTwo, nil),
	// 	// // "S":    createParseReplacement("S", "frac_second", nil, nil),
	// 	// // "SS":   createParseReplacement("SS", "frac_second", nil, nil),
	// 	// // "SSS":  createParseReplacement("SSS", "frac_second", nil, nil),

	addParseReplacement("D", dateIdx, func(input string, locale locales.LocaleDetails) (string, string) {
		return findRegexString(input, regexps.MatchOneToTwo)
	})
	addParseReplacement("Do", parseOrdinalDate, func(input string, locale locales.LocaleDetails) (string, string) {
		return findRegexString(input, locale.DayOfMonthOrdinalRegex)
	})
	addParseReplacement("DD", dateIdx, func(input string, locale locales.LocaleDetails) (string, string) {
		return findRegexString(input, regexps.MatchOneToTwo)
	})
	addParseReplacement("DDD", parseDayOfYear, func(input string, locale locales.LocaleDetails) (string, string) {
		return findRegexString(input, regexps.MatchOneToThree)
	})
	addParseReplacement("DDDD", parseDayOfYear, func(input string, locale locales.LocaleDetails) (string, string) {
		return findRegexString(input, regexps.MatchThree)
	})
	addParseReplacement("ddd", parseShortDayName, func(input string, locale locales.LocaleDetails) (string, string) {
		return findRegexString(input, locale.WeekdaysShortRegex)
	})
	addParseReplacement("dddd", parseLongDayName, func(input string, locale locales.LocaleDetails) (string, string) {
		return findRegexString(input, locale.WeekdaysRegex)
	})

	addParseReplacement("M", monthIdx, func(input string, locale locales.LocaleDetails) (string, string) {
		return findRegexString(input, regexps.MatchOneToTwo)
	})
	addParseReplacement("MM", monthIdx, func(input string, locale locales.LocaleDetails) (string, string) {
		return findRegexString(input, regexps.MatchOneToTwo)
	})
	addParseReplacement("MMM", parseShortMonth, func(input string, locale locales.LocaleDetails) (string, string) {
		return findRegexString(input, locale.MonthsShortRegex)
	})
	addParseReplacement("MMMM", parseLongMonth, func(input string, locale locales.LocaleDetails) (string, string) {
		return findRegexString(input, locale.MonthsRegex)
	})

	addParseReplacement("Y", parseSingleDigitYear, func(input string, locale locales.LocaleDetails) (string, string) {
		return findRegexString(input, regexps.MatchSigned)
	})
	addParseReplacement("YY", parseTwoDigitYear, func(input string, locale locales.LocaleDetails) (string, string) {
		return findRegexString(input, regexps.MatchOneToTwo)
	})
	addParseReplacement("YYYY", yearIdx, func(input string, locale locales.LocaleDetails) (string, string) {
		return findRegexString(input, regexps.MatchOneToFour)
	})
	// addParseReplacement("YYYYY", func(input string, locale locales.LocaleDetails) (string, string) {
	// 	return findRegexString(input, regexps.MatchOneToFour)
	// })
	// addParseReplacement("YYYYYY", func(input string, locale locales.LocaleDetails) (string, string) {
	// 	return findRegexString(input, regexps.MatchOneToFour)
	// })

	addParseReplacement("Q", parseQuarter, func(input string, locale locales.LocaleDetails) (string, string) {
		return findRegexString(input, regexps.MatchOne)
	})

	addParseReplacement("X", parseTimestamp, func(input string, locale locales.LocaleDetails) (string, string) {
		return findRegexString(input, regexps.MatchTimestamp)
	})

	addParseReplacement("H", hourIdx, func(input string, locale locales.LocaleDetails) (string, string) {
		return findRegexString(input, regexps.MatchOneToTwo)
	})
	addParseReplacement("HH", hourIdx, func(input string, locale locales.LocaleDetails) (string, string) {
		return findRegexString(input, regexps.MatchOneToTwo)
	})
	addParseReplacement("h", hourIdx, func(input string, locale locales.LocaleDetails) (string, string) {
		return findRegexString(input, regexps.MatchOneToTwo)
	})
	addParseReplacement("hh", hourIdx, func(input string, locale locales.LocaleDetails) (string, string) {
		return findRegexString(input, regexps.MatchOneToTwo)
	})
	addParseReplacement("k", parseOneToTwentyFourTime, func(input string, locale locales.LocaleDetails) (string, string) {
		return findRegexString(input, regexps.MatchOneToTwo)
	})
	addParseReplacement("kk", parseOneToTwentyFourTime, func(input string, locale locales.LocaleDetails) (string, string) {
		return findRegexString(input, regexps.MatchOneToTwo)
	})
	addParseReplacement("a", parseMeridiem, func(input string, locale locales.LocaleDetails) (string, string) {
		return findRegexString(input, regexps.MatchMeridiem)
	})
	addParseReplacement("A", parseMeridiem, func(input string, locale locales.LocaleDetails) (string, string) {
		return findRegexString(input, regexps.MatchMeridiem)
	})
	addParseReplacement("m", minuteIdx, func(input string, locale locales.LocaleDetails) (string, string) {
		return findRegexString(input, regexps.MatchOneToTwo)
	})
	addParseReplacement("mm", minuteIdx, func(input string, locale locales.LocaleDetails) (string, string) {
		return findRegexString(input, regexps.MatchOneToTwo)
	})
	addParseReplacement("s", secondIdx, func(input string, locale locales.LocaleDetails) (string, string) {
		return findRegexString(input, regexps.MatchOneToTwo)
	})
	addParseReplacement("ss", secondIdx, func(input string, locale locales.LocaleDetails) (string, string) {
		return findRegexString(input, regexps.MatchOneToTwo)
	})

	addParseReplacement("Z", parseOffset, func(input string, locale locales.LocaleDetails) (string, string) {
		return findRegexString(input, regexps.MatchShortOffset)
	})
	addParseReplacement("ZZ", parseOffset, func(input string, locale locales.LocaleDetails) (string, string) {
		return findRegexString(input, regexps.MatchShortOffset)
	})
}

func findRegexString(input string, rx *regexp.Regexp) (string, string) {
	found := rx.FindStringIndex(input)
	if found != nil {
		return input[found[0]:found[1]], input[found[1]:len(input)]
	}

	return "", input
}

func addParseReplacement(args ...interface{}) {
	if len(args) == 3 {
		match := args[0].(string)
		tf := args[2].(func(string, locales.LocaleDetails) (string, string))

		switch v := args[1].(type) {
		case func(string, *parseConfig, locales.LocaleDetails):
			parseReplacements[match] = parseReplacement{
				v,
				tf,
			}
		case int:
			parseReplacements[match] = parseReplacement{
				func(input string, config *parseConfig, locale locales.LocaleDetails) {
					config.parsedArray[v] = parseNumber(input)
				},
				tf,
			}
		}
	}
}

func createISOTimeFormat(format, regex string) isoTimeFormat {
	return isoTimeFormat{
		format,
		regexp.MustCompile(regex),
	}
}

func createISODateFormat(format, regex string, allowTime ...bool) isoDateFormat {
	shouldAllowTime := true
	if len(allowTime) > 0 {
		shouldAllowTime = allowTime[0]
	}

	return isoDateFormat{
		format,
		regexp.MustCompile(regex),
		shouldAllowTime,
	}
}

func parseFromFormat(date string, format string, locale locales.LocaleDetails) (time.Time, error) {
	parsedDate, err := parseToGoment(date, format, locale)
	if err != nil {
		return time.Time{}, err
	}

	return parsedDate.ToTime(), nil
}

func parseISOString(date string) (time.Time, error) {
	match := regexps.ExtendedISORegex.FindStringSubmatch(date)
	if match == nil {
		match = regexps.BasicISORegex.FindStringSubmatch(date)
	}

	allowTime := true
	dateFormat, timeFormat, tzFormat := "", "", ""

	if len(match) == 5 {
		// match[0] = matched
		// match[1] = date
		// match[2] = T
		// match[3] = time
		// match[4] = timezone

		for _, i := range isoDates {
			if i.regex.MatchString(match[1]) {
				dateFormat = i.format
				allowTime = i.allowTime
				break
			}
		}

		if dateFormat == "" {
			return time.Time{}, errors.New("No matching date format found")
		}

		if match[3] != "" {
			for _, i := range isoTimes {
				if i.regex.MatchString(match[3]) {
					separator := match[2]
					if separator == "" {
						separator = " "
					}

					timeFormat = separator + i.format
					break
				}
			}

			if timeFormat == "" {
				return time.Time{}, errors.New("No matching time format found")
			}
		}

		if !allowTime && timeFormat != "" {
			return time.Time{}, errors.New("Time part not allowed")
		}

		if match[4] != "" {
			if regexps.TimeZoneRegex.MatchString(match[4]) {
				tzFormat = "-0700"
			} else {
				return time.Time{}, errors.New("Invalid timezone format")
			}
		}

		finalFormat := dateFormat + timeFormat + tzFormat
		return time.Parse(finalFormat, date)
	}

	return time.Time{}, errors.New("Not a matching ISO-8601 date")
}

func parseToGoment(date, format string, locale locales.LocaleDetails) (*Goment, error) {
	format = expandLocaleFormats(format, locale)

	bracketMatch := regexps.BracketRegex.FindAllStringIndex(format, -1)
	bracketsFound := len(bracketMatch) > 0

	// Remove bracketed fields before parsing for tokens.
	if bracketsFound {
		for i := range bracketMatch {
			format = format[0:bracketMatch[i][0]] + format[bracketMatch[i][1]:len(format)]
		}
	}

	var match [][]string
	if match = regexps.TokenRegex.FindAllStringSubmatch(format, -1); match == nil {
		return nil, errors.New("No matches found in format")
	}

	config := &parseConfig{
		false,
		-99999,
		-1,
		false,
		"",
		nil,
		map[int]int{},
		nil,
	}

	var found = ""
	var remaining = date

	for i := range match {
		token := match[i][0]

		if rep, ok := parseReplacements[token]; ok {
			// Find the input value matching the token.
			found, remaining = rep.findTokenFunction(remaining, locale)

			if found != "" {
				rep.configFunction(found, config, locale)
			}
		}
	}

	// If we were able to build a complete date from the format (X), return now.
	if config.date != nil {
		return config.date, nil
	}

	// Update the config values based on the meridiem.
	fixForMeridiem(config)

	// Get the current date's values.
	currentDate := currentDateArray(config)

	// If the date hasn't been set by a parsing function, update config values if needed.
	if config.dayOfYear > -1 {
		yearToUse := defaults(config.parsedArray, currentDate, yearIdx)

		if config.dayOfYear > daysInYear(yearToUse) || config.dayOfYear == 0 {
			config.overflowDayOfYear = true
		}

		date := time.Date(yearToUse, 1, config.dayOfYear, 0, 0, 0, 0, time.UTC)
		config.parsedArray[monthIdx] = int(date.Month())
		config.parsedArray[dateIdx] = date.Day()
	}

	// Default to current date.
	// If no year, month, day of month are given, default to today.
	// If day of month is given, default month and year.
	// If month is given, default only year.
	// If year is given, don't default anything.
	for i := 0; i < 3 && !keyExists(i, config.parsedArray); i++ {
		if _, exist := config.parsedArray[i]; !exist {
			config.parsedArray[i] = currentDate[i]
		}
	}

	// Zero out any values that weren't default, including time.
	// If its month or day, set the value to 1.
	for i := 0; i < 7; i++ {
		defaultVal := 0
		if !keyExists(i, config.parsedArray) {
			if i == 1 || i == 2 {
				defaultVal = 1
			}
			config.parsedArray[i] = defaultVal
		}
	}

	createDateFromConfig(config)

	return config.date, nil
}

func fixForMeridiem(config *parseConfig) {
	if config.meridiem == "" {
		return
	}

	hour := config.parsedArray[hourIdx]
	isPM := strings.ToLower(config.meridiem)[0] == 'p'

	if isPM && hour < 12 {
		hour += 12
	}
	if !isPM && hour == 12 {
		hour = 0
	}

	config.parsedArray[hourIdx] = hour
}

func skeyExists(key string, m map[string]int) bool {
	_, exist := m[key]
	return exist
}
func keyExists(key int, m map[int]int) bool {
	_, exist := m[key]
	return exist
}

func createDateFromConfig(config *parseConfig) {
	loc := time.Local
	if config.isUTC {
		loc = time.UTC
	}

	datetime := DateTime{
		config.parsedArray[yearIdx],
		config.parsedArray[monthIdx],
		config.parsedArray[dateIdx],
		config.parsedArray[hourIdx],
		config.parsedArray[minuteIdx],
		config.parsedArray[secondIdx],
		config.parsedArray[nanosecondIdx],
		loc,
	}

	d, _ := New(datetime)
	if config.tzMinutes != -99999 {
		// Call this method rather than SetMinute, since it checks that minute is between 0-59.
		d.subtractMinutes(config.tzMinutes)

		// Set the offset.
		d.SetUTCOffset(config.tzMinutes)
	}

	config.date = d
}

func currentDateArray(config *parseConfig) map[int]int {
	newDate, _ := New()
	if config.isUTC {
		newDate.UTC()
	}
	return map[int]int{0: newDate.Year(), 1: newDate.Month(), 2: newDate.Date()}
}

func defaults(parsed, current map[int]int, idx int) int {
	if value, exist := parsed[idx]; exist {
		return value
	}
	if value, exist := current[idx]; exist {
		return value
	}
	return 0
}

func parseSingleDigitYear(input string, config *parseConfig, locale locales.LocaleDetails) {
	y, _ := strconv.Atoi(input)
	config.parsedArray[yearIdx] = y
}

func parseTwoDigitYear(input string, config *parseConfig, locale locales.LocaleDetails) {
	baseYear := 0
	year := parseNumber(input)

	if year > 68 {
		baseYear = 1900
	} else {
		baseYear = 2000
	}

	config.parsedArray[yearIdx] = year + baseYear
}

func parseOrdinalDate(input string, config *parseConfig, locale locales.LocaleDetails) {
	config.parsedArray[dateIdx] = parseNumber(regexps.MatchOneToTwo.FindAllString(input, -1)[0])
}

func parseDayOfYear(input string, config *parseConfig, locale locales.LocaleDetails) {
	config.dayOfYear = parseNumber(strings.TrimLeft(input, "0"))
}

func parseQuarter(input string, config *parseConfig, locale locales.LocaleDetails) {
	config.parsedArray[monthIdx] = parseNumber(input) * 3
}

func parseTimestamp(input string, config *parseConfig, locale locales.LocaleDetails) {
	ts, _ := strconv.ParseInt(input, 10, 64)
	g, _ := Unix(ts)

	config.date = g
}

func parseOneToTwentyFourTime(input string, config *parseConfig, locale locales.LocaleDetails) {
	config.parsedArray[hourIdx] = parseNumber(input) - 1
}

func parseMeridiem(input string, config *parseConfig, locale locales.LocaleDetails) {
	config.meridiem = input
}

func parseOffset(input string, config *parseConfig, locale locales.LocaleDetails) {
	config.isUTC = true

	match := regexps.MatchShortOffset.FindAllString(input, -1)
	if match != nil {
		offset := match[0]
		parts := regexps.ChunkOffset.FindAllString(offset, -1)

		minutes := (parseNumber(parts[1]) * 60) + parseNumber(parts[2])
		if minutes > 0 {
			if parts[0] == "-" {
				minutes *= -1
			}
		}

		config.tzMinutes = minutes
	}
}

func parseMonth(input string, config *parseConfig, locale locales.LocaleDetails) {
	config.parsedArray[monthIdx] = parseNumber(input)
}

func parseLongMonth(input string, config *parseConfig, locale locales.LocaleDetails) {
	config.parsedArray[monthIdx] = locale.GetMonthNumber(input)
}

func parseShortMonth(input string, config *parseConfig, locale locales.LocaleDetails) {
	config.parsedArray[monthIdx] = locale.GetMonthShortNumber(input)
}

func parseLongDayName(input string, config *parseConfig, locale locales.LocaleDetails) {
	val := locale.GetWeekdayNumber(input)
	if val != -1 {
		createWeekConfig(config)
		config.week["day"] = val
	}
}

func parseShortDayName(input string, config *parseConfig, locale locales.LocaleDetails) {
	val := locale.GetWeekdayShortNumber(input)
	if val != -1 {
		createWeekConfig(config)
		config.week["day"] = val
	}
}

func createWeekConfig(config *parseConfig) {
	if config.week == nil {
		config.week = map[string]int{}
	}
}

func parseNumber(dateText string) int {
	num, _ := strconv.Atoi(dateText)
	return num
}
