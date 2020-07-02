package goment

import (
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
