package env_test

import (
	"os"

	"github.com/steenzout/go-env"
	"github.com/stretchr/testify/suite"
)

// ResqueTestSuite test suite for functions in resque.go.
type ResqueTestSuite struct {
	suite.Suite
}

// SetupTest sets test environment variables.
func (s ResqueTestSuite) SetupTest() {
	os.Clearenv()
	os.Setenv(env.EnvResqueQueue, "file_serve")
}

// TearDownTest clears all environment variables.
func (s ResqueTestSuite) TearDownTest() {
	os.Clearenv()
}

// Test check behavior of GetResque*() functions.
func (s ResqueTestSuite) TestGetResque() {
	s.Equal("file_serve", env.GetResqueQueue())
}
