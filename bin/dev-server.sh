#!/bin/bash -xe

cd "$(dirname "$0")/.."
ROOT=$(pwd)

cd server
dep ensure
go get github.com/cespare/reflex
reflex -s go run cmd/server.go
