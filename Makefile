.PHONY: all get get-dev build run run-dev lint test inspect
.PHONY: entrypoint restart ps logs

PROJECT = goldensplit-server
BIN = $(PROJECT)

SRC_PATH = .
SRC = $(SRC_PATH)/...

COVERAGE = coverage.txt

all: build

_copy:
	@cp go.mod.prod go.mod

_copy-dev: _copy
	@cp docker-compose.override.yml.dev docker-compose.override.yml

get:
	@go mod download

get-dev: get
	@go get -u -v github.com/golangci/golangci-lint/cmd/golangci-lint@v1.10.2
	@go get -u -v github.com/githubnemo/CompileDaemon@v1.0.0

install: _copy get

install-dev: _copy-dev get-dev

build:
	@go build

run:
	@./$(BIN)

run-dev:
	@CompileDaemon -build="make build" -command="make run"

lint:
	@golangci-lint run --enable-all $(SRC_PATH)

test:
	@go test -v --race --covermode=atomic --coverprofile=$(COVERAGE) $(SRC)

inspect: lint test

# ========= Docker =========

entrypoint: build run

restart:
	@docker-compose up --remove-orphans -d --force-recreate $(c)

ps:
	@docker-compose ps

logs:
	@docker-compose logs -f $(c)
