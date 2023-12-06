SHELL := /bin/bash

.PHONY: \
	help \
	clean \
	tools \
	deps \
	test \
	errors \
	lint \
	imports \
	fmt \
	version

all: imports fmt lint errors build

help:
	@echo 'Usage: make <OPTIONS> ... <TARGETS>'
	@echo ''
	@echo 'Available targets are:'
	@echo ''
	@echo '    help               Show this help screen.'
	@echo '    clean              Remove binaries and artifacts.'
	@echo '    tools              Install tools needed by the project.'
	@echo '    deps               Download and install build time dependencies.'
	@echo '    test               Run unit tests.'
	@echo '    errors             Run errcheck.'
	@echo '    lint               Run golint.'
	@echo '    imports            Run goimports.'
	@echo '    fmt                Run go fmt.'
	@echo '    version            Display Go version.'
	@echo ''
	@echo 'Targets run by default are: imports, fmt, lint, vet, errors and build.'
	@echo ''

deps:
	go mod tidy

clean:
	go clean -i ./...
	rm -vf \
	  $(CURDIR)/coverage.* \
	  $(CURDIR)/$(TARGET)*

tools:
	go install golang.org/x/tools/cmd/goimports@latest
	go install github.com/kisielk/errcheck@latest
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest

test:
	go test ./... -covermode=count -coverprofile=coverage.out

fmt:
	gofmt -s -w .

imports:
	goimports -l -w .

errors:
	errcheck -ignoretests -blank ./...

lint:
	go run github.com/golangci/golangci-lint/cmd/golangci-lint@latest --timeout 300s  run ./...

generate:
	go generate ./...

version:
	@go version