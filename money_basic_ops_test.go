package money_test

import (
	"flag"
	"github.com/jbowles/money"
	"math/rand"
	"testing"
	"time"
)

var ovf int

func init() {
	flag.IntVar(&ovf, "ovf", 0, "pass 1 to run overflow test")
}

func random(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}

func TestMoneyOverflow(t *testing.T) {
	flag.Parse()
	if ovf == 0 {
		t.Skip("Skip overflow no 'overflow' cli flag set: 'ovf=1 will run it'")
	}

	r := random(1, 25)
	m1 := money.Money{M: money.MaxInt}
	m2 := money.Money{M: money.MinInt}
	if (r % 2) == 1 {
		_ = m1.Add(&m1)
	} else {
		_ = m2.Sub(&m1)
	}

	/*
		** this was testing using fields on the struct to check for overlfow errors
		// overflow error should return value of struct attempting op
		if res.Valuei() != int64(0) {
			t.Error("expected overflow value to be zero", res.Valuei())
		}

		if res.OvfErr != true {
			t.Error("expected ok to be false", res.OvfErr)
		}

		if len(res.Ovf) < 1 {
			t.Error("need overflow message", res.Ovf)
		}
	*/
}

// check addition is good but also that original values are not modified
func TestMoneyAdd(t *testing.T) {
	m1 := money.Money{M: int64(67)}
	m2 := money.Money{M: int64(6700)}
	res := m1.Add(&m2)
	finResi := int64(6767)
	finResf := float64(67.67)

	if res.Valuei() != finResi {
		t.Error("expected '6767' got: ", res.Valuei())
	}

	if res.Valuef() != finResf {
		t.Error("expected '6767' got: ", res.Valuef())
	}

	if m1.Valuei() != int64(67) {
		t.Error("expected '67' got: ", m1.Valuei())
	}

	if m2.Valuei() != int64(6700) {
		t.Error("expected '6700' got: ", m2.Valuei())
	}

	if m1.Valuef() != float64(0.67) {
		t.Error("expected '0.67' got: ", m1.Valuef())
	}

	if m2.Valuef() != float64(67.00) {
		t.Error("expected '67.00' got: ", m2.Valuef())
	}

	if res.StringD() != "67.67" {
		t.Error("expected '67.67' got: ", res.StringD())
	}

	if res.StringC() != "67,67" {
		t.Error("expected '67,67' got: ", res.StringC())
	}
}

// check multiplication is good but also that original values are not modified
func TestMoneyMul(t *testing.T) {
	m1 := money.Money{M: 67}   //67 cents!!
	m2 := money.Money{M: 6700} // 67 dollars
	res := m2.Mul(&m1)
	finResi := int64(4489)    // 44 dollars and 89 cents
	finResf := float64(44.89) // 44 dollars and 89 cents

	if res.Valuei() != finResi {
		t.Error("expected '4489' got: ", res.Valuei())
	}

	if res.Valuef() != finResf {
		t.Error("expected '44.89' got: ", res.Valuef())
	}

	if m1.Valuei() != int64(67) {
		t.Error("expected '67' got: ", m1.Valuei())
	}

	if m2.Valuei() != int64(6700) {
		t.Error("expected '6700' got: ", m2.Valuei())
	}

	if m1.Valuef() != float64(0.67) {
		t.Error("expected '0.67' got: ", m1.Valuef())
	}

	if m2.Valuef() != float64(67.00) {
		t.Error("expected '67.00' got: ", m2.Valuef())
	}

	if res.StringD() != "44.89" {
		t.Error("expected '44.89' got: ", res.StringD())
	}
}

// check subtraction is good but also that original values are not modified
func TestMoneySub(t *testing.T) {
	m1 := money.Money{M: 67}
	m2 := money.Money{M: 6700}
	res := m2.Sub(&m1)
	finResi := int64(6633)
	finResf := float64(66.33)

	if res.Valuei() != finResi {
		t.Error("expected '6633' got: ", res.Valuei())
	}

	if res.Valuef() != finResf {
		t.Error("expected '66.33' got: ", res.Valuef())
	}

	if m1.Valuei() != int64(67) {
		t.Error("expected '67' got: ", m1.Valuei())
	}

	if m2.Valuei() != int64(6700) {
		t.Error("expected '6700' got: ", m2.Valuei())
	}

	if m1.Valuef() != float64(0.67) {
		t.Error("expected '0.67' got: ", m1.Valuef())
	}

	if m2.Valuef() != float64(67.00) {
		t.Error("expected '67.00' got: ", m2.Valuef())
	}

	if res.StringD() != "66.33" {
		t.Error("expected '66.33' got: ", res.StringD())
	}
}

