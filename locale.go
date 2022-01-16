package goment

import (
	"errors"
	"strings"

	"github.com/nleeper/goment/locales"
)

// DefaultLocaleCode is the default locale used by Goment if not set.
const DefaultLocaleCode = "en"

var supportedLocales = map[string]locales.LocaleDetails{
	"en":    locales.EnLocale,
	"es":    locales.EsLocale,
	"fr":    locales.FrLocale,
	"pt-br": locales.PtBRLocale,
	"id":    locales.IdLocale,
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

// MonthByNumber returns the month name by number.
func (g *Goment) MonthByNumber(num int) string {
	if num < 1 || num > 12 {
		return ""
	}
	return g.locale.Months[num-1]
}

// MonthsShort returns the list of abbreviated month names in the current locale.
func (g *Goment) MonthsShort() []string {
	return g.locale.MonthsShort
}

// MonthShortByNumber returns the month short name by number.
func (g *Goment) MonthShortByNumber(num int) string {
	if num < 1 || num > 12 {
		return ""
	}
	return g.locale.MonthsShort[num-1]
}

// Weekdays returns the list of weekdays in the current locale. If the bool parameter is true, the list will be shifted to make the
// locale's first day of the week the first value. If it is not provided or false, Sunday for the locale will be the first value.
func (g *Goment) Weekdays(args ...interface{}) []string {
	shifted := false
	if len(args) == 1 {
		shifted = args[0].(bool)
	}
	return g.locale.GetWeekdays(shifted)
}

// WeekdaysShort returns the list of abbreviated weekday names in the current locale. If the bool parameter is true, the list will be shifted to make the
// locale's first day of the week the first value. If it is not provided or false, Sunday for the locale will be the first value.
func (g *Goment) WeekdaysShort(args ...interface{}) []string {
	shifted := false
	if len(args) == 1 {
		shifted = args[0].(bool)
	}
	return g.locale.GetWeekdaysShort(shifted)
}

// WeekdaysMin returns the list of weekdays in the current locale. If the bool parameter is true, the list will be shifted to make the
// locale's first day of the week the first value. If it is not provided or false, Sunday for the locale will be the first value.
func (g *Goment) WeekdaysMin(args ...interface{}) []string {
	shifted := false
	if len(args) == 1 {
		shifted = args[0].(bool)
	}
	return g.locale.GetWeekdaysMin(shifted)
}

// WeekdayByNumber returns the weekday name by number. If the first parameter is true, the day number parameter will take the locale's
// first day of week into account when returning the day.
func (g *Goment) WeekdayByNumber(args ...interface{}) string {
	count := len(args)
	if count < 1 || count > 2 {
		return ""
	}

	shifted := false
	num := -1

	switch count {
	case 1:
		num = args[0].(int)
	case 2:
		shifted = args[0].(bool)
		num = args[1].(int)
	}

	if num < 0 || num > 6 {
		return ""
	}
	return g.locale.GetWeekdays(shifted)[num]
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
