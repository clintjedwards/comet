APP_NAME = comet
GIT_COMMIT := $(shell git rev-parse --short HEAD)
SHELL = /bin/bash
VERSION=$(shell date +%s)


GO_LDFLAGS := '-X "github.com/clintjedwards/${APP_NAME}/cmd.appVersion=$(VERSION) $(GIT_COMMIT)" \
			   -X "github.com/clintjedwards/${APP_NAME}/service.appVersion=$(VERSION) $(GIT_COMMIT)"'

## build-protos: build required protobuf files
build-protos:
	protoc --go_out=plugins=grpc:. proto/*.proto
	protoc --go_out=plugins=grpc:. backend/proto/*.proto

## build: run tests and compile application
build: check-path-included
	protoc --go_out=plugins=grpc:. proto/*.proto
	protoc --go_out=plugins=grpc:. backend/proto/*.proto
	go mod tidy
	go test ./utils
	go build -ldflags $(GO_LDFLAGS) -o $(path)

## run: build application and run server
run:
	protoc --go_out=plugins=grpc:. proto/*.proto
	protoc --go_out=plugins=grpc:. backend/proto/*.proto
	go mod tidy
	go build -ldflags $(GO_LDFLAGS) -o /tmp/${APP_NAME} && /tmp/${APP_NAME} server

## install: build application and install on system
install:
	protoc --go_out=plugins=grpc:. proto/*.proto
	protoc --go_out=plugins=grpc:. backend/proto/*.proto
	go mod tidy
	go build -ldflags $(GO_LDFLAGS) -o /tmp/${APP_NAME}
	sudo mv /tmp/${APP_NAME} /usr/local/bin/
	chmod +x /usr/local/bin/${APP_NAME}

## help: prints this help message
help:
	@echo "Usage: "
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

check-path-included:
ifndef path
	$(error path is undefined; ex. path=/tmp/${APP_NAME})
endif
