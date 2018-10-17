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

	lib, _ := New()

	details := lib.LocaleDetails()
	assert.Equal(t, "en", details.Code)
	assert.Equal(t, months, details.Resources.Months)
	assert.Equal(t, shortMonths, details.Resources.ShortMonths)
}
