package hi_IN

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type hi_IN struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'hi_IN' locale
func New() locales.Translator {
	return &hi_IN{
		locale:  "hi_IN",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *hi_IN) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'hi_IN'
func (t *hi_IN) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'hi_IN'
func (t *hi_IN) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	i := int64(n)

	if (i == 0) || (n == 1) {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
