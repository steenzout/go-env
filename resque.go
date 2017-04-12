package env

import (
	"time"
)

const (
	// EnvResqueQueue name of the environment variable that contains the name of a single Resque queue.
	EnvResqueQueue = "RESQUE_QUEUE"
	// EnvResqueCount name of the environment variable that contains the number of Resque workers to be spawned.
	EnvResqueCount = "RESQUE_COUNT"
	// EnvResqueInterval name of the environment variable to define, in seconds, the polling frequency.
	EnvResqueInterval = "RESQUE_INTERVAL"
	// EnvResquePIDFile name of the environment variable that contains the PID file location.
	EnvResquePIDFile = "RESQUE_PIDFILE"
	// EnvResqueQueues name of the environment variable that contains a list of Resque queue names.
	EnvResqueQueues = "RESQUE_QUEUES"
	// ResqueCount the default value for the number of Resque workers.
	ResqueCount = 1
	// ResqueInterval the default value for the Resque polling interval.
	ResqueInterval = 5 * time.Second
)

// GetResqueQueue returns the Resque queue.
func GetResqueQueue() string {
	return GetString(EnvResqueQueue)
}

// GetResqueCount returns the number of Resque workers.
func GetResqueCount() int {
	return GetOptionalInt(EnvResqueCount, ResqueCount)
}

// GetResqueInterval returns the Resque polling frequency.
func GetResqueInterval() *time.Duration {
	d := ResqueInterval
	return GetOptionalDuration(EnvResqueInterval, &d)
}

// GetResquePIDFile returns the Resque PID file location.
func GetResquePIDFile() string {
	return GetString(EnvResquePIDFile)
}

// GetResqueQueues returns the list of Resque queue names.
func GetResqueQueues() []string {
	return GetOptionalStringSlice(EnvResqueQueues, ",", []string{"*"})
}
