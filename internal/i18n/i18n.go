// Package i18n provides minimal string localisation for aka (EN / ES).
//
// A parity test (TestCatalogParity) enforces that every key present in one
// language also exists in the other, so contributors cannot land a half-
// translated change.
package i18n

import (
	"fmt"
	"os"
	"strings"
)

// Lang is a BCP-47-like language identifier.
type Lang string

const (
	EN Lang = "en"
	ES Lang = "es"
)

var current = EN

// Set forces the active language.
func Set(l Lang) {
	switch l {
	case EN, ES:
		current = l
	default:
		current = EN
	}
}

// Get returns the active language.
func Get() Lang { return current }

// Detect resolves the language from (in order): an explicit flag value
// ("en"/"es"), the config file ("en"/"es"/"auto"), and finally $LANG.
// Unknown values fall back to English.
func Detect(cfgLang, flagLang string) Lang {
	if l := parse(flagLang); l != "" {
		return l
	}
	if l := parse(cfgLang); l != "" {
		return l
	}
	if cfgLang == "auto" || cfgLang == "" {
		if strings.HasPrefix(strings.ToLower(os.Getenv("LANG")), "es") {
			return ES
		}
	}
	return EN
}

func parse(s string) Lang {
	switch strings.ToLower(s) {
	case "en":
		return EN
	case "es":
		return ES
	}
	return ""
}

// T returns the translation for the given key in the currently active
// language. Missing keys return a debug-friendly placeholder so a missing
// translation is obvious during development.
func T(key string) string {
	m := en
	if current == ES {
		m = es
	}
	if v, ok := m[key]; ok {
		return v
	}
	return "!" + key + "!"
}

// Tf is a convenience wrapper for T + fmt.Sprintf.
func Tf(key string, args ...any) string {
	return fmt.Sprintf(T(key), args...)
}
