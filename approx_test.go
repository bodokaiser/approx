package approx

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ContFracSuite struct {
	suite.Suite

	cf1 ContFrac
	cf2 ContFrac
	cf3 ContFrac
	cf4 ContFrac
	cf5 ContFrac
}

func (s *ContFracSuite) SetupTest() {
	s.cf1 = ContFrac{1, 2}
	s.cf2 = ContFrac{4, 2, 6, 7}
	s.cf3 = ContFrac{3, 4, 12, 4}
	s.cf4 = ContFrac{0, 1, 5, 2, 2}
	s.cf5 = ContFrac{3, 7, 15, 1}
}

func (s *ContFracSuite) TestFloat() {
	assert.Equal(s.T(), 1.5, s.cf1.Float())
	assert.InEpsilon(s.T(), 415.0/93.0, s.cf2.Float(), 1e-6)
	assert.InEpsilon(s.T(), 3.245, s.cf3.Float(), 1e-6)
}

func (s *ContFracSuite) TestNew() {
	cf := NewContFrac(0.84375)
	assert.Equal(s.T(), s.cf4, cf)
}

func (s *ContFracSuite) TestConvergent() {
	p0, q0 := s.cf4.Convergent(0)
	assert.Equal(s.T(), uint(0), p0)
	assert.Equal(s.T(), uint(1), q0)

	p1, q1 := s.cf4.Convergent(1)
	assert.Equal(s.T(), uint(1), p1)
	assert.Equal(s.T(), uint(1), q1)

	p2, q2 := s.cf4.Convergent(2)
	assert.Equal(s.T(), uint(5), p2)
	assert.Equal(s.T(), uint(6), q2)

	p3, q3 := s.cf4.Convergent(3)
	assert.Equal(s.T(), uint(11), p3)
	assert.Equal(s.T(), uint(13), q3)

	p4, q4 := s.cf4.Convergent(4)
	assert.Equal(s.T(), uint(27), p4)
	assert.Equal(s.T(), uint(32), q4)

	p, q := s.cf5.Convergent(3)
	assert.Equal(s.T(), uint(355), p)
	assert.Equal(s.T(), uint(113), q)
}

func (s *ContFracSuite) TestRatio() {
	p, q := s.cf4.Ratio()
	assert.Equal(s.T(), uint(27), p)
	assert.Equal(s.T(), uint(32), q)
}

func (s *ContFracSuite) TestRatioConstr() {
	p, q := s.cf4.RatioConstr(10)
	assert.Equal(s.T(), uint(5), p)
	assert.Equal(s.T(), uint(6), q)
}

func (s *ContFracSuite) TestRatioConstr2() {
	p, q := s.cf4.RatioConstr2(30, 20)
	assert.Equal(s.T(), uint(11), p)
	assert.Equal(s.T(), uint(13), q)
}

func TestRatio(t *testing.T) {
	p, q := Ratio(1 / math.Pi)
	assert.Equal(t, 1/math.Pi, float64(p)/float64(q))
}

func TestRatioConstr(t *testing.T) {
	p, q := RatioConstr(1/math.Pi, 92000)
	assert.True(t, q < 92000)
	assert.Equal(t, uint(113), p)
	assert.Equal(t, uint(355), q)
	assert.InEpsilon(t, 1/math.Pi, float64(p)/float64(q), 1e-7)
}

func TestRatioConstr2(t *testing.T) {
	p, q := RatioConstr2(1/math.Pi, 1<<32-1, 1<<16-1)

	assert.True(t, q < 1<<32-1)
	assert.True(t, p < 1<<16-1)
	assert.InEpsilon(t, 1/math.Pi, float64(p)/float64(q), 1e-9)
}

func TestContFracTestSuite(t *testing.T) {
	suite.Run(t, new(ContFracSuite))
}
