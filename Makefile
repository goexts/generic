# Makefile for the goexts/generic project

.PHONY: help all test lint

# Define the default target. This will run when you just type `make`.
all: test lint

help: ## ✨ Show this help message
	@echo "Usage: make <target>"
	@echo ""
	@echo "Targets:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2}'

test: ##  Run all Go tests
	@echo "Running tests..."
	@go test -v -race -cover ./...

#  清理生成的文件
clean:
	@rm -f coverage.txt coverage.html

#  生成文档
docs:
	@echo "Installing gomarkdoc..."
	@go get github.com/princjef/gomarkdoc/cmd/gomarkdoc
	@echo "Generating documentation..."
	@go generate ./...

#  安装文档工具
docs-tools:
	@go install github.com/princjef/gomarkdoc/cmd/gomarkdoc@latest/...

lint: ##  Lint the codebase with golangci-lint
	@echo "Running linter..."
	@golangci-lint run ./...

# See: https://golangci-lint.run/usage/install/

# To ensure that targets like 'test' and 'lint' always run, even if files
# with those names exist, we declare them as .PHONY.
.PHONY: all help test lint docs docs-tools clean