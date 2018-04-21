#!/bin/bash -e
GOPATH=$HOME
go generate ./...
go test -cover -coverprofile /tmp/c.out
go tool cover -o /tmp/coverage.html -html /tmp/c.out
gofmt -w .
