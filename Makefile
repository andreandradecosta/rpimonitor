.PHONY: build lint test vet install bench cover

PKGS=$(shell go list ./... | grep -v /vendor/)

LDFLAGS=-ldflags "-X main.commit=`git rev-parse --short HEAD` -X main.builtAt=`date +%FT%T%z`"

all: install test vet build build_arm

install:
	go get golang.org/x/tools/cmd/cover
	go get github.com/mattn/goveralls
	go get github.com/go-playground/overalls

lint:
	@go list ./... | grep -v /vendor/ | xargs -L1 golint

test:
	@go test -v $(PKGS)

cover:
	overalls -project=github.com/andreandradecosta/rpimonitor -covermode=count -ignore=.git,vendor,cmd -debug
	goveralls -coverprofile=overalls.coverprofile -service=travis-ci

bench:
	@go test -v $(PKGS) -bench=.

vet:
	@go vet -v $(PKGS)

build:
	@go build ${LDFLAGS} -o bin/rpimonitor cmd/rpimonitor/main.go

build_arm:
	@env GOOS=linux GOARCH=arm GOARM=7 go build ${LDFLAGS} -o bin/rpimonitor cmd/rpimonitor/main.go

run:
	go run cmd/rpimonitor/main.go -config=.env
