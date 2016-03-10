#!/bin/bash

env GOOS=linux GOARCH=arm GOARM=7 go build -o bin/rpimonitor monitor/cmd/main.go
