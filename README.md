# go-env

[![License](https://img.shields.io/badge/license-Apache%20License%202.0-blue.svg?style=flat)][license]
[![Build Status](https://travis-ci.org/steenzout/go-env.svg?branch=master)](https://travis-ci.org/steenzout/go-env/)
[![Coverage Status](https://coveralls.io/repos/steenzout/go-env/badge.svg?branch=master&service=github)](https://coveralls.io/github/steenzout/go-env?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/steenzout/go-env)](https://goreportcard.com/report/github.com/steenzout/go-env)

Library that uses convention to read configuration parameters from the environment.


## environment variables

### app:MySQL

- **MYSQL_DB**: environment variable that contains the MySQL database.
- **MYSQL_HOST**: environment variable that contains the MySQL host.
- **MYSQL_PASSWORD**: environment variable that contains the MySQL user's password.
- **MYSQL_PORT**: environment variable that contains the MySQL port (default: `3306`).
- **MYSQL_USER**: environment variable that contains the MySQL user.

### app:PostgreSQL

- **POSTGRESQL_DB**: environment variable that contains the PostgreSQL database.
- **POSTGRESQL_HOST**: environment variable that contains the PostgreSQL host.
- **POSTGRESQL_PASSWORD**: environment variable that contains the PostgreSQL user's password.
- **POSTGRESQL_PORT**: environment variable that contains the PostgreSQL port (default: `5432`).
- **POSTGRESQL_USER**: environment variable that contains the PostgreSQL user.

### app:Redis

- **REDIS_DB**: environment variable that contains the Redis database.
- **REDIS_PASSWORD**: environment variable that contains the Redis password.
- **REDIS_PORT**: environment variable that contains the Redis port (default: `6379`).
- **REDIS_HOST**: environment variable that contains the Redis host.


### app:Resque

- **RESQUE_COUNT**: environment variable that contains the number of Resque workers to be spawned (default: `1`).
- **RESQUE_INTERVAL**: environment variable to define, in seconds, the polling frequency (default: `5s`).
- **RESQUE_PIDFILE**: environment variable that contains the PID file location..
- **RESQUE_QUEUE**: environment variable that contains the name of a single Resque queue.
- **RESQUE_QUEUES**: environment variable that contains a list of Resque queue names (default: `*`).


### cloud:AWS

- **AWS_ACCESS_KEY_ID**: environment variable that contains the AWS access key ID.
- **AWS_BUCKET**: name of the bucket.
- **AWS_PATH**: path in the bucket.
- **AWS_REGION**: environment variable that contains the AWS region.
- **AWS_SECRET_ACCESS_KEY**: environment variable that contains the AWS access key secret.


[license]:  https://raw.githubusercontent.com/steenzout/go-env/master/LICENSE   "Apache License 2.0"
