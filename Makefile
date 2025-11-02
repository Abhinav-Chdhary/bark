.PHONY: build test clean install run lint help

# Binary name
BINARY_NAME=bark
# Build directory
BUILD_DIR=out

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod
GOVET=$(GOCMD) vet
GOFMT=$(GOCMD) fmt

help: ## Display this help screen
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

build: ## Build the binary
	$(GOBUILD) -o $(BUILD_DIR)/$(BINARY_NAME) ./cmd/bark

build-all: ## Build for multiple platforms
	GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BUILD_DIR)/$(BINARY_NAME)-linux-amd64 ./cmd/bark
	GOOS=darwin GOARCH=amd64 $(GOBUILD) -o $(BUILD_DIR)/$(BINARY_NAME)-darwin-amd64 ./cmd/bark
	GOOS=darwin GOARCH=arm64 $(GOBUILD) -o $(BUILD_DIR)/$(BINARY_NAME)-darwin-arm64 ./cmd/bark
	GOOS=windows GOARCH=amd64 $(GOBUILD) -o $(BUILD_DIR)/$(BINARY_NAME)-windows-amd64.exe ./cmd/bark

test: ## Run tests
	$(GOTEST) -v ./...

test-coverage: ## Run tests with coverage
	$(GOTEST) -v -cover ./...

test-coverage-html: ## Generate HTML coverage report
	$(GOTEST) -coverprofile=coverage.out ./...
	$(GOCMD) tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

lint: ## Run linters
	$(GOVET) ./...
	$(GOFMT) ./...

clean: ## Clean build artifacts
	$(GOCLEAN)
	rm -f $(BUILD_DIR)/$(BINARY_NAME)
	rm -f $(BUILD_DIR)/$(BINARY_NAME)-*
	rm -f coverage.out coverage.html

install: ## Install the binary
	$(GOCMD) install ./cmd/bark

uninstall: ## Uninstall the binary
	rm -f $(shell $(GOBIN))/bin/$(BINARY_NAME)

reinstall: ## Reinstall the binary
	$(MAKE) uninstall
	$(MAKE) install

run: ## Run the application
	$(GOBUILD) -o $(BUILD_DIR)/$(BINARY_NAME) ./cmd/bark
	./$(BINARY_NAME) .

run-testdata: ## Run on testdata directory
	$(GOBUILD) -o $(BUILD_DIR)/$(BINARY_NAME) ./cmd/bark
	./$(BINARY_NAME) testdata

run-json: ## Run on testdata with JSON output
	$(GOBUILD) -o $(BUILD_DIR)/$(BINARY_NAME) ./cmd/bark
	./$(BINARY_NAME) -format json testdata

deps: ## Download dependencies
	$(GOMOD) download
	$(GOMOD) tidy

update-deps: ## Update dependencies
	$(GOGET) -u ./...
	$(GOMOD) tidy

verify: ## Verify dependencies
	$(GOMOD) verify

