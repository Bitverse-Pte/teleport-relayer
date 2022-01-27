#!/usr/bin/make -f
BINARY = teleport-relayer
VERSION ?= $(shell echo $(shell git describe --tags `git rev-list --tags="v*" --max-count=1`) | sed 's/^v//')
COMMIT := $(shell git log -1 --format='%H')
build:
ifeq ($(OS),Windows_NT)
	go build -ldflags '$(ldflags)' -o build/relayer.exe .
else
	go build -ldflags '$(ldflags)' -o build/relayer .
endif

build-linux: go.sum
	LEDGER_ENABLED=false GOOS=linux GOARCH=amd64 $(MAKE) build

go.sum: go.mod
	@echo "--> Ensure dependencies have not been modified"
	@go mod verify

install:
	go build -ldflags '$(ldflags)'  -o relayer && mv relayer $(GOPATH)/bin

format:
	find . -name '*.go' -type f -not -path "./vendor*" -not -path "*.git*" | xargs gofmt -w -s
	find . -name '*.go' -type f -not -path "./vendor*" -not -path "*.git*" | xargs misspell -w
	find . -name '*.go' -type f -not -path "./vendor*" -not -path "*.git*" -not -path "./lite/*/statik.go" -not -path "*.pb.go" | xargs goimports -w -local github.com/teleport-network/teleport-relayer


setup: build-linux
	@docker build -ldflags '$(ldflags)'  -t relayer .
	@rm -rf ./build

ldflags = -X github.com/teleport-network/teleport-relayer/version.Name=teleport-relayer \
		  -X github.com/teleport-network/teleport-relayer/version.AppName=$(BINARY) \
		  -X github.com/teleport-network/teleport-relayer/version.Version=$(VERSION) \
		  -X github.com/teleport-network/teleport-relayer/version.Commit=$(COMMIT)

