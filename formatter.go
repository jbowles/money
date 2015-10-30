package money

import (
	"fmt"
	"golang.org/x/text/currency"
	//"golang.org/x/text/language"
)

/*
// see language Tag, https://godoc.org/golang.org/x/text/language#Tag
var (
	en     = language.English
	en_GB  = language.BritishEnglish
	en_US  = language.AmericanEnglish
	en_AU  = language.MustParse("en-AU")
	es     = language.Spanish
	es_ES  = language.EuropeanSpanish
	es_419 = language.LatinAmericanSpanish
	fr     = language.French
	fr_CA  = language.CanadianFrench
	und    = language.Und //undefined
)
*/

func (m *Money) FormatA(isoCode string) (string, string) {
	code, _ := currency.ParseISO(isoCode)
	var symbol = currency.NarrowSymbol
	// symbol(code) is a currency.Value and has a Format() function but it seems easier to use fmt.
	symString := fmt.Sprintf("%v", symbol(code))
	return symString, code.String()
}
