#! /usr/bin/make -f

VERSION?=$(shell git describe --tags --dirty | cut -c 2-)
BASE_IMAGE?=lgtm-action
BINARY_NAME?=lgtm-action

.PHONY: all
all: package

.PHONY: test
test:
	go test -tags=unit -timeout 30s -short -v `go list ./...  | grep -v /vendor/`

.PHONY: package
package: build 

.PHONY: build
build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w -X main.version=$(VERSION)" -o ${BINARY_NAME}

.PHONY: createImage
createImage:
	docker build -t $(BASE_IMAGE) .

.PHONY: clean
clean:
	rm ${BINARY_NAME}
