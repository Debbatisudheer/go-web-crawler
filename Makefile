# Project variables
APP_NAME=crawler
CMD_DIR=cmd/crawler
BIN_DIR=bin

# Default target
all: build

# Build the binary
build:
	@echo "🔨 Building..."
	go build -o $(BIN_DIR)/$(APP_NAME) ./$(CMD_DIR)

# Run the app
run:
	@echo "🚀 Running..."
	go run ./$(CMD_DIR)

# Install globally
install:
	@echo "📦 Installing..."
	go install ./$(CMD_DIR)

# Run tests
test:
	@echo "🧪 Running tests..."
	go test -v ./...

# Lint / Vet
vet:
	@echo "🔍 Vetting..."
	go vet ./...

# Clean build artifacts
clean:
	@echo "🧹 Cleaning..."
	rm -rf $(BIN_DIR)

# Jenkins integration (just an example, runs test + build)
jenkins: vet test install build
	@echo "✅ Jenkins pipeline complete"

.PHONY: all build run install test vet clean jenkins
