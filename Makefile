NAME := app

# Docker
DOCKER_APP_FILENAME ?= deployments/docker/Dockerfile
DOCKER_COMPOSE_FILE ?= deployments/docker-compose/docker-compose.yml

# Build
BUILD_CMD ?= CGO_ENABLED=0 go build -o bin/${NAME} -ldflags '-v -w -s' ./cmd/${NAME}

.PHONY: run
run: 
	go run cmd/main.go 

.PHONY: build
build:
	echo "building"
	${BUILD_CMD}
	echo "build done"

.PHONY: lint
lint:
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.29.0
	./bin/golangci-lint run -v

.PHONY: check
check: lint test
	echo "check done"

.PHONY: test
test:
	go test ./... -cover