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

	"github.com/steenzout/go-env"
	"github.com/stretchr/testify/suite"
)

// AWSTestSuite test suite for functions in aws.go.
type AWSTestSuite struct {
	suite.Suite
}

// SetupTest sets test environment variables.
func (s AWSTestSuite) SetupTest() {
	os.Clearenv()
	os.Setenv(env.EnvAWSAccessKeyID, "ID")
	os.Setenv(env.EnvAWSBucket, "test.example.com")
	os.Setenv(env.EnvAWSPath, "/backup/database")
	os.Setenv(env.EnvAWSSecretAccessKey, "secret")
	os.Setenv(env.EnvAWSRegion, "us-east-1")
}

// TearDownTest clears all environment variables.
func (s AWSTestSuite) TearDownTest() {
	os.Clearenv()
}

// Test check behavior of GetAWS*() functions.
func (s AWSTestSuite) TestGetAWS() {
	s.Equal("ID", env.GetAWSAccessKeyID())
	s.Equal("test.example.com", env.GetAWSBucket())
	s.Equal("/backup/database", env.GetAWSPath())
	s.Equal("secret", env.GetAWSSecretAccessKey())
	s.Equal("us-east-1", env.GetAWSRegion())
}
