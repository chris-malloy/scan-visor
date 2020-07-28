#!/usr/bin/env bash
docker-compose -f ./deployments/local/docker-compose.local.yml up -d

./scripts/wait-for-it.sh -s -t 150 localhost:9000

go test -v ./test/integration/...

docker-compose -f ./deployments/local/docker-compose.local.yml down