package gocheck_one_package_test

import (
	"testing"

	. "gopkg.in/check.v1"
)

type TestPackageSuite struct {
}

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

func init() {
	Suite(&TestPackageSuite{})
}

func (s *TestPackageSuite) TestPackageTest(c *C) {
	c.Assert(true, Equals, true)
}
