package pt_GW

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type pt_GW struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'pt_GW' locale
func New() locales.Translator {
	return &pt_GW{
		locale:  "pt_GW",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *pt_GW) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'pt_GW'
func (t *pt_GW) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'pt_GW'
func (t *pt_GW) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n >= 0 && n <= 2 && n != 2 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
