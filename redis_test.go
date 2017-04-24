package env_test

//
// Copyright 2017 Pedro Salgado
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

import (
	"os"

	"github.com/stretchr/testify/suite"

	"github.com/steenzout/go-env"
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

// TestGetRedis check default value behavior of GetRedis*().
func (s ClearEnvSuite) TestGetRedis() {
	s.Equal(6379, env.GetRedisPort())
}
