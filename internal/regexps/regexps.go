package regexps

import "regexp"

// TokenRegex is used to parse tokens out of formats.
var TokenRegex = regexp.MustCompile("(LT[S]?|LL?L?L?|l{1,4}|Mo|MM?M?M?|Do|DDDo|DD?D?D?|ddd?d?|do?|w[o|w]?|W[o|W]?|YYYYY|YYYY|YY?|gg(ggg?)?|GG(GGG?)?|e|E|a|A|hh?|HH?|kk?|mm?|ss?|SS?S?|X|zz?|ZZ?|Q)")

// BracketRegex is used to find brackets in formats.
var BracketRegex = regexp.MustCompile(`\[([^\[\]]*)\]`)

// BasicISORegex is used to parse simple ISO 8601 dates.
var BasicISORegex = regexp.MustCompile(`^\s*((?:[+-]\d{6}|\d{4})(?:\d\d\d\d|W\d\d\d|W\d\d|\d\d\d|\d\d))(?:(T| )(\d\d(?:\d\d(?:\d\d(?:[.,]\d+)?)?)?)([\+\-]\d\d(?::?\d\d)?|\s*Z)?)?$`)

// ExtendedISORegex is used to parse extends ISO 8601 dates.
// 0000-00-00 0000-W00 or 0000-W00-0 + T + 00 or 00:00 or 00:00:00 or 00:00:00.000 + +00:00 or +0000 or +00)
var ExtendedISORegex = regexp.MustCompile(`^\s*((?:[+-]\d{6}|\d{4})-(?:\d\d-\d\d|W\d\d-\d|W\d\d|\d\d\d|\d\d))(?:(T| )(\d\d(?::\d\d(?::\d\d(?:[.,]\d+)?)?)?)([\+\-]\d\d(?::?\d\d)?|\s*Z)?)?$`)

// MonthsRegex is used to parse month names.
var MonthsRegex = regexp.MustCompile(`(?i)(january|february|march|april|may|june|july|august|september|october|november|december|jan\.?|feb\.?|mar\.?|apr\.?|may\.?|jun\.?|jul\.?|aug\.?|sep\.?|oct\.?|nov\.?|dec\.?)`)

// WeekdaysRegex is used to parse weekday names.
var WeekdaysRegex = regexp.MustCompile(`(?i)(sunday|monday|tuesday|wednesday|thursday|friday|saturday|sun\.?|mon\.?|tue\.?|wed\.?|thu\.?|fri\.?|sat\.?)`)

// TimeZoneRegex is used to parse timezones.
var TimeZoneRegex = regexp.MustCompile(`Z|[+-]\d\d(?::?\d\d)?`)

// MatchOne is used to match a single digit.
var MatchOne = regexp.MustCompile(`\d`)

// MatchOneToTwo is used to match a one to two digits.
var MatchOneToTwo = regexp.MustCompile(`\d\d?`)

// MatchOneToFour is used to match a one to four digits.
var MatchOneToFour = regexp.MustCompile(`\d{1,4}`)

// MatchUnsigned is used to match unsigned digits.
var MatchUnsigned = regexp.MustCompile(`\d+`)

// MatchSigned is used to match signed digits.
var MatchSigned = regexp.MustCompile(`[+-]?\d+`)

// MatchOneToThree is used to match one to three digits.
var MatchOneToThree = regexp.MustCompile(`\d{1,3}`)

// MatchThree is used to match exactly 3 digits.
var MatchThree = regexp.MustCompile(`\d{3}`)

// DayOfMonthOrdinal is used to match ordinal dates.
var DayOfMonthOrdinal = regexp.MustCompile(`\d{1,2}(th|st|nd|rd)`)

// MatchMeridiem is used to match meridiem field.
var MatchMeridiem = regexp.MustCompile(`(?i)(am|pm)`)

// MatchTimestamp is used to match unix timestamps.
var MatchTimestamp = regexp.MustCompile(`[+-]?\d+(\.\d{1,3})?`)

// MatchShortOffset is used to match short timezone offsets.
var MatchShortOffset = regexp.MustCompile(`(?i)(Z|[+-]\d\d(?::?\d\d)?)`)

// ChunkOffset is used to parse timezone offset.
var ChunkOffset = regexp.MustCompile(`(?i)([\+\-]|\d\d)`)
