# If you just run 'make' we default to the 'up' task.
.DEFAULT_GOAL := run


# For our Windows freinds.
ifeq ($(OS),Windows_NT)
	SHELL := powershell.exe
	.SHELLFLAGS := -NoProfile -Command
endif

VERSION := $(shell cat version)

.PHONY: build
build:
	@docker-compose build
	@docker tag d4bot dillonad/d4bot:${VERSION}

.PHONY: run
run:
	docker-compose up -d --build

.PHONY: publish
publish: build
	docker push dillonad/${NAME}:${VERSION}
	docker push dillonad/${NAME}:latest

.PHONY: lint
lint:
	golangci-lint run

.PHONY: test
test:
	@docker run --rm -v $(pwd):/app -w /app golang:1.20 go test -bench=. -cover ./...