.PHONY: build lint test vet install bench cover

PKGS=$(shell go list ./... | grep -v /vendor/)
WATCH_FILES=find . -type f -not -path '*/\.*' | grep -i '.*[.]go' 2> /dev/null

LDFLAGS=-ldflags "-X main.commit=`git rev-parse --short HEAD` -X main.builtAt=`date +%FT%T%z`"

all: install test vet build build_arm

install:
	go get -u golang.org/x/tools/cmd/cover
	go get -u github.com/mattn/goveralls
	go get -u github.com/go-playground/overalls
	go get -u github.com/kyoh86/richgo

lint:
	@go list ./... | grep -v /vendor/ | xargs -L1 golint

test:
	@richgo test -v $(PKGS)

watch_test:
	@if command -v entr > /dev/null; then ${WATCH_FILES} | entr -c $(MAKE) test; \
	else $(MAKE) test entr_warn; \
	fi


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

start:
	go run cmd/rpimonitor/main.go -config=.env

watch_start:
	@if command -v entr > /dev/null; then ${WATCH_FILES} | entr -r -c $(MAKE) start; \
	else $(MAKE) run entr_warn; \
	fi

entr_warn:
	@echo "----------------------------------------------------------"
	@echo "     ! File watching functionality non-operational !      "
	@echo "                                                          "
	@echo "Install entr(1) to automatically run tasks on file change."
	@echo "See http://entrproject.org/                               "
	@echo "----------------------------------------------------------"
