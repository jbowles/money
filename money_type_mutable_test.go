package money_test

import (
	"fmt"
	"github.com/jbowles/money"
	"testing"
)

func TestMoneyTypePrint(t *testing.T) {
	m := money.Money{}

	printValueDefault := fmt.Sprintf("%v", m)
	//if printValueDefault != "{0 <nil>}" {
	if printValueDefault != "{0}" {
		//t.Error("default value of money struct should be zero '{0 <nil>}'", printValueDefault)
		t.Error("default value of money struct should be zero '{0}'", printValueDefault)
	}

	printValued := fmt.Sprintf("%d", m)
	//if printValued != "{0 0}" {
	if printValued != "{0}" {
		//t.Error("base 10 print format of money struct should be zero '{0 0}'", printValued)
		t.Error("base 10 print format of money struct should be zero '{0}'", printValued)
	}
}

func TestMoneyValueIntAndFloat(t *testing.T) {
	val := money.Money{M: 123456}

	if val.Valuei() != int64(123456) {
		t.Error("Valuei() should be int64 '123456'", val.Valuei())
	}

	if val.Valuef() != float64(1234.56) {
		t.Error("Valuef() should be float64 '123456'", val.Valuef())
	}

	if val.ValueiTrunc() != int64(1234) {
		t.Error("ValueiTrunc() should be int64 '1234'", val.ValueiTrunc())
	}

	if val.StringD() != "1234.56" {
		t.Error("money struct init StringD() should be value '1234.56'", val.StringD())
	}

	if val.StringC() != "1234,56" {
		t.Error("money struct init StringC() should be value '1234,56'", val.StringC())
	}
}

// Neg value make sure 'val' is not updated
func TestMoneyNegNotMutable(t *testing.T) {
	val := money.Money{M: 123456}
	neg := val.Neg()

	if val.Valuei() != int64(123456) {
		t.Error("val should be int64 '123456'", val.Valuei())
	}

	if val.Valuef() != float64(1234.56) {
		t.Error("val should be int64 '1234.56'", val.Valuef())
	}

	if neg.Valuei() != int64(-123456) {
		t.Error("neg.Valuei should be int64 '-123456'", neg.Valuei())
	}

	if neg.Valuef() != float64(-1234.56) {
		t.Error("neg.Valuef should be float64 '-1234.56'", neg.Valuef())
	}

	if val.StringD() != "1234.56" {
		t.Error("val.StringD() should be '1234.56'", val.StringD())
	}
	// TODO oddity for negative money values!!!!
	if neg.StringD() != "-1234.56" {
		t.Error("neg.StringD() should be '-1234.56'", neg.StringD())
	}
}

func TestMoneyUpdateIsMoneyType(t *testing.T) {
	m := money.Money{}
	var val1 int64 = 7868
	res := m.Updatei(val1)

	if res.Valuei() != val1 {
		t.Error("should be '7868'", res.Valuei())
	}

	typ := fmt.Sprintf("%T", m)
	typPointer := fmt.Sprintf("%T", res)

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
	//if printValued != "{7868 0}" {
	if printValued != "{7868}" {
		//t.Error("base 10 print format of money struct should be base value '{7868 0}'", printValued)
		t.Error("base 10 print format of money struct should be base value '{7868}'", printValued)
	}
}

func TestMoneyStringWithUpdate(t *testing.T) {
	m := money.Money{M: 67}
	if m.StringD() != "0.67" {
		t.Error("wanted to see '0.67' cents but got: ", m.StringD())
	}

	var val2 int64 = 6700
	m.Updatei(val2)
	if m.StringD() != "67.00" {
		t.Error("wanted to see '67.00' dollars but got: ", m.StringD())
	}
}

func TestMoneyUpdateChangesOriginalValue(t *testing.T) {
	m := money.Money{M: 67}
	m.Updatei(int64(6700))
	if m.Valuei() != int64(6700) {
		t.Error("original vlaue should be updated to '67.00' but got: ", m.Valuei())
	}
}

func TestMoneyVarsNotChanging(t *testing.T) {
	var val1 int64 = 67
	var val2 int64 = 6700
	m1 := money.Money{M: val1}
	m2 := money.Money{M: val2}
	if m1.Valuei() != val1 {
		t.Error("expected '67' got: ", m1.Valuei())
	}

	if m2.Valuei() != val2 {
		t.Error("expected '6700' got: ", m2.Valuei())
	}
}
