VERSION := $(shell git describe --tags --exact-match 2>/dev/null || echo latest)
DOCKERHUB_NAMESPACE ?= sumelms
IMAGE := ${DOCKERHUB_NAMESPACE}/microservice-account:${VERSION}

run: build-proto
	export SUMELMS_CONFIG_PATH="./config/config.yml" && \
	go run cmd/server/main.go
.PHONY: run

build: build-proto
	go build -o bin/sumelms-account cmd/server/main.go
.PHONY: build

test-unit:
	go test $$(go list ./... | grep -v /test/) $(TEST_OPTIONS)
.PHONY: test-unit

build-proto:
	protoc proto/**/*.proto --go_out=plugins=grpc:.
.PHONY: build-proto

lint:
	golint $$(go list ./... | grep -v /vendor/)

docker-build:
	docker build -t ${IMAGE} .

docker-push: docker-build
	docker push ${IMAGE}

docker-run: docker-build
	docker run -p 8080:8080 ${IMAGE}