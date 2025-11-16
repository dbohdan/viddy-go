all: build test

viddy-go: config.go generator.go go.mod go.sum main.go run.go run_windows.go snapshot.go viddy.go
	CGO_ENABLED=0 go build

.PHONY: build
build: viddy-go ## build binary

.PHONY: clean
clean: ## delete binary
	-rm viddy-go

.PHONY: help
help: ## list makefile targets
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: test
test: ## run go tests
	go test -v ./...

.PHONY: cover
cover: ## display test coverage
	go test -v -race $(shell go list ./... | grep -v /vendor/) -v -coverprofile=coverage.out
	go tool cover -func=coverage.out

.PHONY: fmt
fmt: ## format go files
	gofumpt -w .
	gci write .

.PHONY: lint
lint: ## lint go files
	golangci-lint run -c .golang-ci.yml

.PHONY: release
release: build ## build, checksum, and sign release binaries
	VERSION="$$(./viddy-go --version | awk '{ print $$2 }')" go run script/release.go
