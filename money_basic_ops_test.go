package money_test

import (
	"fmt"
	"github.com/jbowles/money"
	"testing"
)

func TestMoneyTypePrint(t *testing.T) {
	m := money.Money{}

	printValueDefault := fmt.Sprintf("%v", m)
	if printValueDefault != "{0}" {
		t.Error("default value of money struct should be zero '{0}'", printValueDefault)
	}

	printValued := fmt.Sprintf("%d", m)
	if printValued != "{0}" {
		t.Error("base 10 print format of money struct should be zero '{0}'", printValued)
	}
}

func TestMoneyStructInitialize(t *testing.T) {
	val := money.Money{123456}

	if val.Value() != 123456 {
		t.Error("money struct init Value() should be value '123456'", val.Value())
	}

	if val.String() != "1234.56" {
		t.Error("money struct init String() should be value '1234.56'", val.String())
	}

	// modify value in place
	val.Neg()
	if val.Value() != int64(-123456) {
		t.Error("val.Neg() should be value '-123456'", val.Value())
	}

	if val.String() != "-1234.56" {
		t.Error("val.Neg() as String() should be '-1234.56'", val.String())
	}
}

func TestMoneySetIsMoneyType(t *testing.T) {
	m := money.Money{}
	var val1 int64 = 7868
	m1 := m.Set(val1)
	typ := fmt.Sprintf("%T", m1)
	printValue := fmt.Sprintf("%v", m1)
	if typ != "*money.Money" {
		t.Error("should be money type, instead it is: ", typ)
	}

	if printValue != "78.68" {
		t.Error("default print value for money Set() on int64 '7868' should be '78.68', but got", printValue)
	}

	printValued := fmt.Sprintf("%d", m1)
	if printValued != "&{7868}" {
		t.Error("base 10 print format of money struct should be base value '&{7868}'", printValued)
	}
}

func TestMoneyString(t *testing.T) {
	m := money.Money{}
	var val1 int64 = 67
	m1 := m.Set(val1)
	if m1.String() != "0.67" {
		t.Error("wanted to see '0.67' cents but got: ", m1.String())
	}

	var val2 int64 = 6700
	m2 := m.Set(val2)
	if m2.String() != "67.00" {
		t.Error("wanted to see '67.00' dollars but got: ", m2.String())
	}
}

func TestMoneyVarsNotChanging(t *testing.T) {
	var val1 int64 = 67
	var val2 int64 = 6700
	m1 := money.Money{}
	m2 := money.Money{}
	m1Set := m1.Set(val1)
	m2Set := m2.Set(val2)
	if m1Set.String() != "0.67" {
		t.Error("expected '0.67' got: ", m1Set.String())
	}

	if m2Set.String() != "67.00" {
		t.Error("expected '67.00' got: ", m2Set.String())
	}
}

func TestMoneyAdd(t *testing.T) {
	var val1 int64 = 67   //0.67
	var val2 int64 = 6700 //67.00
	m1 := money.Money{}
	m2 := money.Money{}
	m1Set := m1.Set(val1)
	m2Set := m2.Set(val2)
	res := m1Set.Add(m2Set)

	if res.String() != "67.67" {
		t.Error("expected '67.67' got: ", res.String())
	}
}

func TestMoneySub(t *testing.T) {
	var val1 int64 = 67   //0.67
	var val2 int64 = 6700 //67.00
	m1 := money.Money{}
	m2 := money.Money{}
	m1Set := m1.Set(val1)
	m2Set := m2.Set(val2)
	res := m2Set.Sub(m1Set)

	if res.String() != "66.33" {
		t.Error("expected '66.33' got: ", res.String())
	}
}

func TestMoneySmallNumberSubtractLarger(t *testing.T) {
	var val1 int64 = 67   //0.67
	var val2 int64 = 6700 //67.00
	m1 := money.Money{}
	m2 := money.Money{}
	m1Set := m1.Set(val1)
	m2Set := m2.Set(val2)

	badRes := m1Set.Sub(m2Set)
	if badRes.Value() != int64(-6633) {
		t.Error("expected negative '-6633' got: ", badRes.Value())
	}
}

func TestMoneyNegativeSubValue(t *testing.T) {
	var val1Neg int64 = -1   //-0.1
	var val2Neg int64 = 6700 //67.00
	m1Neg := money.Money{}
	m2Neg := money.Money{}
	m1SetNeg := m1Neg.Set(val1Neg)
	m2SetNeg := m2Neg.Set(val2Neg)
	resNeg := m1SetNeg.Sub(m2SetNeg)

	val := resNeg.Value()

	if val != int64(-6701) {
		t.Error("expected negaitve '-6701' got: ", val)
	}
}

func TestMoneyAddLargeSum(t *testing.T) {
	var val1 int64 = 1239067865932 //12390678659.32
	var val2 int64 = 893767008436  // 8937670084.36
	m1 := money.Money{}
	m2 := money.Money{}
	m1Set := m1.Set(val1)
	m2Set := m2.Set(val2)
	res := m1Set.Add(m2Set)

	//12390678659.32 + 8937670084.36 = 21,328,348,743.68
	if res.String() != "21328348743.68" {
		t.Error("expected '21328348743.68' got: ", res.String())
	}
}
