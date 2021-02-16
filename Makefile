# GoLang Commands

GOCMD ?= go
GORUN := ${GOCMD} run
GOBUILD := ${GOCMD} build
GOBUILDFLAGS ?= -ldflags="-s -w"
GOCLEAN := ${GOCMD} clean
GOTEST := ${GOCMD} test -v -race
GOGET := ${GOCMD} get
GOFMT ?= gofmt
LINTER ?= golangci-lint

CONTAINERCMD ?= podman

# Project configuration

VERSION := $(shell git describe --tags --exact-match 2>/dev/null || echo latest)
DOCKERHUB_NAMESPACE ?= sumelms
MICROSERVICE_NAME := account
BINARY_NAME := sumelms-${MICROSERVICE_NAME}
IMAGE := ${DOCKERHUB_NAMESPACE}/microservice-${MICROSERVICE_NAME}:${VERSION}

##############################################################

all: test build

# Runner

run:
	export SUMELMS_CONFIG_PATH="./config/config.yml" && \
	${GORUN} cmd/server/main.go
.PHONY: run

# Builders

build: build-proto
	${GOBUILD} ${GOBUILDFLAGS} -o bin/${BINARY_NAME} cmd/server/main.go
.PHONY: build

build-proto:
	protoc proto/**/*.proto --go_out=plugins=grpc:.
.PHONY: build-proto

# Quality tools

test:
	${GOTEST} $$(go list ./... | grep -v /test/) $(TEST_OPTIONS)
.PHONY: test

lint:
	${LINTER} run

format:
	${GOFMT} -d .

# Container stuff (podman/docker)

container-build:
	${CONTAINERCMD} build -t ${IMAGE} .

container-push: container-build
	${CONTAINERCMD} push ${IMAGE}

container-run: container-build
	${CONTAINERCMD} run -p 8080:8080 ${IMAGE}