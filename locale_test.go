package goment

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func resetLocale() {
	SetLocale(DefaultLocaleCode)
}

func TestDefaultGlobalLocale(t *testing.T) {
	assert.Equal(t, DefaultLocaleCode, Locale(), "Default global locale")
}
func TestDefaultGlobalLocaleUsedForNew(t *testing.T) {
	lib, _ := New()
	assert.Equal(t, DefaultLocaleCode, lib.Locale(), "Default global locale used on new")
}

func TestSetGlobalLocale(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(DefaultLocaleCode, Locale())
	SetLocale("es")
	assert.Equal("es", Locale(), "Set global locale")

	resetLocale()
}

func TestSetGlobalLocaleUsedForNew(t *testing.T) {
	SetLocale("es")
	lib, _ := New()

	assert.Equal(t, "es", lib.Locale(), "Use set global locale")

	resetLocale()
}

func TestSetLocale(t *testing.T) {
	assert := assert.New(t)

	lib, _ := New()
	assert.Equal(DefaultLocaleCode, lib.Locale())

	lib.SetLocale("fr")
	assert.Equal("fr", lib.Locale())
}

func TestLocaleNotChangedForExistingGoment(t *testing.T) {
	assert := assert.New(t)

	lib, _ := New()
	lib.SetLocale("fr")
	assert.Equal("fr", lib.Locale())

	SetLocale("es")
	assert.Equal("fr", lib.Locale())
	assert.Equal("es", Locale())

	lib2, _ := New()
	assert.Equal("es", lib2.Locale())

	resetLocale()
}

func TestEnLocale(t *testing.T) {
	assert := assert.New(t)

	longDays := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	shortDays := []string{"Sun", "Mon", "Tue", "Wed", "Thu", "Fri", "Sat"}
	minDays := []string{"Su", "Mo", "Tu", "We", "Th", "Fr", "Sa"}
	shortMonths := []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}
	longMonths := []string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}

	lib := simpleNow()

	assert.Equal("en", lib.Locale())
	assert.Equal(longDays, lib.Weekdays())
	assert.Equal(shortDays, lib.WeekdaysShort())
	assert.Equal(minDays, lib.WeekdaysMin())
	assert.Equal(longMonths, lib.Months())
	assert.Equal(shortMonths, lib.MonthsShort())
}

func TestEnMonthByNumber(t *testing.T) {
	assert := assert.New(t)

	lib := simpleNow()

	assert.Equal("January", lib.MonthByNumber(1))
	assert.Equal("June", lib.MonthByNumber(6))
	assert.Equal("December", lib.MonthByNumber(12))
}

func TestEnMonthByNumberInvalid(t *testing.T) {
	assert := assert.New(t)

	lib := simpleNow()

	assert.Equal("", lib.MonthByNumber(0))
	assert.Equal("", lib.MonthByNumber(13))
}

func TestEnMonthShortByNumber(t *testing.T) {
	assert := assert.New(t)

	lib := simpleNow()

	assert.Equal("Jan", lib.MonthShortByNumber(1))
	assert.Equal("Jun", lib.MonthShortByNumber(6))
	assert.Equal("Dec", lib.MonthShortByNumber(12))
}

func TestEnMonthShortByNumberInvalid(t *testing.T) {
	assert := assert.New(t)

	lib := simpleNow()

	assert.Equal("", lib.MonthShortByNumber(0))
	assert.Equal("", lib.MonthShortByNumber(13))
}

func TestEnWeekdayByNumber(t *testing.T) {
	assert := assert.New(t)

	lib := simpleNow()

	assert.Equal("Monday", lib.WeekdayByNumber(1))
	assert.Equal("Wednesday", lib.WeekdayByNumber(3))
	assert.Equal("Saturday", lib.WeekdayByNumber(6))
}

func TestEnWeekdayByNumberLocale(t *testing.T) {
	assert := assert.New(t)

	lib := simpleNow()

	assert.Equal("Sunday", lib.WeekdayByNumber(true, 0))
	assert.Equal("Sunday", lib.WeekdayByNumber(0))
	assert.Equal("Wednesday", lib.WeekdayByNumber(true, 3))
	assert.Equal("Wednesday", lib.WeekdayByNumber(3))
	assert.Equal("Saturday", lib.WeekdayByNumber(6))
}

func TestEnWeekdayByNumberInvalid(t *testing.T) {
	assert := assert.New(t)

	lib := simpleNow()

	assert.Equal("", lib.WeekdayByNumber())
	assert.Equal("", lib.WeekdayByNumber(6, 1, true))
}

func TestEnWeekdaysWithLocaleDow(t *testing.T) {
	longDays := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}

	lib := simpleNow()

	assert.Equal(t, longDays, lib.Weekdays(true))
}

func TestEsLocale(t *testing.T) {
	assert := assert.New(t)

	longDays := []string{"domingo", "lunes", "martes", "miércoles", "jueves", "viernes", "sábado"}
	shortDays := []string{"dom.", "lun.", "mar.", "mié.", "jue.", "vie.", "sáb."}
	minDays := []string{"do", "lu", "ma", "mi", "ju", "vi", "sá"}
	shortMonths := []string{"ene", "feb", "mar", "abr", "may", "jun", "jul", "ago", "sep", "oct", "nov", "dic"}
	longMonths := []string{"enero", "febrero", "marzo", "abril", "mayo", "junio", "julio", "agosto", "septiembre", "octubre", "noviembre", "diciembre"}

	lib, _ := New()
	lib.SetLocale("es")

	assert.Equal("es", lib.Locale())
	assert.Equal(longDays, lib.Weekdays())
	assert.Equal(shortDays, lib.WeekdaysShort())
	assert.Equal(minDays, lib.WeekdaysMin())
	assert.Equal(longMonths, lib.Months())
	assert.Equal(shortMonths, lib.MonthsShort())
}

func TestEsWeekdaysWithLocaleDow(t *testing.T) {
	longDays := []string{"lunes", "martes", "miércoles", "jueves", "viernes", "sábado", "domingo"}

	lib := simpleNow()
	lib.SetLocale("es")

	assert.Equal(t, longDays, lib.Weekdays(true))
}

func TestEsWeekdaysShortWithLocaleDow(t *testing.T) {
	longDays := []string{"lun.", "mar.", "mié.", "jue.", "vie.", "sáb.", "dom."}

	lib := simpleNow()
	lib.SetLocale("es")

	assert.Equal(t, longDays, lib.WeekdaysShort(true))
}

func TestEsWeekdaysMinWithLocaleDow(t *testing.T) {
	longDays := []string{"lu", "ma", "mi", "ju", "vi", "sá", "do"}

	lib := simpleNow()
	lib.SetLocale("es")

	assert.Equal(t, longDays, lib.WeekdaysMin(true))
}

func TestEsWeekdayByNumberLocale(t *testing.T) {
	assert := assert.New(t)

	lib := simpleNow()
	lib.SetLocale("es")

	assert.Equal("lunes", lib.WeekdayByNumber(true, 0))  // first dow per locale, Monday
	assert.Equal("domingo", lib.WeekdayByNumber(0))      // first dow is Sunday, Sunday
	assert.Equal("jueves", lib.WeekdayByNumber(true, 3)) // first dow per locale, Thursday
	assert.Equal("miércoles", lib.WeekdayByNumber(3))    // first dow is Sunday, Wednesday
	assert.Equal("sábado", lib.WeekdayByNumber(6))       // first dow is Sunday, Saturday
}

func TestEsFormat(t *testing.T) {
	assert := assert.New(t)

	formats := map[string]string{
		"dddd, MMMM Do YYYY, h:mm:ss a": "domingo, febrero 14º 2010, 3:25:50 pm",
		"ddd, hA":                       "dom., 3PM",
		"M Mo MM MMMM MMM":              "2 2º 02 febrero feb",
		"YYYY YY":                       "2010 10",
		"D Do DD":                       "14 14º 14",
		"d do dddd ddd dd":              "0 0º domingo dom. do",
		"DDD DDDo DDDD":                 "45 45º 045",
		"w wo ww":                       "6 6º 06",
		"YYYY-MMM-DD":                   "2010-feb-14",
		"h hh":                          "3 03",
		"H HH":                          "15 15",
		"m mm":                          "25 25",
		"s ss":                          "50 50",
		"a A":                           "pm PM",
		"[the] DDDo [day of the year]":  "the 45º day of the year",
		"[the] DDDo [day of the year is after January]": "the 45º day of the year is after January",
		"LT":            "15:25",
		"LTS":           "15:25:50",
		"L":             "14/02/2010",
		"LL":            "14 de febrero de 2010",
		"[today is] LL": "today is 14 de febrero de 2010",
		"LLL":           "14 de febrero de 2010 15:25",
		"LLLL":          "domingo, 14 de febrero de 2010 15:25",
		"l":             "14/2/2010",
		"ll":            "14 de feb de 2010",
		"lll":           "14 de feb de 2010 15:25",
		"llll":          "dom., 14 de feb de 2010 15:25",
	}

	lib := simpleTime(time.Date(2010, 2, 14, 15, 25, 50, 125000000, chicagoLocation()))
	lib.SetLocale("es")

	for p, r := range formats {
		assert.Equal(r, lib.Format(p), r)
	}
}

func TestEsShortDayFormatWithUnicode(t *testing.T) {
	lib := simpleTime(time.Date(2010, 2, 20, 15, 25, 50, 125000000, chicagoLocation()))
	lib.SetLocale("es")

	assert.Equal(t, "sá", lib.Format("dd"))
}

func TestEsRelativeTime(t *testing.T) {
	assert := assert.New(t)

	testTime := time.Date(2007, 1, 28, 0, 0, 0, 0, chicagoLocation())
	lib := simpleTime(testTime)

	lib.SetLocale("es")

	assert.Equal("hace unos segundos", lib.From(simpleTime(testTime).Add(44, "s")), "44 seconds = a few seconds ago")
	assert.Equal("unos segundos", lib.From(simpleTime(testTime).Add(44, "s"), true), "44 seconds = a few seconds")
	assert.Equal("hace un minuto", lib.From(simpleTime(testTime).Add(1, "m")), "1 minute = a minute ago")
	assert.Equal("un minuto", lib.From(simpleTime(testTime).Add(1, "m"), true), "1 minute = a minute")
	assert.Equal("hace 44 minutos", lib.From(simpleTime(testTime).Add(44, "m")), "44 minutes = 44 minutes ago")
	assert.Equal("44 minutos", lib.From(simpleTime(testTime).Add(44, "m"), true), "44 minutes = 44 minutes")
	assert.Equal("hace una hora", lib.From(simpleTime(testTime).Add(1, "h")), "1 hour = an hour ago")
	assert.Equal("una hora", lib.From(simpleTime(testTime).Add(1, "h"), true), "1 hour = an hour")
	assert.Equal("hace 2 horas", lib.From(simpleTime(testTime).Add(2, "h")), "2 hours = 2 hours ago")
	assert.Equal("2 horas", lib.From(simpleTime(testTime).Add(2, "h"), true), "2 hours = 2 hours")
	assert.Equal("hace un día", lib.From(simpleTime(testTime).Add(1, "d")), "1 day = a day ago")
	assert.Equal("un día", lib.From(simpleTime(testTime).Add(1, "d"), true), "1 day = a day")
	assert.Equal("hace 5 días", lib.From(simpleTime(testTime).Add(5, "d")), "5 days = 5 days ago")
	assert.Equal("5 días", lib.From(simpleTime(testTime).Add(5, "d"), true), "5 days = 5 days")
	assert.Equal("hace un mes", lib.From(simpleTime(testTime).Add(1, "M")), "1 month = a month ago")
	assert.Equal("un mes", lib.From(simpleTime(testTime).Add(1, "M"), true), "1 month = a month")
	assert.Equal("hace 5 meses", lib.From(simpleTime(testTime).Add(5, "M")), "5 months = 5 months ago")
	assert.Equal("5 meses", lib.From(simpleTime(testTime).Add(5, "M"), true), "5 months = 5 months")
	assert.Equal("hace un año", lib.From(simpleTime(testTime).Add(1, "y")), "1 year = a year ago")
	assert.Equal("un año", lib.From(simpleTime(testTime).Add(1, "y"), true), "1 year = a year")
	assert.Equal("hace 5 años", lib.From(simpleTime(testTime).Add(5, "y")), "5 years = 5 years ago")
	assert.Equal("5 años", lib.From(simpleTime(testTime).Add(5, "y"), true), "5 years = 5 years")
}

