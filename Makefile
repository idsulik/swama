# Variables for Go paths
GOPATH ?= $(shell go env GOPATH)
GOBIN ?= $(or $(shell go env GOBIN),$(GOPATH)/bin)
GOOS ?= $(shell go env GOOS)
GOARCH ?= $(shell go env GOARCH)

# Directories and project
BUILD_DIR ?= ./bin
PROJECT = swama

# Define the target binary name
BINARY = $(BUILD_DIR)/$(PROJECT)

# Create the build directory
$(BUILD_DIR):
	mkdir -p $(BUILD_DIR)

# PHONY targets are not actual files, they are commands.
.PHONY: build
build: $(BINARY)

# Build the binary
$(BINARY): $(BUILD_DIR)
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o $(BINARY)

.PHONY: install
install: $(BINARY)
	mkdir -p $(GOBIN)
	rm -f $(GOBIN)/$(PROJECT)
	cp $(BINARY) $(GOBIN)/$(PROJECT)

.PHONY: clean
clean:
	rm -rf $(BUILD_DIR)