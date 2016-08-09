package zh_Hans

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type zh_Hans struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'zh_Hans' locale
func New() locales.Translator {
	return &zh_Hans{
		locale:  "zh_Hans",
		plurals: []locales.PluralRule{6},
	}
}

// Locale returns the current translators string locale
func (t *zh_Hans) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'zh_Hans'
func (t *zh_Hans) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'zh_Hans'
func (t *zh_Hans) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	return locales.PluralRuleOther
}
