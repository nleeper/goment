package goment

import (
	"testing"

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
