.DEFAULT_GOAL := help

GOLANGCI_LINT_VERSION ?= v1.19.1
BENCH ?= .

.PHONY: help
help:
	@grep -E '^[a-zA-Z-]+:.*?## .*$$' Makefile | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "[32m%-14s[0m %s\n", $$1, $$2}'

.PHONY: test
test: ## run tests
	go test -race -tags="$(TAGS)" $$(go list ./... | grep -v /vendor/)

.PHONY: bench
bench: ## run benchmarks
	go test -bench="$(BENCH)" $$(go list ./... | grep -v /vendor/)

.PHONY: install-collections-gen
install-collections-gen:
	@command -v collections-gen > /dev/null 2>&1 || \
	  GO111MODULE=off go get github.com/martinohmann/collections-gen

.PHONY: generate
generate: install-collections-gen ## run go generate
	go generate $$(go list ./... | grep -v /vendor/)

.PHONY: verify-codegen
verify-codegen: ## verify generated code is uptodate
	scripts/verify-codegen

.PHONY: coverage
coverage: ## generate code coverage
	scripts/coverage

.PHONY: lint
lint: ## run golangci-lint
	command -v golangci-lint > /dev/null 2>&1 || \
	  curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $$(go env GOPATH)/bin $(GOLANGCI_LINT_VERSION)
	golangci-lint run
