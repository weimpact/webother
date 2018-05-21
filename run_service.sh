#!/bin/bash
source .env
#golint ./...
rm ./cmd/server/server
go vet ./...
pushd ./cmd/server/
echo "building and running"
go build 
./server
popd
