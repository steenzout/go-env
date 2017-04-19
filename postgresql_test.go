package env_test

import (
	"os"

	"github.com/stretchr/testify/suite"

	"github.com/steenzout/go-env"
)

// PostgreSQLTestSuite test suite for functions in postgresql.go.
type PostgreSQLTestSuite struct {
	suite.Suite
}

// SetupTest sets test environment variables.
func (s PostgreSQLTestSuite) SetupTest() {
	os.Clearenv()
	os.Setenv(env.EnvPostgreSQLDatabase, "db")
	os.Setenv(env.EnvPostgreSQLHost, "localhost")
	os.Setenv(env.EnvPostgreSQLPort, "5432")
	os.Setenv(env.EnvPostgreSQLPassword, "secret")
	os.Setenv(env.EnvPostgreSQLUser, "user")
}

// TearDownTest clears all environment variables.
func (s PostgreSQLTestSuite) TearDownTest() {
	os.Clearenv()
}

// Test check behavior of GetPostgreSQL*() functions.
func (s PostgreSQLTestSuite) TestGetPostgreSQL() {
	s.Equal("db", env.GetPostgreSQLDatabase())
	s.Equal("localhost", env.GetPostgreSQLHost())
	s.Equal("secret", env.GetPostgreSQLPassword())
	s.Equal(5432, env.GetPostgreSQLPort())
	s.Equal("user", env.GetPostgreSQLUser())
}

// TestGetPostgreSQL check default value behavior of GetPostgreSQL*().
func (s ClearEnvSuite) TestGetPostgreSQL() {
	s.Equal(5432, env.GetPostgreSQLPort())
}
