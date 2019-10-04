.DEFAULT_GOAL := help

.PHONY: help
help:
	@grep -E '^[a-zA-Z-]+:.*?## .*$$' Makefile | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "[32m%-12s[0m %s\n", $$1, $$2}'

.PHONY: build
build: ## build collections-gen
	go build \
		-ldflags "-s -w" \
		./cmd/collections-gen

.PHONY: install
install: build ## install collections-gen
	cp collections-gen $(GOPATH)/bin/

.PHONY: test
test: ## run tests
	go test -race -tags="$(TAGS)" $$(go list ./... | grep -v /vendor/)

.PHONY: vet
vet: ## run go vet
	go vet $$(go list ./... | grep -v /vendor/)

.PHONY: generate
generate: ## run go generate
	go generate $$(go list ./... | grep -v /vendor/)

.PHONY: coverage
coverage: ## generate code coverage
	scripts/coverage

.PHONY: misspell
misspell: ## check spelling in go files
	misspell *.go

.PHONY: lint
lint: ## lint go files
	golint ./...
