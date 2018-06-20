#!/bin/bash

source docker.env
export \
    MYSQL_DATABASE=test \
    MYSQL_HOST=127.0.0.1 \
    MYSQL_PASSWORD= \
    MYSQL_PORT=3306 \
    MYSQL_PROTOCOL=tcp \
    MYSQL_ROOT_PASSWORD= \
    MYSQL_USER=travis
