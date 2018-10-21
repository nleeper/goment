package goment

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"regexp"
	"strings"

	"github.com/pkg/errors"
)

// DefaultLocaleCode is the default locale used by Goment if not set.
const DefaultLocaleCode = "en"

var globalLocale = loadKnownLocale(DefaultLocaleCode)
var cachedLocales = map[string]LocaleDetails{}

// LocaleDetails contains the details of the loaded locale.
type LocaleDetails struct {
	Code      string
	Resources LocaleResources
}

// LocaleResources contains the different locale values.
type LocaleResources struct {
	Months        map[string]int
	ShortMonths   map[string]int
	Days          map[string]int
	ShortDays     map[string]int
	MonthsRegex   *regexp.Regexp
	WeekdaysRegex *regexp.Regexp
}

// Locale gets the current global locale code.
func Locale() string {
	return globalLocale.Code
}

// Locale gets the locale code for the current Goment instance.
func (g *Goment) Locale() string {
	return g.locale.Code
}

// LocaleDetails gets the locale details for the current Goment instance.
func (g *Goment) LocaleDetails() LocaleDetails {
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

func getGlobalLocaleDetails() LocaleDetails {
	return globalLocale
}

func loadKnownLocale(localeCode string) LocaleDetails {
	locale, _ := loadLocale(localeCode)
	return locale
}

func loadLocale(localeCode string) (LocaleDetails, error) {
	var tempResources map[string]string

	normalizedCode := strings.ToLower(localeCode)
	if cached, exist := cachedLocales[normalizedCode]; exist {
		return cached, nil
	}

	if !localeExists(normalizedCode) {
		return LocaleDetails{}, errors.New("Locale " + normalizedCode + "is not supported")
	}

	// TODO - make this more easily testable
	localeJSON, err := ioutil.ReadFile(buildLocalePath(normalizedCode))
	if err != nil {
		return LocaleDetails{}, errors.New("Error reading in locale file")
	}

	err = json.Unmarshal(localeJSON, &tempResources)
	if err != nil {
		return LocaleDetails{}, errors.New("Invalid locale JSON file")
	}

	locale := LocaleDetails{
		Code:      normalizedCode,
		Resources: loadLocaleResources(tempResources),
	}

	cachedLocales[normalizedCode] = locale

	return locale, nil
}

func loadLocaleResources(resources map[string]string) LocaleResources {
	localeResources := LocaleResources{}

	if temp, ok := resources["months"]; ok {
		parsed := strings.Split(strings.ToLower(temp), "_")
		localeResources.Months = loadToMap(parsed, 1)
	}

	if temp, ok := resources["shortMonths"]; ok {
		parsed := strings.Split(strings.ToLower(temp), "_")
		localeResources.ShortMonths = loadToMap(parsed, 1)
	}

	if temp, ok := resources["days"]; ok {
		parsed := strings.Split(strings.ToLower(temp), "_")
		localeResources.Days = loadToMap(parsed, 0)
	}

	if temp, ok := resources["shortDays"]; ok {
		parsed := strings.Split(strings.ToLower(temp), "_")
		localeResources.ShortDays = loadToMap(parsed, 0)
	}

	if temp, ok := resources["monthsRegex"]; ok {
		localeResources.MonthsRegex = regexp.MustCompile(temp)
	}

	if temp, ok := resources["weekdaysRegex"]; ok {
		localeResources.WeekdaysRegex = regexp.MustCompile(temp)
	}

	return localeResources
}

func loadToMap(vals []string, offset int) map[string]int {
	mapValues := make(map[string]int)
	for v := range vals {
		mapValues[vals[v]] = v + offset
	}
	return mapValues
}

func localeExists(localeCode string) bool {
	if _, err := os.Stat(buildLocalePath(localeCode)); os.IsNotExist(err) {
		return false
	}
	return true
}

func buildLocalePath(localeCode string) string {
	return "./locales/" + strings.ToLower(localeCode) + ".json"
}
