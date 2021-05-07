WORKDIR      := $(shell pwd)

NATIVEOS	 := $(shell go version | awk -F '[ /]' '{print $$4}')
NATIVEARCH	 := $(shell go version | awk -F '[ /]' '{print $$5}')
INTEGRATION  := elasticsearch
BINARY_NAME   = nri-$(INTEGRATION)
GO_PKGS      := $(shell go list ./... | grep -v "/vendor/")
GO_FILES     := $(shell find src -type f -name "*.go")
GOFLAGS			 = -mod=readonly
GOLANGCI_LINT	 = github.com/golangci/golangci-lint/cmd/golangci-lint
GOCOV            = github.com/axw/gocov/gocov
GOCOV_XML		 = github.com/AlekSi/gocov-xml
LINTERS_CFG_URL  = https://raw.githubusercontent.com/alvarocabanas/static-analysis-configs-action/main

all: build

build: check-version clean validate test compile

clean:
	@echo "=== $(INTEGRATION) === [ clean ]: Removing binaries and coverage file..."
	@rm -rfv bin coverage.xml

validate:
ifeq ($(strip $(GO_FILES)),)
	@echo "=== $(INTEGRATION) === [ validate ]: no Go files found. Skipping validation."
else
	@printf "=== $(INTEGRATION) === [ validate ]: running golangci-lint & semgrep... "
	@go run  $(GOFLAGS) github.com/golangci/golangci-lint/cmd/golangci-lint run --verbose
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
	@go run $(GOFLAGS) $(GOCOV) test -race ./... | go run $(GOFLAGS) $(GOCOV_XML) > coverage.xml

integration-test:
	@echo "=== $(INTEGRATION) === [ test ]: running integration tests..."
	@docker-compose -f tests/integration/docker-compose.yml up -d --build
	@go test -v -tags=integration ./tests/integration/. -count=1 ; (ret=$$?; docker-compose -f tests/integration/docker-compose.yml down && exit $$ret)

# Include thematic Makefiles
include $(CURDIR)/build/ci.mk
include $(CURDIR)/build/release.mk

check-version:
ifdef GOOS
ifneq "$(GOOS)" "$(NATIVEOS)"
	$(error GOOS is not $(NATIVEOS). Cross-compiling is only allowed for 'clean' target)
endif
endif
ifdef GOARCH
ifneq "$(GOARCH)" "$(NATIVEARCH)"
	$(error GOARCH variable is not $(NATIVEARCH). Cross-compiling is only allowed for 'clean' target)
endif
endif

.PHONY: all build clean validate compile test integration-test check-version
