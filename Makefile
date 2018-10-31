.PHONY: entrypoint restart ps logs

PROJECT = gs-server
BIN = $(PROJECT)

SRC_PATH = .
SRC = $(SRC_PATH)/...

COVERAGE = coverage.txt

all: build

# ================== Settings ==================

.PHONY: get get-dev install install-dev db

_copy:
	@cp go.mod.prod go.mod

get:
	@go get -u -v .

get-dev: get
	@go get -u -v github.com/golangci/golangci-lint/cmd/golangci-lint@v1.10.2
	@go get -u -v github.com/go-playground/justdoit@v0.0.0-20180413125108-f398c5dd9bd7

install: _copy get

install-dev: _copy get-dev

db:
	@go run provider/init/main.go

# ================== Server ==================

.PHONY: build run run-dev lint test inspect

build:
	@go build -o $(BIN)

run: build
	@./$(BIN)

run-dev:
	@justdoit -build="make build" -run="make run"

lint:
	@golangci-lint run --enable-all $(SRC_PATH)

test:
	@go test -v --race --covermode=atomic --coverprofile=$(COVERAGE) $(SRC)

inspect: lint test
