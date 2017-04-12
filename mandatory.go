package env

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
