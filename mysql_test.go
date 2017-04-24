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