func TestEsCalendarDay(t *testing.T) {
	assert := assert.New(t)

	testTime := time.Date(2000, 12, 15, 12, 0, 0, 0, time.UTC)
	timeNow = func() time.Time {
		return testTime
	}

	SetLocale("es")

	refTime := simpleTime(testTime)

	assert.Equal("hoy a las 12:00", simpleGoment(refTime).Calendar(), "today at the same time")
	assert.Equal("hoy a las 12:25", simpleGoment(refTime).Add(25, "m").Calendar(), "now plus 25 min")
	assert.Equal("hoy a las 13:00", simpleGoment(refTime).Add(1, "h").Calendar(), "now plus 1 hour")
	assert.Equal("mañana a las 12:00", simpleGoment(refTime).Add(1, "d").Calendar(), "tomorrow at the same time")
	assert.Equal("hoy a las 11:00", simpleGoment(refTime).Subtract(1, "h").Calendar(), "now minus 1 hour")
	assert.Equal("ayer a las 12:00", simpleGoment(refTime).Subtract(1, "d").Calendar(), "yesterday at the same time")

	refTime = simpleTime(testTime)

	assert.Equal("domingo a las 12:00", refTime.Add(2, "d").Calendar(), "Today + 2 days current time")
	refTime.StartOf("day")
	assert.Equal("domingo a las 0:00", refTime.Calendar(), "Today + 2 days beginning of day")
	refTime.EndOf("day")
	assert.Equal("domingo a las 23:59", refTime.Calendar(), "Today + 2 days end of day")

	refTime = simpleTime(testTime)

	assert.Equal("el miércoles pasado a las 12:00", refTime.Subtract(2, "d").Calendar(), "Today - 2 days current time")
	refTime.StartOf("day")
	assert.Equal("el miércoles pasado a las 0:00", refTime.Calendar(), "Today - 2 days beginning of day")
	refTime.EndOf("day")
	assert.Equal("el miércoles pasado a las 23:59", refTime.Calendar(), "Today - 2 days end of day")

	weeksAgo := simpleTime(testTime).Subtract(1, "w")
	weeksFromNow := simpleTime(testTime).Add(1, "w")

	assert.Equal("08/12/2000", weeksAgo.Calendar())
	assert.Equal("22/12/2000", weeksFromNow.Calendar())

	weeksAgo = simpleTime(testTime).Subtract(2, "w")
	weeksFromNow = simpleTime(testTime).Add(2, "w")

	assert.Equal("01/12/2000", weeksAgo.Calendar())
	assert.Equal("29/12/2000", weeksFromNow.Calendar())

	// Reset timeNow.
	timeNow = time.Now

	SetLocale("en")
}

func TestEsFormatParsing(t *testing.T) {
	assert := assert.New(t)

	formats := map[string][]string{
		"YYYY-Q":                    {"2014-4"},
		"MM-DD-YYYY":                {"12-02-1999"},
		"DD-MM-YYYY":                {"12-02-1999"},
		"DD/MM/YYYY":                {"12/02/1999"},
		"DD_MM_YYYY":                {"12_02_1999"},
		"DD:MM:YYYY":                {"12:02:1999"},
		"D-M-YY":                    {"2-2-99"},
		"Y":                         {"-0025"},
		"YY":                        {"99"},
		"DDD-YYYY":                  {"300-1999"},
		"YYYY-DDD":                  {"1999-300"},
		"YYYY MM Do":                {"2014 01 3º", "2015 11 21º"},
		"MMM":                       {"abr"},
		"MMMM":                      {"septiembre"},
		"DD MMMM":                   {"11 septiembre"},
		"Do MMMM":                   {"3º septiembre"},
		"YYYY MMMM":                 {"2018 octubre"},
		"D":                         {"3", "27"},
		"DD":                        {"04", "23"},
		"DDD":                       {"7", "300"},
		"DDDD":                      {"008", "211", "312"},
		"h":                         {"4"},
		"H":                         {"1", "10", "23"},
		"DD-MM-YYYY h:m:s":          {"12-02-1999 2:45:10"},
		"DD-MM-YYYY h:m:s a":        {"12-02-1999 2:45:10 am", "12-02-1999 2:45:10 pm"},
		"h:mm a":                    {"12:00 pm", "12:30 pm", "12:00 am", "12:30 am"},
		"HH:mm":                     {"12:00"},
		"kk:mm":                     {"12:00"},
		"YYYY-MM-DDTHH:mm:ss":       {"2011-11-11T11:11:11"},
		"MM-DD-YYYY [M]":            {"12-02-1999 M"},
		"ddd MMM DD HH:mm:ss YYYY":  {"mié. abr 08 22:52:51 2009"},
		"dddd MMM DD HH:mm:ss YYYY": {"sábado abr 11 22:52:51 2009"},
		"HH:mm:ss":                  {"12:00:00", "12:30:11", "00:00:00"},
		"kk:mm:ss":                  {"12:00:10", "12:30:42", "24:00:00", "09:00:00"},
		"YYYY-MM-DD HH:mm:ss ZZ":    {"2000-05-15 17:08:00 -0700"},
		"YYYY-MM-DD HH:mm Z":        {"2010-10-20 04:30 +00:00"},
		"e":                         {"0", "5"},
		"E":                         {"1", "7"},
		// "HH:mm:ss.S": []string{"00:30:00.1"},
		// "HH:mm:ss SS":               "00:30:00 12",
		// "HH:mm:ss SSS":              "00:30:00 123",
		// "HH:mm:ss S":                "00:30:00 7",
		// "HH:mm:ss SS":               "00:30:00 78",
		// "HH:mm:ss SSS":              "00:30:00 789",
		// "kk:mm:ss S":   "24:30:00 1",
		// "kk:mm:ss SS":  "24:30:00 12",
		// "kk:mm:ss SSS": "24:30:00 123",
		// "kk:mm:ss S":   "24:30:00 7",
		// "kk:mm:ss SS":  "24:30:00 78",
		// "kk:mm:ss SSS": "24:30:00 789",
		"X":    {"1234567890"},
		"H Z":  {"6 -06:00"},
		"H ZZ": {"5 -0700"},
		"LT":   {"12:30"},
		"LTS":  {"12:30:29"},
		"L":    {"09/02/1999"},
		"l":    {"9/2/1999"},
		"LL":   {"2 de septiembre de 1999"},
		"ll":   {"2 de sep de 1999"},
		"LLL":  {"2 de septiembre de 1999 12:30"},
		"lll":  {"2 de sep de 1999 12:30"},
		"LLLL": {"jueves, 2 de septiembre de 1999 12:30"},
		"llll": {"jue., 2 de sep de 1999 12:30"},
	}

	for format, dates := range formats {
		for _, date := range dates {
			lib, _ := New(date, format, "es")
			assert.Equal(date, lib.Format(format), fmt.Sprintf("%v: %v", format, date))
		}
	}
}

func TestEsInvalidOrdinal(t *testing.T) {
	date := "2014 01 4th"
	format := "YYYY MM Do"

	lib, _ := New(date, format, "es")
	assert.Equal(t, "2014 01 1º", lib.Format(format))
}

func TestEsFormatWeekdayParsing(t *testing.T) {
	assert := assert.New(t)

	lib := simpleFormatLocale("mié. abr 08 22:52:51 2009", "ddd MMM DD HH:mm:ss YYYY", "es")
	assert.Equal("mié. abr 08 22:52:51 2009", lib.Format("ddd MMM DD HH:mm:ss YYYY"))

	lib2 := simpleFormatLocale("sábado abr 11 22:52:51 2009", "dddd MMM DD HH:mm:ss YYYY", "es")
	assert.Equal("sábado abr 11 22:52:51 2009", lib2.Format("dddd MMM DD HH:mm:ss YYYY"))
}

func TestEsDayOfWeekParsing(t *testing.T) {
	assert := assert.New(t)

	outputFormat := "YYYY MM DD"

	date := simpleNow().Format(outputFormat)

	assert.Equal(
		simpleFormatLocale(date, outputFormat, "es").SetDay("lunes").Format(outputFormat),
		simpleFormatLocale("0", "e", "es").Format(outputFormat),
	)
	assert.Equal(
		simpleFormatLocale(date, outputFormat, "es").SetDay("jueves").Format(outputFormat),
		simpleFormatLocale("jue. 0", "ddd e", "es").Format(outputFormat),
	)
}

func TestEsWeekYearParsing(t *testing.T) {
	assert := assert.New(t)

	outputFormat := "YYYY-MM-DD"

	assert.Equal("2007-01-01", simpleFormatLocale("2007-01", "gggg-ww", "es").Format(outputFormat))
	assert.Equal("2007-12-31", simpleFormatLocale("2008-01", "gggg-ww", "es").Format(outputFormat))
	assert.Equal("2002-12-30", simpleFormatLocale("2003-01", "gggg-ww", "es").Format(outputFormat))
	assert.Equal("2008-12-29", simpleFormatLocale("2009-01", "gggg-ww", "es").Format(outputFormat))
	assert.Equal("2010-01-04", simpleFormatLocale("2010-01", "gggg-ww", "es").Format(outputFormat))
	assert.Equal("2011-01-03", simpleFormatLocale("2011-01", "gggg-ww", "es").Format(outputFormat))
	assert.Equal("2012-01-02", simpleFormatLocale("2012-01", "gggg-ww", "es").Format(outputFormat))
}

func TestEsWeekWeekdayParsing(t *testing.T) {
	assert := assert.New(t)

	outputFormat := "YYYY MM DD"

	assert.Equal("1999 09 16", simpleFormatLocale("1999 37 4", "GGGG WW EE", "es").Format(outputFormat), "iso ignores locale")
	assert.Equal("1999 09 19", simpleFormatLocale("1999 37 7", "GGGG WW EE", "es").Format(outputFormat), "iso ignores locale")

	assert.Equal("1999 09 13", simpleFormatLocale("1999 37 0", "gggg ww e", "es").Format(outputFormat), "localized e uses local doy and dow: 0 = monday")
	assert.Equal("1999 09 17", simpleFormatLocale("1999 37 4", "gggg ww e", "es").Format(outputFormat), "localized e uses local doy and dow: 4 = friday")

	assert.Equal("1999 09 13", simpleFormatLocale("1999 37 1", "gggg ww d", "es").Format(outputFormat), "localized d uses 0-indexed days: 1 = monday")
	assert.Equal("1999 09 13", simpleFormatLocale("1999 37 Lu", "gggg ww dd", "es").Format(outputFormat), "localized d uses 0-indexed days: Mo")
	assert.Equal("1999 09 13", simpleFormatLocale("1999 37 lun.", "gggg ww ddd", "es").Format(outputFormat), "localized d uses 0-indexed days: Mon")
	assert.Equal("1999 09 13", simpleFormatLocale("1999 37 lunes", "gggg ww dddd", "es").Format(outputFormat), "localized d uses 0-indexed days: Monday")

	assert.Equal("1999 09 16", simpleFormatLocale("1999 37 4", "gggg ww d", "es").Format(outputFormat), "localized d uses 0-indexed days: 4")

	// Sunday goes at the end of the week
	assert.Equal("1999 09 19", simpleFormatLocale("1999 37 0", "gggg ww d", "es").Format(outputFormat), "localized d uses 0-indexed days: 0 = sund")
	assert.Equal("1999 09 19", simpleFormatLocale("1999 37 Do", "gggg ww dd", "es").Format(outputFormat), "localized d uses 0-indexed days: 0 = sund")
}

