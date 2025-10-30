# Makefile for the goexts/generic project.

# Set the default target. This will run when you just type `make`.
.DEFAULT_GOAL := help

# Phony targets prevent conflicts with files of the same name.
.PHONY: all help test lint clean docs docs-tools

# ==============================================================================
# MAIN TARGETS
# ==============================================================================

all: test lint ## Run all checks (test and lint).

help: ## Show this help message.
	@echo "Usage: make [target]"
	@echo ""
	@echo "Targets:"
	@awk 'BEGIN {FS = ":.*?## "; printf "  \033[36m%-22s\033[0m %s\n", "Target", "Description"; printf "  ----------------------   ----------\n"} /^[a-zA-Z0-9_-]+:.*?## / {printf "  \033[36m%-22s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

# ==============================================================================
# DEVELOPMENT TARGETS
# =================================m=============================================

test: ## Run all Go tests.
	@echo "Running tests..."
	@go test -v -race -cover ./...

lint: ## Lint the codebase with golangci-lint.
	@echo "Running linter..."
	@golangci-lint run ./...

clean: ## Clean up generated files.
	@echo "Cleaning up..."
	@rm -f coverage.txt coverage.html
ifeq ($(OS),Windows_NT)
	@if exist .\docs\api rmdir /s /q .\docs\api
else
	@rm -rf docs/api
endif

# ==============================================================================
# DOCUMENTATION TARGETS
# ==============================================================================

docs: docs-tools ## Generate API documentation.
	@echo "Generating documentation..."
ifeq ($(OS),Windows_NT)
	@if not exist .\docs\api mkdir .\docs\api
	@gomarkdoc --exclude-dirs ./docs/... ./... > .\docs\api\api.md
else
	@mkdir -p docs/api
	@gomarkdoc --exclude-dirs ./docs/... ./... > docs/api/api.md
endif
	@echo "Documentation generated in docs/api/api.md"

docs-tools: ## Install documentation generation tools.
	@echo "Checking for gomarkdoc..."
ifeq ($(OS),Windows_NT)
	@where gomarkdoc >nul 2>nul || (echo "gomarkdoc not found. Installing..." && go install github.com/princjef/gomarkdoc/cmd/gomarkdoc@latest)
else
	@command -v gomarkdoc >/dev/null 2>&1 || (echo "gomarkdoc not found. Installing..." && go install github.com/princjef/gomarkdoc/cmd/gomarkdoc@latest)
endif
