package env

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
	"fmt"
	"os"
	"strconv"
)

const (
	// ErrMandatoryVariable error message format for the case when an environment variable is missing.
	ErrMandatoryVariable = "Environment variable %s is mandatory"
	// ErrValueType error message format for the case when an environment variable has an unexpected type.
	ErrValueType = "Environment variable %s is not of type %s"
)

// GetInt returns the int value of a mandatory environment variable.
func GetInt(v string) int {
	value, found := os.LookupEnv(v)
	if !found {
		panic(fmt.Sprintf(ErrMandatoryVariable, v))
	}

	ivalue, err := strconv.Atoi(value)
	if err != nil {
		panic(fmt.Sprintf(ErrValueType, v, "int"))
	}
	return ivalue
}

// GetString returns the string value of a mandatory environment variable.
func GetString(v string) string {

	value, found := os.LookupEnv(v)
	if !found {
		panic(fmt.Sprintf(ErrMandatoryVariable, v))
	}
	return value
}
