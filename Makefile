SHELL=/bin/bash

GOVERSION:=$(shell go version)

ARCHS:=linux/amd64
CUR_DIR=$(shell pwd)

GO_NOVENDOR:=$(shell find . -type f -name \*.go -not -path ./vendor/\* -not -path ./.glide/\*)
PROJECT_NOVENDOR:=$(shell glide novendor)

VERSION:=$(shell echo devel)
GIT_HASH=$(shell git rev-parse HEAD)

.PHONY: build fmt vet

default: build fmt vet

build:
	CGO_ENABLED=0	gox -osarch "${ARCHS}" -output pkg/{{.OS}}_{{.Arch}}/recorder ./cmd/recorder

deps: ${GO_NOVENDOR}
	glide -q install

fmt:
	diff -u <(echo -n) <(gofmt -s -d ${GO_NOVENDOR})

qtest:
	go test -v ${PROJECT_NOVENDOR}
test: deps
	go test ${PROJECT_NOVENDOR}

vet:
	go vet ${PROJECT_NOVENDOR}
