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

func TestMoneyValueIntAndFloat(t *testing.T) {
	val := money.Money{123456}

	if val.Valuei() != int64(123456) {
		t.Error("Valuei() should be int64 '123456'", val.Valuei())
	}

	if val.Valuef() != float64(1234.56) {
		t.Error("Valuef() should be float64 '123456'", val.Valuef())
	}

	if val.StringP() != "1234.56" {
		t.Error("money struct init StringP() should be value '1234.56'", val.StringP())
	}

	if val.StringC() != "1234,56" {
		t.Error("money struct init StringC() should be value '1234,56'", val.StringC())
	}
}

func TestMoneyNegNotMutable(t *testing.T) {
	val := money.Money{123456}
	// Neg value make sure 'val' is not updated
	neg := val.Neg()

	if val.Valuei() != int64(123456) {
		t.Error("val should be int64 '123456'", val.Valuei())
	}

	if val.Valuef() != float64(1234.56) {
		t.Error("val should be int64 '1234.56'", val.Valuef())
	}

	if val.StringP() != "1234.56" {
		t.Error("val.StringP() should be '1234.56'", val.StringP())
	}

	if neg.Valuei() != int64(-123456) {
		t.Error("neg.Valuei should be int64 '-123456'", neg.Valuei())
	}

	if neg.Valuef() != float64(-1234.56) {
		t.Error("neg.Valuef should be float64 '-1234.56'", neg.Valuef())
	}

	// TODO oddity for negative money values!!!!
	if neg.StringP() != "-1234.-56" {
		t.Error("neg.StringP() should be '-1234.-56'", neg.StringP())
	}
}

func TestMoneyUpdateIsMoneyType(t *testing.T) {
	m := money.Money{}
	var val1 int64 = 7868
	pointerVal := m.Updatei(val1)

	typ := fmt.Sprintf("%T", m)
	typPointer := fmt.Sprintf("%T", pointerVal)

	printValue := fmt.Sprintf("%v", m.M)

	if typ != "money.Money" {
		t.Error("should be money type, instead it is: ", typ)
	}

	if typPointer != "*money.Money" {
		t.Error("should be money type, instead it is: ", typPointer)
	}

	if printValue != "7868" {
		t.Error("default print value for money Updatei() int64 '7868' should be '7868', but got", printValue)
	}

	printValued := fmt.Sprintf("%d", m)
	if printValued != "{7868}" {
		t.Error("base 10 print format of money struct should be base value '{7868}'", printValued)
	}
}

func TestMoneyStringWithUpdate(t *testing.T) {
	m := money.Money{67}
	if m.StringP() != "0.67" {
		t.Error("wanted to see '0.67' cents but got: ", m.StringP())
	}

	var val2 int64 = 6700
	m.Updatei(val2)
	if m.StringP() != "67.00" {
		t.Error("wanted to see '67.00' dollars but got: ", m.StringP())
	}
}

func TestMoneyVarsNotChanging(t *testing.T) {
	var val1 int64 = 67
	var val2 int64 = 6700
	m1 := money.Money{val1}
	m2 := money.Money{val2}
	if m1.StringP() != "0.67" {
		t.Error("expected '0.67' got: ", m1.StringP())
	}

	if m2.StringP() != "67.00" {
		t.Error("expected '67.00' got: ", m2.StringP())
	}

	if m1.StringC() != "0,67" {
		t.Error("expected '0,67' got: ", m1.StringC())
	}

	if m2.StringC() != "67,00" {
		t.Error("expected '67,00' got: ", m2.StringC())
	}
}

func TestMoneyAdd(t *testing.T) {
	m1 := money.Money{int64(67)}
	m2 := money.Money{int64(6700)}
	res := m1.Add(&m2)

	if res.StringP() != "67.67" {
		t.Error("expected '67.67' got: ", res.StringP())
	}

	if res.StringC() != "67,67" {
		t.Error("expected '67,67' got: ", res.StringC())
	}
}

func TestMoneyAddDoesNotMutate(t *testing.T) {
	m1 := money.Money{}
	m2 := money.Money{}

	if m1.StringP() != "0.00" {
		t.Error("expected '0.00' got: ", m1.StringP())
	}

	if m2.StringP() != "0.00" {
		t.Error("expected '0.00' got: ", m2.StringP())
	}

	var val1 int64 = 67   //0.67
	var val2 int64 = 6700 //67.00
	m1Set := m1.Updatei(val1)
	m2Set := m2.Updatei(val2)
	res := m1Set.Add(m2Set)

	if res.StringP() != "67.67" {
		t.Error("expected '67.67' got: ", res.StringP())
	}

	if m1.StringP() != "0.67" {
		t.Error("expected '0.67' got: ", m1.StringP())
	}

	if m1Set.StringP() != "0.67" {
		t.Error("expected '0.67' got: ", m1Set.StringP())
	}
}

func TestMoneySub(t *testing.T) {
	m1 := money.Money{67}
	m2 := money.Money{6700}
	res := m2.Sub(&m1)

	if res.StringP() != "66.33" {
		t.Error("expected '66.33' got: ", res.StringP())
	}
}

func TestMoneySmallNumberSubtractLarger(t *testing.T) {
	var val1 int64 = 67   //0.67
	var val2 int64 = 6700 //67.00
	m1 := money.Money{}
	m2 := money.Money{}
	m1Set := m1.Updatei(val1)
	m2Set := m2.Updatei(val2)

	badRes := m1Set.Sub(m2Set)
	if badRes.Valuei() != int64(-6633) {
		t.Error("expected negative '-6633' got: ", badRes.Valuei())
	}
}

func TestMoneyNegativeSubValue(t *testing.T) {
	var val1Neg int64 = -1   //-0.1
	var val2Neg int64 = 6700 //67.00
	m1Neg := money.Money{}
	m2Neg := money.Money{}
	m1SetNeg := m1Neg.Updatei(val1Neg)
	m2SetNeg := m2Neg.Updatei(val2Neg)
	resNeg := m1SetNeg.Sub(m2SetNeg)

	val := resNeg.Valuei()

	if val != int64(-6701) {
		t.Error("expected negaitve '-6701' got: ", val)
	}
}

func TestMoneyAddLargeSumUpdateFloat(t *testing.T) {
	var val1 float64 = 12390678659.32 //12390678659.32
	var val2 float64 = 8937670084.36  // 8937670084.36
	m1 := money.Money{}
	m2 := money.Money{}
	m1Set := m1.Updatef(val1)
	m2Set := m2.Updatef(val2)
	res := m1Set.Add(m2Set)

	//12390678659.32 + 8937670084.36 = 21,328,348,743.68
	if res.StringP() != "21328348743.68" {
		t.Error("expected '21328348743.68' got: ", res.StringP())
	}

	//12390678659.32 + 8937670084.36 = 21,328,348,743.68
	if res.StringC() != "21328348743,68" {
		t.Error("expected '21328348743,68' got: ", res.StringC())
	}
}
