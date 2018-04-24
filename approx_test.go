package approx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPartialQuotients(t *testing.T) {
	assert.Equal(t, ContFrac{3, 4, 12, 4}, PartialQuotients(649, 200))
}

func TestRational(t *testing.T) {
	p, q := Rational(649, 200, 200)

	assert.Equal(t, uint(649), p)
	assert.Equal(t, uint(200), q)
}

func TestRational2(t *testing.T) {
	//p, q := Rational2(1/math.Pi, 500)

	//assert.Equal(t, uint(31830989), p)
	//assert.Equal(t, uint(100000000), q)
}
