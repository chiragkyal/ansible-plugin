# Ansible Plugin for Kubebuilder
# Build configuration for ansible-cli tool

# Variables
BINARY_NAME = ansible-cli
BINARY_DIR = bin
CMD_DIR = ./cmd
MAIN_PACKAGE = $(CMD_DIR)
PROJECT_NAME = ansible-plugin
DOMAIN = example.com
OPERATOR_NAME = memcached-operator

# Build flags
LDFLAGS = -s -w
BUILD_FLAGS = -ldflags="$(LDFLAGS)"

# Colors for output
RED = \033[0;31m
GREEN = \033[0;32m
YELLOW = \033[1;33m
NC = \033[0m # No Color

.DEFAULT_GOAL := help

##@ General

.PHONY: help
help: ## Display this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

##@ Development

.PHONY: build
build: ## Build the ansible-cli binary.
	@echo "$(GREEN)Building $(BINARY_NAME)...$(NC)"
	@mkdir -p $(BINARY_DIR)
	@go build $(BUILD_FLAGS) -o $(BINARY_DIR)/$(BINARY_NAME) $(MAIN_PACKAGE)
	@echo "$(GREEN)✓ Built $(BINARY_NAME) successfully$(NC)"

.PHONY: build-local
build-local: ## Build the ansible-cli binary in the current directory.
	@echo "$(GREEN)Building $(BINARY_NAME) in current directory...$(NC)"
	@go build $(BUILD_FLAGS) -o $(BINARY_NAME) $(MAIN_PACKAGE)
	@echo "$(GREEN)✓ Built $(BINARY_NAME) successfully$(NC)"

.PHONY: install
install: ## Install the ansible-cli binary to $GOPATH/bin.
	@echo "$(GREEN)Installing $(BINARY_NAME)...$(NC)"
	@go install $(BUILD_FLAGS) $(MAIN_PACKAGE)
	@echo "$(GREEN)✓ Installed $(BINARY_NAME) successfully$(NC)"

.PHONY: clean
clean: ## Clean build artifacts.
	@echo "$(YELLOW)Cleaning build artifacts...$(NC)"
	@rm -f $(BINARY_NAME)
	@rm -rf $(BINARY_DIR)
	@rm -rf $(OPERATOR_NAME)
	@echo "$(GREEN)✓ Cleaned successfully$(NC)"

##@ Code Quality

.PHONY: fmt
fmt: ## Run go fmt against code.
	@echo "$(GREEN)Formatting code...$(NC)"
	@go fmt ./...
	@echo "$(GREEN)✓ Code formatted$(NC)"

##@ Dependencies

.PHONY: deps
deps: ## Update dependencies.
	@echo "$(GREEN)Updating dependencies...$(NC)"
	@go mod tidy
	@go mod vendor
	@echo "$(GREEN)✓ Dependencies updated$(NC)"

##@ Sample Generation

.PHONY: generate-sample
generate-sample: build-local ## Generate a sample memcached operator.
	@echo "$(GREEN)Generating sample $(OPERATOR_NAME) operator...$(NC)"
	@rm -rf $(OPERATOR_NAME)
	@mkdir -p $(OPERATOR_NAME)
	@cd $(OPERATOR_NAME) && \
		../$(BINARY_NAME) init --domain $(DOMAIN) && \
		../$(BINARY_NAME) create api --group cache --version v1alpha1 --kind Memcached --generate-role
	@echo "$(GREEN)✓ Sample operator generated in $(OPERATOR_NAME)/ directory$(NC)"

.PHONY: generate
generate: generate-sample

