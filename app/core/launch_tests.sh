#!/bin/bash

docker exec -it app_backend go test -v ./server/repository
docker exec -it app_backend go test -v ./server/e2e_test.go