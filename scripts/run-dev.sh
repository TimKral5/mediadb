#!/bin/sh

cd /src
go test -v ./...
./mediadb

