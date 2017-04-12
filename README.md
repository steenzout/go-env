# go-env

[![License](https://img.shields.io/badge/license-Apache%20License%202.0-blue.svg?style=flat)][license]
[![Build Status](https://travis-ci.org/steenzout/go-env.svg?branch=master)](https://travis-ci.org/steenzout/go-env/)
[![Coverage Status](https://coveralls.io/repos/steenzout/go-env/badge.svg?branch=master&service=github)](https://coveralls.io/github/steenzout/go-env?branch=master)

Library that uses convention to read configuration paramaters from the environment.


## environment variables

## app:Redis

- **REDIS_DB**: environment variable that contains the Redis database.
- **REDIS_PASSWORD**: environment variable that contains the Redis password.
- **REDIS_PORT**: environment variable that contains the Redis port.
- **REDIS_HOST**: environment variable that contains the Redis host.


## cloud:AWS

- **AWS_ACCESS_KEY_ID**: environment variable that contains the AWS access key ID.
- **AWS_SECRET_ACCESS_KEY**: environment variable that contains the AWS access key secret.
- **AWS_REGION**: environment variable that contains the AWS region.


[license]:  https://raw.githubusercontent.com/steenzout/go-env/master/LICENSE   "Apache License 2.0"
