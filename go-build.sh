#!/bin/sh
go env -w GOOS=linux
go env -w GOARCH=amd64
go build -o docker-build/