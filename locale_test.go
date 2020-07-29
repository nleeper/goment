package goment

import (
	"fmt"
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
	assert.Equal(t, DefaultLocaleCode, Locale())
	SetLocale("es")
	assert.Equal(t, "es", Locale(), "Set global locale")

	resetLocale()
}

func TestSetGlobalLocaleUsedForNew(t *testing.T) {
	SetLocale("es")
	lib, _ := New()
	assert.Equal(t, "es", lib.Locale(), "Use set global locale")

	resetLocale()
}

func TestSetLocale(t *testing.T) {
	lib, _ := New()
	assert.Equal(t, DefaultLocaleCode, lib.Locale())
	lib.SetLocale("fr")
	assert.Equal(t, "fr", lib.Locale())
}

func TestLocaleNotChangedForExistingGoment(t *testing.T) {
	lib, _ := New()
	lib.SetLocale("fr")
	assert.Equal(t, "fr", lib.Locale())

	SetLocale("es")
	assert.Equal(t, "fr", lib.Locale())
	assert.Equal(t, "es", Locale())

	lib2, _ := New()
	assert.Equal(t, "es", lib2.Locale())

	resetLocale()
}

func TestEnLocale(t *testing.T) {
	longDays := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	shortDays := []string{"Sun", "Mon", "Tue", "Wed", "Thu", "Fri", "Sat"}
	minDays := []string{"Su", "Mo", "Tu", "We", "Th", "Fr", "Sa"}
	shortMonths := []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}
	longMonths := []string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}

	lib, _ := New()

	assert.Equal(t, "en", lib.Locale())
	assert.Equal(t, longDays, lib.Weekdays())
	assert.Equal(t, shortDays, lib.WeekdaysShort())
	assert.Equal(t, minDays, lib.WeekdaysMin())
	assert.Equal(t, longMonths, lib.Months())
	assert.Equal(t, shortMonths, lib.MonthsShort())
}

func TestEsLocale(t *testing.T) {
	longDays := []string{"domingo", "lunes", "martes", "miércoles", "jueves", "viernes", "sábado"}
	shortDays := []string{"dom.", "lun.", "mar.", "mié.", "jue.", "vie.", "sáb."}
	minDays := []string{"do", "lu", "ma", "mi", "ju", "vi", "sá"}
	shortMonths := []string{"ene", "feb", "mar", "abr", "may", "jun", "jul", "ago", "sep", "oct", "nov", "dic"}
	longMonths := []string{"enero", "febrero", "marzo", "abril", "mayo", "junio", "julio", "agosto", "septiembre", "octubre", "noviembre", "diciembre"}

	lib, _ := New()
	lib.SetLocale("es")

	assert.Equal(t, "es", lib.Locale())
	assert.Equal(t, longDays, lib.Weekdays())
	assert.Equal(t, shortDays, lib.WeekdaysShort())
	assert.Equal(t, minDays, lib.WeekdaysMin())
	assert.Equal(t, longMonths, lib.Months())
	assert.Equal(t, shortMonths, lib.MonthsShort())
}

func TestEsFormat(t *testing.T) {
	formats := map[string]string{
		"dddd, MMMM Do YYYY, h:mm:ss a": "domingo, febrero 14º 2010, 3:25:50 pm",
		"ddd, hA":                       "dom., 3PM",
		"M Mo MM MMMM MMM":              "2 2º 02 febrero feb",
		"YYYY YY":                       "2010 10",
		"D Do DD":                       "14 14º 14",
		"d do dddd ddd dd":              "0 0º domingo dom. do",
		"DDD DDDo DDDD":                 "45 45º 045",
		// "w wo ww":                       "6 6º 06",
		"YYYY-MMM-DD":                  "2010-feb-14",
		"h hh":                         "3 03",
		"H HH":                         "15 15",
		"m mm":                         "25 25",
		"s ss":                         "50 50",
		"a A":                          "pm PM",
		"[the] DDDo [day of the year]": "the 45º day of the year",
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
		assert.Equal(t, r, lib.Format(p), r)
	}
}

