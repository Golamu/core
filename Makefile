.PHONY: build install test

build:
	go build ./...

install:
	go install ./...

test:
	go test ./...
