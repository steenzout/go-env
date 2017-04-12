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
