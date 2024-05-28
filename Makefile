SHELL := bash

ifndef NO_COLOR
YELLOW=\033[0;33m
CYAN=\033[1;36m
RED=\033[31m
# no color
NC=\033[0m
endif


ARTIFACT_NAME := external-dns-infoblox-webhook

TESTPARALLELISM := 4

WORKING_DIR := $(shell pwd)

.PHONY: clean
clean::
	rm -rf $(WORKING_DIR)/bin

.PHONY: build
build::
	go build -o $(WORKING_DIR)/bin/${ARTIFACT_NAME} ./cmd/webhook
	chmod +x $(WORKING_DIR)/bin/${ARTIFACT_NAME}

.PHONY: test
test::
	go test -v -tags=all -parallel ${TESTPARALLELISM} -timeout 2h -covermode atomic -coverprofile=covprofile ./...

.PHONY: lint-init
lint-init:
	@echo -e "\n$(CYAN)Check for lint dependencies$(NC)"
	brew install golangci-lint
	brew install gitleaks
	brew install yamllint

.PHONY: lint
lint: test
	@echo -e "\n$(YELLOW)Running the linters$(NC)"
	@echo -e "\n$(CYAN)golangci-lint$(NC)"
	goimports -w ./
	golangci-lint run -c ./.golangci.toml
	@echo -e "\n$(CYAN)yamllint$(NC)"
	yamllint .
	@echo -e "\n$(CYAN)gitleaks$(NC)"
	gitleaks detect . --no-git --verbose --config=.gitleaks.toml
