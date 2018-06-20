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
	"time"

	"github.com/steenzout/go-env"
)

// OptionalTestSuite test suite for functions in optional.go.
type OptionalTestSuite struct {
	PackageTestSuite
}

// TestGetOptionalDuration check behavior of GetOptionalDuration().
func (s OptionalTestSuite) TestGetOptionalDuration() {
	d := 10 * time.Second

	s.Equal(d, env.GetOptionalDuration(EnvUnknown, d))
	s.Equal(d, env.GetOptionalDuration(EnvStr, d))
	s.Equal(d, env.GetOptionalDuration(EnvInt, d))
	s.Equal(1*time.Second, env.GetOptionalDuration(EnvDuration, d))
}

// TestGetOptionalInt check behavior of GetOptionalInt().
func (s OptionalTestSuite) TestGetOptionalInt() {
	s.Equal(0, env.GetOptionalInt(EnvUnknown, 0))
	s.Equal(0, env.GetOptionalInt(EnvStr, 0))
	s.Equal(1, env.GetOptionalInt(EnvInt, 0))
}

// TestGetOptionalString check behavior of GetOptionalString().
func (s OptionalTestSuite) TestGetOptionalString() {
	s.Equal("default", env.GetOptionalString(EnvUnknown, "default"))
	s.Equal("str", env.GetOptionalString(EnvStr, "default"))
	s.Equal("1", env.GetOptionalString(EnvInt, "default"))
}

// TestGetOptionalStringSlice check behavior of GetOptionalStringSlice().
func (s OptionalTestSuite) TestGetOptionalStringSlice() {
	d := []string{"A", "B"}

	s.Equal(d, env.GetOptionalStringSlice(EnvUnknown, EnvStrSliceDelimiter, d))
	s.Equal([]string{"str"}, env.GetOptionalStringSlice(EnvStr, EnvStrSliceDelimiter, d))
	s.Equal([]string{"1"}, env.GetOptionalStringSlice(EnvInt, EnvStrSliceDelimiter, d))
	s.Equal([]string{"element1", "elementN"}, env.GetOptionalStringSlice(EnvStrSlice, EnvStrSliceDelimiter, d))
}