func TestEsWeekday(t *testing.T) {
	assert := assert.New(t)

	// es 1st day of week is Monday.
	testTime := time.Date(2020, 7, 28, 18, 0, 0, 0, time.UTC) // Tuesday, July 28 2020

	lib := simpleTime(testTime)
	lib.SetLocale("es")

	assert.Equal(lib.Weekday(), 1)

	lib2 := simpleString("2016-09-05") // Monday, September 5 2016
	lib2.SetLocale("es")

	assert.Equal(0, lib2.Weekday()) // Monday
	assert.Equal(5, lib2.Date())

	lib2.SetWeekday(3) // Thursday
	assert.Equal(8, lib2.Date())

	lib2.SetWeekday(6) // Sunday
	assert.Equal(11, lib2.Date())

	lib2.SetDay("miércoles") // Wednesday
	assert.Equal(14, lib2.Date())
}

func TestEsSetWeekdayRangeExceeded(t *testing.T) {
	format := "YYYY-MM-DD"

	lib := simpleFormatLocale("2020-07-26", format, "es")
	assert.Equal(t, lib.Format(format), "2020-07-26")

	lib.SetWeekday(-7)
	assert.Equal(t, lib.Format(format), "2020-07-13")

	lib.SetWeekday(9)
	assert.Equal(t, lib.Format(format), "2020-07-22")
}

func TestEsSetDayByName(t *testing.T) {
	format := "YYYY-MM-DD"

	lib := simpleFormatLocale("2016-09-04", format, "es")
	assert.Equal(t, 6, lib.Weekday())

	lib.SetDay("miércoles") // Wednesday
	assert.Equal(t, 7, lib.Date())

	lib.SetDay("sábado") // Saturday
	assert.Equal(t, 10, lib.Date())
}

func TestEsStartEndOfWeek(t *testing.T) {
	now := simpleNow()
	now.SetLocale("es")

	assert.Equal(t, "lunes", now.StartOf("week").Format("dddd"))
	assert.Equal(t, "domingo", now.EndOf("week").Format("dddd"))
}

func TestGetWeekWithDow1Doy4(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(52, simpleLocale(DateTime{Year: 2012, Month: 1, Day: 1}, "es").Week(), "Jan 1 2012 should be week 52")
	assert.Equal(1, simpleLocale(DateTime{Year: 2012, Month: 1, Day: 2}, "es").Week(), "Jan 2 2012 should be week 1")
	assert.Equal(1, simpleLocale(DateTime{Year: 2012, Month: 1, Day: 8}, "es").Week(), "Jan 8 2012 should be week 1")
	assert.Equal(2, simpleLocale(DateTime{Year: 2012, Month: 1, Day: 9}, "es").Week(), "Jan 9 2012 should be week 2")
	assert.Equal(2, simpleLocale(DateTime{Year: 2012, Month: 1, Day: 15}, "es").Week(), "Jan 15 2012 should be week 2")
	assert.Equal(1, simpleLocale(DateTime{Year: 2007, Month: 1, Day: 1}, "es").Week(), "Jan 1 2007 should be week 1")
	assert.Equal(1, simpleLocale(DateTime{Year: 2007, Month: 1, Day: 7}, "es").Week(), "Jan 7 2007 should be week 1")
	assert.Equal(2, simpleLocale(DateTime{Year: 2007, Month: 1, Day: 8}, "es").Week(), "Jan 8 2007 should be week 2")
	assert.Equal(2, simpleLocale(DateTime{Year: 2007, Month: 1, Day: 14}, "es").Week(), "Jan 14 2007 should be week 2")
	assert.Equal(3, simpleLocale(DateTime{Year: 2007, Month: 1, Day: 15}, "es").Week(), "Jan 15 2007 should be week 3")
	assert.Equal(1, simpleLocale(DateTime{Year: 2007, Month: 12, Day: 31}, "es").Week(), "Dec 31 2007 should be week 1")
	assert.Equal(1, simpleLocale(DateTime{Year: 2008, Month: 1, Day: 1}, "es").Week(), "Jan 1 2008 should be week 1")
	assert.Equal(1, simpleLocale(DateTime{Year: 2008, Month: 1, Day: 6}, "es").Week(), "Jan 6 2008 should be week 1")
	assert.Equal(2, simpleLocale(DateTime{Year: 2008, Month: 1, Day: 7}, "es").Week(), "Jan 7 2008 should be week 2")
	assert.Equal(2, simpleLocale(DateTime{Year: 2008, Month: 1, Day: 13}, "es").Week(), "Jan 13 2008 should be week 2")
	assert.Equal(3, simpleLocale(DateTime{Year: 2008, Month: 1, Day: 14}, "es").Week(), "Jan 14 2008 should be week 3")
	assert.Equal(1, simpleLocale(DateTime{Year: 2002, Month: 12, Day: 30}, "es").Week(), "Dec 30 2002 should be week 1")
	assert.Equal(1, simpleLocale(DateTime{Year: 2003, Month: 1, Day: 1}, "es").Week(), "Jan 1 2003 should be week 1")
	assert.Equal(1, simpleLocale(DateTime{Year: 2003, Month: 1, Day: 5}, "es").Week(), "Jan 5 2003 should be week 1")
	assert.Equal(2, simpleLocale(DateTime{Year: 2003, Month: 1, Day: 6}, "es").Week(), "Jan 6 2003 should be week 2")
	assert.Equal(2, simpleLocale(DateTime{Year: 2003, Month: 1, Day: 12}, "es").Week(), "Jan 12 2003 should be week 2")
	assert.Equal(3, simpleLocale(DateTime{Year: 2003, Month: 1, Day: 13}, "es").Week(), "Jan 13 2003 should be week 3")
	assert.Equal(1, simpleLocale(DateTime{Year: 2008, Month: 12, Day: 29}, "es").Week(), "Dec 29 2008 should be week 1")
	assert.Equal(1, simpleLocale(DateTime{Year: 2009, Month: 1, Day: 1}, "es").Week(), "Jan 1 2009 should be week 1")
	assert.Equal(1, simpleLocale(DateTime{Year: 2009, Month: 1, Day: 4}, "es").Week(), "Jan 4 2009 should be week 1")
	assert.Equal(2, simpleLocale(DateTime{Year: 2009, Month: 1, Day: 5}, "es").Week(), "Jan 5 2009 should be week 2")
	assert.Equal(2, simpleLocale(DateTime{Year: 2009, Month: 1, Day: 11}, "es").Week(), "Jan 11 2009 should be week 2")
	assert.Equal(3, simpleLocale(DateTime{Year: 2009, Month: 1, Day: 13}, "es").Week(), "Jan 13 2009 should be week 3")
	assert.Equal(53, simpleLocale(DateTime{Year: 2009, Month: 12, Day: 28}, "es").Week(), "Dec 28 2009 should be week 53")
	assert.Equal(53, simpleLocale(DateTime{Year: 2010, Month: 1, Day: 1}, "es").Week(), "Jan 1 2010 should be week 53")
	assert.Equal(53, simpleLocale(DateTime{Year: 2010, Month: 1, Day: 3}, "es").Week(), "Jan 3 2010 should be week 53")
	assert.Equal(1, simpleLocale(DateTime{Year: 2010, Month: 1, Day: 4}, "es").Week(), "Jan 4 2010 should be week 1")
	assert.Equal(1, simpleLocale(DateTime{Year: 2010, Month: 1, Day: 10}, "es").Week(), "Jan 10 2010 should be week 1")
	assert.Equal(2, simpleLocale(DateTime{Year: 2010, Month: 1, Day: 11}, "es").Week(), "Jan 11 2010 should be week 2")
	assert.Equal(52, simpleLocale(DateTime{Year: 2010, Month: 12, Day: 27}, "es").Week(), "Dec 27 2010 should be week 52")
	assert.Equal(52, simpleLocale(DateTime{Year: 2011, Month: 1, Day: 1}, "es").Week(), "Jan 1 2011 should be week 52")
	assert.Equal(52, simpleLocale(DateTime{Year: 2011, Month: 1, Day: 2}, "es").Week(), "Jan 2 2011 should be week 52")
	assert.Equal(1, simpleLocale(DateTime{Year: 2011, Month: 1, Day: 3}, "es").Week(), "Jan 3 2011 should be week 1")
	assert.Equal(1, simpleLocale(DateTime{Year: 2011, Month: 1, Day: 9}, "es").Week(), "Jan 9 2011 should be week 1")
	assert.Equal(2, simpleLocale(DateTime{Year: 2011, Month: 1, Day: 10}, "es").Week(), "Jan 10 2011 should be week 2")
}

func TestGetWeekYearWithDow1Doy4(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(2004, simpleLocale(DateTime{Year: 2005, Month: 1, Day: 1}, "es").WeekYear())
	assert.Equal(2004, simpleLocale(DateTime{Year: 2005, Month: 1, Day: 2}, "es").WeekYear())
	assert.Equal(2005, simpleLocale(DateTime{Year: 2005, Month: 1, Day: 3}, "es").WeekYear())
	assert.Equal(2005, simpleLocale(DateTime{Year: 2005, Month: 12, Day: 31}, "es").WeekYear())
	assert.Equal(2005, simpleLocale(DateTime{Year: 2006, Month: 1, Day: 1}, "es").WeekYear())
	assert.Equal(2006, simpleLocale(DateTime{Year: 2006, Month: 1, Day: 2}, "es").WeekYear())
	assert.Equal(2007, simpleLocale(DateTime{Year: 2007, Month: 1, Day: 1}, "es").WeekYear())
	assert.Equal(2007, simpleLocale(DateTime{Year: 2007, Month: 12, Day: 30}, "es").WeekYear())
	assert.Equal(2008, simpleLocale(DateTime{Year: 2007, Month: 12, Day: 31}, "es").WeekYear())
	assert.Equal(2008, simpleLocale(DateTime{Year: 2008, Month: 1, Day: 1}, "es").WeekYear())
	assert.Equal(2008, simpleLocale(DateTime{Year: 2008, Month: 12, Day: 28}, "es").WeekYear())
	assert.Equal(2009, simpleLocale(DateTime{Year: 2008, Month: 12, Day: 29}, "es").WeekYear())
	assert.Equal(2009, simpleLocale(DateTime{Year: 2008, Month: 12, Day: 30}, "es").WeekYear())
	assert.Equal(2009, simpleLocale(DateTime{Year: 2008, Month: 12, Day: 31}, "es").WeekYear())
	assert.Equal(2009, simpleLocale(DateTime{Year: 2009, Month: 1, Day: 1}, "es").WeekYear())
	assert.Equal(2009, simpleLocale(DateTime{Year: 2010, Month: 1, Day: 1}, "es").WeekYear())
	assert.Equal(2009, simpleLocale(DateTime{Year: 2010, Month: 1, Day: 2}, "es").WeekYear())
	assert.Equal(2009, simpleLocale(DateTime{Year: 2010, Month: 1, Day: 3}, "es").WeekYear())
	assert.Equal(2010, simpleLocale(DateTime{Year: 2010, Month: 1, Day: 4}, "es").WeekYear())
}

