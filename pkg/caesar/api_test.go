package caesar

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type CaesarUnitSuite struct {
	suite.Suite
	A *assert.Assertions
}

func (s *CaesarUnitSuite) SetupSuite() {
	s.A = assert.New(s.T())
	s.A.True(true)
}

func (s *CaesarUnitSuite) TestCaesar() {
	cases := []struct {
		In, Out string
		Shift   int
	}{
		{"OEHGR SBEPR", "BRUTE FORCE", -13},
		{"BRUTE FORCE", "OEHGR SBEPR", 13}, // why?
		{"KEY", "MGA", 2},
		{"QEB NRFZH YOLTK CLU GRJMP LSBO QEB IXWV ALD",
			"THE QUICK BROWN FOX JUMPS OVER THE LAZY DOG", 3},
	}

	for _, c := range cases {
		s.A.Equal(c.Out, Apply(c.In, c.Shift))
	}
}

func TestCaesarUnit(t *testing.T) {
	suite.Run(t, new(CaesarUnitSuite))
}
