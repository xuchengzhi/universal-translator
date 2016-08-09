package lo_LA

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type lo_LA struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'lo_LA' locale
func New() locales.Translator {
	return &lo_LA{
		locale:  "lo_LA",
		plurals: []locales.PluralRule{6},
	}
}

// Locale returns the current translators string locale
func (t *lo_LA) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'lo_LA'
func (t *lo_LA) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'lo_LA'
func (t *lo_LA) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	return locales.PluralRuleOther
}
