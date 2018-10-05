.PHONY: entrypoint restart ps logs

PROJECT = gs-server
BIN = $(PROJECT)

SRC_PATH = .
SRC = $(SRC_PATH)/...

COVERAGE = coverage.txt

all: build

# ================== Settings ==================
.PHONY: get get-dev install install-dev

_copy:
	@cp go.mod.prod go.mod

get:
	@go mod download

get-dev: get
	@go get -u -v github.com/golangci/golangci-lint/cmd/golangci-lint@v1.10.2
	@go get -u -v github.com/githubnemo/CompileDaemon@v1.0.0

install: _copy get

install-dev: _copy get-dev

# ================== Server ==================
.PHONY: build run run-dev lint test inspect

build:
	@go build

run: build
	@./$(BIN)

run-dev:
	@CompileDaemon -build="make build" -command="make run"

lint:
	@golangci-lint run --enable-all $(SRC_PATH)

test:
	@go test -v --race --covermode=atomic --coverprofile=$(COVERAGE) $(SRC)

inspect: lint test
