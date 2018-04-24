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
