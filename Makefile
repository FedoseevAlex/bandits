.PHONY: all proto_generate lint lint-new

BIN := "bandits.app"
PACKAGE_PATH := "github.com/FedoseevAlex/bandits"
LDFLAGS := -X $(PACKAGE_PATH)/internal/version.gitHash=$(shell git rev-parse --short HEAD)
LDFLAGS += -X $(PACKAGE_PATH)/internal/version.buildDate=$(shell date -u +%Y-%m-%dT%H:%M:%S)

proto_generate:
	buf generate

build: proto_generate
	go build -v -o $(BIN) -ldflags "$(LDFLAGS)" ./main.go

lint: 
	golangci-lint run ./...
	buf lint

lint-new: 
	golangci-lint run --new ./...
	buf lint