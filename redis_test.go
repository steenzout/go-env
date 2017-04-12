package env_test

import (
	"os"

	"github.com/steenzout/go-env"
	"github.com/stretchr/testify/suite"
)

// RedisTestSuite test suite for functions in redis.go.
type RedisTestSuite struct {
	suite.Suite
}

// SetupTest sets test environment variables.
func (s RedisTestSuite) SetupTest() {
	os.Clearenv()
	os.Setenv(env.EnvRedisDatabase, "db")
	os.Setenv(env.EnvRedisHost, "localhost")
	os.Setenv(env.EnvRedisPort, "6379")
	os.Setenv(env.EnvRedisPassword, "secret")
}

// TearDownTest clears all environment variables.
func (s RedisTestSuite) TearDownTest() {
	os.Clearenv()
}

// Test check behavior of GetRedis*() functions.
func (s RedisTestSuite) TestGetRedis() {
	s.Equal("db", env.GetRedisDatabase())
	s.Equal("localhost", env.GetRedisHost())
	s.Equal(6379, env.GetRedisPort())
	s.Equal("secret", env.GetRedisPassword())
}

// Test check behavior of GetOptionalString().
func (s ClearEnvSuite) TestDefaultGetRedisPort() {
	s.Equal(6379, env.GetRedisPort())
}
