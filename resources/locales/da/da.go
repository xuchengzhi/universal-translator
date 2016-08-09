package da

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type da struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'da' locale
func New() locales.Translator {
	return &da{
		locale:  "da",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *da) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'da'
func (t *da) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'da'
func (t *da) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	i := int64(n)
	t := locales.T(n, v)

	if (n == 1) || (t != 0 && (i == 0 || i == 1)) {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
