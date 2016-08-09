package ig_NG

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type ig_NG struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'ig_NG' locale
func New() locales.Translator {
	return &ig_NG{
		locale:  "ig_NG",
		plurals: []locales.PluralRule{6},
	}
}

// Locale returns the current translators string locale
func (t *ig_NG) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'ig_NG'
func (t *ig_NG) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'ig_NG'
func (t *ig_NG) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	return locales.PluralRuleOther
}
