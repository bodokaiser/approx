package approx

import "math"

type ContFrac []uint

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

func (f ContFrac) Float() float64 {
	if len(f) == 0 {
		return 0
	}
	if len(f) == 1 {
		return float64(f[0])
	}

	return float64(f[0]) + 1/f[1:].Float()
}

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
func (f ContFrac) Ratio() (uint, uint) {
	return f.Convergent(len(f) - 1)
}

func (f ContFrac) RatioLimited(kmax uint) (p uint, q uint) {
	for k := 0; q <= kmax; k++ {
		p, q = f.Convergent(k)
	}

	return
}
