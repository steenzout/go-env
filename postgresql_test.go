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
