package mk

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type mk struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'mk' locale
func New() locales.Translator {
	return &mk{
		locale:  "mk",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *mk) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'mk'
func (t *mk) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'mk'
func (t *mk) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	i := int64(n)
	f := locales.F(n, v)

	if (v == 0 && i%10 == 1) || (f%10 == 1) {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
