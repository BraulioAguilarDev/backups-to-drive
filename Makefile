# Binary names
BINARY_NAME=go-backup
BINARY_UNIX=$(BINARY_NAME)_unix
BINARY_OSX=$(BINARY_NAME)_osx
GOFILES=$(wildcard *.go)

build:
	@echo "Building $(GOFILES) to ./bin"
	GOOS=linux GOARCH=amd64 go build -v -o bin/$(BINARY_UNIX)
	GOOS=darwin GOARCH=amd64 go build -v -o bin/$(BINARY_OSX)
clear:
	rm -rf ./bin

.PHONY: build clean
