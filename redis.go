//
// Copyright 2017 Pedro Salgado
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
