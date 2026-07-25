# If you just run 'make' we default to the 'up' task.
.DEFAULT_GOAL := up


# For our Windows friends.
ifeq ($(OS),Windows_NT)
	SHELL := powershell.exe
	.SHELLFLAGS := -NoProfile -Command
	PWD := $(shell (Get-Item -Path .).FullName)
else
	PWD := $(shell pwd)
endif

VERSION := $(shell cat version)

.PHONY: build
build:
	@docker-compose build
	@docker tag d4bot dillonad/d4bot:${VERSION}

.PHONY: up
up:
	@docker-compose up --watch

.PHONY: publish
publish: build
	@docker push dillonad/${NAME}:${VERSION}
	@docker push dillonad/${NAME}:latest

.PHONY: lint
lint:
	@docker run --rm -v ${PWD}:/app -w /app golangci/golangci-lint:latest golangci-lint run

.PHONY: test
test:
	@docker run --rm -v ${PWD}:/app -w /app golang:1.26 go test -bench=. -cover ./...