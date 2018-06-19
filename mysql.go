package env

//
// Copyright 2017-2018 Pedro Salgado
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

const (
	// EnvMySQLDatabase name of the environment variable that contains the MySQL database.
	EnvMySQLDatabase = "MYSQL_DB"
	// EnvMySQLHost name of the environment variable that contains the MySQL host.
	EnvMySQLHost = "MYSQL_HOST"
	// EnvMySQLPassword name of the environment variable that contains the MySQL user's password.
	EnvMySQLPassword = "MYSQL_PASSWORD"
	// EnvMySQLPort name of the environment variable that contains the MySQL port.
	EnvMySQLPort = "MYSQL_PORT"
	// EnvMySQLRootPassword name of the environment variable that contains the MySQL root's password.
	EnvMySQLRootPassword = "MYSQL_ROOT_PASSWORD"
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
	return GetOptionalString(EnvMySQLPassword, "")
}

// GetMySQLPort returns the MySQL port.
func GetMySQLPort() int {
	return GetOptionalInt(EnvMySQLPort, 3306)
}

// GetMySQLUser returns the MySQL user.
func GetMySQLUser() string {
	return GetString(EnvMySQLUser)
}

// GetMySQLRootPassword returns the MySQL root password.
func GetMySQLRootPassword() string {
	return GetOptionalString(EnvMySQLRootPassword, "")
}
