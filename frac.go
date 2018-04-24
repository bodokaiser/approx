package approx

type ContFrac []uint64

func (f ContFrac) Float() float64 {
	if len(f) == 0 {
		return 0
	}
	if len(f) == 1 {
		return float64(f[0])
	}

	return float64(f[0]) + 1/f[1:].Float()
}
