package env_test

import (
	"github.com/steenzout/go-env"
)

// MandatoryTestSuite test suite for functions in mandatory.go.
type MandatoryTestSuite struct {
	PackageTestSuite
}

// Test check behavior of GetInt().
func (s MandatoryTestSuite) TestGetInt() {
	s.Panics(func() { env.GetInt(EnvUnknown) })
	s.Panics(func() { env.GetInt(EnvStr) })
	s.Equal(1, env.GetInt(EnvInt))

}

// Test check behavior of GetString().
func (s MandatoryTestSuite) TestGetString() {
	s.Panics(func() { env.GetString(EnvUnknown) })
	s.Equal("str", env.GetString(EnvStr))
	s.Equal("1", env.GetString(EnvInt))
}
