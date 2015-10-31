package money_test

import (
	"github.com/jbowles/money"
	"testing"
)

func TestMoneyFormat(t *testing.T) {
	symbolUSD := "$"
	codeUSD := "USD"
	valExpect := "1234.56"
	m := money.Money{123456}
	symbol, code, val := m.Format("USD")

	if symbol != symbolUSD {
		t.Error("wanted '$' but got", symbol)
	}
	if code != codeUSD {
		t.Error("wanted 'USD' but got", code)
	}
	if val != valExpect {
		t.Error("wanted '1234.56' but got", valExpect)
	}
}
