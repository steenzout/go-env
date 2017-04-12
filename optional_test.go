package env_test

import (
	"github.com/steenzout/go-env"
)

const (
	// EnvInt name of the environment variable that contains an integer value.
	EnvInt = "ENV_INT"
	// EnvUnknown inexistent environment variable.
	EnvUnknown = "ENV_UNKNOWN"
	// EnvStr name of the environment variable that contains a string value.
	EnvStr = "ENV_STR"
)

// OptionalTestSuite test suite for functions in optional.go.
type OptionalTestSuite struct {
	PackageTestSuite
}

// Test check behavior of GetOptionalInt().
func (s OptionalTestSuite) TestGetOptionalInt() {
	s.Equal(0, env.GetOptionalInt(EnvUnknown, 0))
	s.Equal(0, env.GetOptionalInt(EnvStr, 0))
	s.Equal(1, env.GetOptionalInt(EnvInt, 0))
}

// Test check behavior of GetOptionalString().
func (s OptionalTestSuite) TestGetOptionalString() {
	s.Equal("default", env.GetOptionalString(EnvUnknown, "default"))
	s.Equal("str", env.GetOptionalString(EnvStr, "default"))
	s.Equal("1", env.GetOptionalString(EnvInt, "default"))
}
