#!/bin/bash

env GOOS=linux GOARCH=arm GOARM=7 go build -ldflags "-X main.commit `git rev-parse --short HEAD` -X main.builtAt `date +%FT%T%z`" -o bin/rpimonitor monitor/cmd/main.go
