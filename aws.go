package env

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

const (
	// EnvAWSAccessKeyID name of the environment variable that contains the AWS access key ID.
	EnvAWSAccessKeyID = "AWS_ACCESS_KEY_ID"
	// EnvAWSBucket name of the bucket.
	EnvAWSBucket = "AWS_BUCKET"
	// EnvAWSPath path in the bucket.
	EnvAWSPath = "AWS_PATH"
	// EnvAWSSecretAccessKey name of the environment variable that contains the AWS access key secret.
	EnvAWSSecretAccessKey = "AWS_SECRET_ACCESS_KEY"
	// EnvAWSRegion name of the environment variable that contains the AWS region.
	EnvAWSRegion = "AWS_REGION"
)

// GetAWSAccessKeyID returns the AWS access key ID.
func GetAWSAccessKeyID() string {
	return GetString(EnvAWSAccessKeyID)
}

// GetAWSBucket returns the AWS bucket.
func GetAWSBucket() string {
	return GetString(EnvAWSBucket)
}

// GetAWSPath returns the AWS path.
func GetAWSPath() string {
	return GetString(EnvAWSPath)
}

// GetAWSRegion returns the AWS region.
func GetAWSRegion() string {
	return GetString(EnvAWSRegion)
}

// GetAWSSecretAccessKey returns the AWS access key secret.
func GetAWSSecretAccessKey() string {
	return GetString(EnvAWSSecretAccessKey)
}
