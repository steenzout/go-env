package env

const (
	// EnvResqueQueue name of the environment variable that contains the Resque queue configuration parameter.
	EnvResqueQueue = "RESQUE_QUEUE"
)

// GetResqueQueue returns the Resque queue configuration parameter.
func GetResqueQueue() string {
	return GetString(EnvResqueQueue)
}
