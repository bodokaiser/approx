package approx

func PartialQuotients(a, b uint) (f ContFrac) {
	var dividend uint
	var divisor, remainder uint = a, b

	for remainder != 0 {
		dividend = divisor
		divisor = remainder

		f = append(f, dividend/divisor)
		remainder = dividend % divisor
	}

	return f
}

func Rational(a, b, kmax uint) (uint, uint) {
	cf := ContFrac{}
	p, q := cf.Convergent(-1)

	var dividend uint
	var divisor, remainder uint = a, b

	for k := 0; q <= kmax && remainder != 0; k++ {
		dividend = divisor
		divisor = remainder

		cf = append(cf, dividend/divisor)
		remainder = dividend % divisor

		p, q = cf.Convergent(k)
	}

	return p, q
}
