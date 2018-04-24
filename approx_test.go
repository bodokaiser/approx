package approx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPartialQuotients(t *testing.T) {
	assert.Equal(t, ContFrac{3, 4, 12, 4}, PartialQuotients(649, 200))
}
