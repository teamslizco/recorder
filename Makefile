SHELL=/bin/bash

RECORDER_DB_HOST ?= localhost

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
	CGO_ENABLED=0	gox -osarch "${ARCHS}" -ldflags "-X github.com/teamslizco/recorder/main.version=${VERSION} -X github.com/teamslizco/recorder/main.gitHash=${GIT_HASH}"  -output pkg/{{.OS}}_{{.Arch}}/recorder

deps: ${GO_NOVENDOR}
	glide -q install

fmt:
	diff -u <(echo -n) <(gofmt -s -d ${GO_NOVENDOR})

envs:
	RECORDER_DB_HOST=$(RECORDER_DB_HOST) \
	RECORDER_DB_PORT=5432 \
	RECORDER_DB_NAME=recorder \
	RECORDER_DB_USER=recorder \
	RECORDER_DB_PASSWORD=redrocer

test: envs deps
	go test ./... -v ${PROJECT_NOVENDOR}
vet:
	go vet -x ${PROJECT_NOVENDOR}
