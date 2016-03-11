#!/bin/bash

go build -ldflags "-X main.commit=`git rev-parse --short HEAD` -X main.builtAt=`date +%FT%T%z`" -o bin/rpiserver server/cmd/main.go