// check division is good but also that original values are not modified
func TestMoneyDiv(t *testing.T) {
	m1 := money.Money{M: 67}   //67 cents!!
	m2 := money.Money{M: 6700} // 67 dollars
	res := m2.Div(&m1)
	finResi := int64(10000)
	finResf := float64(100.00)

	if res.Valuei() != finResi {
		t.Error("expected '10000' got: ", res.Valuei())
	}

	if res.Valuef() != finResf {
		t.Error("expected '100.00' got: ", res.Valuef())
	}

	if m1.Valuei() != int64(67) {
		t.Error("expected '67' got: ", m1.Valuei())
	}

	if m2.Valuei() != int64(6700) {
		t.Error("expected '6700' got: ", m2.Valuei())
	}

	if m1.Valuef() != float64(0.67) {
		t.Error("expected '0.67' got: ", m1.Valuef())
	}

	if m2.Valuef() != float64(67.00) {
		t.Error("expected '67.00' got: ", m2.Valuef())
	}

	if res.StringD() != "100.00" {
		t.Error("expected '100.00' got: ", res.StringD())
	}
}

/////////////////////////////////////
///// More testing on basic ops with negatives, large sums, etc...
/////////////////////////////////////

func TestMoneySmallNumberSubtractLarger(t *testing.T) {
	var val1 int64 = 67   //0.67
	var val2 int64 = 6700 //67.00
	m1 := money.Money{}
	m2 := money.Money{}
	m1Set := m1.Updatei(val1)
	m2Set := m2.Updatei(val2)

	res := m1Set.Sub(m2Set)

	if res.Valuei() != int64(-6633) {
		t.Error("expected negative '-6633' got: ", res.Valuei())
	}

	if res.Valuef() != float64(-66.33) {
		t.Error("expected negative '-66.33' got: ", res.Valuef())
	}
}

// verify subtracting negatives with updates does UPDATE and change original values
// and returns correct results
func TestMoneyNegativeSubUpdate(t *testing.T) {
	var val1Neg int64 = -1   //-0.01 cents
	var val2Pos int64 = 6700 //-67.00 dollars
	m1Neg := money.Money{}
	m2Pos := money.Money{}
	m1SetNeg := m1Neg.Updatei(val1Neg)
	m2SetPos := m2Pos.Updatei(val2Pos)
	resNeg := m1SetNeg.Sub(m2SetPos)

	vali := resNeg.Valuei()
	valf := resNeg.Valuef()

	if vali != int64(-6701) {
		t.Error("expected negative '-6701' got: ", vali)
	}

	if valf != float64(-67.01) {
		t.Error("expected negative '-67.01' got: ", valf)
	}

	// verify return values from update were set correctly
	if m1SetNeg.Valuei() != val1Neg {
		t.Error("expected negative '-1' got: ", m1SetNeg.Valuei())
	}

	if m2SetPos.Valuei() != val2Pos {
		t.Error("expected negative '6700' got: ", m2SetPos.Valuei())
	}

	// verify original values for an update were updated
	if m1Neg.Valuei() != val1Neg {
		t.Error("expected negative '-1' got: ", m1Neg.Valuei())
	}

	if m2Pos.Valuei() != val2Pos {
		t.Error("expected negative '6700' got: ", m2Pos.Valuei())
	}
}

// adding large numbers floats via an `Updatef`
func TestMoneyAddLarge(t *testing.T) {
	var val1 float64 = 12390678659.32 //12390678659.32
	var val2 float64 = 8937670084.36  // 8937670084.36
	m1 := money.Money{}
	m2 := money.Money{}
	m1Set := m1.Updatef(val1)
	m2Set := m2.Updatef(val2)
	res := m1Set.Add(m2Set)

	//12390678659.32 + 8937670084.36 = 21,328,348,743.68
	if res.Valuef() != float64(21328348743.68) {
		t.Error("expected negaitve '21328348743.68' got: ", res.Valuef())
	}

	if res.Valuei() != int64(2132834874368) {
		t.Error("expected negaitve '2132834874368' got: ", res.Valuei())
	}

	if res.StringD() != "21328348743.68" {
		t.Error("expected '21328348743.68' got: ", res.StringD())
	}

	//12390678659.32 + 8937670084.36 = 21,328,348,743.68
	if res.StringC() != "21328348743,68" {
		t.Error("expected '21328348743,68' got: ", res.StringC())
	}
}

