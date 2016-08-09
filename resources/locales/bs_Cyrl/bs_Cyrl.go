package bs_Cyrl

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type bs_Cyrl struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'bs_Cyrl' locale
func New() locales.Translator {
	return &bs_Cyrl{
		locale:  "bs_Cyrl",
		plurals: []locales.PluralRule{2, 4, 6},
	}
}

// Locale returns the current translators string locale
func (t *bs_Cyrl) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'bs_Cyrl'
func (t *bs_Cyrl) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'bs_Cyrl'
func (t *bs_Cyrl) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	i := int64(n)
	f := locales.F(n, v)

	if (v == 0 && i%10 == 1 && i%100 != 11) || (f%10 == 1 && f%100 != 11) {
		return locales.PluralRuleOne
	} else if (v == 0 && i%10 >= 2 && i%10 <= 4 && i%100 < 12 && i%100 > 14) || (f%10 >= 2 && f%10 <= 4 && f%100 < 12 && f%100 > 14) {
		return locales.PluralRuleFew
	}

	return locales.PluralRuleOther
}
