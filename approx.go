package approx

import (
	"math"
)

// ContFrac is a continued fraction representation.
//
// [a0; a1, a2, a3] is the same as ContFrac{a0, a1, a2, a3}.
type ContFrac []uint

// NewContFrac creates a continued fraction representation from a float.
func NewContFrac(x float64) (f ContFrac) {
	var dividend float64
	var divisor, remainder float64 = x, 1

	for remainder != 0 {
		dividend = divisor
		divisor = remainder

		f = append(f, uint(math.Floor(dividend/divisor)))
		remainder = math.Mod(dividend, divisor)
	}

	return f
}

// Float returns float representation of a continued fraction.
func (f ContFrac) Float() float64 {
	if len(f) == 0 {
		return 0
	}
	if len(f) == 1 {
		return float64(f[0])
	}

	return float64(f[0]) + 1/f[1:].Float()
}

// Convergent returns the k-th convergent.
//
// Definitions about the convergent may differ and have offset 1.
func (f ContFrac) Convergent(k int) (p uint, q uint) {
	if k == -1 {
		p, q = 1, 0

		return
	}
	if k == 0 {
		p, q = f[0], 1

		return
	}
	if k > len(f)-1 {
		panic("convergent exceeds number of fractions")
	}

	pp, pq := f.Convergent(k - 1)
	ppp, ppq := f.Convergent(k - 2)

	p = f[k]*pp + ppp
	q = f[k]*pq + ppq

	return
}

// Ratio returns an integer ratio with value equal to the continued fraction.
func (f ContFrac) Ratio() (uint, uint) {
	return f.Convergent(len(f) - 1)
}

// RatioConstr is like Ratio but constraints the denumerator to be less than
// dmax.
func (f ContFrac) RatioConstr(dmax uint) (p uint, q uint) {
	var np, nq uint

	for k := 0; k < len(f); k++ {
		np, nq = f.Convergent(k)
		if nq > dmax {
			break
		} else {
			p, q = np, nq
		}
	}

	return
}

// RatioConstr2 is like Ratio but constraints the denumerator to be less than
// dmax and the numerator less than nmax.
func (f ContFrac) RatioConstr2(dmax, nmax uint) (p uint, q uint) {
	var np, nq uint

	for k := 0; k < len(f); k++ {
		np, nq = f.Convergent(k)
		if nq > dmax || np > nmax {
			break
		} else {
			p, q = np, nq
		}
	}

	return
}

// Ratio returns a integer ratio with the same value as x.
func Ratio(x float64) (uint, uint) {
	return NewContFrac(x).Ratio()
}

// RatioConstr returns a integer ratio with the same value as x but with the
// denumerator being less than dmax.
func RatioConstr(x float64, dmax uint) (uint, uint) {
	return NewContFrac(x).RatioConstr(dmax)
}

// RatioConstr2 returns a integer ratio with the same value as x but with the
// denumerator being less than dmax and the numerator being less than nmax.
func RatioConstr2(x float64, dmax, nmax uint) (uint, uint) {
	return NewContFrac(x).RatioConstr2(dmax, nmax)
}
