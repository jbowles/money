package money

/*

 I borowed and modified from the original project:
	https://github.com/Confunctionist/finance
 See the license here:
	https://github.com/Confunctionist/finance/blob/master/LICENSE

The package contains type Money...

type Money struct {
	M	int64
}

...which usese a fixed-length guard for precision arithmetic: the
int64 variable Guard (and its float64 and int-related variables Guardf
and Guardi.

Rounding is done on float64 to int64 by	the Rnd() function truncating
at values less than (.5 + (1 / Guardf))	or greater than -(.5 + (1 / Guardf))
in the case of negative numbers. The Guard adds four decimal places
of protection to rounding.

DP is the decimal precision, which can be changed in the DecimalPrecision()
function.  DP hold the places after the decimalplace in teh active money struct field M

The following functions are available

Abs Returns the absolute value of Money
	(m *Money) Abs() *Money
Add Adds two Money types
	(m *Money) Add(n *Money) *Money
Cov Covariance
	Cov(x, y []float64) float64
DecimalChange resets the package-wide decimal place (default is 2 decimal places)
	DecimalChange(d int)
Div Divides one Money type from another
	(m *Money) Div(n *Money) *Money
Gett gets value of money truncating after DP (see Value() for no truncation)
	(m *Money) Gett() int64
Get gets the float64 value of money (see Value() for int64)
	(m *Money) Get() float64
Mean Average
	Mean(a []float64) float64
Mul Multiplies two Money types
	(m *Money) Mul(n *Money) *Money
Mulf Multiplies a Money with a float to return a money-stored type
	(m *Money) Mulf(f float64) *Money
Neg Returns the negative value of Money
	(m *Money) Neg() *Money
R Regression
	R(x, y []float64) (a, b, r float64)
RND Rounds int64 remainder if greater than Round
	Rnd(r int64, trunc float64) int64
SD Standard Deviation
	SD(a []float64) float64
SDs Standard Deviation of a sample
	SDs(a []float64) float64
Set sets the Money field M
	(m *Money) Set(x int64) *Money
Setf sets a float 64 into a Money type for precision calculations
	(m *Money) Setf(f float64) *Money
Sign returns the Sign of Money 1 if positive, -1 if negative
	(m *Money) Sign() int
String for money type representation in basic monetary unit (DOLLARS CENTS)
	(m *Money) String() string
Sub subtracts one Money type from another
	(m *Money) Sub(n *Money) *Money
Value returns in int64 the value of Money (also see Gett, See Get() for float64)
	(m *Money) Value() int64
*/

import (
	"fmt"
	"math"
)

type Money struct {
	M int64 // value of the integer64 Money
}

var (
	Guardi int     = 100
	Guard  int64   = int64(Guardi)
	Guardf float64 = float64(Guardi)
	DP     int64   = 100         // for default of 2 decimal places => 10^2 (can be reset)
	DPf    float64 = float64(DP) // for default of 2 decimal places => 10^2 (can be reset)
	Round          = .5
	//	Round  = .5 + (1 / Guardf)
	Roundn = Round * -1
)

const (
	DTL    = "Decimal places too large"
	DLZ    = "Decimal places cannot be less than zero"
	NOOR   = "Number out of range"
	OVFL   = "Overflow"
	MAXDEC = 18
)

//////////////////////////////////////////////////////////////
///////// GET AND SET BASIC VALUES ///////////////////////////
//////////////////////////////////////////////////////////////

// Get gets the float64 value of money (see Value() for int64)
func (m *Money) Get() float64 {
	return float64(m.M) / DPf
}

// Set sets the Money field M
func (m *Money) Set(x int64) *Money {
	m.M = x
	return m
}

// Gett gets value of money truncating after DP (see Value() for no truncation)
func (m *Money) Gett() int64 {
	return m.M / DP
}

// Setf sets a float64 into a Money type for precision calculations
func (m *Money) Setf(f float64) *Money {
	fDPf := f * DPf
	r := int64(f * DPf)
	return m.Set(Rnd(r, fDPf-float64(r)))
}

//////////////////////////////////////////////////////////////
///////// BASIC OPERATIONS '+', '-', '*', '/'  ///////////////////////////
//////////////////////////////////////////////////////////////

// Add Adds two Money types
func (m *Money) Add(n *Money) *Money {
	r := m.M + n.M
	if (r^m.M)&(r^n.M) < 0 {
		panic(OVFL)
	}
	m.M = r
	return m
}

// Sub subtracts one Money type from another
func (m *Money) Sub(n *Money) *Money {
	r := m.M - n.M
	if (r^m.M)&^(r^n.M) < 0 {
		panic(OVFL)
	}
	m.M = r
	return m
}

// Mul Multiplies two Money types
func (m *Money) Mul(n *Money) *Money {
	return m.Set(m.M * n.M / DP)
}

// Div Divides one Money type from another
func (m *Money) Div(n *Money) *Money {
	f := Guardf * DPf * float64(m.M) / float64(n.M) / Guardf
	i := int64(f)
	return m.Set(Rnd(i, f-float64(i)))
}

//////////////////////////////////////////////////////////////
///////////////////////// FORMATTING, ROUNDING ///////////////////////////////////
//////////////////////////////////////////////////////////////

// String for money type representation in basic monetary unit (DOLLARS CENTS)
func (m *Money) String() string {
	return fmt.Sprintf("%d.%02d", m.Value()/DP, m.Abs().Value()%DP)
}

// Value returns in int64 the value of Money (also see Gett(), See Get() for float64)
func (m *Money) Value() int64 {
	return m.M
}

// Abs Returns the absolute value of Money
func (m *Money) Abs() *Money {
	if m.M < 0 {
		m.Neg()
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

//////////////////////////////////////////////////////////////
///////// NOT SO BASIC FUNCTIONS ///////////////////////////
//////////////////////////////////////////////////////////////

// Neg Returns the negative value of Money
func (m *Money) Neg() *Money {
	if m.M != 0 {
		m.M *= -1
	}
	return m
}

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
	return m.Set(Rnd(r, float64(i)/Guardf/DPf-float64(r)))
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
