.PHONY: build
build:
	go build -v ./cmd/apiserver

.PHONY: test
test:
	make build
	go test -v -race -timeout 30s ./...

.DEFAULT_GOAL := build
