#!/bin/bash

set -ebu

# generate code first
# generates the stubs
# generates the graphql server and resolvers
go generate ./...

#run the docker-compose and exit the script with the integration test containers result
docker-compose -f ./integrationTests/docker-compose.yaml up --build --exit-code-from integration_test_runner