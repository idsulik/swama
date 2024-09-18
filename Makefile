# Variables for Go paths
GOPATH ?= $(shell go env GOPATH)
GOBIN ?= $(or $(shell go env GOBIN),$(GOPATH)/bin)
GOOS ?= $(shell go env GOOS)
GOARCH ?= $(shell go env GOARCH)

# Directories and project settings
BUILD_DIR ?= ./bin
PROJECT_NAME = swama
BINARY = $(BUILD_DIR)/$(PROJECT_NAME)

# PHONY targets
.PHONY: all build install clean

# Default target
all: build

# Ensure the build directory exists and build the binary
build: | $(BUILD_DIR)
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o $(BINARY)

$(BUILD_DIR):
	mkdir -p $(BUILD_DIR)

# Install the binary to GOBIN
install: build
	install -d $(GOBIN)
	install -m 755 $(BINARY) $(GOBIN)/$(PROJECT_NAME)

# Clean build artifacts
clean:
	rm -rf $(BUILD_DIR)