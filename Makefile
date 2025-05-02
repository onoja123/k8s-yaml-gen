# Makefile for k8s-yaml-gen

BINARY_NAME := k8s-yaml-gen
BUILD_DIR   := bin
VERSION     := $(shell git describe --tags --always --dirty)
GOFILES     := $(shell find . -type f -name '*.go' -not -path "./vendor/*")

.PHONY: all build fmt lint test clean run release

all: build

## build the binary for the host OS/ARCH
build:
	@echo "→ Building $(BINARY_NAME) (version: $(VERSION))"
	@go build -ldflags="-X main.version=$(VERSION)" -o $(BUILD_DIR)/$(BINARY_NAME) ./cmd/$(BINARY_NAME)

## format all Go code
fmt:
	@echo "→ go fmt"
	@go fmt $(GOFILES)

## lint (basic vet)
lint:
	@echo "→ go vet"
	@go vet ./...

## run tests with verbose output
test:
	@echo "→ go test"
	@go test ./... -v

## remove artifacts
clean:
	@echo "→ Cleaning"
	@rm -rf $(BUILD_DIR) coverage.out

## run the CLI against flags
run: build
	@echo "→ Running $(BINARY_NAME)…"
	@$(BUILD_DIR)/$(BINARY_NAME) generate --replicas 2 --memory 256Mi --port 31000

## build for all major platforms via goreleaser
release:
	@echo "→ Releasing via goreleaser"
	@goreleaser --rm-dist

## generate coverage report
coverage:
	@echo "→ Running tests with coverage"
	@go test ./... -coverprofile=coverage.out
	@go tool cover -html=coverage.out

