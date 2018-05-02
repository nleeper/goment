package goment

import (
	"regexp"
	"time"

	"github.com/pkg/errors"
)

// iso 8601 regex
// 0000-00-00 0000-W00 or 0000-W00-0 + T + 00 or 00:00 or 00:00:00 or 00:00:00.000 + +00:00 or +0000 or +00)
var extendedISORegex = regexp.MustCompile(`^\s*((?:[+-]\d{6}|\d{4})-(?:\d\d-\d\d|W\d\d-\d|W\d\d|\d\d\d|\d\d))(?:(T| )(\d\d(?::\d\d(?::\d\d(?:[.,]\d+)?)?)?)([\+\-]\d\d(?::?\d\d)?|\s*Z)?)?$`)
var basicISORegex = regexp.MustCompile(`^\s*((?:[+-]\d{6}|\d{4})(?:\d\d\d\d|W\d\d\d|W\d\d|\d\d\d|\d\d))(?:(T| )(\d\d(?:\d\d(?:\d\d(?:[.,]\d+)?)?)?)([\+\-]\d\d(?::?\d\d)?|\s*Z)?)?$`)

var tzRegex = regexp.MustCompile(`Z|[+-]\d\d(?::?\d\d)?`)

var isoDates = []isoDateFormat{
	// newISOFormat("+002006-01-02", `\+\d{6}-\d\d-\d\d`),
	// newISOFormat("-002006-01-02", `-\d{6}-\d\d-\d\d`),
	newISODateFormat("2006-01-02", `\d{4}-\d\d-\d\d`),
	newISODateFormat("2006-01", `\d{4}-\d\d`, false),
	newISODateFormat("20060102", `\d{8}`),
}

var isoTimes = []isoTimeFormat{
	newISOTimeFormat("15:04:05.9999", `\d\d:\d\d:\d\d\.\d+`),
	// newISOTimeFormat("15:04:05,9999", `\d\d:\d\d:\d\d,\d+`), // TODO - will need to replace the comma with a period in date string
	newISOTimeFormat("15:04:05", `\d\d:\d\d:\d\d`),
	newISOTimeFormat("15:04", `\d\d:\d\d`),
	newISOTimeFormat("150405.9999", `\d\d\d\d\d\d\.\d+`),
	// newISOTimeFormat("150405,9999", `\d\d\d\d\d\d,\d+`), // TODO - replace comma
	newISOTimeFormat("150405", `\d\d\d\d\d\d`),
	newISOTimeFormat("1504", `\d\d\d\d`),
	newISOTimeFormat("15", `\d\d`),
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

type isoTimeFormat struct {
	Format string
	Regex  *regexp.Regexp
}

type isoDateFormat struct {
	Format    string
	Regex     *regexp.Regexp
	AllowTime bool
}

func newISOTimeFormat(format, regex string) isoTimeFormat {
	return isoTimeFormat{
		format,
		regexp.MustCompile(regex),
	}
}

func newISODateFormat(format, regex string, allowTime ...bool) isoDateFormat {
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

func parseISOString(date string) (time.Time, error) {
	match := extendedISORegex.FindStringSubmatch(date)
	if match == nil {
		match = basicISORegex.FindStringSubmatch(date)
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
			if tzRegex.MatchString(match[4]) {
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
