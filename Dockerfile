FROM golang:1.11.0-alpine3.8

RUN apk add --update --no-cache \
    git=2.18.0-r0 \
    build-base=0.5-r1

ARG PROJECT_DIR=/gs-api

WORKDIR ${PROJECT_DIR}
COPY go.mod.prod ${PROJECT_DIR}/go.mod
COPY Makefile ${PROJECT_DIR}/Makefile

RUN make get-dev
