.PHONY: default all build run fmt test vet clean

default: all

all: fmt build vet test run

build:
	go build -o leetcode .

run:
	go run .

fmt:
	go fmt ./...

test:
	go test ./...

vet:
	go vet ./...

clean:
	rm -f leetcode
