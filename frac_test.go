package approx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContFrac(t *testing.T) {
	cf1 := ContFrac{1, 2}
	assert.Equal(t, 1.5, cf1.Float())

	cf2 := ContFrac{4, 2, 6, 7}
	assert.InEpsilon(t, 415.0/93.0, cf2.Float(), 1e-6)

	cf3 := ContFrac{3, 4, 12, 4}
	assert.InEpsilon(t, 3.245, cf3.Float(), 1e-6)
}

func TestContFracConvergent(t *testing.T) {
	cf := ContFrac{0, 1, 5, 2, 2}

	p0, q0 := cf.Convergent(0)
	assert.Equal(t, uint(0), p0)
	assert.Equal(t, uint(1), q0)

	p1, q1 := cf.Convergent(1)
	assert.Equal(t, uint(1), p1)
	assert.Equal(t, uint(1), q1)

	p2, q2 := cf.Convergent(2)
	assert.Equal(t, uint(5), p2)
	assert.Equal(t, uint(6), q2)

	p3, q3 := cf.Convergent(3)
	assert.Equal(t, uint(11), p3)
	assert.Equal(t, uint(13), q3)

	p4, q4 := cf.Convergent(4)
	assert.Equal(t, uint(27), p4)
	assert.Equal(t, uint(32), q4)
}
