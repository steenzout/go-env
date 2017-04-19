package env

const (
	// EnvPostgreSQLDatabase name of the environment variable that contains the PostgreSQL database.
	EnvPostgreSQLDatabase = "POSTGRESQL_DB"
	// EnvPostgreSQLHost name of the environment variable that contains the PostgreSQL host.
	EnvPostgreSQLHost = "POSTGRESQL_HOST"
	// EnvPostgreSQLPassword name of the environment variable that contains the PostgreSQL user's password.
	EnvPostgreSQLPassword = "POSTGRESQL_PASSWORD"
	// EnvPostgreSQLPort name of the environment variable that contains the PostgreSQL port.
	EnvPostgreSQLPort = "POSTGRESQL_PORT"
	// EnvPostgreSQLUser name of the environment variable that contains the PostgreSQL user.
	EnvPostgreSQLUser = "POSTGRESQL_USER"
)

// GetPostgreSQLDatabase returns the PostgreSQL database.
func GetPostgreSQLDatabase() string {
	return GetString(EnvPostgreSQLDatabase)
}

// GetPostgreSQLHost returns the PostgreSQL host.
func GetPostgreSQLHost() string {
	return GetString(EnvPostgreSQLHost)
}

// GetPostgreSQLPassword returns the PostgreSQL user password.
func GetPostgreSQLPassword() string {
	return GetString(EnvPostgreSQLPassword)
}

// GetPostgreSQLPort returns the PostgreSQL port.
func GetPostgreSQLPort() int {
	return GetOptionalInt(EnvPostgreSQLPort, 5432)
}

// GetPostgreSQLUser returns the PostgreSQL user.
func GetPostgreSQLUser() string {
	return GetString(EnvPostgreSQLUser)
}
