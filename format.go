package money

import (
	"fmt"
	"github.com/jbowles/i18n"
	"golang.org/x/text/currency"
)

var FormatTranslator, _ = i18n.NewTranslatorFactory(
	[]string{"i18_data/rules"},
	[]string{"i18_data/messages"},
	"en",
)

type MoneyFormat struct {
	MoneyVal *Money
	IsoCode  string
	Symbol   string
	Display  string
}

var currencySymbolFunc = currency.NarrowSymbol
var symbKindCash = currency.Symbol.Kind(currency.Cash)

func (m *Money) Format(isoCode string) (MoneyFormat, error) {
	var mf = MoneyFormat{MoneyVal: m}
	code, err := currency.ParseISO(isoCode)
	if err != nil {
		return mf, err
	}
	mf.IsoCode = code.String()
	// symbol(code) is a currency.Value and has a Format() function but it seems easier to use fmt.
	mf.Symbol = fmt.Sprintf("%v", currencySymbolFunc(code))
	return mf, nil
}

func (m *Money) Formati18Display(isoCode, language string) (MoneyFormat, error) {
	mf, err := m.Format(isoCode)
	if err != nil {
		return mf, err
	}

	tlang, _ := FormatTranslator.GetTranslator(language)
	d, err := tlang.FormatCurrency(mf.MoneyVal.Valuef(), isoCode)
	if err != nil {
		return mf, err
	}

	mf.Display = (d + " " + mf.IsoCode)

	return mf, nil
}
