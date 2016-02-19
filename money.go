// Package money implements the Money type, which uses a fixed-length guard for precision arithmetic: the int64 variable Guard (and its float64 and int-related variables Guardf and Guardi).
// DP is the decimal precision, which can be changed in the DecimalPrecision() function.  DP hold the places after the decimal place in the active money struct field M.
//
// Money also uses the text/currency package to format money types for currency codes and symbols.
//
// Money is heavily modified from the original project (see the external directory of this project for a copy of the original code):
//	https://github.com/Confunctionist/finance
// See the license here:
//	https://github.com/Confunctionist/finance/blob/master/LICENSE
//
package money

import (
	"fmt"
)

type Overflow struct {
	Fn func(int64, int64) string
}

var ovf = map[string]Overflow{
	"add": {func(x, y int64) string { return fmt.Sprintf("%s for %d Add() %d", OVFL, x, y) }},
	"sub": {func(x, y int64) string { return fmt.Sprintf("%s for %d Sub() %d", OVFL, x, y) }},
}

type Money struct {
	M int64 // value of the integer64 Money
	//Ovf    string //overflow message
	//OvfErr bool
}

var (
	DP    int64   = 100         // for default of 2 decimal places => 10^2 (can be reset)
	DPf   float64 = float64(DP) // for default of 2 decimal places => 10^2 (can be reset)
	Round         = float64(0.5)
	//Round  = .5 + (1 / Guardf)
	Roundn = Round * float64(-1)
)

const (
	Guardi int     = 100
	Guard  int64   = int64(Guardi)
	Guardf float64 = float64(Guardi)
	//instead of using math package `math.MaxInt64` just define ourselves
	MaxInt = int64(^uint(0) >> 1) // max largest int 9223372036854775807
	MinInt = (-MaxInt - 1)        // max largest int -9223372036854775808
	OVFL   = "Overflow"
	//DTL    = "Decimal places too large"
	//DLZ    = "Decimal places cannot be less than zero"
	//NOOR   = "Number out of range"
	//MAXDEC = 18
)

//////////////////////////////////////////////////////////////
///////// GET AND UPDATE MONEY TYPE ///////////////////////////
//////////////////////////////////////////////////////////////

// Valuei returns int64 value of Money
func (m *Money) Valuei() int64 {
	return m.M
}

// Valuef gets the float64 value of money (see Value() for int64)
func (m *Money) Valuef() float64 {
	return float64(m.M) / DPf
}

// ValueiTrunc gets value of money truncating after DP (see Value() for no truncation)
func (m *Money) ValueiTrunc() int64 {
	return m.M / DP
}

// Set sets the Money field M, this is destructive
func (m *Money) Updatei(x int64) *Money {
	m.M = x
	return m
}

// Setf sets a float64 into a Money type for precision calculations
func (m *Money) Updatef(f float64) *Money {
	fDPf := f * DPf
	r := int64(f * DPf)
	return m.Updatei(Rnd(r, fDPf-float64(r)))
}

//////////////////////////////////////////////////////////////
///////// BASIC OPERATIONS '+', '-', '*', '/'  ///////////////////////////
//////////////////////////////////////////////////////////////

// Add Adds two Money types
func (m *Money) Add(n *Money) *Money {
	r := m.M + n.M
	if (r^m.M)&(r^n.M) < int64(0) {
		f := ovf["add"]
		/*
			return &Money{
				M:      int64(0),
				OvfErr: true,
				Ovf:    f.Fn(m.M, n.M),
			}
		*/
		panic(f.Fn(m.M, n.M))
	}
	return &Money{M: r}
}

// Sub subtracts one Money type from another
func (m *Money) Sub(n *Money) *Money {
	r := m.M - n.M
	if (r^m.M)&^(r^n.M) < int64(0) {
		f := ovf["sub"]
		panic(f.Fn(m.M, n.M))
	}
	return &Money{M: r}
}

// Mul Multiplies two Money types
func (m *Money) Mul(n *Money) *Money {
	return &Money{M: ((m.M * n.M) / DP)}
}

// Div Divides one Money type from another
// Division by zero will return:
//  int64 = -9223372036854775807
//  float64 = 9223372036854775807
func (m *Money) Div(n *Money) *Money {
	f := ((((Guardf * DPf) * float64(m.M)) / float64(n.M)) / Guardf)
	i := int64(f)
	return &Money{M: Rnd(i, f-float64(i))}
}

//////////////////////////////////////////////////////////////
///////////////////////// FORMATTING, ROUNDING ///////////////////////////////////
//////////////////////////////////////////////////////////////

// StringP for money type representation in basic monetary unit (DOLLARS POINT CENTS)
func (m *Money) StringD() string {
	return fmt.Sprintf("%d.%02d", m.Valuei()/DP, m.Abs().Valuei()%DP)
}

// StringC for money type representation in basic monetary unit (DOLLARS COMMA CENTS)
func (m *Money) StringC() string {
	return fmt.Sprintf("%d,%02d", m.Valuei()/DP, m.Abs().Valuei()%DP)
}

// Neg Returns the negative value of Money
func (m *Money) Neg() *Money {
	r := m.M
	if m.M != 0 {
		r *= -1
	}
	return &Money{M: r}
}