func TestFormatWeekYearWithDow1Doy4(t *testing.T) {
	assert := assert.New(t)

	cases := map[string]string{
		"2005-01-02": "2004-53",
		"2005-12-31": "2005-52",
		"2007-01-01": "2007-01",
		"2007-12-30": "2007-52",
		"2007-12-31": "2008-01",
		"2008-01-01": "2008-01",
		"2008-12-28": "2008-52",
		"2008-12-29": "2009-01",
		"2008-12-30": "2009-01",
		"2008-12-31": "2009-01",
		"2009-01-01": "2009-01",
		"2009-12-31": "2009-53",
		"2010-01-01": "2009-53",
		"2010-01-02": "2009-53",
		"2010-01-03": "2009-53",
		"404-12-31":  "0404-53",
		"405-12-31":  "0405-52",
	}

	for date, iso := range cases {
		isoWeekYear := strings.Split(iso, "-")[0]
		assert.Equal("0"+isoWeekYear, simpleFormatLocale(date, "YYYY-MM-DD", "es").Format("ggggg"))
		assert.Equal(isoWeekYear, simpleFormatLocale(date, "YYYY-MM-DD", "es").Format("gggg"))
		assert.Equal(isoWeekYear[2:4], simpleFormatLocale(date, "YYYY-MM-DD", "es").Format("gg"))
	}
}

func TestFrLocale(t *testing.T) {
	assert := assert.New(t)

	longDays := []string{"dimanche", "lundi", "mardi", "mercredi", "jeudi", "vendredi", "samedi"}
	shortDays := []string{"dim.", "lun.", "mar.", "mer.", "jeu.", "ven.", "sam."}
	minDays := []string{"di", "lu", "ma", "me", "je", "ve", "sa"}
	shortMonths := []string{"janv.", "févr.", "mars", "avr.", "mai", "juin", "juil.", "août", "sept.", "oct.", "nov.", "déc."}
	longMonths := []string{"janvier", "février", "mars", "avril", "mai", "juin", "juillet", "août", "septembre", "octobre", "novembre", "décembre"}

	lib := simpleNow()
	lib.SetLocale("fr")

	assert.Equal("fr", lib.Locale())
	assert.Equal(longDays, lib.Weekdays())
	assert.Equal(shortDays, lib.WeekdaysShort())
	assert.Equal(minDays, lib.WeekdaysMin())
	assert.Equal(longMonths, lib.Months())
	assert.Equal(shortMonths, lib.MonthsShort())
}

func TestFrWeekdaysWithLocaleDow(t *testing.T) {
	longDays := []string{"lundi", "mardi", "mercredi", "jeudi", "vendredi", "samedi", "dimanche"}

	lib := simpleNow()
	lib.SetLocale("fr")

	assert.Equal(t, longDays, lib.Weekdays(true))
}

func TestFrWeekdaysShortWithLocaleDow(t *testing.T) {
	longDays := []string{"lun.", "mar.", "mer.", "jeu.", "ven.", "sam.", "dim."}

	lib := simpleNow()
	lib.SetLocale("fr")

	assert.Equal(t, longDays, lib.WeekdaysShort(true))
}

func TestFrWeekdaysMinWithLocaleDow(t *testing.T) {
	longDays := []string{"lu", "ma", "me", "je", "ve", "sa", "di"}

	lib := simpleNow()
	lib.SetLocale("fr")

	assert.Equal(t, longDays, lib.WeekdaysMin(true))
}

func TestFrWeekdayByNumberLocale(t *testing.T) {
	assert := assert.New(t)

	lib := simpleNow()
	lib.SetLocale("fr")

	assert.Equal("lundi", lib.WeekdayByNumber(true, 0)) // first dow per locale, Monday
	assert.Equal("dimanche", lib.WeekdayByNumber(0))    // first dow is Sunday, Sunday
	assert.Equal("jeudi", lib.WeekdayByNumber(true, 3)) // first dow per locale, Thursday
	assert.Equal("mercredi", lib.WeekdayByNumber(3))    // first dow is Sunday, Wednesday
	assert.Equal("samedi", lib.WeekdayByNumber(6))      // first dow is Sunday, Saturday
}

func TestFrFormat(t *testing.T) {
	assert := assert.New(t)

	formats := map[string]string{
		"dddd, MMMM Do YYYY, h:mm:ss a": "dimanche, février 14 2010, 3:25:50 pm",
		"ddd, hA":                       "dim., 3PM",
		"M Mo MM MMMM MMM":              "2 2e 02 février févr.",
		"YYYY YY":                       "2010 10",
		"D Do DD":                       "14 14 14",
		"d do dddd ddd dd":              "0 0e dimanche dim. di",
		"DDD DDDo DDDD":                 "45 45e 045",
		"w wo ww":                       "6 6e 06",
		"h hh":                          "3 03",
		"H HH":                          "15 15",
		"m mm":                          "25 25",
		"s ss":                          "50 50",
		"a A":                           "pm PM",
		"[le] Do [jour du mois]":        "le 14 jour du mois",
		"[le] DDDo [jour de l’année]":   "le 45e jour de l’année",
		"LTS":                           "15:25:50",
		"L":                             "14/02/2010",
		"LL":                            "14 février 2010",
		"LLL":                           "14 février 2010 15:25",
		"LLLL":                          "dimanche 14 février 2010 15:25",
		"l":                             "14/2/2010",
		"ll":                            "14 févr. 2010",
		"lll":                           "14 févr. 2010 15:25",
		"llll":                          "dim. 14 févr. 2010 15:25",
	}

	lib := simpleTime(time.Date(2010, 2, 14, 15, 25, 50, 125000000, chicagoLocation()))
	lib.SetLocale("fr")

	for p, r := range formats {
		assert.Equal(r, lib.Format(p), r)
	}
}

func TestFrOrdinal(t *testing.T) {
	assert := assert.New(t)

	SetLocale("fr")

	assert.Equal("1er", simple(DateTime{Year: 2017, Month: 1, Day: 1}).Format("Mo"))
	assert.Equal("2e", simple(DateTime{Year: 2017, Month: 2, Day: 1}).Format("Mo"))
	assert.Equal("1er", simple(DateTime{Year: 2017, Month: 1, Day: 1}).Format("Qo"))
	assert.Equal("2e", simple(DateTime{Year: 2017, Month: 4, Day: 1}).Format("Qo"))
	assert.Equal("1er", simple(DateTime{Year: 2017, Month: 1, Day: 1}).Format("Do"))
	assert.Equal("2", simple(DateTime{Year: 2017, Month: 1, Day: 2}).Format("Do"))
	assert.Equal("1er", simple(DateTime{Year: 2011, Month: 1, Day: 1}).Format("DDDo"))
	assert.Equal("2e", simple(DateTime{Year: 2011, Month: 1, Day: 2}).Format("DDDo"))
	assert.Equal("3e", simple(DateTime{Year: 2011, Month: 1, Day: 3}).Format("DDDo"))
	assert.Equal("4e", simple(DateTime{Year: 2011, Month: 1, Day: 4}).Format("DDDo"))
	assert.Equal("5e", simple(DateTime{Year: 2011, Month: 1, Day: 5}).Format("DDDo"))
	assert.Equal("6e", simple(DateTime{Year: 2011, Month: 1, Day: 6}).Format("DDDo"))
	assert.Equal("7e", simple(DateTime{Year: 2011, Month: 1, Day: 7}).Format("DDDo"))
	assert.Equal("8e", simple(DateTime{Year: 2011, Month: 1, Day: 8}).Format("DDDo"))
	assert.Equal("9e", simple(DateTime{Year: 2011, Month: 1, Day: 9}).Format("DDDo"))
	assert.Equal("10e", simple(DateTime{Year: 2011, Month: 1, Day: 10}).Format("DDDo"))
	assert.Equal("11e", simple(DateTime{Year: 2011, Month: 1, Day: 11}).Format("DDDo"))
	assert.Equal("12e", simple(DateTime{Year: 2011, Month: 1, Day: 12}).Format("DDDo"))
	assert.Equal("13e", simple(DateTime{Year: 2011, Month: 1, Day: 13}).Format("DDDo"))
	assert.Equal("14e", simple(DateTime{Year: 2011, Month: 1, Day: 14}).Format("DDDo"))
	assert.Equal("15e", simple(DateTime{Year: 2011, Month: 1, Day: 15}).Format("DDDo"))
	assert.Equal("16e", simple(DateTime{Year: 2011, Month: 1, Day: 16}).Format("DDDo"))
	assert.Equal("17e", simple(DateTime{Year: 2011, Month: 1, Day: 17}).Format("DDDo"))
	assert.Equal("18e", simple(DateTime{Year: 2011, Month: 1, Day: 18}).Format("DDDo"))
	assert.Equal("19e", simple(DateTime{Year: 2011, Month: 1, Day: 19}).Format("DDDo"))
	assert.Equal("20e", simple(DateTime{Year: 2011, Month: 1, Day: 20}).Format("DDDo"))
	assert.Equal("21e", simple(DateTime{Year: 2011, Month: 1, Day: 21}).Format("DDDo"))
	assert.Equal("22e", simple(DateTime{Year: 2011, Month: 1, Day: 22}).Format("DDDo"))
	assert.Equal("23e", simple(DateTime{Year: 2011, Month: 1, Day: 23}).Format("DDDo"))
	assert.Equal("24e", simple(DateTime{Year: 2011, Month: 1, Day: 24}).Format("DDDo"))
	assert.Equal("25e", simple(DateTime{Year: 2011, Month: 1, Day: 25}).Format("DDDo"))
	assert.Equal("26e", simple(DateTime{Year: 2011, Month: 1, Day: 26}).Format("DDDo"))
	assert.Equal("27e", simple(DateTime{Year: 2011, Month: 1, Day: 27}).Format("DDDo"))
	assert.Equal("28e", simple(DateTime{Year: 2011, Month: 1, Day: 28}).Format("DDDo"))
	assert.Equal("29e", simple(DateTime{Year: 2011, Month: 1, Day: 29}).Format("DDDo"))
	assert.Equal("30e", simple(DateTime{Year: 2011, Month: 1, Day: 30}).Format("DDDo"))
	assert.Equal("31e", simple(DateTime{Year: 2011, Month: 1, Day: 31}).Format("DDDo"))
	assert.Equal("0e", simple(DateTime{Year: 2017, Month: 1, Day: 1}).Format("do"))
	assert.Equal("1er", simple(DateTime{Year: 2017, Month: 1, Day: 2}).Format("do"))
	assert.Equal("1re 1re", simple(DateTime{Year: 2017, Month: 1, Day: 4}).Format("wo Wo"))
	assert.Equal("2e 2e", simple(DateTime{Year: 2017, Month: 1, Day: 11}).Format("wo Wo"))
	assert.Equal("1re", simple(DateTime{Year: 2017, Month: 1, Day: 4}).Format("Wo"))
	assert.Equal("2e", simple(DateTime{Year: 2017, Month: 1, Day: 11}).Format("Wo"))

	SetLocale("en")
}

