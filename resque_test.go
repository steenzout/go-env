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
	"time"

	"github.com/stretchr/testify/suite"

	"github.com/steenzout/go-env"
)

// ResqueTestSuite test suite for functions in resque.go.
type ResqueTestSuite struct {
	suite.Suite
}

// SetupTest sets test environment variables.
func (s ResqueTestSuite) SetupTest() {
	os.Clearenv()
	os.Setenv(env.EnvResqueCount, "5")
	os.Setenv(env.EnvResqueInterval, "3s")
	os.Setenv(env.EnvResquePIDFile, "/var/run/resque.workerA.pid")
	os.Setenv(env.EnvResqueQueue, "file_serve")
	os.Setenv(env.EnvResqueQueues, "queue_A, queue_B ")
}

// TearDownTest clears all environment variables.
func (s ResqueTestSuite) TearDownTest() {
	os.Clearenv()
}

// Test check behavior of GetResque*() functions.
func (s ResqueTestSuite) TestGetResque() {
	s.Equal("file_serve", env.GetResqueQueue())
	s.Equal(5, env.GetResqueCount())
	s.Equal(3*time.Second, env.GetResqueInterval())
	s.Equal("/var/run/resque.workerA.pid", env.GetResquePIDFile())
	s.Equal([]string{"queue_A", "queue_B"}, env.GetResqueQueues())
}

// TestGetResque check default value behavior of GetResque*().
func (s ClearEnvSuite) TestGetResque() {
	s.Equal(env.ResqueCount, env.GetResqueCount())
	s.Equal(env.ResqueInterval, env.GetResqueInterval())
	s.Equal([]string{"*"}, env.GetResqueQueues())
}
