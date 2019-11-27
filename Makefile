SHELL = /bin/bash
VERSION=$(shell date +%s)
GIT_COMMIT := $(shell git rev-parse --short HEAD)


GO_LDFLAGS := '-X "github.com/clintjedwards/comet/cmd.appVersion=$(VERSION) $(GIT_COMMIT)" \
			   -X "github.com/clintjedwards/comet/service.appVersion=$(VERSION) $(GIT_COMMIT)"'

build-protos:
	protoc --go_out=plugins=grpc:. proto/*.proto

build: check-path-included
	protoc --go_out=plugins=grpc:. proto/*.proto
	go mod tidy
	go test ./utils
	go build -ldflags $(GO_LDFLAGS) -o $(path)

run:
	protoc --go_out=plugins=grpc:. proto/*.proto
	go mod tidy
	go build -ldflags $(GO_LDFLAGS) -o /tmp/comet && /tmp/comet server

install:
	protoc --go_out=plugins=grpc:. proto/*.proto
	go mod tidy
	go build -ldflags $(GO_LDFLAGS) -o /tmp/comet
	sudo mv /tmp/comet /usr/local/bin/
	chmod +x /usr/local/bin/comet

check-path-included:
ifndef path
	$(error path is undefined; ex. path=/tmp/comet)
endif
