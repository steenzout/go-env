package env

const (
	// EnvAWSAccessKeyID name of the environment variable that contains the AWS access key ID configuration parameter.
	EnvAWSAccessKeyID = "AWS_ACCESS_KEY_ID"
	// EnvAWSSecretAccessKey name of the environment variable that contains the AWS access key secret configuration parameter.
	EnvAWSSecretAccessKey = "AWS_SECRET_ACCESS_KEY"
	// EnvAWSRegion name of the environment variable that contains the AWS region configuration parameter.
	EnvAWSRegion = "AWS_REGION"
)

// GetAWSAccessKeyID returns the AWS access key ID configuration parameter.
func GetAWSAccessKeyID() string {
	return GetString(EnvAWSAccessKeyID)
}

// GetAWSSecretAccessKey returns the AWS access key secret configuration parameter.
func GetAWSSecretAccessKey() string {
	return GetString(EnvAWSSecretAccessKey)
}

// GetAWSRegion returns the AWS region configuration parameter.
func GetAWSRegion() string {
	return GetString(EnvAWSRegion)
}
