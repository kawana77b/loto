.PHONY: help init build test clean fmt vet run install credits release

# Default target
.DEFAULT_GOAL := help

# Binary name
BINARY_NAME=loto

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOFMT=$(GOCMD) fmt
GOVET=$(GOCMD) vet
GOINSTALL=$(GOCMD) install

help: ## Display this help screen
	@echo "Available targets:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2}'

init: ## Install necessary tools
	@echo "Installing gocredits..."
	$(GOINSTALL) github.com/Songmu/gocredits/cmd/gocredits@latest

build: ## Build the application
	@echo "Building $(BINARY_NAME)..."
	$(GOBUILD) -o $(BINARY_NAME) -v .
	@echo "Build complete: $(BINARY_NAME)"

test: ## Run all tests
	@echo "Running tests..."
	$(GOTEST) -v ./...
	@echo "Tests complete"

test-coverage: ## Run tests with coverage report
	@echo "Running tests with coverage..."
	$(GOTEST) -v -coverprofile=coverage.out ./...
	$(GOCMD) tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

clean: ## Remove build artifacts and cache
	@echo "Cleaning..."
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f coverage.out coverage.html
	rm -rf dist
	@echo "Clean complete"

fmt: ## Format Go source code
	@echo "Formatting code..."
	$(GOFMT) ./...
	@echo "Format complete"

vet: ## Run go vet for static analysis
	@echo "Running go vet..."
	$(GOVET) ./...
	@echo "Vet complete"

run: build ## Build and run the application
	@echo "Running $(BINARY_NAME)..."
	./$(BINARY_NAME)

install: ## Install the binary to $GOPATH/bin
	@echo "Installing $(BINARY_NAME)..."
	$(GOINSTALL) .
	@echo "Install complete"

release: ## Create a release using goreleaser (snapshot)
	goreleaser release --snapshot

credits: ## Add credit information to source files
	@echo "Adding credits to source files..."
	gocredits . > CREDITS

all: fmt vet test build ## Run format, vet, test, and build
	@echo "All tasks complete"
