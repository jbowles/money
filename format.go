package money

import (
	"fmt"
	"golang.org/x/text/currency"
)

func (m *Money) Format(isoCode string) (string, string, string) {
	code, _ := currency.ParseISO(isoCode)
	var symbol = currency.NarrowSymbol
	// symbol(code) is a currency.Value and has a Format() function but it seems easier to use fmt.
	symString := fmt.Sprintf("%v", symbol(code))
	return symString, code.String(), m.StringP()
}

/*
  	m := money.Money{123456}
	FormatFloat64(m.Valuef(), "DKK")
func FormatFloat64(m float64, isoCode string) {
	c, _ := currency.ParseISO(isoCode)
	var nsymbol = currency.NarrowSymbol
	var symbol = currency.Symbol
	var iso = currency.ISO
	var symbKindCash = currency.Symbol.Kind(currency.Cash)
	fmt.Printf("c.String() with c: %v, c.Value(m) with c: %v\n", c.String(), c.Value(m))

	val := c.Value(m)
	fmt.Printf("narrow symbol formatter with v: %v\n", nsymbol(val))
	fmt.Printf("symbol formatter with v: %v\n", symbol(val))
	fmt.Printf("iso formatter with v: %v\n", iso(val))
	fmt.Printf("symbol kind cash formatter with value: %v\n", symbKindCash(val))
}
*/
