# Makefile for building the grantmaster project

# Variables
BINARY_NAME = gm
BIN_DIR = bin
GOFLAGS = GOOS=linux GOARCH=amd64  # Change as needed for target OS/arch

ASSETS_TEMP_DIR=code/modules/template/assets
TEMPLATES_DIR=code/share/templates

# Default target
all: build

# Build the Go project
build: copy_assets
	@echo "Building the project..."
	@mkdir -p $(BIN_DIR)
	$(GOFLAGS) go build -C code/ -o ../$(BIN_DIR)/$(BINARY_NAME) -v

# Copy assets
copy_assets:
	@rm -rf $(ASSETS_TEMP_DIR)
	@mkdir -p $(ASSETS_TEMP_DIR)
	@cp -prf $(TEMPLATES_DIR)/* $(ASSETS_TEMP_DIR)/

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
