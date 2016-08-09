package ug

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type ug struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'ug' locale
func New() locales.Translator {
	return &ug{
		locale:  "ug",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *ug) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'ug'
func (t *ug) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'ug'
func (t *ug) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
