package env

import (
	"fmt"
)

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
	// EnvInfluxDBDatabase name of the environment variable that
	// contains the name of the InfluxDB database.
	EnvInfluxDBDatabase = "INFLUXDB_DB"
	// EnvInfluxDBHost name of the environment variable that contains the InfluxDB host.
	EnvInfluxDBHost = "INFLUXDB_HOST"
	// EnvInfluxDBProtocol name of the environment variable that contains the protocol to be used when connecting to InfluxDB.
	EnvInfluxDBProtocol = "INFLUXDB_PROTOCOL"

	// EnvInfluxDBPort name of the environment variable that contains the InfluxDB HTTP API port.
	EnvInfluxDBPort = "INFLUXDB_PORT"
	// EnvInfluxDBAdminPort name of the environment variable that contains the InfluxDB administrator interface port.
	// The administrator interface is deprecated as of 1.1.0 and will be removed in 1.3.0.
	EnvInfluxDBAdminPort = "INFLUXDB_ADMIN_PORT"
	// EnvInfluxDBGraphitePort name of the environment variable that contains the Graphite support port.
	EnvInfluxDBGraphitePort = "INFLUXDB_GRAPHITE_PORT"

	// EnvInfluxDBAdminUser name of the environment variable that
	// contains the InfluxDB administrator account.
	EnvInfluxDBAdminUser = "INFLUXDB_ADMIN_USER"
	// EnvInfluxDBAdminPassword name of the environment variable that
	// contains the name of the InfluxDB administrator account password.
	EnvInfluxDBAdminPassword = "INFLUXDB_ADMIN_PASSWORD"

	// EnvInfluxDBUser name of the environment variable that
	// contains the name of the InfluxDB database read/write account.
	EnvInfluxDBUser = "INFLUXDB_USER"
	// EnvInfluxDBPassword name of the environment variable that
	// contains the name of the InfluxDB database read/write account password.
	EnvInfluxDBPassword = "INFLUXDB_USER_PASSWORD"

	// EnvInfluxDBReadUser name of the environment variable that
	// contains the name of the InfluxDB database read account.
	EnvInfluxDBReadUser = "INFLUXDB_READ_USER"
	// EnvInfluxDBReadPassword name of the environment variable that
	// contains the name of the InfluxDB database read account password.
	EnvInfluxDBReadPassword = "INFLUXDB_READ_USER_PASSWORD"

	// EnvInfluxDBWriteUser name of the environment variable that
	// contains the name of the InfluxDB database write account.
	EnvInfluxDBWriteUser = "INFLUXDB_WRITE_USER"
	// EnvInfluxDBWritePassword name of the environment variable that
	// contains the name of the InfluxDB database write account password.
	EnvInfluxDBWritePassword = "INFLUXDB_WRITE_USER_PASSWORD"
)

// GetInfluxDBDatabase returns the InfluxDB database.
func GetInfluxDBDatabase() string {
	return GetString(EnvInfluxDBDatabase)
}

// GetInfluxDBHost returns the InfluxDB host.
func GetInfluxDBHost() string {
	return GetString(EnvInfluxDBHost)
}

// GetInfluxDBPort returns the InfluxDB HTTP API port.
func GetInfluxDBPort() int {
	return GetOptionalInt(EnvInfluxDBPort, 8086)
}

// GetInfluxDBAddress returns the InfluxDB HTTP address.
func GetInfluxDBAddress() string {
	return fmt.Sprintf(
		"%s://%s:%d/",
		GetInfluxDBProtocol(),
		GetInfluxDBHost(),
		GetInfluxDBPort(),
	)
}

// GetInfluxDBAdminPort returns the InfluxDB administrator interface port.
func GetInfluxDBAdminPort() int {
	return GetOptionalInt(EnvInfluxDBAdminPort, 8083)
}

// GetInfluxDBGraphitePort returns the Graphite support port.
func GetInfluxDBGraphitePort() int {
	return GetOptionalInt(EnvInfluxDBGraphitePort, 2003)
}

// GetInfluxDBAdminUser returns the InfluxDB administrator account.
func GetInfluxDBAdminUser() string {
	return GetString(EnvInfluxDBAdminUser)
}

// GetInfluxDBAdminPassword returns the InfluxDB administrator password.
func GetInfluxDBAdminPassword() string {
	return GetString(EnvInfluxDBAdminPassword)
}

// GetInfluxDBUser returns the InfluxDB database read/write account.
func GetInfluxDBUser() string {
	return GetString(EnvInfluxDBUser)
}

// GetInfluxDBPassword returns the InfluxDB database read/write account password.
func GetInfluxDBPassword() string {
	return GetString(EnvInfluxDBPassword)
}

// GetInfluxDBProtocol returns the protocol to use when connecting to InfluxDB (default: http).
func GetInfluxDBProtocol() string {
	return GetOptionalString(EnvInfluxDBProtocol, "http")
}

// GetInfluxDBReadUser returns the InfluxDB database read account.
func GetInfluxDBReadUser() string {
	return GetString(EnvInfluxDBReadUser)
}

// GetInfluxDBReadPassword returns the InfluxDB database read account password.
func GetInfluxDBReadPassword() string {
	return GetString(EnvInfluxDBReadPassword)
}

// GetInfluxDBWriteUser returns the InfluxDB database write account.
func GetInfluxDBWriteUser() string {
	return GetString(EnvInfluxDBWriteUser)
}

// GetInfluxDBWritePassword returns the InfluxDB database write account password.
func GetInfluxDBWritePassword() string {
	return GetString(EnvInfluxDBWritePassword)
}
