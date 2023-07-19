WORKDIR      := $(shell pwd)
INTEGRATION  := elasticsearch
BINARY_NAME   = nri-$(INTEGRATION)
GO_PKGS      := $(shell go list ./... | grep -v "/vendor/")
GO_FILES     := $(shell find src -type f -name "*.go")

all: build

build: clean test compile

clean:
	@echo "=== $(INTEGRATION) === [ clean ]: Removing binaries and coverage file..."
	@rm -rfv bin coverage.xml

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

# rt-update-changelog runs the release-toolkit run.sh script by piping it into bash to update the CHANGELOG.md.
# It also passes down to the script all the flags added to the make target. To check all the accepted flags,
# see: https://github.com/newrelic/release-toolkit/blob/main/contrib/ohi-release-notes/run.sh
#  e.g. `make rt-update-changelog -- -v`
rt-update-changelog:
	curl "https://raw.githubusercontent.com/newrelic/release-toolkit/v1/contrib/ohi-release-notes/run.sh" | bash -s -- $(filter-out $@,$(MAKECMDGOALS))

# Include thematic Makefiles
include $(CURDIR)/build/ci.mk
include $(CURDIR)/build/release.mk

.PHONY: all build clean compile test integration-test rt-update-changelog
