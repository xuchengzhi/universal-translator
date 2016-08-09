package am_ET

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type am_ET struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'am_ET' locale
func New() locales.Translator {
	return &am_ET{
		locale:  "am_ET",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *am_ET) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'am_ET'
func (t *am_ET) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'am_ET'
func (t *am_ET) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	i := int64(n)

	if (i == 0) || (n == 1) {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