// Abs Returns the absolute value of Money
func (m *Money) Abs() *Money {
	if m.M < int64(0) {
		return m.Neg()
	}
	return m
}

// RND rounds int64 remainder rounded half towards plus infinity
// trunc = the remainder of the float64 calc
// r     = the result of the int64 cal
func Rnd(r int64, trunc float64) int64 {
	//fmt.Printf("RND 1 r = % v, trunc = %v Round = %v\n", r, trunc, Round)
	if trunc > 0 {
		if trunc >= Round {
			r++
		}
	} else {
		if trunc < Roundn {
			r--
		}
	}
	//fmt.Printf("RND 2 r = % v, trunc = %v Round = %v\n", r, trunc, Round)
	return r
}

/*
** these functions are not supported
//////////////////////////////////////////////////////////////
///////// NOT SO BASIC FUNCTIONS ///////////////////////////
//////////////////////////////////////////////////////////////

// Sign returns the Sign of Money 1 if positive, -1 if negative
func (m *Money) Sign() int {
	if m.M < 0 {
		return -1
	}
	return 1
}

// DecimalChange resets the package-wide decimal place (default is 2 decimal places)
func DecimalChange(d int) {
	if d < 0 {
		panic(DLZ)
	}
	if d > MAXDEC {
		panic(DTL)
	}
	var newDecimal int
	if d > 0 {
		newDecimal++
		for i := 0; i < d; i++ {
			newDecimal *= 10
		}
	}
	DPf = float64(newDecimal)
	DP = int64(newDecimal)
	return
}

// SDs Standard Deviation of a sample
// sd = sqrt(SIGMA ((a[i] - mean) ^ 2) / (len(a)-1))
// SIGMA a total of all of the elements of a
// a[i] is the ith elemant of a
// len(a) = the number of elements in the slice a adjusted for sample
func SDs(a []float64) float64 {
	var sum float64
	m := Mean(a)
	for _, v := range a {
		sum += math.Pow(v-m, 2)
	}
	return math.Sqrt(sum / float64(len(a)-1))
}

// SD Standard Deviation
// sd = sqrt(SIGMA ((a[i] - mean) ^ 2) / len(a))
// SIGMA a total of all of the elements of a
// a[i] is the ith elemant of a
// len(a) = the number of elements in the slice a
func SD(a []float64) float64 {
	var sum float64
	m := Mean(a)
	for _, v := range a {
		sum += math.Pow(v-m, 2)
	}
	return math.Sqrt(sum / float64(len(a)))
}

// Mulf Multiplies a Money with a float to return a money-stored type
func (m *Money) Mulf(f float64) *Money {
	i := m.M * int64(f*Guardf*DPf)
	r := i / Guard / DP
	return &Money{Rnd(r, float64(i)/Guardf/DPf-float64(r))}
}

// Mean Average
// mean = SIGMA a / len(a)
// SIGMA a total of all of the elements of a
// len(a) = the number of values
func Mean(a []float64) float64 {
	lenA := float64(len(a))
	if lenA == 0 {
		panic(NOOR)
	}
	var sum float64
	for _, v := range a {
		sum += v
	}
	return sum / lenA
}

// Cov Covariance
// Cov(x,y) = SIGMA(XY) - (SIGMA(X) * SIGMA(Y))
// SIGMA a total of all of the elements of a
// n is the number of x,y data points
func Cov(x, y []float64) float64 {
	if len(x) == 0 {
		panic(NOOR)
	}
	if len(x) != len(y) {
		panic(NOOR)
	}
	xy := make([]float64, len(x))
	for i, _ := range x {
		xy[i] = x[i] * y[i]
	}
	xysl := xy[:]
	return Mean(xysl) - (Mean(x) * Mean(y))
}

// R Regression
// slope(b) = (n * SIGMAXY - (SIGMA X)(SIGMA Y))) / (n * SIGMAX^2) - (SIGMAX)^2)
// Intercept(a) = (SIGMA Y - b(SIGMA X)) / n
// r-squared = (Cov(x,y) / SD(x) * SD(y))^2
// r-squared = s1 / (p1' * q1')
// s1 = n('XY) - ('X)('Y)
// p1 = (n('X2) -- ('X)2)^1/2
// q1 = (n('X2) -- ('X)2)^1/2
func R(x, y []float64) (a, b, r float64) {
	n := float64(len(x))
	var (
		sumX, sumY, sumXY, sumXsq, sumYsq float64
	)
	for _, v := range x {
		sumX += v
	}
	for _, v := range y {
		sumY += v
	}
	for i, v := range x {
		sumXY += v * y[i]
	}
	for _, v := range x {
		sumXsq += math.Pow(v, 2)
	}
	b = ((n * sumXY) - (sumX * sumY)) / ((n * sumXsq) - math.Pow(sumX, 2))
	a = (sumY - (b * sumX)) / n
	s1 := (n * sumXY) - (sumX * sumY)
	p1 := n*sumXsq - math.Pow(sumX, 2)
	p1sqrt := math.Sqrt(p1)
	for _, v := range y {
		sumYsq += math.Pow(v, 2)
	}
	q1 := n*sumYsq - math.Pow(sumY, 2)
	q1sqrt := math.Sqrt(q1)
	r = s1 / (p1sqrt * q1sqrt)
	return a, b, r
}
*/
