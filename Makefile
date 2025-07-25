GOLANGCI_LINT_VERSION ?= v2.1.5

init:
	go mod tidy

install-lint-deps:
	(which golangci-lint > /dev/null) || curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(shell go env GOPATH)/bin $(GOLANGCI_LINT_VERSION)

lint: install-lint-deps
	golangci-lint run ./...

lint-fix: install-lint-deps
	golangci-lint run ./... --fix

.PHONY: init run generate migrate-create lint lint-fix install-lint-deps