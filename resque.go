package env

const (
	// EnvResqueQueue name of the environment variable that contains the name of a single Resque queue.
	EnvResqueQueue = "RESQUE_QUEUE"
)

// GetResqueQueue returns the Resque queue.
func GetResqueQueue() string {
	return GetString(EnvResqueQueue)
}
