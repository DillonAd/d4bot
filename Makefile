NAME = d4bot
VERSION := $(shell cat version)

.PHONY: build
build:
	docker build -t ${NAME} .
	docker tag ${NAME} dillonad/${NAME}:${VERSION}
	docker tag ${NAME} dillonad/${NAME}:latest

.PHONY: run
run: build
	docker run -it --rm dillonad/${NAME}:latest

.PHONY: publish
publish: build
	docker push dillonad/${NAME}:${VERSION}
	docker push dillonad/${NAME}:latest

.PHONY: lint
lint:
	golangci-lint run

.PHONY: test
test:
	go test -bench=. -cover ./...