// division with a large denominator
func TestMoneyDivLargeDenom(t *testing.T) {
	m1 := money.Money{M: 67}   //67 cents!!
	m2 := money.Money{M: 6700} // 67 dollars
	res := m1.Div(&m2)
	finResi := int64(1)
	finResf := float64(0.01)

	if res.Valuei() != finResi {
		t.Error("expected '1' got: ", res.Valuei())
	}

	if res.Valuef() != finResf {
		t.Error("expected '0.01' got: ", res.Valuef())
	}

	if res.StringD() != "0.01" {
		t.Error("expected '100' got: ", res.StringD())
	}
}

func TestMoneyDivByZero(t *testing.T) {
	m1 := money.Money{}
	m2 := money.Money{M: int64(1)}
	m3 := money.Money{M: int64(72684839824903281)}
	m4 := money.Money{}
	m4.Updatef(float64(478326437489327489327838462381))

	res1 := m2.Div(&m1)
	res2 := m3.Div(&m1)
	res3 := m4.Div(&m1)

	if res1.Valuei() != int64(-9223372036854775807) {
		t.Error("expected '-9223372036854775807' got: ", res1.Valuei())
	}

	if res2.Valuei() != int64(-9223372036854775807) {
		t.Error("expected '-9223372036854775807' got: ", res2.Valuei())
	}

	if res3.Valuei() != int64(9223372036854775807) {
		t.Error("expected '9223372036854775807' got: ", res3.Valuei())
	}

	if res1.Valuef() != float64(-9.223372036854776e+16) {
		t.Error("expected '-9223372036854775807' got: ", res1.Valuef())
	}

	if res2.Valuef() != float64(-9.223372036854776e+16) {
		t.Error("expected '-9223372036854775807' got: ", res2.Valuef())
	}

	if res3.Valuef() != float64(9.223372036854776e+16) {
		t.Error("expected '9223372036854775807' got: ", res3.Valuef())
	}
}

func TestMoneyAbs(t *testing.T) {
	v1 := int64(7839)
	v2 := float64(38261748.09)
	m1 := money.Money{M: -v1}
	m2 := money.Money{}
	m2.Updatef(-v2)

	m1abs := m1.Abs()
	m2abs := m2.Abs()

	m1pos := money.Money{M: v1}
	m2pos := money.Money{}
	m2pos.Updatef(v2)

	if m1abs.Valuei() != m1pos.Valuei() {
		t.Error("expected '7839' got: ", m1abs.Valuei())
	}

	if m2abs.Valuei() != m2pos.Valuei() {
		t.Error("expected '3826174809' got: ", m2abs.Valuei())
	}

	if m1.Abs().Valuei() != m1pos.Valuei() {
		t.Error("expected '7839' got: ", m1.Abs().Valuei())
	}

	if m2.Abs().Valuei() != m2pos.Valuei() {
		t.Error("expected '3826174809' got:", m2.Abs().Valuei())
	}

	//verify original values are still negative
	if m1.Valuei() != -m1pos.Valuei() {
		t.Error("expected '-7839' got: ", m1.Valuei())
	}

	if m2.Valuei() != -m2pos.Valuei() {
		t.Error("expected '-3826174809' got: ", m2.Valuei())
	}
}

func TestMoneyRnd(t *testing.T) {
	r := int64(123500)
	r2 := int64(46783629)
	trunc1 := float64(0.01)
	trunc2 := float64(1.0)
	trunc3 := float64(2000.0)
	trunc4 := float64(0.00000005)
	//trunc3 := ((((money.Guardf * money.DPf) * float64(60)) / float64(600)) / money.Guardf)

	res1 := money.Rnd(r, trunc1)
	if res1 != int64(123500) {
		t.Error("expected '123500' got: ", res1)
	}

	res2 := money.Rnd(r, trunc2)
	if res2 != int64(123501) {
		t.Error("expected '123501' got: ", res2)
	}

	res3 := money.Rnd(r2, trunc3)
	if res3 != int64(46783630) {
		t.Error("expected '46783630' got: ", res3)
	}

	res4 := money.Rnd(r2, trunc4)
	if res4 != int64(46783629) {
		t.Error("expected '46783629' got: ", res4)
	}
}
