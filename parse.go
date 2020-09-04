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

type configFunc func(string, *parseConfig, locales.LocaleDetails, string)
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
	overflowWeekday   bool
	overflowWeeks     bool
	meridiem          string
	week              map[string]int
	parsedArray       map[int]int
	date              *Goment
	locale            locales.LocaleDetails
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

	// 	// // "S":    createParseReplacement("S", "frac_second", nil, nil),
	// 	// // "SS":   createParseReplacement("SS", "frac_second", nil, nil),
	// 	// // "SSS":  createParseReplacement("SSS", "frac_second", nil, nil),

	addParseReplacement([]string{"D", "DD"}, dateIdx, regexps.MatchOneToTwo)
	addParseReplacement("Do", handleOrdinalDate, func(input string, locale locales.LocaleDetails) (string, string) {
		return findRegexString(input, locale.DayOfMonthOrdinalRegex)
	})
	addParseReplacement("DDD", handleDayOfYear, regexps.MatchOneToThree)
	addParseReplacement("DDDD", handleDayOfYear, regexps.MatchThree)

	addParseReplacement([]string{"M", "MM"}, monthIdx, regexps.MatchOneToTwo)
	addParseReplacement("MMM", handleShortMonth, func(input string, locale locales.LocaleDetails) (string, string) {
		return findRegexString(input, locale.MonthsShortRegex)
	})
	addParseReplacement("MMMM", handleLongMonth, func(input string, locale locales.LocaleDetails) (string, string) {
		return findRegexString(input, locale.MonthsRegex)
	})

	addParseReplacement("Y", handleSingleDigitYear, regexps.MatchSigned)
	addParseReplacement("YY", handleTwoDigitYear, regexps.MatchOneToTwo)
	addParseReplacement("YYYY", handleFourDigitYear, regexps.MatchOneToFour)
	addParseReplacement([]string{"YYYYY", "YYYYYY"}, yearIdx, regexps.MatchOneToSix)

	addParseReplacement("Q", handleQuarter, regexps.MatchOne)

	addParseReplacement("X", handleTimestamp, regexps.MatchTimestamp)

	addParseReplacement([]string{"h", "hh", "H", "HH"}, hourIdx, regexps.MatchOneToTwo)
	addParseReplacement([]string{"k", "kk"}, handleOneToTwentyFourTime, regexps.MatchOneToTwo)
	addParseReplacement([]string{"m", "mm"}, minuteIdx, regexps.MatchOneToTwo)
	addParseReplacement([]string{"s", "ss"}, secondIdx, regexps.MatchOneToTwo)
	addParseReplacement([]string{"a", "A"}, handleMeridiem, regexps.MatchMeridiem)

	addParseReplacement([]string{"Z", "ZZ"}, handleOffset, regexps.MatchShortOffset)

	// Week & weekday parsing
	addWeekParseReplacement("dd", handleMinDayName, func(input string, locale locales.LocaleDetails) (string, string) {
		return findRegexString(input, locale.WeekdaysMinRegex)
	})
	addWeekParseReplacement("ddd", handleShortDayName, func(input string, locale locales.LocaleDetails) (string, string) {
		return findRegexString(input, locale.WeekdaysShortRegex)
	})
	addWeekParseReplacement("dddd", handleLongDayName, func(input string, locale locales.LocaleDetails) (string, string) {
		return findRegexString(input, locale.WeekdaysRegex)
	})

	addWeekParseReplacement([]string{"d", "e", "E"}, regexps.MatchOneToTwo)

	addWeekParseReplacement([]string{"w", "ww", "W", "WW"}, handleWeekNumber, regexps.MatchOneToTwo)

	addWeekParseReplacement([]string{"gg", "GG"}, handleTwoDigitWeekYear, regexps.MatchOneToTwo)
	addWeekParseReplacement([]string{"gggg", "GGGG"}, handleWeekYear, regexps.MatchOneToFour)
	addWeekParseReplacement([]string{"ggggg", "GGGGG"}, handleWeekYear, regexps.MatchOneToSix)
}

func findRegexString(input string, rx *regexp.Regexp) (string, string) {
	found := rx.FindStringIndex(input)
	if found != nil {
		return input[found[0]:found[1]], input[found[1]:len(input)]
	}
	return "", input
}

func buildStringSlice(arg interface{}) []string {
	switch v := arg.(type) {
	case string:
		return []string{v}
	case []string:
		return v
	}
	return []string{}
}

func buildFindTokenFunction(arg interface{}) findTokenFunc {
	switch v := arg.(type) {
	case func(input string, locale locales.LocaleDetails) (string, string):
		return v
	case *regexp.Regexp:
		return func(input string, locale locales.LocaleDetails) (string, string) {
			return findRegexString(input, v)
		}
	}

	// Return a no-op findTokenFunc if the type doesn't match.
	return func(input string, locale locales.LocaleDetails) (string, string) {
		return "", input
	}
}

