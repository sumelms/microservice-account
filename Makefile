VERSION := $(shell git describe --tags --exact-match 2>/dev/null || echo latest)
DOCKERHUB_NAMESPACE ?= sumelms
IMAGE := ${DOCKERHUB_NAMESPACE}/user:${VERSION}

run:
	export SUMELMS_CONFIG_PATH="./config/config.yml" && \
	go run cmd/server/main.go

build:
	go build -o bin/sumelms-user cmd/server/main.go

test:
	go test

docker-build:
	docker build -t ${IMAGE} .

docker-push: docker-build
	docker push ${IMAGE}

docker-run: docker-build
	docker run -p 8080:8080 ${IMAGE}