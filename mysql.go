package env

const (
	// EnvMySQLDatabase name of the environment variable that contains the MySQL database.
	EnvMySQLDatabase = "MYSQL_DB"
	// EnvMySQLHost name of the environment variable that contains the MySQL host.
	EnvMySQLHost = "MYSQL_HOST"
	// EnvMySQLPassword name of the environment variable that contains the MySQL user's password.
	EnvMySQLPassword = "MYSQL_PASSWORD"
	// EnvMySQLPort name of the environment variable that contains the MySQL port.
	EnvMySQLPort = "MYSQL_PORT"
	// EnvMySQLUser name of the environment variable that contains the MySQL user.
	EnvMySQLUser = "MYSQL_USER"
)

// GetMySQLDatabase returns the MySQL database.
func GetMySQLDatabase() string {
	return GetString(EnvMySQLDatabase)
}

// GetMySQLHost returns the MySQL host.
func GetMySQLHost() string {
	return GetString(EnvMySQLHost)
}

// GetMySQLPassword returns the MySQL user password.
func GetMySQLPassword() string {
	return GetString(EnvMySQLPassword)
}

// GetMySQLPort returns the MySQL port.
func GetMySQLPort() int {
	return GetOptionalInt(EnvMySQLPort, 3306)
}

// GetMySQLUser returns the MySQL user.
func GetMySQLUser() string {
	return GetString(EnvMySQLUser)
}
