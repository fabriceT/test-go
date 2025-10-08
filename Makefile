.PHONY: build test version

build:
	go build -ldflags="-X main.version=$(shell git rev-parse --short HEAD)" -o app .

test:
	go test -v -cover ./...

version:
	@echo "Version: $(shell git rev-parse --short HEAD)"
