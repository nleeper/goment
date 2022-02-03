package locales

import (
	"fmt"
	"strings"
)

// FaLocale is the FA Persian language locale.
var FaLocale = newLocale(
	"fa",
	strings.Split("یکشنبه_دوشنبه_سه‌شنبه_چهارشنبه_پنج‌شنبه_جمعه_شنبه", "_"),
	strings.Split("یک_دو_سه_چهار_پنج_جمعه_شنبه", "_"),
	strings.Split("یک_دو_سه_چه_پن_جم_شن", "_"),
	strings.Split("ژانویه_فوریه_مارس_آوریل_مه_ژوئن_ژوئیه_اوت_سپتامبر_اکتبر_نوامبر_دسامبر", "_"),
	strings.Split("ژان_فور_مار_آور_مه_ژون_ژوئ_اوت_سپت_اکت_نوا_دسا", "_"),
	func(num int, period string) string {
		return fmt.Sprintf("%d روز", num)
	},
	nil,
	week{Dow: 0, Doy: 6},
	longDateFormats{
		"LTS":  "h:mm:ss A",
		"LT":   "h:mm A",
		"L":    "MM/DD/YYYY",
		"LL":   "MMMM D, YYYY",
		"LLL":  "MMMM D, YYYY h:mm A",
		"LLLL": "dddd, MMMM D, YYYY h:mm A",
	},
	relativeTimeFormats{
		"future": "در %s",
		"past":   "%s پیش",
		"s":      "چند ثانیه پیش",
		"ss":     "%d ثانیه",
		"m":      "یک دقیقه",
		"mm":     "%d دقیقه",
		"h":      "یک ساعت",
		"hh":     "%d ساعت",
		"d":      "یک روز",
		"dd":     "%d روز",
		"M":      "یک ماه",
		"MM":     "%d ماه",
		"y":      "یک سال",
		"yy":     "%d سال",
	},
	calendarFunctions{
		"sameDay": func(hours int, day int) string {
			return "[امروز] LT"
		},
		"nextDay": func(hours int, day int) string {
			return "[فردا] LT"
		},
		"nextWeek": func(hours int, day int) string {
			return "dddd [در ساعت] LT"
		},
		"lastDay": func(hours int, day int) string {
			return "[دیروز] LT"
		},
		"lastWeek": func(hours int, day int) string {
			return "[هفته پیش] dddd [در ساعت] LT"
		},
		"sameElse": func(hours int, day int) string {
			return "L"
		},
	},
	`(?i)(ژانویه|فوریه|مارس|آوریل|مه|ژوئن|ژوئیه|اوت|سپتامبر|اکتبر|نوامبر|دسامبر)`,
	`(?i)(ژان|فور|مار|آور|مه|ژون|ژوئ|اوت|سپت|اکت|نوا|دسا)`,
	`(?i)(یکشنبه|دوشنبه|سه‌شنبه|چهارشنبه|پنج‌شنبه|جمعه|شنبه)`,
	`(?i)(یک|دو|سه|چهار|پنج|جمعه|شنبه)`,
	`(?i)(یک|دو|سه|چه|پن|جم|شن)`,
	`\d{1,2}روز`,
)
