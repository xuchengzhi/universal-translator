package lrc

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type lrc struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'lrc' locale
func New() locales.Translator {
	return &lrc{
		locale:  "lrc",
		plurals: nil,
	}
}

// Locale returns the current translators string locale
func (t *lrc) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'lrc'
func (t *lrc) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'lrc'
func (t *lrc) cardinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown, nil
}
