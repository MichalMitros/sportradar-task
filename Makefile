pwd = $(shell pwd)
app_name = "score-board"

# help - list of available commands

.PHONY: help
help:              ## show this help
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'

# build / run
.PHONY: build
build:             ## build docker image (e.g. to rebuild image cache)
	@docker build --tag $(app_name) .

.PHONY: run-demo
run-demo: build    ## runs demo presentation of score-board
	@docker run --rm -it $(app_name) demo

.PHONY: run-ui
run-ui: build      ## runs score-board in interactive mode
	@docker run --rm -it $(app_name) ui

.PHONY: run
run: run-ui        ## runs score-board in default (interactive) mode

# testing

.PHONY: test
test:              ## run tests
	@docker run --rm -v $(pwd):/app -w /app golang:1.22.2-alpine3.18 go test ./... -v

# Linters and formatters

.PHONY: golangci
golangci:          ## run golangci-lint
	@docker run --rm \
		-v $(pwd):/src \
		-w /src \
		golangci/golangci-lint:v1.57.2-alpine golangci-lint run