func TestFrRelativeTime(t *testing.T) {
	assert := assert.New(t)

	testTime := time.Date(2007, 1, 28, 0, 0, 0, 0, chicagoLocation())
	lib := simpleTime(testTime)

	lib.SetLocale("fr")

	assert.Equal("il y a quelques secondes", lib.From(simpleTime(testTime).Add(44, "s")), "44 seconds = a few seconds ago")
	assert.Equal("quelques secondes", lib.From(simpleTime(testTime).Add(44, "s"), true), "44 seconds = a few seconds")
	assert.Equal("il y a une minute", lib.From(simpleTime(testTime).Add(1, "m")), "1 minute = a minute ago")
	assert.Equal("une minute", lib.From(simpleTime(testTime).Add(1, "m"), true), "1 minute = a minute")
	assert.Equal("il y a 44 minutes", lib.From(simpleTime(testTime).Add(44, "m")), "44 minutes = 44 minutes ago")
	assert.Equal("44 minutes", lib.From(simpleTime(testTime).Add(44, "m"), true), "44 minutes = 44 minutes")
	assert.Equal("il y a une heure", lib.From(simpleTime(testTime).Add(1, "h")), "1 hour = an hour ago")
	assert.Equal("une heure", lib.From(simpleTime(testTime).Add(1, "h"), true), "1 hour = an hour")
	assert.Equal("il y a 2 heures", lib.From(simpleTime(testTime).Add(2, "h")), "2 hours = 2 hours ago")
	assert.Equal("2 heures", lib.From(simpleTime(testTime).Add(2, "h"), true), "2 hours = 2 hours")
	assert.Equal("il y a un jour", lib.From(simpleTime(testTime).Add(1, "d")), "1 day = a day ago")
	assert.Equal("un jour", lib.From(simpleTime(testTime).Add(1, "d"), true), "1 day = a day")
	assert.Equal("il y a 5 jours", lib.From(simpleTime(testTime).Add(5, "d")), "5 days = 5 days ago")
	assert.Equal("5 jours", lib.From(simpleTime(testTime).Add(5, "d"), true), "5 days = 5 days")
	assert.Equal("il y a un mois", lib.From(simpleTime(testTime).Add(1, "M")), "1 month = a month ago")
	assert.Equal("un mois", lib.From(simpleTime(testTime).Add(1, "M"), true), "1 month = a month")
	assert.Equal("il y a 5 mois", lib.From(simpleTime(testTime).Add(5, "M")), "5 months = 5 months ago")
	assert.Equal("5 mois", lib.From(simpleTime(testTime).Add(5, "M"), true), "5 months = 5 months")
	assert.Equal("il y a un an", lib.From(simpleTime(testTime).Add(1, "y")), "1 year = a year ago")
	assert.Equal("un an", lib.From(simpleTime(testTime).Add(1, "y"), true), "1 year = a year")
	assert.Equal("il y a 5 ans", lib.From(simpleTime(testTime).Add(5, "y")), "5 years = 5 years ago")
	assert.Equal("5 ans", lib.From(simpleTime(testTime).Add(5, "y"), true), "5 years = 5 years")
}

func TestFrCalendarDay(t *testing.T) {
	assert := assert.New(t)

	testTime := time.Date(2000, 12, 15, 12, 0, 0, 0, time.UTC)
	timeNow = func() time.Time {
		return testTime
	}

	SetLocale("fr")

	refTime := simpleTime(testTime)

	assert.Equal("Aujourd’hui à 12:00", simpleGoment(refTime).Calendar(), "today at the same time")
	assert.Equal("Aujourd’hui à 12:25", simpleGoment(refTime).Add(25, "m").Calendar(), "now plus 25 min")
	assert.Equal("Aujourd’hui à 13:00", simpleGoment(refTime).Add(1, "h").Calendar(), "now plus 1 hour")
	assert.Equal("Demain à 12:00", simpleGoment(refTime).Add(1, "d").Calendar(), "tomorrow at the same time")
	assert.Equal("Aujourd’hui à 11:00", simpleGoment(refTime).Subtract(1, "h").Calendar(), "now minus 1 hour")
	assert.Equal("Hier à 12:00", simpleGoment(refTime).Subtract(1, "d").Calendar(), "yesterday at the same time")

	refTime = simpleTime(testTime)

	assert.Equal("dimanche à 12:00", refTime.Add(2, "d").Calendar(), "Today + 2 days current time")
	refTime.StartOf("day")
	assert.Equal("dimanche à 00:00", refTime.Calendar(), "Today + 2 days beginning of day")
	refTime.EndOf("day")
	assert.Equal("dimanche à 23:59", refTime.Calendar(), "Today + 2 days end of day")

	refTime = simpleTime(testTime)

	assert.Equal("mercredi dernier à 12:00", refTime.Subtract(2, "d").Calendar(), "Today - 2 days current time")
	refTime.StartOf("day")
	assert.Equal("mercredi dernier à 00:00", refTime.Calendar(), "Today - 2 days beginning of day")
	refTime.EndOf("day")
	assert.Equal("mercredi dernier à 23:59", refTime.Calendar(), "Today - 2 days end of day")

	weeksAgo := simpleTime(testTime).Subtract(1, "w")
	weeksFromNow := simpleTime(testTime).Add(1, "w")

	assert.Equal("08/12/2000", weeksAgo.Calendar())
	assert.Equal("22/12/2000", weeksFromNow.Calendar())

	weeksAgo = simpleTime(testTime).Subtract(2, "w")
	weeksFromNow = simpleTime(testTime).Add(2, "w")

	assert.Equal("01/12/2000", weeksAgo.Calendar())
	assert.Equal("29/12/2000", weeksFromNow.Calendar())

	// Reset timeNow.
	timeNow = time.Now

	SetLocale("en")
}

func TestFrFormatParsing(t *testing.T) {
	assert := assert.New(t)

	formats := map[string][]string{
		"YYYY-Q":                    {"2014-4"},
		"MM-DD-YYYY":                {"12-02-1999"},
		"DD-MM-YYYY":                {"12-02-1999"},
		"DD/MM/YYYY":                {"12/02/1999"},
		"DD_MM_YYYY":                {"12_02_1999"},
		"DD:MM:YYYY":                {"12:02:1999"},
		"D-M-YY":                    {"2-2-99"},
		"Y":                         {"-0025"},
		"YY":                        {"99"},
		"DDD-YYYY":                  {"300-1999"},
		"YYYY-DDD":                  {"1999-300"},
		"YYYY MM Do":                {"2014 01 1er", "2015 11 21"},
		"MMM":                       {"avr."},
		"MMMM":                      {"septembre"},
		"DD MMMM":                   {"11 septembre"},
		"Do MMMM":                   {"3 septembre"},
		"YYYY MMMM":                 {"2018 octobre"},
		"D":                         {"3", "27"},
		"DD":                        {"04", "23"},
		"DDD":                       {"7", "300"},
		"DDDD":                      {"008", "211", "312"},
		"h":                         {"4"},
		"H":                         {"1", "10", "23"},
		"DD-MM-YYYY h:m:s":          {"12-02-1999 2:45:10"},
		"DD-MM-YYYY h:m:s a":        {"12-02-1999 2:45:10 am", "12-02-1999 2:45:10 pm"},
		"h:mm a":                    {"12:00 pm", "12:30 pm", "12:00 am", "12:30 am"},
		"HH:mm":                     {"12:00"},
		"kk:mm":                     {"12:00"},
		"YYYY-MM-DDTHH:mm:ss":       {"2011-11-11T11:11:11"},
		"MM-DD-YYYY [M]":            {"12-02-1999 M"},
		"ddd MMM DD HH:mm:ss YYYY":  {"mer. avr. 08 22:52:51 2009"},
		"dddd MMM DD HH:mm:ss YYYY": {"samedi avr. 11 22:52:51 2009"},
		"HH:mm:ss":                  {"12:00:00", "12:30:11", "00:00:00"},
		"kk:mm:ss":                  {"12:00:10", "12:30:42", "24:00:00", "09:00:00"},
		"YYYY-MM-DD HH:mm:ss ZZ":    {"2000-05-15 17:08:00 -0700"},
		"YYYY-MM-DD HH:mm Z":        {"2010-10-20 04:30 +00:00"},
		"e":                         {"0", "5"},
		"E":                         {"1", "7"},
		// "HH:mm:ss.S": []string{"00:30:00.1"},
		// "HH:mm:ss SS":               "00:30:00 12",
		// "HH:mm:ss SSS":              "00:30:00 123",
		// "HH:mm:ss S":                "00:30:00 7",
		// "HH:mm:ss SS":               "00:30:00 78",
		// "HH:mm:ss SSS":              "00:30:00 789",
		// "kk:mm:ss S":   "24:30:00 1",
		// "kk:mm:ss SS":  "24:30:00 12",
		// "kk:mm:ss SSS": "24:30:00 123",
		// "kk:mm:ss S":   "24:30:00 7",
		// "kk:mm:ss SS":  "24:30:00 78",
		// "kk:mm:ss SSS": "24:30:00 789",
		"X":    {"1234567890"},
		"H Z":  {"6 -06:00"},
		"H ZZ": {"5 -0700"},
		"LT":   {"12:30"},
		"LTS":  {"12:30:29"},
		"L":    {"09/02/1999"},
		"l":    {"9/2/1999"},
		"LL":   {"2 septembre 1999"},
		"ll":   {"2 sept. 1999"},
		"LLL":  {"2 septembre 1999 12:30"},
		"lll":  {"2 sept. 1999 12:30"},
		"LLLL": {"jeudi 2 septembre 1999 12:30"},
		"llll": {"jeu. 2 sept. 1999 12:30"},
	}

	for format, dates := range formats {
		for _, date := range dates {
			lib, _ := New(date, format, "fr")
			assert.Equal(date, lib.Format(format), fmt.Sprintf("%v: %v", format, date))
		}
	}
}

func TestFrWeekday(t *testing.T) {
	assert := assert.New(t)

	// fr 1st day of week is Monday.
	testTime := time.Date(2020, 7, 28, 18, 0, 0, 0, time.UTC) // Tuesday, July 28 2020

	lib := simpleTime(testTime)
	lib.SetLocale("fr")

	assert.Equal(lib.Weekday(), 1)

	lib2 := simpleString("2016-09-05") // Monday, September 5 2016
	lib2.SetLocale("fr")

	assert.Equal(0, lib2.Weekday()) // Monday
	lib2.SetWeekday(3)              // Thursday
	assert.Equal(8, lib2.Date())
	lib2.SetWeekday(6) // Sunday
	assert.Equal(11, lib2.Date())
}

func TestFrSetWeekdayRangeExceeded(t *testing.T) {
	assert := assert.New(t)

	format := "YYYY-MM-DD"

	lib := simpleFormatLocale("2020-07-26", format, "fr")
	assert.Equal(lib.Format(format), "2020-07-26")

	lib.SetWeekday(-7)
	assert.Equal(lib.Format(format), "2020-07-13")

	lib.SetWeekday(9)
	assert.Equal(lib.Format(format), "2020-07-22")
}

func TestFrSetDayByName(t *testing.T) {
	assert := assert.New(t)

	format := "YYYY-MM-DD"

	lib := simpleFormatLocale("2016-09-04", format, "fr")

	lib.SetDay("mercredi") // Wednesday
	assert.Equal(7, lib.Date())

	lib.SetDay("samedi") // Saturday
	assert.Equal(10, lib.Date())
}

func TestFrStartEndOfWeek(t *testing.T) {
	now := simpleNow()
	now.SetLocale("fr")

	assert.Equal(t, "lundi", now.StartOf("week").Format("dddd"))
	assert.Equal(t, "dimanche", now.EndOf("week").Format("dddd"))
}

func TestPtBRLocale(t *testing.T) {
	assert := assert.New(t)

	longDays := []string{"domingo", "segunda-feira", "terça-feira", "quarta-feira", "quinta-feira", "sexta-feira", "sábado"}
	shortDays := []string{"dom", "seg", "ter", "qua", "qui", "sex", "sab"}
	minDays := []string{"do", "se", "te", "qa", "qi", "se", "sa"}
	shortMonths := []string{"jan", "fev", "mar", "abr", "mai", "jun", "jul", "ago", "set", "out", "nov", "dez"}
	longMonths := []string{"janeiro", "fevereiro", "março", "abril", "maio", "junho", "julho", "agosto", "setembro", "outubro", "novembro", "dezembro"}

	lib, _ := New()
	lib.SetLocale("pt-br")

	assert.Equal("pt-br", lib.Locale())
	assert.Equal(longDays, lib.Weekdays())
	assert.Equal(shortDays, lib.WeekdaysShort())
	assert.Equal(minDays, lib.WeekdaysMin())
	assert.Equal(longMonths, lib.Months())
	assert.Equal(shortMonths, lib.MonthsShort())
}

