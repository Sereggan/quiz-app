.PHONY: build
build:
	go build -o server.exe -v ./cmd

.PHONY: test
test:
	go test -v -race -timeout 30s ./...

.DEFAULT_GOAL := build
