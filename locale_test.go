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
	months := map[string]int{"january": 1, "february": 2, "march": 3, "april": 4, "may": 5, "june": 6, "july": 7, "august": 8, "september": 9, "october": 10, "november": 11, "december": 12}
	shortMonths := map[string]int{"jan": 1, "feb": 2, "mar": 3, "apr": 4, "may": 5, "jun": 6, "jul": 7, "aug": 8, "sep": 9, "oct": 10, "nov": 11, "dec": 12}
	days := map[string]int{"sunday": 0, "monday": 1, "tuesday": 2, "wednesday": 3, "thursday": 4, "friday": 5, "saturday": 6}
	shortDays := map[string]int{"sun": 0, "mon": 1, "tue": 2, "wed": 3, "thu": 4, "fri": 5, "sat": 6}

	lib, _ := New()

	details := lib.LocaleDetails()
	assert.Equal(t, "en", details.Code)
	assert.Equal(t, months, details.Resources.Months)
	assert.Equal(t, shortMonths, details.Resources.ShortMonths)
	assert.Equal(t, days, details.Resources.Days)
	assert.Equal(t, shortDays, details.Resources.ShortDays)
}

func TestEsLocale(t *testing.T) {
	months := map[string]int{"enero": 1, "febrero": 2, "marzo": 3, "abril": 4, "mayo": 5, "junio": 6, "julio": 7, "agosto": 8, "septiembre": 9, "octubre": 10, "noviembre": 11, "diciembre": 12}
	shortMonths := map[string]int{"ene": 1, "feb": 2, "mar": 3, "abr": 4, "may": 5, "jun": 6, "jul": 7, "ago": 8, "sep": 9, "oct": 10, "nov": 11, "dic": 12}
	days := map[string]int{"domingo": 0, "lunes": 1, "martes": 2, "miércoles": 3, "jueves": 4, "viernes": 5, "sábado": 6}
	shortDays := map[string]int{"dom.": 0, "lun.": 1, "mar.": 2, "mié.": 3, "jue.": 4, "vie.": 5, "sáb.": 6}

	lib, _ := New()
	lib.SetLocale("es")

	details := lib.LocaleDetails()
	assert.Equal(t, "es", details.Code)
	assert.Equal(t, months, details.Resources.Months)
	assert.Equal(t, shortMonths, details.Resources.ShortMonths)
	assert.Equal(t, days, details.Resources.Days)
	assert.Equal(t, shortDays, details.Resources.ShortDays)
}

func TestEsFormat(t *testing.T) {

	formats := map[string]string{
		"dddd, MMMM Do YYYY, h:mm:ss a": "domingo, febrero 14º 2010, 3:25:50 pm",
		"ddd, hA":                       "dom., 3PM",
		// ['M Mo MM MMMM MMM',                   '2 2º 02 febrero feb.'],
		// ['YYYY YY',                            '2010 10'],
		// ['D Do DD',                            '14 14º 14'],
		// ['d do dddd ddd dd',                   '0 0º domingo dom. do'],
		// ['DDD DDDo DDDD',                      '45 45º 045'],
		// ['w wo ww',                            '6 6º 06'],
		// ['YYYY-MMM-DD',                        '2010-feb-14'],
		// ['h hh',                               '3 03'],
		// ['H HH',                               '15 15'],
		// ['m mm',                               '25 25'],
		// ['s ss',                               '50 50'],
		// ['a A',                                'pm PM'],
		// ['[the] DDDo [day of the year]',       'the 45º day of the year'],
		// ['LTS',                                '15:25:50'],
		// ['L',                                  '14/02/2010'],
		// ['LL',                                 '14 de febrero de 2010'],
		// ['LLL',                                '14 de febrero de 2010 15:25'],
		// ['LLLL',                               'domingo, 14 de febrero de 2010 15:25'],
		// ['l',                                  '14/2/2010'],
		// ['ll',                                 '14 de feb. de 2010'],
		// ['lll',                                '14 de feb. de 2010 15:25'],
		// ['llll',                               'dom., 14 de feb. de 2010 15:25']
	}

	lib := simpleTime(time.Date(2010, 2, 14, 15, 25, 50, 125000000, chicagoLocation()))
	lib.SetLocale("es")

	for p, r := range formats {
		assert.Equal(t, r, lib.Format(p), r)
	}
}

func TestEsParse(t *testing.T) {

}

func TestFrLocale(t *testing.T) {
	months := map[string]int{"janvier": 1, "février": 2, "mars": 3, "avril": 4, "mai": 5, "juin": 6, "juillet": 7, "août": 8, "septembre": 9, "octobre": 10, "novembre": 11, "décembre": 12}
	shortMonths := map[string]int{"janv.": 1, "févr.": 2, "mars": 3, "avr.": 4, "mai": 5, "juin": 6, "juil.": 7, "août": 8, "sept.": 9, "oct.": 10, "nov.": 11, "déc.": 12}
	days := map[string]int{"dimanche": 0, "lundi": 1, "mardi": 2, "mercredi": 3, "jeudi": 4, "vendredi": 5, "samedi": 6}
	shortDays := map[string]int{"dim.": 0, "lun.": 1, "mar.": 2, "mer.": 3, "jeu.": 4, "ven.": 5, "sam.": 6}

	lib, _ := New()
	lib.SetLocale("fr")

	details := lib.LocaleDetails()
	assert.Equal(t, "fr", details.Code)
	assert.Equal(t, months, details.Resources.Months)
	assert.Equal(t, shortMonths, details.Resources.ShortMonths)
	assert.Equal(t, days, details.Resources.Days)
	assert.Equal(t, shortDays, details.Resources.ShortDays)
}

func TestFrFormat(t *testing.T) {

}

func TestFrParse(t *testing.T) {

}
func TestDeLocale(t *testing.T) {
	months := map[string]int{"Januar": 1, "Februar": 2, "März": 3, "April": 4, "Mai": 5, "Juni": 6, "Juli": 7, "August": 8, "September": 9, "Oktober": 10, "November": 11, "Dezember": 12}
	shortMonths := map[string]int{"Jan.": 1, "Feb.": 2, "März": 3, "Apr.": 4, "Mai": 5, "Juni": 6, "Juli": 7, "Aug.": 8, "Sep.": 9, "Okt.": 10, "Nov.": 11, "Dez.": 12}
	days := map[string]int{"Sonntag": 0, "Montag": 1, "Dienstag": 2, "Mittwoch": 3, "Donnerstag": 4, "Freitag": 5, "Samstag": 6}
	shortDays := map[string]int{"So.": 0, "Mo.": 1, "Di.": 2, "Mi.": 3, "Do.": 4, "Fr.": 5, "Sa.": 6}

	lib, _ := New()
	assert.NoError(t, lib.SetLocale("de"))

	details := lib.LocaleDetails()
	assert.Equal(t, "de", details.Code)
	assert.Equal(t, months, details.Resources.Months)
	assert.Equal(t, shortMonths, details.Resources.ShortMonths)
	assert.Equal(t, days, details.Resources.Days)
	assert.Equal(t, shortDays, details.Resources.ShortDays)
}