func TestPtBRWeekdaysWithLocaleDow(t *testing.T) {
	longDays := []string{"segunda-feira", "terça-feira", "quarta-feira", "quinta-feira", "sexta-feira", "sábado", "domingo"}

	lib := simpleNow()
	lib.SetLocale("pt-br")

	assert.Equal(t, longDays, lib.Weekdays(true))
}

func TestPtBRWeekdaysShortWithLocaleDow(t *testing.T) {
	shortDays := []string{"seg", "ter", "qua", "qui", "sex", "sab", "dom"}

	lib := simpleNow()
	lib.SetLocale("pt-br")

	assert.Equal(t, shortDays, lib.WeekdaysShort(true))
}

func TestPtBRWeekdaysMinWithLocaleDow(t *testing.T) {
	minDays := []string{"se", "te", "qa", "qi", "se", "sa", "do"}

	lib := simpleNow()
	lib.SetLocale("pt-br")

	assert.Equal(t, minDays, lib.WeekdaysMin(true))
}

func TestPtBRWeekdayByNumberLocale(t *testing.T) {
	assert := assert.New(t)

	lib := simpleNow()
	lib.SetLocale("pt-br")

	assert.Equal("segunda-feira", lib.WeekdayByNumber(true, 0)) // first dow per locale, Monday
	assert.Equal("domingo", lib.WeekdayByNumber(0))             // first dow is Sunday, Sunday
	assert.Equal("quinta-feira", lib.WeekdayByNumber(true, 3))  // first dow per locale, Thursday
	assert.Equal("quarta-feira", lib.WeekdayByNumber(3))        // first dow is Sunday, Wednesday
	assert.Equal("sábado", lib.WeekdayByNumber(6))              // first dow is Sunday, Saturday
}

func TestPtBRFormat(t *testing.T) {
	assert := assert.New(t)

	formats := map[string]string{
		"dddd, MMMM Do YYYY, h:mm:ss a": "domingo, fevereiro 14º 2010, 3:25:50 pm",
		"ddd, hA":                       "dom, 3PM",
		"M Mo MM MMMM MMM":              "2 2º 02 fevereiro fev",
		"YYYY YY":                       "2010 10",
		"D Do DD":                       "14 14º 14",
		"d do dddd ddd dd":              "0 0º domingo dom do",
		"DDD DDDo DDDD":                 "45 45º 045",
		"w wo ww":                       "6 6º 06",
		"YYYY-MMM-DD":                   "2010-fev-14",
		"h hh":                          "3 03",
		"H HH":                          "15 15",
		"m mm":                          "25 25",
		"s ss":                          "50 50",
		"a A":                           "pm PM",
		"[the] DDDo [day of the year]":  "the 45º day of the year",
		"[the] DDDo [day of the year is after January]": "the 45º day of the year is after January",
		"LT":            "15:25",
		"LTS":           "15:25:50",
		"L":             "14/02/2010",
		"LL":            "14 de fevereiro de 2010",
		"[today is] LL": "today is 14 de fevereiro de 2010",
		"LLL":           "14 de fevereiro de 2010 15:25",
		"LLLL":          "domingo, 14 de fevereiro de 2010 15:25",
		"l":             "14/2/2010",
		"ll":            "14 de fev de 2010",
		"lll":           "14 de fev de 2010 15:25",
		"llll":          "dom, 14 de fev de 2010 15:25",
	}

	lib := simpleTime(time.Date(2010, 2, 14, 15, 25, 50, 125000000, chicagoLocation()))
	lib.SetLocale("pt-br")

	for p, r := range formats {
		assert.Equal(r, lib.Format(p), r)
	}
}

func TestPtBRShortDayFormatWithUnicode(t *testing.T) {
	lib := simpleTime(time.Date(2010, 2, 20, 15, 25, 50, 125000000, chicagoLocation()))
	lib.SetLocale("es")

	assert.Equal(t, "sá", lib.Format("dd"))
}

func TestPtBRRelativeTime(t *testing.T) {
	assert := assert.New(t)

	testTime := time.Date(2007, 1, 28, 0, 0, 0, 0, chicagoLocation())
	lib := simpleTime(testTime)

	lib.SetLocale("pt-br")

	assert.Equal("há alguns segundos", lib.From(simpleTime(testTime).Add(44, "s")), "44 seconds = a few seconds ago")
	assert.Equal("alguns segundos", lib.From(simpleTime(testTime).Add(44, "s"), true), "44 seconds = a few seconds")
	assert.Equal("há um minuto", lib.From(simpleTime(testTime).Add(1, "m")), "1 minute = a minute ago")
	assert.Equal("um minuto", lib.From(simpleTime(testTime).Add(1, "m"), true), "1 minute = a minute")
	assert.Equal("há 44 minutos", lib.From(simpleTime(testTime).Add(44, "m")), "44 minutes = 44 minutes ago")
	assert.Equal("44 minutos", lib.From(simpleTime(testTime).Add(44, "m"), true), "44 minutes = 44 minutes")
	assert.Equal("há uma hora", lib.From(simpleTime(testTime).Add(1, "h")), "1 hour = an hour ago")
	assert.Equal("uma hora", lib.From(simpleTime(testTime).Add(1, "h"), true), "1 hour = an hour")
	assert.Equal("há 2 horas", lib.From(simpleTime(testTime).Add(2, "h")), "2 hours = 2 hours ago")
	assert.Equal("2 horas", lib.From(simpleTime(testTime).Add(2, "h"), true), "2 hours = 2 hours")
	assert.Equal("há um dia", lib.From(simpleTime(testTime).Add(1, "d")), "1 day = a day ago")
	assert.Equal("um dia", lib.From(simpleTime(testTime).Add(1, "d"), true), "1 day = a day")
	assert.Equal("há 5 dias", lib.From(simpleTime(testTime).Add(5, "d")), "5 days = 5 days ago")
	assert.Equal("5 dias", lib.From(simpleTime(testTime).Add(5, "d"), true), "5 days = 5 days")
	assert.Equal("há um mês", lib.From(simpleTime(testTime).Add(1, "M")), "1 month = a month ago")
	assert.Equal("um mês", lib.From(simpleTime(testTime).Add(1, "M"), true), "1 month = a month")
	assert.Equal("há 5 meses", lib.From(simpleTime(testTime).Add(5, "M")), "5 months = 5 months ago")
	assert.Equal("5 meses", lib.From(simpleTime(testTime).Add(5, "M"), true), "5 months = 5 months")
	assert.Equal("há um ano", lib.From(simpleTime(testTime).Add(1, "y")), "1 year = a year ago")
	assert.Equal("um ano", lib.From(simpleTime(testTime).Add(1, "y"), true), "1 year = a year")
	assert.Equal("há 5 anos", lib.From(simpleTime(testTime).Add(5, "y")), "5 years = 5 years ago")
	assert.Equal("5 anos", lib.From(simpleTime(testTime).Add(5, "y"), true), "5 years = 5 years")
}

func TestPtBRCalendarDay(t *testing.T) {
	assert := assert.New(t)

	testTime := time.Date(2000, 12, 15, 12, 0, 0, 0, time.UTC)
	timeNow = func() time.Time {
		return testTime
	}

	SetLocale("pt-br")

	refTime := simpleTime(testTime)

	assert.Equal("hoje às 12:00", simpleGoment(refTime).Calendar(), "today at the same time")
	assert.Equal("hoje às 12:25", simpleGoment(refTime).Add(25, "m").Calendar(), "now plus 25 min")
	assert.Equal("hoje às 13:00", simpleGoment(refTime).Add(1, "h").Calendar(), "now plus 1 hour")
	assert.Equal("amanhã às 12:00", simpleGoment(refTime).Add(1, "d").Calendar(), "tomorrow at the same time")
	assert.Equal("hoje às 11:00", simpleGoment(refTime).Subtract(1, "h").Calendar(), "now minus 1 hour")
	assert.Equal("ontem às 12:00", simpleGoment(refTime).Subtract(1, "d").Calendar(), "yesterday at the same time")

	refTime = simpleTime(testTime)

	assert.Equal("domingo às 12:00", refTime.Add(2, "d").Calendar(), "Today + 2 days current time")
	refTime.StartOf("day")
	assert.Equal("domingo às 0:00", refTime.Calendar(), "Today + 2 days beginning of day")
	refTime.EndOf("day")
	assert.Equal("domingo às 23:59", refTime.Calendar(), "Today + 2 days end of day")

	refTime = simpleTime(testTime)

	assert.Equal("na quarta-feira passada às 12:00", refTime.Subtract(2, "d").Calendar(), "Today - 2 days current time")
	refTime.StartOf("day")
	assert.Equal("na quarta-feira passada às 0:00", refTime.Calendar(), "Today - 2 days beginning of day")
	refTime.EndOf("day")
	assert.Equal("na quarta-feira passada às 23:59", refTime.Calendar(), "Today - 2 days end of day")

	weeksAgo := simpleTime(testTime).Subtract(1, "w")
	weeksFromNow := simpleTime(testTime).Add(1, "w")

	assert.Equal("08/12/2000", weeksAgo.Calendar())
	assert.Equal("22/12/2000", weeksFromNow.Calendar())

	weeksAgo = simpleTime(testTime).Subtract(2, "w")
	weeksFromNow = simpleTime(testTime).Add(2, "w")

	assert.Equal("01/12/2000", weeksAgo.Calendar())
	assert.Equal("29/12/2000", weeksFromNow.Calendar())

	// Reset timeNow.
	timeNow = time.Now

	SetLocale("en")
}

func TestPtBRFormatParsing(t *testing.T) {
	assert := assert.New(t)

	formats := map[string][]string{
		"YYYY-Q":                    {"2014-4"},
		"MM-DD-YYYY":                {"12-02-1999"},
		"DD-MM-YYYY":                {"12-02-1999"},
		"DD/MM/YYYY":                {"12/02/1999"},
		"DD_MM_YYYY":                {"12_02_1999"},
		"DD:MM:YYYY":                {"12:02:1999"},
		"D-M-YY":                    {"2-2-99"},
		"Y":                         {"-0025"},
		"YY":                        {"99"},
		"DDD-YYYY":                  {"300-1999"},
		"YYYY-DDD":                  {"1999-300"},
		"YYYY MM Do":                {"2014 01 3º", "2015 11 21º"},
		"MMM":                       {"abr"},
		"MMMM":                      {"setembro"},
		"DD MMMM":                   {"11 setembro"},
		"Do MMMM":                   {"3º setembro"},
		"YYYY MMMM":                 {"2018 outubro"},
		"D":                         {"3", "27"},
		"DD":                        {"04", "23"},
		"DDD":                       {"7", "300"},
		"DDDD":                      {"008", "211", "312"},
		"h":                         {"4"},
		"H":                         {"1", "10", "23"},
		"DD-MM-YYYY h:m:s":          {"12-02-1999 2:45:10"},
		"DD-MM-YYYY h:m:s a":        {"12-02-1999 2:45:10 am", "12-02-1999 2:45:10 pm"},
		"h:mm a":                    {"12:00 pm", "12:30 pm", "12:00 am", "12:30 am"},
		"HH:mm":                     {"12:00"},
		"kk:mm":                     {"12:00"},
		"YYYY-MM-DDTHH:mm:ss":       {"2011-11-11T11:11:11"},
		"MM-DD-YYYY [M]":            {"12-02-1999 M"},
		"ddd MMM DD HH:mm:ss YYYY":  {"qua abr 08 22:52:51 2009"},
		"dddd MMM DD HH:mm:ss YYYY": {"sábado abr 11 22:52:51 2009"},
		"HH:mm:ss":                  {"12:00:00", "12:30:11", "00:00:00"},
		"kk:mm:ss":                  {"12:00:10", "12:30:42", "24:00:00", "09:00:00"},
		"YYYY-MM-DD HH:mm:ss ZZ":    {"2000-05-15 17:08:00 -0700"},
		"YYYY-MM-DD HH:mm Z":        {"2010-10-20 04:30 +00:00"},
		"e":                         {"0", "5"},
		"E":                         {"1", "7"},
		"X":                         {"1234567890"},
		"H Z":                       {"6 -06:00"},
		"H ZZ":                      {"5 -0700"},
		"LT":                        {"12:30"},
		"LTS":                       {"12:30:29"},
		"L":                         {"09/02/1999"},
		"l":                         {"9/2/1999"},
		"LL":                        {"2 de setembro de 1999"},
		"ll":                        {"2 de set de 1999"},
		"LLL":                       {"2 de setembro de 1999 12:30"},
		"lll":                       {"2 de set de 1999 12:30"},
		"LLLL":                      {"quinta-feira, 2 de setembro de 1999 12:30"},
		"llll":                      {"dom, 14 de fev de 2010 15:25"},
	}

	for format, dates := range formats {
		for _, date := range dates {
			lib, _ := New(date, format, "pt-br")
			assert.Equal(date, lib.Format(format), fmt.Sprintf("%v: %v", format, date))
		}
	}
}

