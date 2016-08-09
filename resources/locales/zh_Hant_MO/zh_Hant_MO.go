package zh_Hant_MO

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type zh_Hant_MO struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'zh_Hant_MO' locale
func New() locales.Translator {
	return &zh_Hant_MO{
		locale:  "zh_Hant_MO",
		plurals: []locales.PluralRule{6},
	}
}

// Locale returns the current translators string locale
func (t *zh_Hant_MO) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'zh_Hant_MO'
func (t *zh_Hant_MO) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'zh_Hant_MO'
func (t *zh_Hant_MO) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	return locales.PluralRuleOther
}
