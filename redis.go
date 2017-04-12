package env

const (
	// EnvRedisDatabase name of the environment variable that contains the Redis database.
	EnvRedisDatabase = "REDIS_DB"
	// EnvRedisPassword name of the environment variable that contains the Redis password.
	EnvRedisPassword = "REDIS_PASSWORD"
	// EnvRedisPort name of the environment variable that contains the Redis port.
	EnvRedisPort = "REDIS_PORT"
	// EnvRedisHost name of the environment variable that contains the Redis host.
	EnvRedisHost = "REDIS_HOST"
)

// GetRedisDatabase returns the Redis database.
func GetRedisDatabase() string {
	return GetString(EnvRedisDatabase)
}

// GetRedisPassword returns the Redis password.
func GetRedisPassword() string {
	return GetString(EnvRedisPassword)
}

// GetRedisPort returns the Redis port.
func GetRedisPort() int {
	return GetOptionalInt(EnvRedisPort, 6379)
}

// GetRedisHost returns the Redis host.
func GetRedisHost() string {
	return GetString(EnvRedisHost)
}
