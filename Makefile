.PHONY: default all build run fmt test lint clean

default: all

all: fmt build lint test run

build:
	go build -o leetcode .

run:
	go run .

fmt:
	go fmt ./...

test:
	go test ./...

lint:
	go run github.com/golangci/golangci-lint/v2/cmd/golangci-lint@v2.14.0 run ./...

clean:
	rm -f leetcode
