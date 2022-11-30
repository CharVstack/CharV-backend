.DEFAULT_GOAL := help
BINDIR := bin
GOBIN := $(abspath tools)

ROOT_PACKAGE := $(shell go list .)
COMMAND_PACKAGES := $(shell go list ./cmd/...)
BINARIES:=$(COMMAND_PACKAGES:$(ROOT_PACKAGE)/cmd/%=$(BINDIR)/%)
GO_FILES:=$(shell find . -type f -name '*.go' -print)

VERSION:=$(shell cat VERSION)
REVISION:=$(shell git rev-parse --short HEAD)

# ldflag
GO_LDFLAGS_VERSION:=-X '${ROOT_PACKAGE}.VERSION=${VERSION}' -X '${ROOT_PACKAGE}.REVISION=${REVISION}'
GO_LDFLAGS:=$(GO_LDFLAGS_VERSION)

# go build
GO_BUILD:=-ldflags "$(GO_LDFLAGS)"

GOFLAGS:="-tags=develop"

.PHONY: help fmt lint test build coverage

help:
	@cat $(MAKEFILE_LIST) | \
	    perl -ne 'print if /^\w+.*##/;' | \
	    perl -pe 's/(.*):.*##\s*/sprintf("%-20s",$$1)/eg;'

get: ## Update dependencies
	go get -u

tidy: ## Optimize go.mod and go.sum
	go mod tidy

fmt: tools ## Format Code
	tools/goimports -w ./

lint: fmt tidy ## Lint Code
	go vet ./...

test: testassets ## Run Test
	go test -v ./...

testassets: test/resources/image/bad.qcow2 test/resources/image/ok.qcow2 ## Generate Test Assets

test/resources/image/bad.qcow2:
	head -c 1024 /dev/urandom > test/resources/image/bad.qcow2

test/resources/image/ok.qcow2:
	qemu-img create -f qcow2 test/resources/image/ok.qcow2 4G

tools: tools/goimports tools/air # Install Tools

tools/goimports:
	GOBIN=$(GOBIN) go install golang.org/x/tools/cmd/goimports@v0.3.0

tools/air:
	GOBIN=$(GOBIN) go install github.com/cosmtrek/air@v1.40.4

dev: tools # Run Development Server
	tools/air

build: $(BINARIES) ## Build Server Binary

$(BINARIES): $(GO_FILES) VERSION .git/HEAD
	@go build $(GOFLAGS) -o $@ $(GO_BUILD) $(@:$(BINDIR)/%=$(ROOT_PACKAGE)/cmd/%)

coverage: testassets tools # Generate Coverage
	go test -cover -coverprofile=coverage.out ./...
