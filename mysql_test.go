package env_test

import (
	"os"

	"github.com/stretchr/testify/suite"

	"github.com/steenzout/go-env"
)

// MySQLTestSuite test suite for functions in postgresql.go.
type MySQLTestSuite struct {
	suite.Suite
}

// SetupTest sets test environment variables.
func (s MySQLTestSuite) SetupTest() {
	os.Clearenv()
	os.Setenv(env.EnvMySQLDatabase, "db")
	os.Setenv(env.EnvMySQLHost, "localhost")
	os.Setenv(env.EnvMySQLPort, "3306")
	os.Setenv(env.EnvMySQLPassword, "secret")
	os.Setenv(env.EnvMySQLUser, "user")
}

// TearDownTest clears all environment variables.
func (s MySQLTestSuite) TearDownTest() {
	os.Clearenv()
}

// Test check behavior of GetMySQL*() functions.
func (s MySQLTestSuite) TestGetMySQL() {
	s.Equal("db", env.GetMySQLDatabase())
	s.Equal("localhost", env.GetMySQLHost())
	s.Equal("secret", env.GetMySQLPassword())
	s.Equal(3306, env.GetMySQLPort())
	s.Equal("user", env.GetMySQLUser())
}

// TestGetMySQL check default value behavior of GetMySQL*().
func (s ClearEnvSuite) TestGetMySQL() {
	s.Equal(3306, env.GetMySQLPort())
}
