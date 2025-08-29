# Makefile for the goexts/generic project

.PHONY: help all test lint

# Define the default target. This will run when you just type `make`.
all: test lint

help: ## ✨ Show this help message
	@echo "Usage: make <target>"
	@echo ""
	@echo "Targets:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2}'

test: ## 🧪 Run all Go tests
	@echo "Running tests..."
	@go test -v -race -cover ./...

lint: ## 🧹 Lint the codebase with golangci-lint
	@echo "Running linter..."
	@golangci-lint run ./...

# Note: To use the 'lint' target, you need to install golangci-lint first.
# See: https://golangci-lint.run/usage/install/

# To ensure that targets like 'test' and 'lint' always run, even if files
# with those names exist, we declare them as .PHONY.
.PHONY: all help test lint