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
