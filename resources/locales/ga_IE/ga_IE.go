package ga_IE

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type ga_IE struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'ga_IE' locale
func New() locales.Translator {
	return &ga_IE{
		locale:  "ga_IE",
		plurals: []locales.PluralRule{2, 3, 4, 5, 6},
	}
}

// Locale returns the current translators string locale
func (t *ga_IE) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'ga_IE'
func (t *ga_IE) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'ga_IE'
func (t *ga_IE) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	} else if n == 2 {
		return locales.PluralRuleTwo
	} else if n >= 3 && n <= 6 {
		return locales.PluralRuleFew
	} else if n >= 7 && n <= 10 {
		return locales.PluralRuleMany
	}

	return locales.PluralRuleOther
}
