.DEFAULT_GOAL := default

.PHONY: build
build:
	go build

.PHONY: install
install:
	go install

.PHONY: check
check:
	go vet .
	golint .
	golint main.go

.PHONY: clean
clean:
	go clean

.PHONY: run
run:
	go run *.go

default:
	make check
	make build
