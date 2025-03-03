# Makefile for building the grantmaster project

# Variables
BINARY_NAME = gm
BIN_DIR = bin
GOFLAGS = GOOS=linux GOARCH=amd64  # Change as needed for target OS/arch

# Default target
all: build

# Build the Go project
build:
	@echo "Building the project..."
	@mkdir -p $(BIN_DIR)
	$(GOFLAGS) go build -C code/ -o ../$(BIN_DIR)/$(BINARY_NAME) -v

# Clean the bin directory
clean:
	@echo "Cleaning up the bin directory..."
	@rm -rf $(BIN_DIR)

# Run the application
run: build
	@echo "Running the application..."
	./$(BIN_DIR)/$(BINARY_NAME)

# Specify who to ignore
.PHONY: all build clean run
