package bs

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type bs struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'bs' locale
func New() locales.Translator {
	return &bs{
		locale:  "bs",
		plurals: []locales.PluralRule{2, 4, 6},
	}
}

// Locale returns the current translators string locale
func (t *bs) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'bs'
func (t *bs) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'bs'
func (t *bs) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

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
