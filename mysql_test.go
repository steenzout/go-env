package env_test

//
// Copyright 2017-2018 Pedro Salgado
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

	_ "github.com/go-sql-driver/mysql"

	"github.com/stretchr/testify/suite"

	"github.com/steenzout/go-env"
)

// MySQLTestSuite test suite for functions in mysql.go.
type MySQLTestSuite struct {
	suite.Suite
}

// TearDownTest clears all environment variables.
func (s MySQLTestSuite) TearDownTest() {
	os.Clearenv()
}

// TestGetMySQL check behavior of GetMySQL*() functions.
func (s MySQLTestSuite) TestGetMySQL() {
	setUp := func() {
		os.Setenv(env.EnvMySQLDatabase, "test")
		os.Setenv(env.EnvMySQLHost, "example.com")
		os.Setenv(env.EnvMySQLPort, "3306")
		os.Setenv(env.EnvMySQLProtocol, "tcp")
		os.Setenv(env.EnvMySQLPassword, "secret1")
		os.Setenv(env.EnvMySQLRootPassword, "secret2")
		os.Setenv(env.EnvMySQLUser, "travis")
	}
	setUp()

	s.Equal("example.com:3306", env.GetMySQLAddress())
	s.Equal("test", env.GetMySQLDatabase())
	s.Equal("example.com", env.GetMySQLHost())
	s.Equal("secret1", env.GetMySQLPassword())
	s.Equal(3306, env.GetMySQLPort())
	s.Equal("secret2", env.GetMySQLRootPassword())
	s.Equal("travis", env.GetMySQLUser())
}

// TestGetMySQLAddressWithUnixDomainSocket check behavior of GetMySQLAddress() function when using Unix domain socket.
func (s MySQLTestSuite) TestGetMySQLAddressWithUnixDomainSocket() {
	setUp := func() {
		os.Setenv(env.EnvMySQLHost, "/tmp/mysql.sock")
		os.Setenv(env.EnvMySQLProtocol, "unix")
	}
	setUp()

	s.Equal("/tmp/mysql.sock", env.GetMySQLAddress())
}

// TestGetMySQLConnection check behavior of GetMySQLConnection().
func (s MySQLTestSuite) TestGetMySQLConnection() {
	setUp := func() {
		// set these just like defined in docker.env
		os.Setenv(env.EnvMySQLDatabase, "mysql")
		os.Setenv(env.EnvMySQLHost, "127.0.0.1")
		os.Setenv(env.EnvMySQLPort, "3306")
		os.Setenv(env.EnvMySQLProtocol, "tcp")
		os.Setenv(env.EnvMySQLPassword, "")
		os.Setenv(env.EnvMySQLUser, "root")
	}
	setUp()

	db, err := env.GetMySQLConnection()
	if s.Nil(err) && s.NotNil(db) {
		s.Nil(db.Ping())
	}
}

// TestWrongCredentials check behavior when wrong credentials are passed to GetMySQLConnection().
func (s MySQLTestSuite) TestWrongCredentials() {
	os.Setenv(env.EnvMySQLDatabase, "mysql")
	os.Setenv(env.EnvMySQLHost, "127.0.0.1")
	os.Setenv(env.EnvMySQLPort, "3306")
	os.Setenv(env.EnvMySQLProtocol, "tcp")
	os.Setenv(env.EnvMySQLPassword, "bad")
	os.Setenv(env.EnvMySQLRootPassword, "bad")
	os.Setenv(env.EnvMySQLUser, "root")

	db, err := env.GetMySQLConnection()
	if s.NotNil(err) {
		s.Nil(db)
	}
}

// TestGetMySQL check default value behavior of GetMySQL*().
func (s ClearEnvSuite) TestGetMySQL() {
	s.Equal("/var/run/mysqld/mysqld.sock", env.GetMySQLHost())
	s.Equal("", env.GetMySQLPassword())
	s.Equal(3306, env.GetMySQLPort())
	s.Equal("unix", env.GetMySQLProtocol())
	s.Equal("", env.GetMySQLRootPassword())
}
