.PHONY: all protogen lint lint-new

BIN := "bandits.app"
PACKAGE_PATH := "github.com/FedoseevAlex/bandits"
LDFLAGS := -X $(PACKAGE_PATH)/internal/version.gitHash=$(shell git rev-parse --short HEAD)
LDFLAGS += -X $(PACKAGE_PATH)/internal/version.buildDate=$(shell date -u +%Y-%m-%dT%H:%M:%S)

protogen:
	protoc --go_out=. proto/*.proto

build:
	go build -v -o $(BIN) -ldflags "$(LDFLAGS)" ./main.go

lint: 
	golangci-lint run ./...

lint-new: 
	golangci-lint run --new ./...