func TestPtBRInvalidOrdinal(t *testing.T) {
	date := "2014 01 4th"
	format := "YYYY MM Do"

	lib, _ := New(date, format, "pt-br")
	assert.Equal(t, "2014 01 1º", lib.Format(format))
}

func TestPtBRFormatWeekdayParsing(t *testing.T) {
	assert := assert.New(t)

	lib := simpleFormatLocale("qua abr 08 22:52:51 2009", "ddd MMM DD HH:mm:ss YYYY", "pt-br")
	assert.Equal("qua abr 08 22:52:51 2009", lib.Format("ddd MMM DD HH:mm:ss YYYY"))

	lib2 := simpleFormatLocale("sábado abr 11 22:52:51 2009", "dddd MMM DD HH:mm:ss YYYY", "pt-br")
	assert.Equal("sábado abr 11 22:52:51 2009", lib2.Format("dddd MMM DD HH:mm:ss YYYY"))
}

func TestPtBRDayOfWeekParsing(t *testing.T) {
	assert := assert.New(t)

	outputFormat := "YYYY MM DD"

	date := simpleNow().Format(outputFormat)

	assert.Equal(
		simpleFormatLocale(date, outputFormat, "pt-br").SetDay("segunda-feira").Format(outputFormat),
		simpleFormatLocale("0", "e", "pt-br").Format(outputFormat),
	)
	assert.Equal(
		simpleFormatLocale(date, outputFormat, "pt-br").SetDay("quinta-feira").Format(outputFormat),
		simpleFormatLocale("qui 0", "ddd e", "pt-br").Format(outputFormat),
	)
}

func TestPtBRWeekYearParsing(t *testing.T) {
	assert := assert.New(t)

	outputFormat := "YYYY-MM-DD"

	assert.Equal("2007-01-01", simpleFormatLocale("2007-01", "gggg-ww", "pt-br").Format(outputFormat))
	assert.Equal("2007-12-31", simpleFormatLocale("2008-01", "gggg-ww", "pt-br").Format(outputFormat))
	assert.Equal("2002-12-30", simpleFormatLocale("2003-01", "gggg-ww", "pt-br").Format(outputFormat))
	assert.Equal("2008-12-29", simpleFormatLocale("2009-01", "gggg-ww", "pt-br").Format(outputFormat))
	assert.Equal("2010-01-04", simpleFormatLocale("2010-01", "gggg-ww", "pt-br").Format(outputFormat))
	assert.Equal("2011-01-03", simpleFormatLocale("2011-01", "gggg-ww", "pt-br").Format(outputFormat))
	assert.Equal("2012-01-02", simpleFormatLocale("2012-01", "gggg-ww", "pt-br").Format(outputFormat))
}

func TestPtBRWeekWeekdayParsing(t *testing.T) {
	assert := assert.New(t)

	outputFormat := "YYYY MM DD"

	assert.Equal("1999 09 16", simpleFormatLocale("1999 37 4", "GGGG WW EE", "pt-br").Format(outputFormat), "iso ignores locale")
	assert.Equal("1999 09 19", simpleFormatLocale("1999 37 7", "GGGG WW EE", "pt-br").Format(outputFormat), "iso ignores locale")

	assert.Equal("1999 09 13", simpleFormatLocale("1999 37 0", "gggg ww e", "pt-br").Format(outputFormat), "localized e uses local doy and dow: 0 = monday")
	assert.Equal("1999 09 17", simpleFormatLocale("1999 37 4", "gggg ww e", "pt-br").Format(outputFormat), "localized e uses local doy and dow: 4 = friday")

	assert.Equal("1999 09 13", simpleFormatLocale("1999 37 1", "gggg ww d", "pt-br").Format(outputFormat), "localized d uses 0-indexed days: 1 = monday")
	assert.Equal("1999 09 13", simpleFormatLocale("1999 37 se", "gggg ww dd", "pt-br").Format(outputFormat), "localized d uses 0-indexed days: Mo")
	assert.Equal("1999 09 13", simpleFormatLocale("1999 37 seg", "gggg ww ddd", "pt-br").Format(outputFormat), "localized d uses 0-indexed days: Mon")
	assert.Equal("1999 09 13", simpleFormatLocale("1999 37 segunda-feira", "gggg ww dddd", "pt-br").Format(outputFormat), "localized d uses 0-indexed days: Monday")

	assert.Equal("1999 09 16", simpleFormatLocale("1999 37 4", "gggg ww d", "pt-br").Format(outputFormat), "localized d uses 0-indexed days: 4")

	// Sunday goes at the end of the week
	assert.Equal("1999 09 19", simpleFormatLocale("1999 37 0", "gggg ww d", "pt-br").Format(outputFormat), "localized d uses 0-indexed days: 0 = sund")
	assert.Equal("1999 09 19", simpleFormatLocale("1999 37 do", "gggg ww dd", "pt-br").Format(outputFormat), "localized d uses 0-indexed days: 0 = sund")
}

func TestPtBRWeekday(t *testing.T) {
	assert := assert.New(t)

	// es 1st day of week is Monday.
	testTime := time.Date(2020, 7, 28, 18, 0, 0, 0, time.UTC) // Tuesday, July 28 2020

	lib := simpleTime(testTime)
	lib.SetLocale("pt-br")

	assert.Equal(lib.Weekday(), 1)

	lib2 := simpleString("2016-09-05") // Monday, September 5 2016
	lib2.SetLocale("pt-br")

	assert.Equal(0, lib2.Weekday()) // Monday
	assert.Equal(5, lib2.Date())

	lib2.SetWeekday(3) // Thursday
	assert.Equal(8, lib2.Date())

	lib2.SetWeekday(6) // Sunday
	assert.Equal(11, lib2.Date())

	lib2.SetDay("quarta-feira") // Wednesday
	assert.Equal(14, lib2.Date())
}

func TestPtBRSetWeekdayRangeExceeded(t *testing.T) {
	format := "YYYY-MM-DD"

	lib := simpleFormatLocale("2020-07-26", format, "pt-br")
	assert.Equal(t, lib.Format(format), "2020-07-26")

	lib.SetWeekday(-7)
	assert.Equal(t, lib.Format(format), "2020-07-13")

	lib.SetWeekday(9)
	assert.Equal(t, lib.Format(format), "2020-07-22")
}

func TestPtBRSetDayByName(t *testing.T) {
	format := "YYYY-MM-DD"

	lib := simpleFormatLocale("2016-09-04", format, "pt-br")
	assert.Equal(t, 6, lib.Weekday())

	lib.SetDay("quarta-feira") // Wednesday
	assert.Equal(t, 7, lib.Date())

	lib.SetDay("sábado") // Saturday
	assert.Equal(t, 10, lib.Date())
}

func TestPtBRStartEndOfWeek(t *testing.T) {
	now := simpleNow()
	now.SetLocale("pt-br")

	assert.Equal(t, "segunda-feira", now.StartOf("week").Format("dddd"))
	assert.Equal(t, "domingo", now.EndOf("week").Format("dddd"))
}

func TestIdLocale(t *testing.T) {
	assert := assert.New(t)

	longDays := []string{"Minggu", "Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu"}
	shortDays := []string{"Min", "Sen", "Sel", "Rab", "Kam", "Jum", "Sab"}
	minDays := []string{"Mg", "Sn", "Sl", "Rb", "Km", "Jm", "Sb"}
	shortMonths := []string{"Jan", "Feb", "Mar", "Apr", "Mei", "Jun", "Jul", "Agt", "Sep", "Okt", "Nov", "Des"}
	longMonths := []string{"Januari", "Februari", "Maret", "April", "Mei", "Juni", "Juli", "Agustus", "September", "Oktober", "November", "Desember"}

	lib := simpleNow()
	lib.SetLocale("id")

	assert.Equal("id", lib.Locale())
	assert.Equal(longDays, lib.Weekdays())
	assert.Equal(shortDays, lib.WeekdaysShort())
	assert.Equal(minDays, lib.WeekdaysMin())
	assert.Equal(longMonths, lib.Months())
	assert.Equal(shortMonths, lib.MonthsShort())
}

func TestIdMonthByNumber(t *testing.T) {
	assert := assert.New(t)

	lib := simpleNow()
	lib.SetLocale("id")

	assert.Equal("Januari", lib.MonthByNumber(1))
	assert.Equal("Juni", lib.MonthByNumber(6))
	assert.Equal("Desember", lib.MonthByNumber(12))
}

func TestIdMonthByNumberInvalid(t *testing.T) {
	assert := assert.New(t)

	lib := simpleNow()

	assert.Equal("", lib.MonthByNumber(0))
	assert.Equal("", lib.MonthByNumber(13))
}

func TestIdMonthShortByNumber(t *testing.T) {
	assert := assert.New(t)

	lib := simpleNow()
	lib.SetLocale("id")

	assert.Equal("Jan", lib.MonthShortByNumber(1))
	assert.Equal("Jun", lib.MonthShortByNumber(6))
	assert.Equal("Des", lib.MonthShortByNumber(12))
}

func TestIdMonthShortByNumberInvalid(t *testing.T) {
	assert := assert.New(t)

	lib := simpleNow()
	lib.SetLocale("id")

	assert.Equal("", lib.MonthShortByNumber(0))
	assert.Equal("", lib.MonthShortByNumber(13))
}

func TestIdWeekdayByNumber(t *testing.T) {
	assert := assert.New(t)

	lib := simpleNow()
	lib.SetLocale("id")

	assert.Equal("Senin", lib.WeekdayByNumber(1))
	assert.Equal("Rabu", lib.WeekdayByNumber(3))
	assert.Equal("Sabtu", lib.WeekdayByNumber(6))
}

func TestIdWeekdayByNumberInvalid(t *testing.T) {
	assert := assert.New(t)

	lib := simpleNow()
	lib.SetLocale("id")

	assert.Equal("", lib.WeekdayByNumber())
	assert.Equal("", lib.WeekdayByNumber(6, 1, true))
}

func TestIdWeekdaysWithLocaleDow(t *testing.T) {
	longDays := []string{"Minggu", "Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu"}

	lib := simpleNow()
	lib.SetLocale("id")

	assert.Equal(t, longDays, lib.Weekdays(true))
}

func TestIdWeekdaysShortWithLocaleDow(t *testing.T) {
	shortDays := []string{"Min", "Sen", "Sel", "Rab", "Kam", "Jum", "Sab"}

	lib := simpleNow()
	lib.SetLocale("id")

	assert.Equal(t, shortDays, lib.WeekdaysShort(true))
}

