package approx

func PartialQuotients(a, b uint) (f ContFrac) {
	var dividend, divisor, remainder uint = 0, a, b

	for remainder != 0 {
		dividend = divisor
		divisor = remainder

		f = append(f, uint64(dividend/divisor))
		remainder = dividend % divisor
	}

	return f
}
