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
	"github.com/steenzout/go-env"
)

// MandatoryTestSuite test suite for functions in mandatory.go.
type MandatoryTestSuite struct {
	PackageTestSuite
}

// Test check behavior of GetInt().
func (s MandatoryTestSuite) TestGetInt() {
	s.Panics(func() { env.GetInt(EnvUnknown) })
	s.Panics(func() { env.GetInt(EnvStr) })
	s.Equal(1, env.GetInt(EnvInt))

}

// Test check behavior of GetString().
func (s MandatoryTestSuite) TestGetString() {
	s.Panics(func() { env.GetString(EnvUnknown) })
	s.Equal("str", env.GetString(EnvStr))
	s.Equal("1", env.GetString(EnvInt))
}
