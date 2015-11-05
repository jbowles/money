package money_test

import (
	"github.com/jbowles/money"
	"testing"
)

func BenchmarkFormatUSD(b *testing.B) {
	mon := money.Money{123456}
	for i := 0; i < b.N; i++ {
		_, _ = mon.Format("USD")
	}
}

func BenchmarkFormatI18USD(b *testing.B) {
	mon := money.Money{123456}
	for i := 0; i < b.N; i++ {
		_, _ = mon.Formati18Display("USD", "en")
	}
}

func TestMoneyFormatUSD(t *testing.T) {
	symbolUSD := "$"
	codeUSD := "USD"
	valStrExpect := "1234.56"
	valStr2Expect := "1234,56"
	//valStr3Expect := "$1,234.56 USD"
	//valStr3Expect := "US$ 1234.56"
	val := int64(123456)
	val2 := float64(1234.56)

	mon := money.Money{123456}
	mf, _ := mon.Format("USD")

	if mf.Symbol != symbolUSD {
		t.Error("wanted '$' but got", mf.Symbol)
	}
	if mf.IsoCode != codeUSD {
		t.Error("wanted 'USD' but got", mf.IsoCode)
	}
	if mf.MoneyVal.StringP() != valStrExpect {
		t.Error("wanted '1234.56' but got", mf.MoneyVal.StringP())
	}
	if mf.MoneyVal.StringC() != valStr2Expect {
		t.Error("wanted '1234,56' but got", mf.MoneyVal.StringC())
	}
	if mf.MoneyVal.Valuei() != val {
		t.Error("wanted int64 '123456' but got", mf.MoneyVal.Valuei())
	}
	if mf.MoneyVal.Valuef() != val2 {
		t.Error("wanted float64 '123456' but got", mf.MoneyVal.Valuef())
	}

	/*
		if mf.Display != valStr3Expect {
			t.Error("wanted '1234.56' but got", mf.Display)
		}
	*/
}

func TestMoneyFormatI18USD(t *testing.T) {
	symbolUSD := "$"
	codeUSD := "USD"
	valStrExpect := "1234.56"
	valStr2Expect := "1234,56"
	//valStr3Expect := "$1,234.56 USD"
	val := int64(123456)
	val2 := float64(1234.56)

	mon := money.Money{123456}
	mf, _ := mon.Formati18Display("USD", "en")

	if mf.Symbol != symbolUSD {
		t.Error("wanted '$' but got", mf.Symbol)
	}
	if mf.IsoCode != codeUSD {
		t.Error("wanted 'USD' but got", mf.IsoCode)
	}
	if mf.MoneyVal.StringP() != valStrExpect {
		t.Error("wanted '1234.56' but got", mf.MoneyVal.StringP())
	}
	if mf.MoneyVal.StringC() != valStr2Expect {
		t.Error("wanted '1234,56' but got", mf.MoneyVal.StringC())
	}
	if mf.MoneyVal.Valuei() != val {
		t.Error("wanted int64 '123456' but got", mf.MoneyVal.Valuei())
	}
	if mf.MoneyVal.Valuef() != val2 {
		t.Error("wanted float64 '123456' but got", mf.MoneyVal.Valuef())
	}

	/*
		if mf.Display != valStr3Expect {
			t.Error("wanted '1234.56' but got", mf.Display)
		}
	*/
}

func TestMoneyFormatI18BRL(t *testing.T) {
	symbolBRL := "R$"
	codeBRL := "BRL"
	//valStrExpect := "R$1.234,56 BRL"

	mon := money.Money{123456}
	mf, _ := mon.Formati18Display("BRL", "pt-br") // "pt-BR" also works

	if mf.Symbol != symbolBRL {
		t.Error("wanted 'R$' but got", mf.Symbol)
	}
	if mf.IsoCode != codeBRL {
		t.Error("wanted 'BRL' but got", mf.IsoCode)
	}
	/*
		if mf.Display != valStrExpect {
			t.Error("wanted '1234.56' but got", mf.Display)
		}
	*/
}
