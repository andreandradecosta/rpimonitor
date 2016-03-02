#!/bin/bash
go run rpimonitor.go -StartServer=true -HTTP_PORT=8081 -HTTPS_PORT=8443 -IsDevelopment=true -SAMPLE_INTERVAL=10s
