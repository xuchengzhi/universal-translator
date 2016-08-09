package as_IN

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type as_IN struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'as_IN' locale
func New() locales.Translator {
	return &as_IN{
		locale:  "as_IN",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *as_IN) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'as_IN'
func (t *as_IN) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'as_IN'
func (t *as_IN) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	i := int64(n)

	if (i == 0) || (n == 1) {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
