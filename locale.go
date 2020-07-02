package goment

import (
	"errors"
	"strings"

	"github.com/nleeper/goment/locales"
)

// TODO - add support for https://momentjs.com/docs/#/i18n/listing-months-weekdays/

// DefaultLocaleCode is the default locale used by Goment if not set.
const DefaultLocaleCode = "en"

var supportedLocales = map[string]locales.LocaleDetails{
	"en": locales.EnLocale,
	"es": locales.EsLocale,
	"fr": locales.FrLocale,
}

var globalLocale = loadKnownLocale(DefaultLocaleCode)

// Locale gets the current global locale code.
func Locale() string {
	return globalLocale.Code
}

// Locale gets the locale code for the current Goment instance.
func (g *Goment) Locale() string {
	return g.locale.Code
}

// LocaleDetails gets the locale details for the current Goment instance.
func (g *Goment) LocaleDetails() locales.LocaleDetails {
	return g.locale
}

// SetLocale sets the global locale for all Coment instances.
func SetLocale(localeCode string) error {
	if strings.ToLower(localeCode) != globalLocale.Code {
		loc, err := loadLocale(localeCode)
		if err != nil {
			return err
		}

		globalLocale = loc
	}
	return nil
}

// SetLocale sets the locale for only the current Goment instance.
func (g *Goment) SetLocale(localeCode string) error {
	if strings.ToLower(localeCode) != g.locale.Code {
		loc, err := loadLocale(localeCode)
		if err != nil {
			return err
		}

		g.locale = loc
	}
	return nil
}

// Months returns the list of months in the current locale.
func (g *Goment) Months() []string {
	return g.locale.Months
}

// MonthsShort returns the list of abbreviated month names in the current locale.
func (g *Goment) MonthsShort() []string {
	return g.locale.MonthsShort
}

// Weekdays returns the list of weekdays in the current locale.
func (g *Goment) Weekdays() []string {
	return g.locale.Weekdays
}

// WeekdaysShort returns the list of abbreviated weekday names in the current locale.
func (g *Goment) WeekdaysShort() []string {
	return g.locale.WeekdaysShort
}

// WeekdaysMin returns the list of weekdays in the current locale.
func (g *Goment) WeekdaysMin() []string {
	return g.locale.WeekdaysMin
}

func getGlobalLocaleDetails() locales.LocaleDetails {
	return globalLocale
}

func loadKnownLocale(localeCode string) locales.LocaleDetails {
	locale, _ := loadLocale(localeCode)
	return locale
}

func loadLocale(localeCode string) (locales.LocaleDetails, error) {
	normalizedCode := strings.ToLower(localeCode)

	if locale, exist := supportedLocales[normalizedCode]; exist {
		return locale, nil
	}

	return locales.LocaleDetails{}, errors.New("Locale " + normalizedCode + " is not supported")
}