func TestIdWeekdaysMinWithLocaleDow(t *testing.T) {
	minDays := []string{"Mg", "Sn", "Sl", "Rb", "Km", "Jm", "Sb"}

	lib := simpleNow()
	lib.SetLocale("id")

	assert.Equal(t, minDays, lib.WeekdaysMin(true))
}

func TestIdWeekdayByNumberLocale(t *testing.T) {
	assert := assert.New(t)

	lib := simpleNow()
	lib.SetLocale("id")

	assert.Equal("Minggu", lib.WeekdayByNumber(true, 0))
	assert.Equal("Minggu", lib.WeekdayByNumber(0))
	assert.Equal("Rabu", lib.WeekdayByNumber(true, 3))
	assert.Equal("Rabu", lib.WeekdayByNumber(3))
	assert.Equal("Sabtu", lib.WeekdayByNumber(6))
}

func TestIdFormat(t *testing.T) {
	assert := assert.New(t)

	formats := map[string]string{
		"dddd, MMMM Do YYYY, h:mm:ss a": "Minggu, Februari 14 2010, 3:25:50 sore",
		"ddd, hA":                       "Min, 3Sore",
		"M Mo MM MMMM MMM":              "2 2 02 Februari Feb",
		"YYYY YY":                       "2010 10",
		"D Do DD":                       "14 14 14",
		"d do dddd ddd dd":              "0 0 Minggu Min Mg",
		"DDD DDDo DDDD":                 "45 45 045",
		"w wo ww":                       "8 8 08",
		"YYYY-MMM-DD":                   "2010-Feb-14",
		"h hh":                          "3 03",
		"H HH":                          "15 15",
		"m mm":                          "25 25",
		"s ss":                          "50 50",
		"a A":                           "sore Sore",
		"[the] DDDo [day of the year]":  "the 45 day of the year",
		"[the] DDDo [day of the year is after January]": "the 45 day of the year is after January",
		"LT":            "15:25",
		"LTS":           "15:25:50",
		"L":             "14/02/2010",
		"LL":            "14 Februari 2010",
		"[today is] LL": "today is 14 Februari 2010",
		"LLL":           "14 Februari 2010, 15:25",
		"LLLL":          "Minggu, 14 Februari 2010, 15:25",
		"l":             "14/2/2010",
		"ll":            "14 Feb 2010",
		"lll":           "14 Feb 2010, 15:25",
		"llll":          "Min, 14 Feb 2010, 15:25",
	}

	lib := simpleTime(time.Date(2010, 2, 14, 15, 25, 50, 125000000, time.UTC))
	lib.SetLocale("id")

	for p, r := range formats {
		assert.Equal(r, lib.Format(p), r)
	}
}

func TestIdRelativeTime(t *testing.T) {
	assert := assert.New(t)

	testTime := time.Date(2007, 1, 28, 0, 0, 0, 0, chicagoLocation())
	lib := simpleTime(testTime)

	lib.SetLocale("id")

	assert.Equal("beberapa detik yang lalu", lib.From(simpleTime(testTime).Add(44, "s")), "44 seconds = a few seconds ago")
	assert.Equal("beberapa detik", lib.From(simpleTime(testTime).Add(44, "s"), true), "44 seconds = a few seconds")
	assert.Equal("semenit yang lalu", lib.From(simpleTime(testTime).Add(1, "m")), "1 minute = a minute ago")
	assert.Equal("semenit", lib.From(simpleTime(testTime).Add(1, "m"), true), "1 minute = a minute")
	assert.Equal("44 menit yang lalu", lib.From(simpleTime(testTime).Add(44, "m")), "44 minutes = 44 minutes ago")
	assert.Equal("44 menit", lib.From(simpleTime(testTime).Add(44, "m"), true), "44 minutes = 44 minutes")
	assert.Equal("sejam yang lalu", lib.From(simpleTime(testTime).Add(1, "h")), "1 hour = an hour ago")
	assert.Equal("sejam", lib.From(simpleTime(testTime).Add(1, "h"), true), "1 hour = an hour")
	assert.Equal("2 jam yang lalu", lib.From(simpleTime(testTime).Add(2, "h")), "2 hours = 2 hours ago")
	assert.Equal("2 jam", lib.From(simpleTime(testTime).Add(2, "h"), true), "2 hours = 2 hours")
	assert.Equal("sehari yang lalu", lib.From(simpleTime(testTime).Add(1, "d")), "1 day = a day ago")
	assert.Equal("sehari", lib.From(simpleTime(testTime).Add(1, "d"), true), "1 day = a day")
	assert.Equal("5 hari yang lalu", lib.From(simpleTime(testTime).Add(5, "d")), "5 days = 5 days ago")
	assert.Equal("5 hari", lib.From(simpleTime(testTime).Add(5, "d"), true), "5 days = 5 days")
	assert.Equal("sebulan yang lalu", lib.From(simpleTime(testTime).Add(1, "M")), "1 month = a month ago")
	assert.Equal("sebulan", lib.From(simpleTime(testTime).Add(1, "M"), true), "1 month = a month")
	assert.Equal("5 bulan yang lalu", lib.From(simpleTime(testTime).Add(5, "M")), "5 months = 5 months ago")
	assert.Equal("5 bulan", lib.From(simpleTime(testTime).Add(5, "M"), true), "5 months = 5 months")
	assert.Equal("setahun yang lalu", lib.From(simpleTime(testTime).Add(1, "y")), "1 year = a year ago")
	assert.Equal("setahun", lib.From(simpleTime(testTime).Add(1, "y"), true), "1 year = a year")
	assert.Equal("5 tahun yang lalu", lib.From(simpleTime(testTime).Add(5, "y")), "5 years = 5 years ago")
	assert.Equal("5 tahun", lib.From(simpleTime(testTime).Add(5, "y"), true), "5 years = 5 years")
}

func TestIdCalendarDay(t *testing.T) {
	assert := assert.New(t)

	testTime := time.Date(2000, 12, 15, 12, 0, 0, 0, time.UTC)
	timeNow = func() time.Time {
		return testTime
	}

	SetLocale("id")

	refTime := simpleTime(testTime)

	assert.Equal("hari ini pukul 12:00", simpleGoment(refTime).Calendar(), "today at the same time")
	assert.Equal("hari ini pukul 12:25", simpleGoment(refTime).Add(25, "m").Calendar(), "now plus 25 min")
	assert.Equal("hari ini pukul 13:00", simpleGoment(refTime).Add(1, "h").Calendar(), "now plus 1 hour")
	assert.Equal("besok pukul 12:00", simpleGoment(refTime).Add(1, "d").Calendar(), "tomorrow at the same time")
	assert.Equal("hari ini pukul 11:00", simpleGoment(refTime).Subtract(1, "h").Calendar(), "now minus 1 hour")
	assert.Equal("kemarin pukul 12:00", simpleGoment(refTime).Subtract(1, "d").Calendar(), "yesterday at the same time")

	refTime = simpleTime(testTime)

	assert.Equal("Minggu pukul 12:00", refTime.Add(2, "d").Calendar(), "Today + 2 days current time")
	refTime.StartOf("day")
	assert.Equal("Minggu pukul 00:00", refTime.Calendar(), "Today + 2 days beginning of day")
	refTime.EndOf("day")
	assert.Equal("Minggu pukul 23:59", refTime.Calendar(), "Today + 2 days end of day")

	refTime = simpleTime(testTime)

	assert.Equal("Rabu lalu pukul 12:00", refTime.Subtract(2, "d").Calendar(), "Today - 2 days current time")
	refTime.StartOf("day")
	assert.Equal("Rabu lalu pukul 00:00", refTime.Calendar(), "Today - 2 days beginning of day")
	refTime.EndOf("day")
	assert.Equal("Rabu lalu pukul 23:59", refTime.Calendar(), "Today - 2 days end of day")

	weeksAgo := simpleTime(testTime).Subtract(1, "w")
	weeksFromNow := simpleTime(testTime).Add(1, "w")

	assert.Equal("08/12/2000", weeksAgo.Calendar())
	assert.Equal("22/12/2000", weeksFromNow.Calendar())

	weeksAgo = simpleTime(testTime).Subtract(2, "w")
	weeksFromNow = simpleTime(testTime).Add(2, "w")

	assert.Equal("01/12/2000", weeksAgo.Calendar())
	assert.Equal("29/12/2000", weeksFromNow.Calendar())

	// Reset timeNow.
	timeNow = time.Now

	SetLocale("en")
}

func TestIdFormatParsing(t *testing.T) {
	assert := assert.New(t)

	formats := map[string][]string{
		"YYYY-Q":                    {"2014-4"},
		"MM-DD-YYYY":                {"12-02-1999"},
		"DD-MM-YYYY":                {"12-02-1999"},
		"DD/MM/YYYY":                {"12/02/1999"},
		"DD_MM_YYYY":                {"12_02_1999"},
		"DD:MM:YYYY":                {"12:02:1999"},
		"D-M-YY":                    {"2-2-99"},
		"Y":                         {"-0025"},
		"YY":                        {"99"},
		"DDD-YYYY":                  {"300-1999"},
		"YYYY-DDD":                  {"1999-300"},
		"YYYY MM Do":                {"2014 01 3", "2015 11 21"},
		"MMM":                       {"Jan"},
		"MMMM":                      {"Januari"},
		"DD MMMM":                   {"11 Januari"},
		"Do MMMM":                   {"3 Januari"},
		"YYYY MMMM":                 {"2018 Januari"},
		"D":                         {"3", "27"},
		"DD":                        {"04", "23"},
		"DDD":                       {"7", "300"},
		"DDDD":                      {"008", "211", "312"},
		"h":                         {"4"},
		"H":                         {"1", "10", "23"},
		"DD-MM-YYYY h:m:s":          {"12-02-1999 2:45:10"},
		"DD-MM-YYYY h:m:s a":        {"12-02-1999 2:45:10 dini hari", "12-02-1999 2:45:10 dini hari"},
		"h:mm a":                    {"12:00 siang", "12:30 siang", "12:00 siang", "12:30 siang"},
		"HH:mm":                     {"12:00"},
		"kk:mm":                     {"12:00"},
		"YYYY-MM-DDTHH:mm:ss":       {"2011-11-11T11:11:11"},
		"MM-DD-YYYY [M]":            {"12-02-1999 M"},
		"ddd MMM DD HH:mm:ss YYYY":  {"Kam Jan 08 22:52:51 2009"},
		"dddd MMM DD HH:mm:ss YYYY": {"Minggu Jan 11 22:52:51 2009"},
		"HH:mm:ss":                  {"12:00:00", "12:30:11", "00:00:00"},
		"kk:mm:ss":                  {"12:00:10", "12:30:42", "24:00:00", "09:00:00"},
		"YYYY-MM-DD HH:mm:ss ZZ":    {"2000-05-15 17:08:00 -0700"},
		"YYYY-MM-DD HH:mm Z":        {"2010-10-20 04:30 +00:00"},
		"X":                         {"1234567890"},
		"H Z":                       {"6 -06:00"},
		"H ZZ":                      {"5 -0700"},
		"LT":                        {"12:30"},
		"LTS":                       {"12:30:29"},
		"L":                         {"09/02/1999"},
		"l":                         {"9/2/1999"},
		"LL":                        {"2 Januari 1999"},
		"ll":                        {"2 Jan 1999"},
		"LLL":                       {"2 Januari 1999, 12:30"},
		"lll":                       {"2 Jan 1999, 12:30"},
		"LLLL":                      {"Sabtu, 2 Januari 1999, 12:30"},
		"llll":                      {"Kam, 14 Jan 2010, 15:25"},
	}

	for format, dates := range formats {
		for _, date := range dates {
			lib, _ := New(date, format, "id")
			assert.Equal(date, lib.Format(format), fmt.Sprintf("%v: %v", format, date))
		}
	}
}
