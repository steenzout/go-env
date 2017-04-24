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
	"testing"

	"fmt"

	"github.com/stretchr/testify/suite"
)

const (
	// EnvStr name of the environment variable that contains a time duration value.
	EnvDuration = "ENV_DURATION"
	// EnvInt name of the environment variable that contains an integer value.
	EnvInt = "ENV_INT"
	// EnvUnknown inexistent environment variable.
	EnvUnknown = "ENV_UNKNOWN"
	// EnvStr name of the environment variable that contains a string value.
	EnvStr = "ENV_STR"
	// EnvStrSlice name of the environment variable that contains a string slice value.
	EnvStrSlice = "ENV_STR_SLICE"
	// EnvStrSliceDelimiter the delimiter used to split the string slice value.
	EnvStrSliceDelimiter = ","
)

// PackageTestSuite this package test suite.
type PackageTestSuite struct {
	suite.Suite
}

// SetupTest sets test environment variables.
func (s PackageTestSuite) SetupTest() {
	os.Clearenv()
	os.Setenv(EnvDuration, "1s")
	os.Setenv(EnvInt, "1")
	os.Setenv(EnvStr, "str")
	os.Setenv(EnvStrSlice, fmt.Sprintf("element1%s elementN ", EnvStrSliceDelimiter))
}

// TearDownTest clears all environment variables.
func (s PackageTestSuite) TearDownTest() {
	os.Clearenv()
}

// TestPackage run all test suites included in the env package.
func TestPackage(t *testing.T) {
	// test generic functions
	suite.Run(t, new(MandatoryTestSuite))
	suite.Run(t, new(OptionalTestSuite))

	// test app specific functions
	suite.Run(t, new(MySQLTestSuite))
	suite.Run(t, new(PostgreSQLTestSuite))
	suite.Run(t, new(RedisTestSuite))
	suite.Run(t, new(ResqueTestSuite))

	// test cloud specific functions
	suite.Run(t, new(AWSTestSuite))

	// test defaults
	suite.Run(t, new(ClearEnvSuite))
}

// ClearEnvSuite test suite with zero configuration set.
// The purpose of this struct is to be extended with
// more test functions.
type ClearEnvSuite struct {
	suite.Suite
}

func (s ClearEnvSuite) SetupTest() {
	os.Clearenv()
}