func TestEsShortDayFormatWithUnicode(t *testing.T) {
	lib := simpleTime(time.Date(2010, 2, 20, 15, 25, 50, 125000000, chicagoLocation()))
	lib.SetLocale("es")

	assert.Equal(t, "sá", lib.Format("dd"))
}

func TestEsRelativeTime(t *testing.T) {
	testTime := time.Date(2007, 1, 28, 0, 0, 0, 0, chicagoLocation())
	lib := simpleTime(testTime)

	lib.SetLocale("es")

	assert.Equal(t, "hace unos segundos", lib.From(simpleTime(testTime).Add(44, "s")), "44 seconds = a few seconds ago")
	assert.Equal(t, "unos segundos", lib.From(simpleTime(testTime).Add(44, "s"), true), "44 seconds = a few seconds")
	assert.Equal(t, "hace un minuto", lib.From(simpleTime(testTime).Add(1, "m")), "1 minute = a minute ago")
	assert.Equal(t, "un minuto", lib.From(simpleTime(testTime).Add(1, "m"), true), "1 minute = a minute")
	assert.Equal(t, "hace 44 minutos", lib.From(simpleTime(testTime).Add(44, "m")), "44 minutes = 44 minutes ago")
	assert.Equal(t, "44 minutos", lib.From(simpleTime(testTime).Add(44, "m"), true), "44 minutes = 44 minutes")
	assert.Equal(t, "hace una hora", lib.From(simpleTime(testTime).Add(1, "h")), "1 hour = an hour ago")
	assert.Equal(t, "una hora", lib.From(simpleTime(testTime).Add(1, "h"), true), "1 hour = an hour")
	assert.Equal(t, "hace 2 horas", lib.From(simpleTime(testTime).Add(2, "h")), "2 hours = 2 hours ago")
	assert.Equal(t, "2 horas", lib.From(simpleTime(testTime).Add(2, "h"), true), "2 hours = 2 hours")
	assert.Equal(t, "hace un día", lib.From(simpleTime(testTime).Add(1, "d")), "1 day = a day ago")
	assert.Equal(t, "un día", lib.From(simpleTime(testTime).Add(1, "d"), true), "1 day = a day")
	assert.Equal(t, "hace 5 días", lib.From(simpleTime(testTime).Add(5, "d")), "5 days = 5 days ago")
	assert.Equal(t, "5 días", lib.From(simpleTime(testTime).Add(5, "d"), true), "5 days = 5 days")
	assert.Equal(t, "hace un mes", lib.From(simpleTime(testTime).Add(1, "M")), "1 month = a month ago")
	assert.Equal(t, "un mes", lib.From(simpleTime(testTime).Add(1, "M"), true), "1 month = a month")
	assert.Equal(t, "hace 5 meses", lib.From(simpleTime(testTime).Add(5, "M")), "5 months = 5 months ago")
	assert.Equal(t, "5 meses", lib.From(simpleTime(testTime).Add(5, "M"), true), "5 months = 5 months")
	assert.Equal(t, "hace un año", lib.From(simpleTime(testTime).Add(1, "y")), "1 year = a year ago")
	assert.Equal(t, "un año", lib.From(simpleTime(testTime).Add(1, "y"), true), "1 year = a year")
	assert.Equal(t, "hace 5 años", lib.From(simpleTime(testTime).Add(5, "y")), "5 years = 5 years ago")
	assert.Equal(t, "5 años", lib.From(simpleTime(testTime).Add(5, "y"), true), "5 years = 5 years")
}

