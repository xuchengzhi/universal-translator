package dav_KE

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type dav_KE struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'dav_KE' locale
func New() locales.Translator {
	return &dav_KE{
		locale:  "dav_KE",
		plurals: nil,
	}
}

// Locale returns the current translators string locale
func (t *dav_KE) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'dav_KE'
func (t *dav_KE) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'dav_KE'
func (t *dav_KE) cardinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown, nil
}
