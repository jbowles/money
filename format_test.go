package money_test

import (
	"github.com/jbowles/money"
	"testing"
)

func TestMoneyFormatUSD(t *testing.T) {
	symbolUSD := "$"
	codeUSD := "USD"
	valStrExpect := "1234.56"
	valStr2Expect := "1234,56"
	val := int64(123456)
	val2 := float64(1234.56)

	mon := money.Money{123456}
	symbol, code, m := mon.Format("USD")

	if symbol != symbolUSD {
		t.Error("wanted '$' but got", symbol)
	}
	if code != codeUSD {
		t.Error("wanted 'USD' but got", code)
	}
	if m.StringP() != valStrExpect {
		t.Error("wanted '1234.56' but got", valStrExpect)
	}
	if m.StringC() != valStr2Expect {
		t.Error("wanted '1234,56' but got", valStr2Expect)
	}
	if m.Valuei() != val {
		t.Error("wanted int64 '123456' but got", val)
	}
	if m.Valuef() != val2 {
		t.Error("wanted float64 '123456' but got", val2)
	}
}

func TestMoneyFormatBRL(t *testing.T) {
	symbolBRL := "R$"
	codeBRL := "BRL"

	mon := money.Money{123456}
	symbol, code, _ := mon.Format("BRL")

	if symbol != symbolBRL {
		t.Error("wanted 'R$' but got", symbol)
	}
	if code != codeBRL {
		t.Error("wanted 'BRL' but got", code)
	}
}
