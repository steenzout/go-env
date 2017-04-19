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
	// EnvAWSAccessKeyID name of the environment variable that contains the AWS access key ID.
	EnvAWSAccessKeyID = "AWS_ACCESS_KEY_ID"
	// EnvAWSSecretAccessKey name of the environment variable that contains the AWS access key secret.
	EnvAWSSecretAccessKey = "AWS_SECRET_ACCESS_KEY"
	// EnvAWSRegion name of the environment variable that contains the AWS region.
	EnvAWSRegion = "AWS_REGION"
)

// GetAWSAccessKeyID returns the AWS access key ID.
func GetAWSAccessKeyID() string {
	return GetString(EnvAWSAccessKeyID)
}

// GetAWSSecretAccessKey returns the AWS access key secret.
func GetAWSSecretAccessKey() string {
	return GetString(EnvAWSSecretAccessKey)
}

// GetAWSRegion returns the AWS region.
func GetAWSRegion() string {
	return GetString(EnvAWSRegion)
}
