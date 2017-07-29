#!/bin/bash

#################################################################
# Script to run functional tests
# usage: ./run_tests.sh
#################################################################
set -o errexit

MY_DIR=$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )
cd "$MY_DIR"

echo "RUNNING TESTS"

docker build -t "docker.uacf.io/svc_friends/friendsapi:CIVERSION" -f ../docker/friendsapi/Dockerfile ../

docker-compose -f functional.yml \
	run --rm -e "GOCONVEY_REPORTER=story" test \
	go test -v friends.uacf.io/...
