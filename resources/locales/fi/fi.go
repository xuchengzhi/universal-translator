package fi

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type fi struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'fi' locale
func New() locales.Translator {
	return &fi{
		locale:  "fi",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *fi) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'fi'
func (t *fi) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'fi'
func (t *fi) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	i := int64(n)

	if i == 1 && v == 0 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