func TestEsCalendarDay(t *testing.T) {
	testTime := time.Date(2000, 12, 15, 12, 0, 0, 0, time.UTC)
	timeNow = func() time.Time {
		return testTime
	}

	SetLocale("es")

	refTime := simpleTime(testTime)

	assert.Equal(t, "hoy a las 12:00", simpleGoment(refTime).Calendar(), "today at the same time")
	assert.Equal(t, "hoy a las 12:25", simpleGoment(refTime).Add(25, "m").Calendar(), "now plus 25 min")
	assert.Equal(t, "hoy a las 13:00", simpleGoment(refTime).Add(1, "h").Calendar(), "now plus 1 hour")
	assert.Equal(t, "mañana a las 12:00", simpleGoment(refTime).Add(1, "d").Calendar(), "tomorrow at the same time")
	assert.Equal(t, "hoy a las 11:00", simpleGoment(refTime).Subtract(1, "h").Calendar(), "now minus 1 hour")
	assert.Equal(t, "ayer a las 12:00", simpleGoment(refTime).Subtract(1, "d").Calendar(), "yesterday at the same time")

	refTime = simpleTime(testTime)

	assert.Equal(t, "domingo a las 12:00", refTime.Add(2, "d").Calendar(), "Today + 2 days current time")
	refTime.StartOf("day")
	assert.Equal(t, "domingo a las 0:00", refTime.Calendar(), "Today + 2 days beginning of day")
	refTime.EndOf("day")
	assert.Equal(t, "domingo a las 23:59", refTime.Calendar(), "Today + 2 days end of day")

	refTime = simpleTime(testTime)

	assert.Equal(t, "el miércoles pasado a las 12:00", refTime.Subtract(2, "d").Calendar(), "Today - 2 days current time")
	refTime.StartOf("day")
	assert.Equal(t, "el miércoles pasado a las 0:00", refTime.Calendar(), "Today - 2 days beginning of day")
	refTime.EndOf("day")
	assert.Equal(t, "el miércoles pasado a las 23:59", refTime.Calendar(), "Today - 2 days end of day")

	weeksAgo := simpleTime(testTime).Subtract(1, "w")
	weeksFromNow := simpleTime(testTime).Add(1, "w")

	assert.Equal(t, "08/12/2000", weeksAgo.Calendar())
	assert.Equal(t, "22/12/2000", weeksFromNow.Calendar())

	weeksAgo = simpleTime(testTime).Subtract(2, "w")
	weeksFromNow = simpleTime(testTime).Add(2, "w")

	assert.Equal(t, "01/12/2000", weeksAgo.Calendar())
	assert.Equal(t, "29/12/2000", weeksFromNow.Calendar())

	// Reset timeNow.
	timeNow = time.Now

	SetLocale("en")
}

