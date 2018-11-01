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

	"github.com/stretchr/testify/suite"

	"github.com/steenzout/go-env"
)

// InfluxDBTestSuite test suite for functions in influxdb.go.
type InfluxDBTestSuite struct {
	suite.Suite
}

// TearDownTest clears all environment variables.
func (s InfluxDBTestSuite) TearDownTest() {
	os.Clearenv()
}

// TestGetInfluxDB check behavior of GetInfluxDB*() functions.
func (s InfluxDBTestSuite) TestGetInfluxDB() {
	setUp := func() {
		os.Setenv(env.EnvInfluxDBDatabase, "test")
		os.Setenv(env.EnvInfluxDBHost, "example.com")
		os.Setenv(env.EnvInfluxDBProtocol, "https")
		os.Setenv(env.EnvInfluxDBPort, "1234")
		os.Setenv(env.EnvInfluxDBAdminPort, "1235")
		os.Setenv(env.EnvInfluxDBGraphitePort, "1236")
		os.Setenv(env.EnvInfluxDBAdminUser, "root")
		os.Setenv(env.EnvInfluxDBAdminPassword, "secret")
		os.Setenv(env.EnvInfluxDBUser, "usr")
		os.Setenv(env.EnvInfluxDBPassword, "pwd")
		os.Setenv(env.EnvInfluxDBReadUser, "reader")
		os.Setenv(env.EnvInfluxDBReadPassword, "only")
		os.Setenv(env.EnvInfluxDBWriteUser, "writer")
		os.Setenv(env.EnvInfluxDBWritePassword, "always")
	}
	setUp()

	s.Equal("test", env.GetInfluxDBDatabase())
	s.Equal("example.com", env.GetInfluxDBHost())
	s.Equal("https", env.GetInfluxDBProtocol())
	s.Equal(1234, env.GetInfluxDBPort())
	s.Equal(1235, env.GetInfluxDBAdminPort)
	s.Equal(1236, env.GetInfluxDBGraphitePort())
	s.Equal("root", env.GetInfluxDBAdminUser())
	s.Equal("secret", env.GetInfluxDBAdminPassword())
	s.Equal("usr", env.GetInfluxDBUser())
	s.Equal("pwd", env.GetInfluxDBPassword())
	s.Equal("reader", env.GetInfluxDBReadUser())
	s.Equal("only", env.GetInfluxDBReadPassword())
	s.Equal("writer", env.GetInfluxDBWriteUser())
	s.Equal("always", env.GetInfluxDBWritePassword())
	s.Equal("https://example.com:1234", env.GetInfluxDBAddress())
}

// TestGetInfluxDBDefaultValues check default value behavior of GetInfluxDB*().
func (s ClearEnvSuite) TestGetInfluxDBDefaultValues() {
	s.Equal("http", env.GetInfluxDBProtocol())
	s.Equal(8086, env.GetInfluxDBPort())
	s.Equal(8083, env.GetInfluxDBAdminPort())
	s.Equal(2003, env.GetInfluxDBGraphitePort())
}
