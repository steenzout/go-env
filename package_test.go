package env_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/suite"
)

// PackageTestSuite this package test suite.
type PackageTestSuite struct {
	suite.Suite
}

// SetupTest sets test environment variables.
func (s PackageTestSuite) SetupTest() {
	os.Clearenv()
	os.Setenv(EnvInt, "1")
	os.Setenv(EnvStr, "str")
}

// TearDownTest clears all environment variables.
func (s PackageTestSuite) TearDownTest() {
	os.Clearenv()
}

// TestPackage run all test suites included in the env package.
func TestPackage(t *testing.T) {
	// test generic functions
	suite.Run(t, new(MandatoryTestSuite))
	suite.Run(t, new(OptionalTestSuite))

	// test app specific functions
	suite.Run(t, new(RedisTestSuite))

	// test cloud specific functions
	suite.Run(t, new(AWSTestSuite))

	// test defaults
	suite.Run(t, new(ClearEnvSuite))
}

// ClearEnvSuite test suite with zero configuration set.
// The purpose of this struct is to be extended with
// more test functions.
type ClearEnvSuite struct {
	suite.Suite
}

func (s ClearEnvSuite) SetupTest() {
	os.Clearenv()
}
