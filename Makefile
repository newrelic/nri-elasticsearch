WORKDIR      := $(shell pwd)

INTEGRATION  := elasticsearch
BINARY_NAME   = nri-$(INTEGRATION)
GO_PKGS      := $(shell go list ./... | grep -v "/vendor/")
GO_FILES     := $(shell find src -type f -name "*.go")
GOFLAGS			 = -mod=readonly
GOLANGCI_LINT	 = github.com/golangci/golangci-lint/cmd/golangci-lint

all: build

build: clean validate test compile

clean:
	@echo "=== $(INTEGRATION) === [ clean ]: Removing binaries and coverage file..."
	@rm -rfv bin coverage.xml

validate:
ifeq ($(strip $(GO_FILES)),)
	@echo "=== $(INTEGRATION) === [ validate ]: no Go files found. Skipping validation."
else
	@printf "=== $(INTEGRATION) === [ validate ]: running golangci-lint & semgrep... "
	@go run  $(GOFLAGS) $(GOLANGCI_LINT) run --verbose
	@if [ -f .semgrep.yml ]; then \
        docker run --rm -v "${PWD}:/src:ro" --workdir /src returntocorp/semgrep -c .semgrep.yml ; \
    else \
    	docker run --rm -v "${PWD}:/src:ro" --workdir /src returntocorp/semgrep -c p/golang ; \
    fi
endif

compile:
	@echo "=== $(INTEGRATION) === [ compile ]: Building $(BINARY_NAME)..."
	@go build -o bin/$(BINARY_NAME) ./src

test:
	@echo "=== $(INTEGRATION) === [ test ]: Running unit tests..."
	@go test -race ./... -count=1

integration-test:
	@echo "=== $(INTEGRATION) === [ test ]: running integration tests..."
	@docker-compose -f tests/integration/docker-compose.yml up -d --build
	@go test -v -tags=integration ./tests/integration/. -count=1 ; (ret=$$?; docker-compose -f tests/integration/docker-compose.yml down && exit $$ret)

# Include thematic Makefiles
include $(CURDIR)/build/ci.mk
include $(CURDIR)/build/release.mk

.PHONY: all build clean validate compile test integration-test
