package goment

import (
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/nleeper/goment/internal/constants"
	"github.com/nleeper/goment/internal/regexps"
	"github.com/pkg/errors"
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

var parseReplacements = map[string]parseReplacement{
	"M":    createParseReplacement("M", "month", regexps.MatchOneToTwo, monthIdx),
	"MM":   createParseReplacement("MM", "month", regexps.MatchOneToTwo, monthIdx),
	"MMM":  createParseReplacement("MMM", "month", regexps.MonthsRegex, parseShortMonth),
	"MMMM": createParseReplacement("MMMM", "month", regexps.MonthsRegex, parseLongMonth),
	"D":    createParseReplacement("D", "day", regexps.MatchOneToTwo, dateIdx),
	"Do":   createParseReplacement("Do", "day", regexps.DayOfMonthOrdinal, parseOrdinalDate),
	"DD":   createParseReplacement("DD", "day", regexps.MatchOneToTwo, dateIdx),
	"DDD":  createParseReplacement("DDD", "dayOfYear", regexps.MatchOneToThree, parseDayOfYear),
	"DDDD": createParseReplacement("DDDD", "dayOfYear", regexps.MatchThree, parseDayOfYear),
	"Y":    createParseReplacement("Y", "year", regexps.MatchSigned, parseSingleDigitYear),
	"YY":   createParseReplacement("YY", "year", regexps.MatchOneToTwo, parseTwoDigitYear),
	"YYYY": createParseReplacement("YYYY", "year", regexps.MatchOneToFour, yearIdx),
	"Q":    createParseReplacement("Q", "quarter", regexps.MatchOne, parseQuarter),
	"X":    createParseReplacement("X", "unix", regexps.MatchTimestamp, parseTimestamp),
	// // "gggg": createParseReplacement("gggg", "isoWeekYear", regexps.MatchOneToFour, nil),
	// // "gg":   createParseReplacement("gg", "isoWeekYear", regexps.MatchOneToTwo, nil),
	// // "w":  createParseReplacement("w", "weekOfYear", regexps.MatchOneToTwo, nil),
	// // "ww": createParseReplacement("ww", "weekOfYear", regexps.MatchOneToTwo, nil),
	// // "e":    createParseReplacement("e", "dayOfWeek", regexps.MatchOneToTwo, nil),
	// "ddd":  createParseReplacement("ddd", "weekday", regexps.WeekdaysRegex, parseShortDayName),
	// "dddd": createParseReplacement("dddd", "weekday", regexps.WeekdaysRegex, parseLongDayName),
	// // "GGGG": createParseReplacement("GGGG", "isoWeekyear", regexps.MatchOneToFour, nil),
	// // "GG":   createParseReplacement("GG", "isoWeekYear", regexps.MatchOneToTwo, nil),
	// // "W": createParseReplacement("W", "isoWeekOfYear", regexps.MatchOneToTwo, nil),
	// // "WW": createParseReplacement("WW", "isoWeekOfYear", regexps.MatchOneToTwo, nil),
	// "E":  createParseReplacement("E", "isoWeekday", regexps.MatchOneToTwo, nil),
	"H":  createParseReplacement("H", "hour", regexps.MatchOneToTwo, hourIdx),
	"HH": createParseReplacement("HH", "hour", regexps.MatchOneToTwo, hourIdx),
	"h":  createParseReplacement("h", "hour", regexps.MatchOneToTwo, hourIdx),
	"hh": createParseReplacement("hh", "hour", regexps.MatchOneToTwo, hourIdx),
	"k":  createParseReplacement("k", "hour", regexps.MatchOneToTwo, parseOneToTwentyFourTime),
	"kk": createParseReplacement("kk", "hour", regexps.MatchOneToTwo, parseOneToTwentyFourTime),
	"a":  createParseReplacement("a", "meridiem", regexps.MatchMeridiem, parseMeridiem),
	"A":  createParseReplacement("A", "meridiem", regexps.MatchMeridiem, parseMeridiem),
	"m":  createParseReplacement("m", "minute", regexps.MatchOneToTwo, minuteIdx),
	"mm": createParseReplacement("mm", "minute", regexps.MatchOneToTwo, minuteIdx),
	"s":  createParseReplacement("s", "second", regexps.MatchOneToTwo, secondIdx),
	"ss": createParseReplacement("ss", "second", regexps.MatchOneToTwo, secondIdx),
	// // "S":    createParseReplacement("S", "frac_second", nil, nil),
	// // "SS":   createParseReplacement("SS", "frac_second", nil, nil),
	// // "SSS":  createParseReplacement("SSS", "frac_second", nil, nil),
	"Z":  createParseReplacement("Z", "utcOffset", regexps.MatchShortOffset, parseOffset),
	"ZZ": createParseReplacement("ZZ", "utcOffset", regexps.MatchShortOffset, parseOffset),
}

type isoTimeFormat struct {
	Format string
	Regex  *regexp.Regexp
}

type isoDateFormat struct {
	Format    string
	Regex     *regexp.Regexp
	AllowTime bool
}

type parseConfig struct {
	IsUTC             bool
	TzMinutes         int
	DayOfYear         int
	OverflowDayOfYear bool
	Meridiem          string
	Week              map[string]int
	ParsedArray       map[int]int
	Date              *Goment
}

type parseReplacement struct {
	Match         string
	Type          string
	Regex         *regexp.Regexp
	ParseFunction func(string, *parseConfig)
}

func createParseReplacement(args ...interface{}) parseReplacement {
	if len(args) == 4 {
		match := args[0].(string)
		replaceType := args[1].(string)
		regex := args[2].(*regexp.Regexp)

		switch v := args[3].(type) {
		case func(string, *parseConfig):
			return parseReplacement{
				match,
				replaceType,
				regex,
				v,
			}
		case int:
			return parseReplacement{
				match,
				replaceType,
				regex,
				func(input string, config *parseConfig) {
					config.ParsedArray[v] = parseNumber(input)
				},
			}
		default:
			return parseReplacement{}
		}
	} else {
		return parseReplacement{}
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

func parseFromFormat(date string, format string) (time.Time, error) {
	parsedDate, err := parseToGoment(date, format)
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
			if i.Regex.MatchString(match[1]) {
				dateFormat = i.Format
				allowTime = i.AllowTime
				break
			}
		}

		if dateFormat == "" {
			return time.Time{}, errors.New("No matching date format found")
		}

		if match[3] != "" {
			for _, i := range isoTimes {
				if i.Regex.MatchString(match[3]) {
					separator := match[2]
					if separator == "" {
						separator = " "
					}

					timeFormat = separator + i.Format
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

func parseToGoment(date, format string) (*Goment, error) {
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

	var parseString = date

	for i := range match {
		token := match[i][0]
		input := ""

		if rep, ok := parseReplacements[token]; ok {
			// Find the input value matching the token.
			found := rep.Regex.FindStringIndex(parseString)
			if found != nil {
				input = parseString[found[0]:found[1]]
				parseString = parseString[found[1]:len(parseString)]
			}

			// Parse the input value.
			rep.ParseFunction(input, config)
		}
	}

	// If we were able to build a complete date from the format (X), return now.
	if config.Date != nil {
		return config.Date, nil
	}

	// Update the config values based on the meridiem.
	fixForMeridiem(config)

	// Get the current date's values.
	currentDate := currentDateArray(config)

	// If the date hasn't been set by a parsing function, update config values if needed.
	if config.DayOfYear > -1 {
		yearToUse := defaults(config.ParsedArray, currentDate, yearIdx)

		if config.DayOfYear > daysInYear(yearToUse) || config.DayOfYear == 0 {
			config.OverflowDayOfYear = true
		}

		date := time.Date(yearToUse, 1, config.DayOfYear, 0, 0, 0, 0, time.UTC)
		config.ParsedArray[monthIdx] = int(date.Month())
		config.ParsedArray[dateIdx] = date.Day()
	}

	// Default to current date.
	// If no year, month, day of month are given, default to today.
	// If day of month is given, default month and year.
	// If month is given, default only year.
	// If year is given, don't default anything.
	for i := 0; i < 3 && !keyExists(i, config.ParsedArray); i++ {
		if _, exist := config.ParsedArray[i]; !exist {
			config.ParsedArray[i] = currentDate[i]
		}
	}

	// Zero out any values that weren't default, including time.
	// If its month or day, set the value to 1.
	for i := 0; i < 7; i++ {
		defaultVal := 0
		if !keyExists(i, config.ParsedArray) {
			if i == 1 || i == 2 {
				defaultVal = 1
			}
			config.ParsedArray[i] = defaultVal
		}
	}

	createDateFromConfig(config)

	return config.Date, nil
}

func fixForMeridiem(config *parseConfig) {
	if config.Meridiem == "" {
		return
	}

	hour := config.ParsedArray[hourIdx]
	isPM := strings.ToLower(config.Meridiem)[0] == 'p'

	if isPM && hour < 12 {
		hour += 12
	}
	if !isPM && hour == 12 {
		hour = 0
	}

	config.ParsedArray[hourIdx] = hour
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
	if config.IsUTC {
		loc = time.UTC
	}

	datetime := DateTime{
		config.ParsedArray[yearIdx],
		config.ParsedArray[monthIdx],
		config.ParsedArray[dateIdx],
		config.ParsedArray[hourIdx],
		config.ParsedArray[minuteIdx],
		config.ParsedArray[secondIdx],
		config.ParsedArray[nanosecondIdx],
		loc,
	}

	d, _ := New(datetime)
	if config.TzMinutes != -99999 {
		// Call this method rather than SetMinute, since it checks that minute is between 0-59.
		d.subtractMinutes(config.TzMinutes)

		// Set the offset.
		d.SetUTCOffset(config.TzMinutes)
	}

	config.Date = d
}

func currentDateArray(config *parseConfig) map[int]int {
	newDate, _ := New()
	if config.IsUTC {
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

func parseMonth(input string, config *parseConfig) {
	config.ParsedArray[monthIdx] = parseNumber(input)
}

func parseLongMonth(input string, config *parseConfig) {
	config.ParsedArray[monthIdx] = constants.LongMonthNames[strings.ToLower(input)]
}

func parseShortMonth(input string, config *parseConfig) {
	config.ParsedArray[monthIdx] = constants.ShortMonthNames[strings.ToLower(input)]
}

func parseSingleDigitYear(input string, config *parseConfig) {
	y, _ := strconv.Atoi(input)
	config.ParsedArray[yearIdx] = y
}

func parseTwoDigitYear(input string, config *parseConfig) {
	baseYear := 0
	year := parseNumber(input)

	if year > 68 {
		baseYear = 1900
	} else {
		baseYear = 2000
	}

	config.ParsedArray[yearIdx] = year + baseYear
}

func parseOrdinalDate(input string, config *parseConfig) {
	config.ParsedArray[dateIdx] = parseNumber(regexps.MatchOneToTwo.FindAllString(input, -1)[0])
}

func parseDayOfYear(input string, config *parseConfig) {
	config.DayOfYear = parseNumber(strings.TrimLeft(input, "0"))
}

func parseQuarter(input string, config *parseConfig) {
	config.ParsedArray[monthIdx] = parseNumber(input) * 3
}

func parseTimestamp(input string, config *parseConfig) {
	ts, _ := strconv.ParseInt(input, 10, 64)
	g, _ := Unix(ts)

	config.Date = g
}

func parseOneToTwentyFourTime(input string, config *parseConfig) {
	config.ParsedArray[hourIdx] = parseNumber(input) - 1
}

func parseMeridiem(input string, config *parseConfig) {
	config.Meridiem = input
}

func parseOffset(input string, config *parseConfig) {
	config.IsUTC = true

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

		config.TzMinutes = minutes
	}
}

func parseLongDayName(input string, config *parseConfig) {
	longDayName := strings.ToLower(input)
	if val, ok := constants.LongDayNames[longDayName]; ok {
		createWeekConfig(config)
		config.Week["day"] = val
	}
}

func parseShortDayName(input string, config *parseConfig) {
	shortDayName := strings.ToLower(input)
	if val, ok := constants.ShortDayNames[shortDayName]; ok {
		createWeekConfig(config)
		config.Week["day"] = val
	}
}

func createWeekConfig(config *parseConfig) {
	if config.Week == nil {
		config.Week = map[string]int{}
	}
}

func parseNumber(dateText string) int {
	num, _ := strconv.Atoi(dateText)
	return num
}
