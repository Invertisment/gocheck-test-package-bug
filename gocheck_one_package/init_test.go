package gocheck_one_package

import (
	"testing"

	. "gopkg.in/check.v1"
)

type MainTestSuite struct {
}

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

func init() {
	Suite(&MainTestSuite{})
}

func (s *MainTestSuite) TestMain(c *C) {
	c.Assert(true, Equals, true)
}
