package goment

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"

	"github.com/pkg/errors"
)

// DefaultLocaleCode is the default locale used by Goment if not set.
const DefaultLocaleCode = "en"

var globalLocale = loadKnownLocale(DefaultLocaleCode)
var cachedLocales = map[string]LocaleDetails{}

// LocaleDetails is a struct containing the details of the loaded locale.
type LocaleDetails struct {
	Code      string
	Resources map[string]string
}

// Locale gets the current global locale code.
func Locale() string {
	return globalLocale.Code
}

// Locale gets the locale code for the current Goment instance.
func (g *Goment) Locale() string {
	return g.locale.Code
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
	localeCode = strings.ToLower(localeCode)
	if cached, exist := cachedLocales[localeCode]; exist {
		return cached, nil
	}

	if !localeExists(localeCode) {
		return LocaleDetails{}, errors.New("Locale " + localeCode + "is not supported")
	}

	locale := LocaleDetails{Code: localeCode}

	// TODO - make this more easily testable
	localeJSON, err := ioutil.ReadFile(buildLocalePath(localeCode))
	if err != nil {
		return LocaleDetails{}, errors.New("Error reading in locale file")
	}

	err = json.Unmarshal(localeJSON, &locale.Resources)
	if err != nil {
		return LocaleDetails{}, errors.New("Invalid locale JSON file")
	}

	cachedLocales[localeCode] = locale

	return locale, nil
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
