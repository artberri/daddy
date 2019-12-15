.PHONY: help all formatcheck format test vet lint qa coverage

GOFMT_FILES?=$$(find . -name '*.go' | grep -v vendor)
FILES_WITHOUT_PROPER_FORMAT?=$$(gofmt -l ${GOFMT_FILES})

help:
	@echo ""
	@echo "The following commands are available:"
	@echo ""
	@echo "    make qa          : Ensure project quality"
	@echo "    make format      : Format the source code"
	@echo "    make formatcheck : Check if the source code has been formatted"
	@echo "    make vet         : Look for suspicious constructs"
	@echo "    make test        : Run tests"
	@echo "    make lint        : Check for style errors"
	@echo ""

all: help

formatcheck:
	@echo "Checking file format..."
	@if [ ! -z "$(FILES_WITHOUT_PROPER_FORMAT)" ]; then \
		echo "The following files have formatting errors:"; \
		echo "$(FILES_WITHOUT_PROPER_FORMAT)"; \
		exit 1; \
	else \
		echo "OK"; \
	fi;

format:
	@echo "Formatting files..."
	@gofmt -w $(GOFMT_FILES)
	@echo "OK"

test:
	@echo "Testing..."
	@go test -i ./... || exit 1
	@go test -timeout=60s -parallel=4 ./...

vet:
	@echo "Looking for suspicious constructs..."
	@go vet ./...
	@echo "OK"

lint:
	@echo "Checking for style errors..."
	@golint ./...
	@test -z "$$(golint ./...)"
	@echo "OK"

qa: formatcheck vet lint test

COVERAGE_MODE    = atomic
COVERAGE_PROFILE = $(COVERAGE_DIR)/profile.out
COVERAGE_XML     = $(COVERAGE_DIR)/coverage.xml
COVERAGE_HTML    = $(COVERAGE_DIR)/index.html
coverage: COVERAGE_DIR := $(CURDIR)/coverage/$(shell date -u +"%Y-%m-%dT%H:%M:%SZ")
coverage:
	@echo "Running coverage..."
	mkdir -p $(COVERAGE_DIR)
	go test \
		-covermode=$(COVERAGE_MODE) \
		-coverprofile="$(COVERAGE_PROFILE)" ./...
	go tool cover -html=$(COVERAGE_PROFILE) -o $(COVERAGE_HTML)