func addWeekParseReplacement(args ...interface{}) {
	count := len(args)

	if count < 2 {
		return
	}

	var cf configFunc
	var ftf findTokenFunc

	tokens := buildStringSlice(args[0])

	switch count {
	case 2:
		tokens = buildStringSlice(args[0])
		ftf = buildFindTokenFunction(args[1])

		cf = func(input string, config *parseConfig, locale locales.LocaleDetails, token string) {
			config.week[token] = parseNumber(input)
		}
	case 3:
		tokens = buildStringSlice(args[0])
		cf = args[1].(func(string, *parseConfig, locales.LocaleDetails, string))
		ftf = buildFindTokenFunction(args[2])
	}

	for _, t := range tokens {
		parseReplacements[t] = parseReplacement{
			func(input string, config *parseConfig, locale locales.LocaleDetails, token string) {
				if config.week == nil {
					config.week = map[string]int{}
				}
				cf(input, config, locale, token)
			},
			ftf,
		}
	}
}

func addParseReplacement(args ...interface{}) {
	if len(args) == 3 {
		tokens := buildStringSlice(args[0])

		var cf configFunc

		switch v := args[1].(type) {
		case func(string, *parseConfig, locales.LocaleDetails, string):
			cf = v
		case int:
			cf = func(input string, config *parseConfig, locale locales.LocaleDetails, token string) {
				config.parsedArray[v] = parseNumber(input)
			}
		}

		ftf := buildFindTokenFunction(args[2])

		for _, t := range tokens {
			parseReplacements[t] = parseReplacement{
				cf,
				ftf,
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
			timezoneMatch := regexps.TimeZoneRegex.FindString(match[4])
			if timezoneMatch == "Z" {
				tzFormat = "Z"
			} else if timezoneMatch != "" {
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
		false,
		false,
		"",
		nil,
		map[int]int{},
		nil,
		locale,
	}

	var found = ""
	var remaining = date

	for i := range match {
		token := match[i][0]

		if rep, ok := parseReplacements[token]; ok {
			// Find the input value matching the token.
			found, remaining = rep.findTokenFunction(remaining, locale)

			if found != "" {
				rep.configFunction(found, config, locale, token)
			}
		}
	}

	return buildFromParseConfig(config)
}

func buildFromParseConfig(config *parseConfig) (*Goment, error) {
	// Update the config values based on the meridiem.
	fixForMeridiem(config)

	// If we were able to build a complete date from the format (X), return now.
	if config.date != nil {
		return config.date, nil
	}

	// Get the current date's values.
	currentDate := currentDateArray(config)

	if config.week != nil && !keyExists(dateIdx, config.parsedArray) && !keyExists(monthIdx, config.parsedArray) {
		dayOfYearFromWeekInfo(config)
	}

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

	expectedWeekday := config.date.Day()
	if config.week != nil && skeyExists("d", config.week) && config.week["d"] != expectedWeekday {
		return nil, errors.New("There is a mismatch between parsed weekday and expected weekday")
	}

	return config.date, nil
}

func dayOfYearFromWeekInfo(config *parseConfig) {
	var weekday, dow, doy, week, wy int
	var currWeek weekYear

	weekdayOverflow := false

	local, _ := New()

	w := config.week

	if skeyExists("GG", w) || skeyExists("W", w) || skeyExists("E", w) {
		dow = 1
		doy = 4

		currWeek = weekOfYear(local, dow, doy)

		wy = getDefaultWeekYear("GG", config, currWeek)

		week = getDefaultWeek("W", w, 1)
		weekday = getDefaultWeekday("E", w, 1)

		if weekday < 1 || weekday > 7 {
			weekdayOverflow = true
		}
	} else {
		dow = config.locale.Week.Dow
		doy = config.locale.Week.Doy

		currWeek = weekOfYear(local, dow, doy)

		wy = getDefaultWeekYear("gg", config, currWeek)

		week = getDefaultWeek("w", w, currWeek.week)

		if skeyExists("d", w) {
			weekday = w["d"]
			if weekday < 0 || weekday > 6 {
				weekdayOverflow = true
			}
		} else if skeyExists("e", w) {
			weekday = w["e"] + dow
			if w["e"] < 0 || w["e"] > 6 {
				weekdayOverflow = true
			}
		} else {
			weekday = dow
		}
	}

	if week < 1 || week > weeksInYear(wy, dow, doy) {
		config.overflowWeeks = true
	} else if weekdayOverflow {
		config.overflowWeekday = true
	} else {
		temp := dayOfYearFromWeeks(wy, week, weekday, dow, doy)
		config.parsedArray[yearIdx] = temp.year
		config.dayOfYear = temp.dayOfYear
	}
}

func getDefaultWeekday(token string, week map[string]int, currWeekday int) int {
	if skeyExists(token, week) {
		return week[token]
	}
	return currWeekday
}

func getDefaultWeek(token string, week map[string]int, currWeek int) int {
	if skeyExists(token, week) {
		return week[token]
	}
	return currWeek
}

func getDefaultWeekYear(token string, config *parseConfig, currWeek weekYear) int {
	if skeyExists(token, config.week) {
		return config.week[token]
	}
	if keyExists(yearIdx, config.parsedArray) {
		return config.parsedArray[yearIdx]
	}
	return currWeek.year
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

func handleSingleDigitYear(input string, config *parseConfig, locale locales.LocaleDetails, token string) {
	y, _ := strconv.Atoi(input)
	config.parsedArray[yearIdx] = y
}

func handleOrdinalDate(input string, config *parseConfig, locale locales.LocaleDetails, token string) {
	config.parsedArray[dateIdx] = parseNumber(regexps.MatchOneToTwo.FindAllString(input, -1)[0])
}

func handleDayOfYear(input string, config *parseConfig, locale locales.LocaleDetails, token string) {
	config.dayOfYear = parseNumber(strings.TrimLeft(input, "0"))
}

func handleQuarter(input string, config *parseConfig, locale locales.LocaleDetails, token string) {
	config.parsedArray[monthIdx] = parseNumber(input) * 3
}

func handleTimestamp(input string, config *parseConfig, locale locales.LocaleDetails, token string) {
	ts, _ := strconv.ParseInt(input, 10, 64)
	g, _ := Unix(ts)

	config.date = g
}

func handleOneToTwentyFourTime(input string, config *parseConfig, locale locales.LocaleDetails, token string) {
	config.parsedArray[hourIdx] = parseNumber(input) - 1
}

func handleMeridiem(input string, config *parseConfig, locale locales.LocaleDetails, token string) {
	config.meridiem = input
}

func handleOffset(input string, config *parseConfig, locale locales.LocaleDetails, token string) {
	config.isUTC = true

	match := regexps.MatchShortOffset.FindAllString(input, -1)
	if match != nil {
		offset := match[0]
		parts := regexps.ChunkOffset.FindAllString(offset, -1)

		minutes := 0
		if len(parts) == 3 {
			minutes = (parseNumber(parts[1]) * 60) + parseNumber(parts[2])
		}
		if minutes > 0 {
			if parts[0] == "-" {
				minutes *= -1
			}
		}

		config.tzMinutes = minutes
	}
}

func handleLongMonth(input string, config *parseConfig, locale locales.LocaleDetails, token string) {
	config.parsedArray[monthIdx] = locale.GetMonthNumber(input)
}

func handleShortMonth(input string, config *parseConfig, locale locales.LocaleDetails, token string) {
	config.parsedArray[monthIdx] = locale.GetMonthShortNumber(input)
}

func handleLongDayName(input string, config *parseConfig, locale locales.LocaleDetails, token string) {
	val := locale.GetWeekdayNumber(input)
	if val != -1 {
		config.week["d"] = val
	}
}

func handleShortDayName(input string, config *parseConfig, locale locales.LocaleDetails, token string) {
	val := locale.GetWeekdayShortNumber(input)
	if val != -1 {
		config.week["d"] = val
	}
}

func handleMinDayName(input string, config *parseConfig, locale locales.LocaleDetails, token string) {
	val := locale.GetWeekdayMinNumber(input)
	if val != -1 {
		config.week["d"] = val
	}
}

func handleWeekNumber(input string, config *parseConfig, locale locales.LocaleDetails, token string) {
	config.week[token[:1]] = parseNumber(input)
}

func handleWeekYear(input string, config *parseConfig, locale locales.LocaleDetails, token string) {
	config.week[token[:2]] = parseNumber(input)
}

func handleTwoDigitWeekYear(input string, config *parseConfig, locale locales.LocaleDetails, token string) {
	config.week[token] = parseTwoDigitYear(input)
}

func handleTwoDigitYear(input string, config *parseConfig, locale locales.LocaleDetails, token string) {
	config.parsedArray[yearIdx] = parseTwoDigitYear(input)
}

func handleFourDigitYear(input string, config *parseConfig, locale locales.LocaleDetails, token string) {
	year := 0
	if len(input) == 2 {
		year = parseTwoDigitYear(input)
	} else {
		year = parseNumber(input)
	}
	config.parsedArray[yearIdx] = year
}

func parseTwoDigitYear(input string) int {
	baseYear := 0
	year := parseNumber(input)

	if year > 68 {
		baseYear = 1900
	} else {
		baseYear = 2000
	}

	return year + baseYear
}

func parseNumber(dateText string) int {
	num, _ := strconv.Atoi(dateText)
	return num
}
