# go-env

[![License](https://img.shields.io/badge/license-Apache%20License%202.0-blue.svg?style=flat)][license]
[![Build Status](https://travis-ci.org/steenzout/go-env.svg?branch=master)](https://travis-ci.org/steenzout/go-env/)
[![Coverage Status](https://coveralls.io/repos/steenzout/go-env/badge.svg?branch=master&service=github)](https://coveralls.io/github/steenzout/go-env?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/steenzout/go-env)](https://goreportcard.com/report/github.com/steenzout/go-env)

[![Project](https://www.openhub.net/p/go-steenzout-env/widgets/project_thin_badge.gif)][project]

Library that uses convention to read configuration parameters from the environment.

## environment variables

### app:InfluxDB

- **INFLUXDB_DB**: name of the InfluxDB database.
- **INFLUXDB_HOST**:  host.
- **INFLUXDB_PORT**: HTTP API port.
- **INFLUXDB_ADMIN_PORT**: administrator interface port.
- **INFLUXDB_GRAPHITE_PORT**: Graphite support port.
- **INFLUXDB_ADMIN_USER**: administrator account.
- **INFLUXDB_ADMIN_PASSWORD**: administrator password.
- **INFLUXDB_USER**: read/write database account.
- **INFLUXDB_USER_PASSWORD**: password.
- **INFLUXDB_READ_USER**: database reader account.
- **INFLUXDB_READ_USER_PASSWORD**: reader password.
- **INFLUXDB_WRITE_USER**: database writer account.
- **INFLUXDB_WRITE_USER_PASSWORD**: writer password.

### app:MySQL

- **MYSQL_DATABASE**: name of the MySQL database.
- **MYSQL_HOST**: host or path to Unix domain socket (default: `/var/run/mysqld/mysqld.sock`)
- **MYSQL_PASSWORD**: password (default: `''`)
- **MYSQL_PORT**: port (default: `3306`).
- **MYSQL_PROTOCOL**: protocol to establish connection (default: `unix`)
- **MYSQL_ROOT_PASSWORD**: root password (default: `''`)
- **MYSQL_USER**: user name.

- [github.com/go-sql-driver/mysql](https://github.com/go-sql-driver/mysql)
- [hub.docker.com/_/mysql](https://hub.docker.com/_/mysql/)

### app:PostgreSQL

- **POSTGRES_DB**: name of the database.
- **POSTGRES_HOST**: host.
- **POSTGRES_PASSWORD**: password.
- **POSTGRES_PORT**: port (default: `5432`).
- **POSTGRES_USER**: user name.

### app:Redis

- **REDIS_DB**: name of the database.
- **REDIS_PASSWORD**: password.
- **REDIS_PORT**: port (default: `6379`).
- **REDIS_HOST**: host.

### app:Resque

- **RESQUE_COUNT**: number of Resque workers to be spawned (default: `1`).
- **RESQUE_INTERVAL**: polling frequency, in seconds (default: `5s`).
- **RESQUE_PIDFILE**: PID file location..
- **RESQUE_QUEUE**: name of a single Resque queue.
- **RESQUE_QUEUES**: list of Resque queue names (default: `*`).

### cloud:AWS

- **AWS_ACCESS_KEY_ID**: AWS access key ID.
- **AWS_BUCKET**: name of the bucket.
- **AWS_CA_BUNDLE**: path to the certificate bundle.
- **AWS_CONFIG_FILE**: path to the file where configuration profiles are stored.
- **AWS_PATH**: path in the bucket.
- **AWS_DEFAULT_REGION**: AWS region.
- **AWS_DEFAULT_OUTPUT**: output format (default: `json`).
- **AWS_PROFILE**: name of the profile (default: `default`).
- **AWS_SECRET_ACCESS_KEY**: the AWS access key secret.
- **AWS_SESSION_TOKEN**: temporary session token (default: ``).
- **AWS_SHARED_CREDENTIALS_FILE**: path to the file where access keys are stored.

- [AWS Documentation » AWS Command Line Interface » User Guide » Configuring the AWS CLI » Environment Variables](http://docs.aws.amazon.com/cli/latest/userguide/cli-environment.html)

[license]:  https://raw.githubusercontent.com/steenzout/go-env/master/LICENSE   "Apache License 2.0"
[project]:  https://www.openhub.net/p/go-steenzout-env/    "OpenHub project page"
