package money_test

import (
	"github.com/jbowles/money"
	"testing"
)

var valueVar = money.Money{M: 89230467}
var m1 = money.Money{M: int64(6700)}
var m1NegSmall = money.Money{M: int64(-6700)}
var m2NegLarge = money.Money{M: int64(-74893627)}

func BenchmarkPrintEmptyMoneyStruct(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = money.Money{}
	}
}

func BenchmarkPrintfMoney(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = money.Money{M: 74983620}
	}
}

func BenchmarkNegValueForMoney(b *testing.B) {
	for i := 0; i < b.N; i++ {
		val := money.Money{M: 123456}
		val.Neg()
	}
}

func BenchmarkValueInt64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		valueVar.Valuei()
	}
}

func BenchmarkValueFloat64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		valueVar.Valuef()
	}
}

func BenchmarkStringC(b *testing.B) {
	for i := 0; i < b.N; i++ {
		valueVar.StringC()
	}
}

func BenchmarkStringD(b *testing.B) {
	for i := 0; i < b.N; i++ {
		valueVar.StringD()
	}
}

func BenchmarkUpdateInt64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		valueVar.Updatei(int64(467832))
	}
}

func BenchmarkUpdateFloat64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		valueVar.Updatef(float64(467832))
	}
}

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		valueVar.Add(&m1)
	}
}

func BenchmarkAddPosNeg(b *testing.B) {
	for i := 0; i < b.N; i++ {
		valueVar.Add(&m1NegSmall)
	}
}

func BenchmarkAddNegPos(b *testing.B) {
	for i := 0; i < b.N; i++ {
		m1NegSmall.Add(&valueVar)
	}
}

func BenchmarkAddBiggerNegSmallerNeg(b *testing.B) {
	for i := 0; i < b.N; i++ {
		m2NegLarge.Add(&m1NegSmall)
	}
}

func BenchmarkAddSmallerNegBiggerNeg(b *testing.B) {
	for i := 0; i < b.N; i++ {
		m1NegSmall.Add(&m2NegLarge)
	}
}

func BenchmarkSub(b *testing.B) {
	for i := 0; i < b.N; i++ {
		valueVar.Sub(&m1)
	}
}

func BenchmarkSubPosNeg(b *testing.B) {
	for i := 0; i < b.N; i++ {
		valueVar.Sub(&m1NegSmall)
	}
}

func BenchmarkSubNegPos(b *testing.B) {
	for i := 0; i < b.N; i++ {
		m1NegSmall.Sub(&valueVar)
	}
}

func BenchmarkSubBiggerNegSmallerNeg(b *testing.B) {
	for i := 0; i < b.N; i++ {
		m2NegLarge.Sub(&m1NegSmall)
	}
}

func BenchmarkSubSmallerNegBiggerNeg(b *testing.B) {
	for i := 0; i < b.N; i++ {
		m1NegSmall.Sub(&m2NegLarge)
	}
}

func BenchmarkMul(b *testing.B) {
	for i := 0; i < b.N; i++ {
		valueVar.Mul(&m1)
	}
}

func BenchmarkMulPosNeg(b *testing.B) {
	for i := 0; i < b.N; i++ {
		valueVar.Mul(&m1NegSmall)
	}
}

func BenchmarkMulNegPos(b *testing.B) {
	for i := 0; i < b.N; i++ {
		m1NegSmall.Mul(&valueVar)
	}
}

func BenchmarkMulBiggerNegSmallerNeg(b *testing.B) {
	for i := 0; i < b.N; i++ {
		m2NegLarge.Mul(&m1NegSmall)
	}
}

func BenchmarkMulSmallerNegBiggerNeg(b *testing.B) {
	for i := 0; i < b.N; i++ {
		m1NegSmall.Mul(&m2NegLarge)
	}
}

func BenchmarkDiv(b *testing.B) {
	for i := 0; i < b.N; i++ {
		valueVar.Div(&m1)
	}
}

func BenchmarkDivPosNeg(b *testing.B) {
	for i := 0; i < b.N; i++ {
		valueVar.Div(&m1NegSmall)
	}
}

func BenchmarkDivNegPos(b *testing.B) {
	for i := 0; i < b.N; i++ {
		m1NegSmall.Div(&valueVar)
	}
}

func BenchmarkDivBiggerNegSmallerNeg(b *testing.B) {
	for i := 0; i < b.N; i++ {
		m2NegLarge.Div(&m1NegSmall)
	}
}

func BenchmarkDivSmallerNegBiggerNeg(b *testing.B) {
	for i := 0; i < b.N; i++ {
		m1NegSmall.Div(&m2NegLarge)
	}
}
