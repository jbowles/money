package money

import (
	"fmt"
	"github.com/jbowles/money/Godeps/_workspace/src/golang.org/x/text/currency"
	//"vendor/currency"
)

type MoneyFormat struct {
	MoneyVal *Money
	IsoCode  string
	Symbol   string
	Display  string
}

var currencySymbolFunc = currency.NarrowSymbol

// Format returns a money format struct with the money int64 value, as well as curency code and symbol
// NOTE that `symbol(code)` is a `currency.Value` and has a `Format()` function but it seems easier to use fmt.
func (m *Money) Format(isoCode string) (MoneyFormat, error) {
	var mf = MoneyFormat{MoneyVal: m}
	code, err := currency.ParseISO(isoCode)
	if err != nil {
		return mf, err
	}
	mf.IsoCode = code.String()
	mf.Symbol = fmt.Sprintf("%v", currencySymbolFunc(code))
	return mf, nil
}
