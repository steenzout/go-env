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
	// EnvAWSCABundle
	EnvAWSCABundle = "AWS_CA_BUNDLE"
	// EnvAWSConfigFile
	EnvAWSConfigFile = "AWS_CONFIG_FILE"
	// EnvAWSOutput
	EnvAWSOutput = "AWS_DEFAULT_OUTPUT"
	// EnvAWSPath path in the bucket.
	EnvAWSPath = "AWS_PATH"
	// EnvAWSProfile
	EnvAWSProfile = "AWS_PROFILE"
	// EnvAWSRegion name of the environment variable that contains the AWS region.
	EnvAWSRegion = "AWS_DEFAULT_REGION"
	// EnvAWSSharedCredentialsFile
	EnvAWSSharedCredentialsFile = "AWS_SHARED_CREDENTIALS_FILE"
	// EnvAWSSecretAccessKey name of the environment variable that contains the AWS access key secret.
	EnvAWSSecretAccessKey = "AWS_SECRET_ACCESS_KEY"
	// EnvAWSSessionToken
	EnvAWSSessionToken = "AWS_SESSION_TOKEN"
)

// GetAWSAccessKeyID returns the AWS access key ID.
func GetAWSAccessKeyID() string {
	return GetString(EnvAWSAccessKeyID)
}

// GetAWSBucket returns the AWS bucket.
func GetAWSBucket() string {
	return GetString(EnvAWSBucket)
}

// GetAWSCABundle returns the path to the certificate bundle.
func GetAWSCABundle() string {
	return GetString(EnvAWSCABundle)
}

// GetAWSConfigFile returns the path to the configuration file.
func GetAWSConfigFile() string {
	return GetString(EnvAWSConfigFile)
}

// GetAWSOutput returns the output format.
func GetAWSOutput() string {
	return GetOptionalString(EnvAWSOutput, "json")
}

// GetAWSPath returns the AWS path.
func GetAWSPath() string {
	return GetString(EnvAWSPath)
}

// GetAWSProfile returns the name of the profile.
func GetAWSProfile() string {
	return GetOptionalString(EnvAWSProfile, "default")
}

// GetAWSRegion returns the AWS region.
func GetAWSRegion() string {
	return GetString(EnvAWSRegion)
}

// GetAWSSecretAccessKey returns the AWS access key secret.
func GetAWSSecretAccessKey() string {
	return GetString(EnvAWSSecretAccessKey)
}

// GetAWSSessionToken temporary session token.
func GetAWSSessionToken() string {
	return GetString(EnvAWSSessionToken)
}

// GetAWSSharedCredentialsFile returns path to the file where access keys are stored.
func GetAWSSharedCredentialsFile() string {
	return GetString(EnvAWSSharedCredentialsFile)
}
