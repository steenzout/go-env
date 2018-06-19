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
	// EnvPostgreSQLDatabase name of the environment variable that contains the PostgreSQL database.
	EnvPostgreSQLDatabase = "POSTGRES_DB"
	// EnvPostgreSQLHost name of the environment variable that contains the PostgreSQL host.
	EnvPostgreSQLHost = "POSTGRES_HOST"
	// EnvPostgreSQLPassword name of the environment variable that contains the PostgreSQL user's password.
	EnvPostgreSQLPassword = "POSTGRES_PASSWORD"
	// EnvPostgreSQLPort name of the environment variable that contains the PostgreSQL port.
	EnvPostgreSQLPort = "POSTGRES_PORT"
	// EnvPostgreSQLUser name of the environment variable that contains the PostgreSQL user.
	EnvPostgreSQLUser = "POSTGRES_USER"
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
