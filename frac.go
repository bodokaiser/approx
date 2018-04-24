package approx

type ContFrac []uint

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

	pp, pq := f.Convergent(k - 1)
	ppp, ppq := f.Convergent(k - 2)

	p = f[k]*pp + ppp
	q = f[k]*pq + ppq

	return
}
