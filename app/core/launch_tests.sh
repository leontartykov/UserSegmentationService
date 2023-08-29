#!/bin/bash

go test -v ./server/repository
go test -v ./server/e2e_test.go