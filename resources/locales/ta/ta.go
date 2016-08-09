package ta

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type ta struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'ta' locale
func New() locales.Translator {
	return &ta{
		locale:  "ta",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *ta) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'ta'
func (t *ta) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'ta'
func (t *ta) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
