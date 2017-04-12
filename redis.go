package env

const (
	// EnvRedisDatabase name of the environment variable that contains the Redis database configuration parameter.
	EnvRedisDatabase = "REDIS_DB"
	// EnvRedisPassword name of the environment variable that contains the Redis password configuration parameter.
	EnvRedisPassword = "REDIS_PASSWORD"
	// EnvRedisPort name of the environment variable that contains the Redis port configuration parameter.
	EnvRedisPort = "REDIS_PORT"
	// EnvRedisHost name of the environment variable that contains the Redis host configuration parameter.
	EnvRedisHost = "REDIS_HOST"
)

// GetRedisDatabase returns the Redis database configuration parameter.
func GetRedisDatabase() string {
	return GetString(EnvRedisDatabase)
}

// GetRedisPassword returns the Redis password configuration parameter.
func GetRedisPassword() string {
	return GetString(EnvRedisPassword)
}

// GetRedisPort returns the Redis port configuration parameter.
func GetRedisPort() int {
	return GetOptionalInt(EnvRedisPort, 6379)
}

// GetRedisHost returns the Redis host configuration parameter.
func GetRedisHost() string {
	return GetString(EnvRedisHost)
}
