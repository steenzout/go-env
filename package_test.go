package env_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/suite"
)

type PackageTestSuite struct {
	suite.Suite
}

func (s PackageTestSuite) SetupTest() {
	os.Clearenv()
	os.Setenv(EnvInt, "1")
	os.Setenv(EnvStr, "str")
}

func (s PackageTestSuite) TearDownTest() {
	os.Clearenv()
}

// TestPackage run all test suites included in the env package.
func TestPackage(t *testing.T) {
	suite.Run(t, new(MandatoryTestSuite))
	suite.Run(t, new(OptionalTestSuite))
}
