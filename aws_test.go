package env_test

import (
	"os"

	"github.com/steenzout/go-env"
	"github.com/stretchr/testify/suite"
)

// AWSTestSuite test suite for functions in aws.go.
type AWSTestSuite struct {
	suite.Suite
}

// SetupTest sets test environment variables.
func (s AWSTestSuite) SetupTest() {
	os.Clearenv()
	os.Setenv(env.EnvAWSAccessKeyID, "ID")
	os.Setenv(env.EnvAWSSecretAccessKey, "secret")
	os.Setenv(env.EnvAWSRegion, "us-east-1")
}

// TearDownTest clears all environment variables.
func (s AWSTestSuite) TearDownTest() {
	os.Clearenv()
}

// Test check behavior of GetAWS*() functions.
func (s AWSTestSuite) TestGetAWS() {
	s.Equal("ID", env.GetAWSAccessKeyID())
	s.Equal("secret", env.GetAWSSecretAccessKey())
	s.Equal("us-east-1", env.GetAWSRegion())
}
