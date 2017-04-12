package env

import (
	"os"
	"strconv"
)

// GetOptionalInt returns the string value of the environment variable if it exists and is an integer;
// otherwise it will return the given default value.
func GetOptionalInt(v string, d int) int {
	value, found := os.LookupEnv(v)
	if !found {
		return d
	}

	ivalue, err := strconv.Atoi(value)
	if err != nil {
		return d
	}
	return ivalue
}

// GetOptionalString returns the string value of the environment variable if it exists;
// otherwise it will return the given default value.
func GetOptionalString(v string, d string) string {
	value, found := os.LookupEnv(v)
	if !found {
		return d
	}
	return value
}