func TestEsFormatParsing(t *testing.T) {
	formats := map[string][]string{
		"YYYY-Q":                    []string{"2014-4"},
		"MM-DD-YYYY":                []string{"12-02-1999"},
		"DD-MM-YYYY":                []string{"12-02-1999"},
		"DD/MM/YYYY":                []string{"12/02/1999"},
		"DD_MM_YYYY":                []string{"12_02_1999"},
		"DD:MM:YYYY":                []string{"12:02:1999"},
		"D-M-YY":                    []string{"2-2-99"},
		"Y":                         []string{"-0025"},
		"YY":                        []string{"99"},
		"DDD-YYYY":                  []string{"300-1999"},
		"YYYY-DDD":                  []string{"1999-300"},
		"YYYY MM Do":                []string{"2014 01 3º", "2015 11 21º"},
		"MMM":                       []string{"abr"},
		"MMMM":                      []string{"septiembre"},
		"DD MMMM":                   []string{"11 septiembre"},
		"Do MMMM":                   []string{"3º septiembre"},
		"YYYY MMMM":                 []string{"2018 octubre"},
		"D":                         []string{"3", "27"},
		"DD":                        []string{"04", "23"},
		"DDD":                       []string{"7", "300"},
		"DDDD":                      []string{"008", "211", "312"},
		"h":                         []string{"4"},
		"H":                         []string{"1", "10", "23"},
		"DD-MM-YYYY h:m:s":          []string{"12-02-1999 2:45:10"},
		"DD-MM-YYYY h:m:s a":        []string{"12-02-1999 2:45:10 am", "12-02-1999 2:45:10 pm"},
		"h:mm a":                    []string{"12:00 pm", "12:30 pm", "12:00 am", "12:30 am"},
		"HH:mm":                     []string{"12:00"},
		"kk:mm":                     []string{"12:00"},
		"YYYY-MM-DDTHH:mm:ss":       []string{"2011-11-11T11:11:11"},
		"MM-DD-YYYY [M]":            []string{"12-02-1999 M"},
		"ddd MMM DD HH:mm:ss YYYY":  []string{"mié. abr 08 22:52:51 2009"},
		"dddd MMM DD HH:mm:ss YYYY": []string{"sábado abr 11 22:52:51 2009"},
		"HH:mm:ss":                  []string{"12:00:00", "12:30:11", "00:00:00"},
		"kk:mm:ss":                  []string{"12:00:10", "12:30:42", "24:00:00", "09:00:00"},
		"YYYY-MM-DD HH:mm:ss ZZ":    []string{"2000-05-15 17:08:00 -0700"},
		"YYYY-MM-DD HH:mm Z":        []string{"2010-10-20 04:30 +00:00"},
		// "e":                         []string{"0", "5"},
		// "E": []string{"1", "7"},
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
		"X":    []string{"1234567890"},
		"H Z":  []string{"6 -06:00"},
		"H ZZ": []string{"5 -0700"},
		"LT":   []string{"12:30"},
		"LTS":  []string{"12:30:29"},
		"L":    []string{"09/02/1999"},
		"l":    []string{"9/2/1999"},
		"LL":   []string{"2 de septiembre de 1999"},
		"ll":   []string{"2 de sep de 1999"},
		"LLL":  []string{"2 de septiembre de 1999 12:30"},
		"lll":  []string{"2 de sep de 1999 12:30"},
		"LLLL": []string{"jueves, 2 de septiembre de 1999 12:30"},
		"llll": []string{"jue., 2 de sep de 1999 12:30"},
	}

	for format, dates := range formats {
		for _, date := range dates {
			lib, _ := New(date, format, "es")
			assert.Equal(t, date, lib.Format(format), fmt.Sprintf("%v: %v", format, date))
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
	lib, _ := New("mié. abr 08 22:52:51 2009", "ddd MMM DD HH:mm:ss YYYY", "es")
	assert.Equal(t, "mié. abr 08 22:52:51 2009", lib.Format("ddd MMM DD HH:mm:ss YYYY"))

	lib2, _ := New("sábado abr 11 22:52:51 2009", "dddd MMM DD HH:mm:ss YYYY", "es")
	assert.Equal(t, "sábado abr 11 22:52:51 2009", lib2.Format("dddd MMM DD HH:mm:ss YYYY"))
}

func TestFrLocale(t *testing.T) {
	longDays := []string{"dimanche", "lundi", "mardi", "mercredi", "jeudi", "vendredi", "samedi"}
	shortDays := []string{"dim.", "lun.", "mar.", "mer.", "jeu.", "ven.", "sam."}
	minDays := []string{"di", "lu", "ma", "me", "je", "ve", "sa"}
	shortMonths := []string{"janv.", "févr.", "mars", "avr.", "mai", "juin", "juil.", "août", "sept.", "oct.", "nov.", "déc."}
	longMonths := []string{"janvier", "février", "mars", "avril", "mai", "juin", "juillet", "août", "septembre", "octobre", "novembre", "décembre"}

	lib, _ := New()
	lib.SetLocale("fr")

	assert.Equal(t, "fr", lib.Locale())
	assert.Equal(t, longDays, lib.Weekdays())
	assert.Equal(t, shortDays, lib.WeekdaysShort())
	assert.Equal(t, minDays, lib.WeekdaysMin())
	assert.Equal(t, longMonths, lib.Months())
	assert.Equal(t, shortMonths, lib.MonthsShort())
}

func TestFrFormat(t *testing.T) {
	formats := map[string]string{
		"dddd, MMMM Do YYYY, h:mm:ss a": "dimanche, février 14 2010, 3:25:50 pm",
		"ddd, hA":                       "dim., 3PM",
		"M Mo MM MMMM MMM":              "2 2e 02 février févr.",
		"YYYY YY":                       "2010 10",
		"D Do DD":                       "14 14 14",
		"d do dddd ddd dd":              "0 0e dimanche dim. di",
		"DDD DDDo DDDD":                 "45 45e 045",
		// "w wo ww":                       "6 6e 06",
		"h hh":                        "3 03",
		"H HH":                        "15 15",
		"m mm":                        "25 25",
		"s ss":                        "50 50",
		"a A":                         "pm PM",
		"[le] Do [jour du mois]":      "le 14 jour du mois",
		"[le] DDDo [jour de l’année]": "le 45e jour de l’année",
		"LTS":                         "15:25:50",
		"L":                           "14/02/2010",
		"LL":                          "14 février 2010",
		"LLL":                         "14 février 2010 15:25",
		"LLLL":                        "dimanche 14 février 2010 15:25",
		"l":                           "14/2/2010",
		"ll":                          "14 févr. 2010",
		"lll":                         "14 févr. 2010 15:25",
		"llll":                        "dim. 14 févr. 2010 15:25",
	}

	lib := simpleTime(time.Date(2010, 2, 14, 15, 25, 50, 125000000, chicagoLocation()))
	lib.SetLocale("fr")

	for p, r := range formats {
		assert.Equal(t, r, lib.Format(p), r)
	}
}

func TestFrOrdinal(t *testing.T) {
	SetLocale("fr")
	assert.Equal(t, "1er", simple(DateTime{Year: 2017, Month: 1, Day: 1}).Format("Mo"))
	assert.Equal(t, "2e", simple(DateTime{Year: 2017, Month: 2, Day: 1}).Format("Mo"))
	assert.Equal(t, "1er", simple(DateTime{Year: 2017, Month: 1, Day: 1}).Format("Qo"))
	assert.Equal(t, "2e", simple(DateTime{Year: 2017, Month: 4, Day: 1}).Format("Qo"))
	assert.Equal(t, "1er", simple(DateTime{Year: 2017, Month: 1, Day: 1}).Format("Do"))
	assert.Equal(t, "2", simple(DateTime{Year: 2017, Month: 1, Day: 2}).Format("Do"))
	assert.Equal(t, "1er", simple(DateTime{Year: 2011, Month: 1, Day: 1}).Format("DDDo"))
	assert.Equal(t, "2e", simple(DateTime{Year: 2011, Month: 1, Day: 2}).Format("DDDo"))
	assert.Equal(t, "3e", simple(DateTime{Year: 2011, Month: 1, Day: 3}).Format("DDDo"))
	assert.Equal(t, "4e", simple(DateTime{Year: 2011, Month: 1, Day: 4}).Format("DDDo"))
	assert.Equal(t, "5e", simple(DateTime{Year: 2011, Month: 1, Day: 5}).Format("DDDo"))
	assert.Equal(t, "6e", simple(DateTime{Year: 2011, Month: 1, Day: 6}).Format("DDDo"))
	assert.Equal(t, "7e", simple(DateTime{Year: 2011, Month: 1, Day: 7}).Format("DDDo"))
	assert.Equal(t, "8e", simple(DateTime{Year: 2011, Month: 1, Day: 8}).Format("DDDo"))
	assert.Equal(t, "9e", simple(DateTime{Year: 2011, Month: 1, Day: 9}).Format("DDDo"))
	assert.Equal(t, "10e", simple(DateTime{Year: 2011, Month: 1, Day: 10}).Format("DDDo"))
	assert.Equal(t, "11e", simple(DateTime{Year: 2011, Month: 1, Day: 11}).Format("DDDo"))
	assert.Equal(t, "12e", simple(DateTime{Year: 2011, Month: 1, Day: 12}).Format("DDDo"))
	assert.Equal(t, "13e", simple(DateTime{Year: 2011, Month: 1, Day: 13}).Format("DDDo"))
	assert.Equal(t, "14e", simple(DateTime{Year: 2011, Month: 1, Day: 14}).Format("DDDo"))
	assert.Equal(t, "15e", simple(DateTime{Year: 2011, Month: 1, Day: 15}).Format("DDDo"))
	assert.Equal(t, "16e", simple(DateTime{Year: 2011, Month: 1, Day: 16}).Format("DDDo"))
	assert.Equal(t, "17e", simple(DateTime{Year: 2011, Month: 1, Day: 17}).Format("DDDo"))
	assert.Equal(t, "18e", simple(DateTime{Year: 2011, Month: 1, Day: 18}).Format("DDDo"))
	assert.Equal(t, "19e", simple(DateTime{Year: 2011, Month: 1, Day: 19}).Format("DDDo"))
	assert.Equal(t, "20e", simple(DateTime{Year: 2011, Month: 1, Day: 20}).Format("DDDo"))
	assert.Equal(t, "21e", simple(DateTime{Year: 2011, Month: 1, Day: 21}).Format("DDDo"))
	assert.Equal(t, "22e", simple(DateTime{Year: 2011, Month: 1, Day: 22}).Format("DDDo"))
	assert.Equal(t, "23e", simple(DateTime{Year: 2011, Month: 1, Day: 23}).Format("DDDo"))
	assert.Equal(t, "24e", simple(DateTime{Year: 2011, Month: 1, Day: 24}).Format("DDDo"))
	assert.Equal(t, "25e", simple(DateTime{Year: 2011, Month: 1, Day: 25}).Format("DDDo"))
	assert.Equal(t, "26e", simple(DateTime{Year: 2011, Month: 1, Day: 26}).Format("DDDo"))
	assert.Equal(t, "27e", simple(DateTime{Year: 2011, Month: 1, Day: 27}).Format("DDDo"))
	assert.Equal(t, "28e", simple(DateTime{Year: 2011, Month: 1, Day: 28}).Format("DDDo"))
	assert.Equal(t, "29e", simple(DateTime{Year: 2011, Month: 1, Day: 29}).Format("DDDo"))
	assert.Equal(t, "30e", simple(DateTime{Year: 2011, Month: 1, Day: 30}).Format("DDDo"))
	assert.Equal(t, "31e", simple(DateTime{Year: 2011, Month: 1, Day: 31}).Format("DDDo"))
	assert.Equal(t, "0e", simple(DateTime{Year: 2017, Month: 1, Day: 1}).Format("do"))
	assert.Equal(t, "1er", simple(DateTime{Year: 2017, Month: 1, Day: 2}).Format("do"))
	// assert.Equal(t, "1re 1re", simple(DateTime{Year: 2017, Month: 1, Day: 4}).Format("wo Wo"))
	// assert.Equal(t, "2e 2e", simple(DateTime{Year: 2017, Month: 1, Day: 11}).Format("wo Wo"))
	assert.Equal(t, "1re", simple(DateTime{Year: 2017, Month: 1, Day: 4}).Format("Wo"))
	assert.Equal(t, "2e", simple(DateTime{Year: 2017, Month: 1, Day: 11}).Format("Wo"))
	SetLocale("en")
}

func TestFrRelativeTime(t *testing.T) {
	testTime := time.Date(2007, 1, 28, 0, 0, 0, 0, chicagoLocation())
	lib := simpleTime(testTime)

	lib.SetLocale("fr")

	assert.Equal(t, "il y a quelques secondes", lib.From(simpleTime(testTime).Add(44, "s")), "44 seconds = a few seconds ago")
	assert.Equal(t, "quelques secondes", lib.From(simpleTime(testTime).Add(44, "s"), true), "44 seconds = a few seconds")
	assert.Equal(t, "il y a une minute", lib.From(simpleTime(testTime).Add(1, "m")), "1 minute = a minute ago")
	assert.Equal(t, "une minute", lib.From(simpleTime(testTime).Add(1, "m"), true), "1 minute = a minute")
	assert.Equal(t, "il y a 44 minutes", lib.From(simpleTime(testTime).Add(44, "m")), "44 minutes = 44 minutes ago")
	assert.Equal(t, "44 minutes", lib.From(simpleTime(testTime).Add(44, "m"), true), "44 minutes = 44 minutes")
	assert.Equal(t, "il y a une heure", lib.From(simpleTime(testTime).Add(1, "h")), "1 hour = an hour ago")
	assert.Equal(t, "une heure", lib.From(simpleTime(testTime).Add(1, "h"), true), "1 hour = an hour")
	assert.Equal(t, "il y a 2 heures", lib.From(simpleTime(testTime).Add(2, "h")), "2 hours = 2 hours ago")
	assert.Equal(t, "2 heures", lib.From(simpleTime(testTime).Add(2, "h"), true), "2 hours = 2 hours")
	assert.Equal(t, "il y a un jour", lib.From(simpleTime(testTime).Add(1, "d")), "1 day = a day ago")
	assert.Equal(t, "un jour", lib.From(simpleTime(testTime).Add(1, "d"), true), "1 day = a day")
	assert.Equal(t, "il y a 5 jours", lib.From(simpleTime(testTime).Add(5, "d")), "5 days = 5 days ago")
	assert.Equal(t, "5 jours", lib.From(simpleTime(testTime).Add(5, "d"), true), "5 days = 5 days")
	assert.Equal(t, "il y a un mois", lib.From(simpleTime(testTime).Add(1, "M")), "1 month = a month ago")
	assert.Equal(t, "un mois", lib.From(simpleTime(testTime).Add(1, "M"), true), "1 month = a month")
	assert.Equal(t, "il y a 5 mois", lib.From(simpleTime(testTime).Add(5, "M")), "5 months = 5 months ago")
	assert.Equal(t, "5 mois", lib.From(simpleTime(testTime).Add(5, "M"), true), "5 months = 5 months")
	assert.Equal(t, "il y a un an", lib.From(simpleTime(testTime).Add(1, "y")), "1 year = a year ago")
	assert.Equal(t, "un an", lib.From(simpleTime(testTime).Add(1, "y"), true), "1 year = a year")
	assert.Equal(t, "il y a 5 ans", lib.From(simpleTime(testTime).Add(5, "y")), "5 years = 5 years ago")
	assert.Equal(t, "5 ans", lib.From(simpleTime(testTime).Add(5, "y"), true), "5 years = 5 years")
}

func TestFrCalendarDay(t *testing.T) {
	testTime := time.Date(2000, 12, 15, 12, 0, 0, 0, time.UTC)
	timeNow = func() time.Time {
		return testTime
	}

	SetLocale("fr")

	refTime := simpleTime(testTime)

	assert.Equal(t, "Aujourd’hui à 12:00", simpleGoment(refTime).Calendar(), "today at the same time")
	assert.Equal(t, "Aujourd’hui à 12:25", simpleGoment(refTime).Add(25, "m").Calendar(), "now plus 25 min")
	assert.Equal(t, "Aujourd’hui à 13:00", simpleGoment(refTime).Add(1, "h").Calendar(), "now plus 1 hour")
	assert.Equal(t, "Demain à 12:00", simpleGoment(refTime).Add(1, "d").Calendar(), "tomorrow at the same time")
	assert.Equal(t, "Aujourd’hui à 11:00", simpleGoment(refTime).Subtract(1, "h").Calendar(), "now minus 1 hour")
	assert.Equal(t, "Hier à 12:00", simpleGoment(refTime).Subtract(1, "d").Calendar(), "yesterday at the same time")

	refTime = simpleTime(testTime)

	assert.Equal(t, "dimanche à 12:00", refTime.Add(2, "d").Calendar(), "Today + 2 days current time")
	refTime.StartOf("day")
	assert.Equal(t, "dimanche à 00:00", refTime.Calendar(), "Today + 2 days beginning of day")
	refTime.EndOf("day")
	assert.Equal(t, "dimanche à 23:59", refTime.Calendar(), "Today + 2 days end of day")

	refTime = simpleTime(testTime)

	assert.Equal(t, "mercredi dernier à 12:00", refTime.Subtract(2, "d").Calendar(), "Today - 2 days current time")
	refTime.StartOf("day")
	assert.Equal(t, "mercredi dernier à 00:00", refTime.Calendar(), "Today - 2 days beginning of day")
	refTime.EndOf("day")
	assert.Equal(t, "mercredi dernier à 23:59", refTime.Calendar(), "Today - 2 days end of day")

	weeksAgo := simpleTime(testTime).Subtract(1, "w")
	weeksFromNow := simpleTime(testTime).Add(1, "w")

	assert.Equal(t, "08/12/2000", weeksAgo.Calendar())
	assert.Equal(t, "22/12/2000", weeksFromNow.Calendar())

	weeksAgo = simpleTime(testTime).Subtract(2, "w")
	weeksFromNow = simpleTime(testTime).Add(2, "w")

	assert.Equal(t, "01/12/2000", weeksAgo.Calendar())
	assert.Equal(t, "29/12/2000", weeksFromNow.Calendar())

	// Reset timeNow.
	timeNow = time.Now

	SetLocale("en")
}

func TestFrFormatParsing(t *testing.T) {
	formats := map[string][]string{
		"YYYY-Q":                    []string{"2014-4"},
		"MM-DD-YYYY":                []string{"12-02-1999"},
		"DD-MM-YYYY":                []string{"12-02-1999"},
		"DD/MM/YYYY":                []string{"12/02/1999"},
		"DD_MM_YYYY":                []string{"12_02_1999"},
		"DD:MM:YYYY":                []string{"12:02:1999"},
		"D-M-YY":                    []string{"2-2-99"},
		"Y":                         []string{"-0025"},
		"YY":                        []string{"99"},
		"DDD-YYYY":                  []string{"300-1999"},
		"YYYY-DDD":                  []string{"1999-300"},
		"YYYY MM Do":                []string{"2014 01 1er", "2015 11 21"},
		"MMM":                       []string{"avr."},
		"MMMM":                      []string{"septembre"},
		"DD MMMM":                   []string{"11 septembre"},
		"Do MMMM":                   []string{"3 septembre"},
		"YYYY MMMM":                 []string{"2018 octobre"},
		"D":                         []string{"3", "27"},
		"DD":                        []string{"04", "23"},
		"DDD":                       []string{"7", "300"},
		"DDDD":                      []string{"008", "211", "312"},
		"h":                         []string{"4"},
		"H":                         []string{"1", "10", "23"},
		"DD-MM-YYYY h:m:s":          []string{"12-02-1999 2:45:10"},
		"DD-MM-YYYY h:m:s a":        []string{"12-02-1999 2:45:10 am", "12-02-1999 2:45:10 pm"},
		"h:mm a":                    []string{"12:00 pm", "12:30 pm", "12:00 am", "12:30 am"},
		"HH:mm":                     []string{"12:00"},
		"kk:mm":                     []string{"12:00"},
		"YYYY-MM-DDTHH:mm:ss":       []string{"2011-11-11T11:11:11"},
		"MM-DD-YYYY [M]":            []string{"12-02-1999 M"},
		"ddd MMM DD HH:mm:ss YYYY":  []string{"mer. avr. 08 22:52:51 2009"},
		"dddd MMM DD HH:mm:ss YYYY": []string{"samedi avr. 11 22:52:51 2009"},
		"HH:mm:ss":                  []string{"12:00:00", "12:30:11", "00:00:00"},
		"kk:mm:ss":                  []string{"12:00:10", "12:30:42", "24:00:00", "09:00:00"},
		"YYYY-MM-DD HH:mm:ss ZZ":    []string{"2000-05-15 17:08:00 -0700"},
		"YYYY-MM-DD HH:mm Z":        []string{"2010-10-20 04:30 +00:00"},
		// "e":                         []string{"0", "5"},
		// "E": []string{"1", "7"},
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
		"X":    []string{"1234567890"},
		"H Z":  []string{"6 -06:00"},
		"H ZZ": []string{"5 -0700"},
		"LT":   []string{"12:30"},
		"LTS":  []string{"12:30:29"},
		"L":    []string{"09/02/1999"},
		"l":    []string{"9/2/1999"},
		"LL":   []string{"2 septembre 1999"},
		"ll":   []string{"2 sept. 1999"},
		"LLL":  []string{"2 septembre 1999 12:30"},
		"lll":  []string{"2 sept. 1999 12:30"},
		"LLLL": []string{"jeudi 2 septembre 1999 12:30"},
		"llll": []string{"jeu. 2 sept. 1999 12:30"},
	}

	for format, dates := range formats {
		for _, date := range dates {
			lib, _ := New(date, format, "fr")
			assert.Equal(t, date, lib.Format(format), fmt.Sprintf("%v: %v", format, date))
		}
	}
}
