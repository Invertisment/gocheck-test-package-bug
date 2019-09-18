package gocheck_one_package_test

import (
	"testing"

	. "gopkg.in/check.v1"
)

type TestPackageSuite struct {}

func init() {
	Suite(&TestPackageSuite{})
}

func (s *TestPackageSuite) TestPackageTest(c *C) {
	c.Assert(true, Equals, true)
}
