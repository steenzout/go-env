package env

import (
	"os"
	"strconv"
	"strings"
	"time"
)

// GetOptionalDuration returns the time.Duration value of the environment variable if it exists and has a valid format;
// otherwise it will return the given default value.
func GetOptionalDuration(v string, d *time.Duration) *time.Duration {
	value, found := os.LookupEnv(v)
	if !found {
		return d
	}

	dvalue, err := time.ParseDuration(value)
	if err != nil {
		return d
	}
	return &dvalue
}

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

// GetOptionalStringSlice returns a string slice of the environment variable, if it exists;
// otherwise it will return the given default value.
// Each element of the slice is trimmed of all whitespace.
func GetOptionalStringSlice(v string, delimiter string, d []string) []string {
	value, found := os.LookupEnv(v)
	if !found {
		return d
	}

	out := splitAndTrimSlice(value, delimiter)
	if len(out) == 0 {
		return d
	}
	return out
}

// splitAndTrimSlice returns a string slice by
// splitting the given string using the given delimiter.
// Each element of the slice is trimmed of all whitespace.
func splitAndTrimSlice(s string, delimiter string) []string {
	var r []string

	for _, str := range strings.Split(s, delimiter) {
		r = append(r, strings.TrimSpace(str))
	}
	return r
}
